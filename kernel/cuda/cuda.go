package main

//#cgo CFLAGS: -I.
//#cgo LDFLAGS: -L. -ltest
//#cgo LDFLAGS: -lcudart
//#include "test.h"
import "C"
import (
	`fmt`
)

func main()  {
	fmt.Println(C.test_add())
}
