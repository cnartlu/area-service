package zip7

/*
#include "main.h"
*/
import "C"

type Zip struct {
	c7zip *C.C7ZIP
}

func (z *Zip) List() error {
	return nil
}

func (z *Zip) Size() int64 {
	return 0
}

func (z *Zip) Extract(path string) error {
	for i := 0; i < int(z.c7zip.db.NumFiles); i++ {
		// C.SzArEx_IsDir(z.c7zip.db, i)
	}
	C.extract(z.c7zip, C.CString(path))
	return nil
}

func (z *Zip) Close() error {
	C.close(z.c7zip)
	if int(z.c7zip.res) > 0 {

	}
	return nil
}

func Open(name string) (*Zip, error) {
	data := &Zip{}
	a := C.open(C.CString(name))
	data.c7zip = &a
	return data, nil
}
