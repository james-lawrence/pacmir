// +build !six
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

func (db *DB) Search(targets []string) IPackageList {
	var needles *C.alpm_list_t

	for _, str := range targets {
		needles = C.alpm_list_add(needles, unsafe.Pointer(C.CString(str)))
	}

	pkglist := (*list)(unsafe.Pointer(C.alpm_db_search(db.ptr, needles)))
	C.alpm_list_free(needles)
	return PackageList{pkglist, db.handle}
}
