// +build six
// db.go - Functions for database handling.
//
// Copyright (c) 2013 The go-alpm Authors
//
// MIT Licensed. See LICENSE for details.

package alpm

/*
#include <alpm.h>
#include <alpm_list.h>
*/
import "C"

import (
	"unsafe"
)

// Search returns a list of packages matching the targets.
// In case of error the Package List will be nil
func (db *DB) Search(targets []string) IPackageList {
	var needles *C.alpm_list_t = nil
	var ret *C.alpm_list_t = nil

	for _, str := range targets {
		needles = C.alpm_list_add(needles, unsafe.Pointer(C.CString(str)))
	}

	ok := C.alpm_db_search(db.ptr, needles, &ret)
	if ok != 0 {
		return PackageList{nil, db.handle}
	}

	C.alpm_list_free(needles)
	return PackageList{(*list)(unsafe.Pointer(ret)), db.handle}
}
