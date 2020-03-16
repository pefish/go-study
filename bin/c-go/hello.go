// Go语言实现 GoSayHello api

package main

import "C"

import "fmt"

//export GoSayHello  // export前面不能有空格
func GoSayHello(s *C.char) {
	fmt.Print(C.GoString(s))
}