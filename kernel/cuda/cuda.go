package main

import (
    `fmt`
    `github.com/Qitmeer/qitmeer/common/hash`
    `github.com/Qitmeer/qitmeer/crypto/cuckoo/siphash`
)

//#cgo CFLAGS: -I. 
//#cgo LDFLAGS: -L. -ltest 
//#cgo LDFLAGS: -lcudart 
//#include "test.h"
import "C" 


func main() { 
    fmt.Printf("Invoking cuda library...\n")
    str := "hello world"
    header := make([]byte,113)
    copy(header[:],[]byte(str))
    deviceID := 0
    h := hash.HashH(header)
    fmt.Println(h)
    sip := siphash.Newsip(h[:])
    s := C.CString(string(header))
    fmt.Printf("%x %x %x %x \n",sip.V[0],sip.V[1],sip.V[2],sip.V[3])
    fmt.Println("Done ", C.test_cuda((C.int)(deviceID),s))
}