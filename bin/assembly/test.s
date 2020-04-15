#include "textflag.h"

TEXT ·Add(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX // 参数 a
    MOVQ b+8(FP), BX // 参数 b
    ADDQ BX, AX    // AX += BX
    MOVQ AX, ret+16(FP) // 返回
    RET




DATA  msg<>+0x00(SB)/8, $"Hello, W"  // 初始化变量
DATA  msg<>+0x08(SB)/8, $"orld!\n"
GLOBL msg<>(SB),NOPTR,$16  // 将变量声明为 global，后面需要跟两个参数，flag 和变量的大小，msg 后面有一个<>，这表示这个全局变量只在当前文件中可以被访问，类似于 C 语言中的 static

// TEXT 表示是汇编中的 .text 分段，
// 注意到 TEXT 和 PrintMe中间除了一个空格以外，还有一个反人类的「中点·」，这个中点在编译以后会被替换为.，同时也会加上包名，比如这里的 helloworld.PrintMe
// NOSPLIT 标志位这里先不介绍
// $0 表示栈帧大小为 0
//
// Mac 下的系统调用参数需要存储在 DI、SI、DX 等寄存器中，系统调用编号存储在 AX 中

TEXT ·PrintMe(SB), NOSPLIT, $0
	MOVL 	$(0x2000000+4), AX 	 // write 系统调用数字编号 4，Mac 下的系统调用数字编号需要加 0x2000000，就是系统约定
	MOVQ 	$1, DI 			      // 第 1 个参数 fd
	LEAQ 	msg<>(SB), SI 		// 第 2 个参数 buffer 指针地址
	MOVL 	$16, DX 		        // 第 3 个参数 count
	SYSCALL
	RET


