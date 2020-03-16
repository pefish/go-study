
// C语言实现 CSayHello api

#include "hello.h"
#include <stdio.h>
#include <errno.h>

void CSayHello(const char* s) {
//    puts(s);
    GoSayHello(s);  // 去调用Go语言实现的函数（回调Go语言）
}

int Div(int a, int b) {
    if(b == 0) {
        errno = EINVAL;
        return 0;
    }
    return a/b;
}
