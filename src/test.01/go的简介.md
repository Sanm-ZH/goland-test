### go语言简介
#### 为什么用GO
- 天生支持并发
- 执行性能好
- 开发效率高
- 21世纪的C语言
- 新领域首选语言

```go
import "fmt"

func main() {
    fmt.Println("人生苦短，Let's Go!")
}
```
#### `go build`
- 会生成一个当前系统可执行的文件
- `go build -o xxxx` xxx为自定义生成文件名

#### `go run`
- 直接执行`.go`文件

#### `go install` （两步）
1. 先编译得到一个可执行文件
2. 将可执行文件拷贝到`GOPATH/bin`目录下

#### 交叉编译(跨平台编译)
```shell script
# Windows
# 编译后可以得到linux平台执行文件
SET CGO_ENABLED=0  # 禁用CGO
SET GOOS=linux  # 目标平台是linux
SET GOARCH=amd64  # 目标处理器架构是amd64
# Windows下编译Mac平台64位可执行程序：
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64

# mac
# Mac 下编译 Linux 和 Windows平台 64位 可执行程序
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

# Linux
# Linux 下编译 Mac 和 Windows 平台64位可执行程序：
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
go build
```

`ps: 程序入口main函数`