#include "Precomp.h"

#include <stdio.h>
#include <string.h>

#include "CpuArch.h"

#include "7z.h"
#include "7zAlloc.h"
#include "7zBuf.h"
#include "7zCrc.h"
#include "7zFile.h"

typedef struct
{
    ISzAlloc *allocImp;

    CFileInStream *archiveStream;

    CLookToRead2 *lookStream;

    CSzArEx *db;
    SRes res;

    UInt16 *temp;
} C7ZIP;

C7ZIP open(char *filename);

void extract(C7ZIP *result, char *fullPath);

void close(C7ZIP *result);