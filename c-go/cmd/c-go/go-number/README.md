
## 将Go实现的函数导出为C静态库

```shell script
go build -buildmode=c-archive -o number.a
```

## 编译C代码
```shell script
gcc -o a.out _test_main.c number.a
```

## 运行C程序
```shell script
./a.out
```


## 将Go实现的函数导出为C动态库

```shell script
go build -buildmode=c-shared -o number.so
```

## 编译C代码
```shell script
gcc -o a.out _test_main.c number.so
```

## 运行C程序
```shell script
./a.out
```
