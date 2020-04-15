package main

import "fmt"

func Add(a, b int) int // 汇编函数声明，同目录下存在一个.s文件即可使用此语法
func PrintMe()

func main() {
	fmt.Println(Add(10, 11))
	PrintMe()
}

//四个伪寄存器
//FP: Frame pointer，用来访问函数的参数
//PC: Program counter: 用于分支和跳转
//SB: Static base pointer: 一般用于声明函数或者全局变量
//SP: Stack pointer：指向当前栈帧的局部变量的开始位置，一般用来引用函数的局部变量

// 函数定义的语法如下
// TEXT symbol(SB), [flags,] $framesize[-argsize]
// 分为 5 个组成部分：TEXT 指令、函数名、可选的 flags 标志、函数帧大小和可选的函数参数大小