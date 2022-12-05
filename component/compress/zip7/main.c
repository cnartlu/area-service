/* 7zMain.c - Test application for 7z Decoder
2021-04-29 : Igor Pavlov : Public domain */

#include "Precomp.h"
#include "main.h"

#include <stdio.h>
#include <string.h>

#include "CpuArch.h"

#include "7z.h"
#include "7zAlloc.h"
#include "7zBuf.h"
#include "7zCrc.h"
#include "7zFile.h"
#include "7zTypes.h"

#ifndef USE_WINDOWS_FILE
/* for mkdir */
#ifdef _WIN32
#include <direct.h>
#else
#include <stdlib.h>
#include <time.h>
#ifdef __GNUC__
#include <sys/time.h>
#endif
#include <fcntl.h>
// #include <utime.h>
#include <sys/stat.h>
#include <errno.h>
#endif
#endif

#define kInputBufSize ((size_t)1 << 18)

static const ISzAlloc g_Alloc = {SzAlloc, SzFree};

static WRes MyCreateDir(const UInt16 *name)
{
#ifdef USE_WINDOWS_FILE

    return CreateDirectoryW((LPCWSTR)name, NULL) ? 0 : GetLastError();

#else

    CBuf buf;
    WRes res;
    Buf_Init(&buf);
    RINOK(Utf16_To_Char(&buf, name MY_FILE_CODE_PAGE_PARAM));

    res =
#ifdef _WIN32
        _mkdir((const char *)buf.data)
#else
        mkdir((const char *)buf.data, 0777)
#endif
                == 0
            ? 0
            : errno;
    Buf_Free(&buf, &g_Alloc);
    return res;

#endif
}

static WRes OutFile_OpenUtf16(CSzFile *p, const UInt16 *name)
{
#ifdef USE_WINDOWS_FILE
    return OutFile_OpenW(p, (LPCWSTR)name);
#else
    CBuf buf;
    WRes res;
    Buf_Init(&buf);
    RINOK(Utf16_To_Char(&buf, name MY_FILE_CODE_PAGE_PARAM));
    res = OutFile_Open(p, (const char *)buf.data);
    Buf_Free(&buf, &g_Alloc);
    return res;
#endif
}

// #define NUM_PARENTS_MAX 128

C7ZIP open(char *filename)
{
    ISzAlloc allocImp;
    ISzAlloc allocTempImp;

    CFileInStream archiveStream;
    CLookToRead2 lookStream;
    CSzArEx db;
    SRes res;

    UInt16 *temp = NULL;

    C7ZIP zip;
    // zip.db = &db;
    // zip.archiveStream = &archiveStream;
    // zip.lookStream = &lookStream;
    // zip.res = &res;
    // zip.allocImp = &allocImp;
    // zip.temp = temp;

#if defined(_WIN32) && !defined(USE_WINDOWS_FILE) && !defined(UNDER_CE)
    g_FileCodePage = AreFileApisANSI() ? CP_ACP : CP_OEMCP;
#endif

    allocImp = g_Alloc;
    allocTempImp = g_Alloc;

    {
        WRes wres = InFile_Open(&archiveStream.file, filename);
        if (wres != 0)
        {
            zip.res = SZ_ERROR_PARAM;
            return zip;
        }
    }

    FileInStream_CreateVTable(&archiveStream);
    archiveStream.wres = 0;
    LookToRead2_CreateVTable(&lookStream, False);
    lookStream.buf = NULL;

    zip.res = SZ_OK;

    {
        lookStream.buf = (Byte *)ISzAlloc_Alloc(&allocImp, kInputBufSize);
        if (!lookStream.buf)
        {
            zip.res = SZ_ERROR_MEM;
            close(&zip);
            return zip;
        }
        lookStream.bufSize = kInputBufSize;
        lookStream.realStream = &archiveStream.vt;
        LookToRead2_Init(&lookStream);
    }

    CrcGenerateTable();
    SzArEx_Init(&db);

    zip.res = SzArEx_Open(&db, &lookStream.vt, &allocImp, &allocTempImp);
    if (zip.res != SZ_OK)
    {
        close(&zip);
        return zip;
    }

    return zip;
}

