/*
Unless explicitly stated otherwise all files in this repository are licensed
under the $license_for_repo License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package main

import (
	"fmt"

	"github.com/DataDog/go-python3"
)

func main() {
	python3.Py_Initialize()
	var err error
	if !python3.Py_IsInitialized() {
		panic("Error initializing the python interpreter")
	}
	err = python3.PySys_SetPath("./")
	if err != nil {
		panic("ERROR: Path set error.")
	}
	pModule := python3.PyImport_ImportModule("testFunc")
	if pModule == nil {
		panic("ERROR: Module not found.")
	}
	pFunc := pModule.GetAttrString("testF")
	if !python3.PyCallable_Check(pFunc) {
		panic("ERROR: Func not found.")
	}
	args := python3.PyTuple_New(1)
	arg := python3.PyLong_FromGoInt(8)
	python3.PyTuple_SetItem(args, 0, arg)
	pRet := pFunc.CallObject(args)
	if pRet != nil {
		fmt.Println("ret=", python3.PyLong_AsLong(pRet))
	}
	pRet = pFunc.CallObject(args)
	if pRet != nil {
		fmt.Println("ret=", python3.PyLong_AsLong(pRet))
	}
	python3.Py_Finalize()
	return
}
