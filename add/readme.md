
## Rust 编译成 wasm 被 Golang 使用


### Rust 安装工具链

```shell
rustup target add wasm32-unknown-unknown

cargo install wasm-gc
```

### 编译成wasm

```shell
cargo build --target wasm32-unknown-unknown --release
```

### wasm 打包到 golang 二进制中

```shell
go get -u github.com/go-bindata/go-bindata/...

go-bindata -o add.go add/target/wasm32-unknown-unknown/release/add.wasm
```