void extract(C7ZIP *result, char *fullPath)
{
    UInt32 i;
    /*
    if you need cache, use these 3 variables.
    if you use external function, you can make these variable as static.
    */
    UInt32 blockIndex = 0xFFFFFFFF; /* it can have any value before first call (if outBuffer = 0) */
    Byte *outBuffer = 0;            /* it must be 0 before first call for each new archive. */
    size_t outBufferSize = 0;       /* it can have any value before first call (if outBuffer = 0) */

    size_t tempSize = 0;
    UInt16 *temp = NULL;

    int fullPaths = 1;
    for (i = 0; i < result->db->NumFiles; i++)
    {
        size_t offset = 0;
        size_t outSizeProcessed = 0;
        // const CSzFileItem *f = db.Files + i;
        size_t len;
        const BoolInt isDir = SzArEx_IsDir(result->db, i);

        // len = SzArEx_GetFullNameLen(result.db, i);
        len = SzArEx_GetFileNameUtf16(result->db, i, NULL);
        if (len > tempSize)
        {
            SzFree(NULL, temp);
            tempSize = len;
            temp = (UInt16 *)SzAlloc(NULL, tempSize * sizeof(temp[0]));
            if (!temp)
            {
                result->res = SZ_ERROR_MEM;
                break;
            }
        }

        SzArEx_GetFileNameUtf16(result->db, i, temp);
        /*
        if (SzArEx_GetFullNameUtf16_Back(result.db, i, temp + len) != temp)
        {
          res = SZ_ERROR_FAIL;
          break;
        }
        */

        // 导出文件
        {
            CSzFile outFile;
            size_t processedSize;
            size_t j;
            UInt16 *name = (UInt16 *)temp;
            const UInt16 *destPath = (const UInt16 *)name;

            for (j = 0; name[j] != 0; j++)
            {
                if (name[j] == '/')
                {
                    if (fullPaths)
                    {
                        name[j] = 0;
                        MyCreateDir(name);
                        name[j] = CHAR_PATH_SEPARATOR;
                    }
                    else
                        destPath = name + j + 1;
                }
            }
            if (isDir)
            {
                MyCreateDir(destPath);
                continue;
            }
            else
            {
                WRes wres = OutFile_OpenUtf16(&outFile, destPath);
                if (wres != 0)
                {
                    result->res = SZ_ERROR_FAIL;
                    break;
                }
            }

            processedSize = outSizeProcessed;

            {
                WRes wres = File_Write(&outFile, outBuffer + offset, &processedSize);
                if (wres != 0 || processedSize != outSizeProcessed)
                {
                    result->res = SZ_ERROR_FAIL;
                    break;
                }
            }

            {
                FILETIME mtime;
                FILETIME *mtimePtr = NULL;

#ifdef USE_WINDOWS_FILE
                FILETIME ctime;
                FILETIME *ctimePtr = NULL;
#endif

                if (SzBitWithVals_Check(&result->db->MTime, i))
                {
                    const CNtfsFileTime *t = &result->db->MTime.Vals[i];
                    mtime.dwLowDateTime = (DWORD)(t->Low);
                    mtime.dwHighDateTime = (DWORD)(t->High);
                    mtimePtr = &mtime;
                }

#ifdef USE_WINDOWS_FILE
                if (SzBitWithVals_Check(&result->db->CTime, i))
                {
                    const CNtfsFileTime *t = &result->db->CTime.Vals[i];
                    ctime.dwLowDateTime = (DWORD)(t->Low);
                    ctime.dwHighDateTime = (DWORD)(t->High);
                    ctimePtr = &ctime;
                }

                if (mtimePtr || ctimePtr)
                    SetFileTime(outFile.handle, ctimePtr, NULL, mtimePtr);
#endif

                {
                    WRes wres = File_Close(&outFile);
                    if (wres != 0)
                    {
                        result->res = SZ_ERROR_FAIL;
                        break;
                    }
                }

#ifndef USE_WINDOWS_FILE
#ifdef _WIN32
                mtimePtr = mtimePtr;
#else
                if (mtimePtr)
                    Set_File_FILETIME(destPath, mtimePtr);
#endif
#endif
            }

#ifdef USE_WINDOWS_FILE
            if (SzBitWithVals_Check(&result->db->Attribs, i))
            {
                UInt32 attrib = result->db->Attribs.Vals[i];
                /* p7zip stores posix attributes in high 16 bits and adds 0x8000 as marker.
                   We remove posix bits, if we detect posix mode field */
                if ((attrib & 0xF0000000) != 0)
                    attrib &= 0x7FFF;
                SetFileAttributesW((LPCWSTR)destPath, attrib);
            }
#endif
        }
    }
    ISzAlloc_Free(result->allocImp, outBuffer);
}

void close(C7ZIP *result)
{
    SzFree(NULL, result->temp);
    SzArEx_Free(result->db, result->allocImp);
    ISzAlloc_Free(result->allocImp, result->lookStream->buf);
    File_Close(&result->archiveStream->file);
}