
```shell script
gcc -c -o number.o number.c  // 编译

ar rcs libnumber.a number.o  // 打包静态链接库

gcc -shared -o libnumber1.so number.c // 构建动态链接库
```
