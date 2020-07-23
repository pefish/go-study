
## Start

```shell script
DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:`pwd`/bin/c-go/number go run ./bin/c-go/  // 添加一个库查找路径
```

或者

```shell script
go build ./bin/c-go/

install_name_tool -change libnumber1.so `pwd`/bin/c-go/number/libnumber1.so ./c-go  // 修改库的查找路径  otool -L ./c-go 可以查看依赖库的查找路径

./c-go

```

## Go语言和C语言类型对比
|  C语言类型  | CGO类型  | Go语言类型 |
|  ----  | ----  | ---- |
| char|	C.char|	byte|
|singed char	|C.schar|	int8|
|unsigned char	|C.uchar|	uint8|
|short	|C.short	|int16|
|unsigned short	|C.ushort|	uint16|
|int	|C.int	|int32|
|unsigned int	|C.uint|	uint32|
|long	|C.long	|int32|
|unsigned long	|C.ulong|	uint32|
|long long int	|C.longlong	|int64|
|unsigned long long int|	C.ulonglong|	uint64|
|float	|C.float	|float32|
|double	|C.double	|float64|
|size_t|	C.size_t|	uint|
|int8_t	|C.int8_t	|int8|
|uint8_t|	C.uint8_t	|uint8|
|int16_t|	C.int16_t	|int16|
|uint16_t|	C.uint16_t	|uint16|
|int32_t|	C.int32_t	|int32|
|uint32_t|	C.uint32_t	|uint32|
|int64_t|	C.int64_t	|int64|
|uint64_t|	C.uint64_t	|uint64|
