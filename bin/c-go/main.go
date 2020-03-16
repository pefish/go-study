// // +build linux,386 darwin,!cgo  // 条件编译。只有在”linux/386“或”darwin平台下非cgo环境“才进行构建
// // +build debug  // build时指定了debug tag时才编译。例如 go build -tags="debug haha"

package main

/*
// #cgo windows CFLAGS: -DCGO_OS_WINDOWS=1  // windows系统下设置CGO_OS_WINDOWS标志
// #cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
// #cgo linux CFLAGS: -DCGO_OS_LINUX=1
//#if defined(CGO_OS_WINDOWS)
//    const char* os = "windows";
//#elif defined(CGO_OS_DARWIN)
//    const char* os = "darwin";
//#elif defined(CGO_OS_LINUX)
//    const char* os = "linux";
//#else
//#error(unknown os)
//#endif

// #cgo CFLAGS: -DPNG_DEBUG=1 -I./include  // #cgo可以指定编译选项
// #cgo LDFLAGS: -L/usr/local/lib -lpng
#include "hello.h"
 */
import "C"  // 表示将调用cgo命令生成对应的中间文件
import (
	"fmt"
	"go-study/bin/c-go/qsort"
	"unsafe"
)

func main() {
	C.CSayHello(C.CString("CSayHello: Hello World!!!\n"))
	fmt.Println("Go: Hello World!!!")
	C.GoSayHello(C.CString("GoSayHello: Hello World!!!\n"))

	// 访问C语言中的struct
	var a C.struct_A
	fmt.Println(a._type)
	fmt.Println(a.i)

	// 访问C语言中的union类型
	var b C.union_B;
	fmt.Println("b.i:", *(*C.int)(unsafe.Pointer(&b)))
	fmt.Println("b.f:", *(*C.float)(unsafe.Pointer(&b)))

	// 访问C语言中的enum类型
	var c C.enum_C = C.TWO
	fmt.Println(c)
	fmt.Println(C.ONE)
	fmt.Println(C.TWO)

	// 调用具有返回值的C函数
	v0, err0 := C.Div(2, 1)
	fmt.Println(v0, err0)
	v1, err1 := C.Div(1, 0)
	fmt.Println(v1, err1)

	// 调用go封装的C函数
	values := []int64{42, 9, 101, 95, 27, 25}
	qsort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	fmt.Println(values)
}


