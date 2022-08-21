视频

# **【马哥教育Golang专题课】**

[华科大大佬50个小时讲完的Golang，Go语言全套](https://www.bilibili.com/video/BV1CU4y1d7Vc?spm_id_from=333.337.search-card.all.click)

[马哥教育2021-Go语言开发实战-Prometheus定制化二次开发【高薪必备】](https://www.bilibili.com/video/BV1Tr4y1P7SB?spm_id_from=333.999.0.0)

[马哥教育2021-Go语言开发框架beego入门到精通必备](https://www.bilibili.com/video/BV1Ti4y1K7Tt?spm_id_from=333.999.0.0)

[马哥教育2021-Go语言开发实战-kubernetes二次开发实战【高薪必备】](https://www.bilibili.com/video/BV1s54y1a7jR?spm_id_from=333.999.0.0)

# **【实战项目】**

[golang在发光：Go高频面试之算法](https://www.bilibili.com/video/BV1QF41147eG?spm_id_from=333.999.0.0)

[golang在发光：Go语言微服务/从入门到实战/gRPC Stream/Protobuf/](https://www.bilibili.com/video/BV1ZY41137jx?spm_id_from=333.999.0.0)

[Golang在发光：Go语言如何使用GRPC构建微服务/RPC/Protobuf/GRPC](https://www.bilibili.com/video/BV1mi4y1d7SL?spm_id_from=333.999.0.0)

[Golang在发光：探索骄傲的Go语言/go并发/MPG并发/Channel](https://www.bilibili.com/video/BV1qP4y1b7g9?spm_id_from=333.999.0.0)

[Golang在发光：Go语言开发企业级DevOps平台](https://www.bilibili.com/video/BV1zU4y1P732?spm_id_from=333.999.0.0)

[Golang开发运维架构项目实战 socker beego websocker 运维 监控](https://www.bilibili.com/video/BV12K4y137sL?spm_id_from=333.999.0.0)

[Golang协程！让服务端的研发飞舞!](https://www.bilibili.com/video/BV1Zf4y1W7mH?spm_id_from=333.999.0.0)

# 介绍

## 应用场景

- 容器编排引擎 kubernetes
- Web应用框架BEEGO
- 高性能分布式监控系统falcon
- 分布式可靠KV存储etcd
- 微服务Go kit
- 高性能消息系统NSQ
- 区块链HYPERLEDGER
- 人工只能monmenTa
- 监控报警系统Prometheus
- 机器学习Golearn
- 消息系统NSQ
- DevOps open-falcon
- 爬虫Pholcus
- 微服务istio

## go微服务开发

- 零依赖，让我们可以最小化镜像，节省存储空间
- Runtime使用更小内存，对于Java的JVM
- 更好的并行能力，当你需要更多的CPU的时候
- 更高的性能，对比解释性语言，在处理数据以及并发方面优势更大
- 简单学习成本低，内部人员可以转入Go
- 使用Go能更接近云原生生态，比如docker，k8s，habor都是用Go开发的

# 环境配置

## 下载

[链接](https://golang.google.cn/dl/)

```bash
wget https://golang.google.cn/dl/go1.19.linux-amd64.tar.gz
tar -zxvf go1.19.linux-amd64.tar.gz -C /usr/local/
```

## 环境变量

```bash
sudo vim /etc/profile  
```

```bash
export GOROOT=/usr/local/go   
export GOPATH=$HOME/go  
export GOBIN=$GOPATH/bin  
export GO111MODULE=on 
export GOPROXY=https://goproxy.io 
export PATH=$GOPATH:$GOBIN:$GOROOT/bin:$PATH 
```

GOROOT是go安装目录，go原生工具在该目录下

GOPATH通常存放自己开发的代码或第三方依赖库

GO111MODULE启用go mod管理第三方依赖库

GOPROXY下载依赖走的哪个镜像代理，可以公司内部自建镜像

在GOPATH目录下新建三个目录bin，src，pkg

```bash
source /etc/profile
```

环境变量生效

```bash
mkdir -p $GOPATH/{bin,src,pkg}
```

建立Go的工作空间（workspace，也就是GOPATH环境变量指向的目录）
GO代码必须在工作空间内。工作空间是一个目录，其中包含三个子目录：
src ---- 里面每一个子目录，就是一个包。包内是Go的源码文件
pkg ---- 编译后生成的，包的目标文件
bin ---- 生成的可执行文件

## 测试

```bash
go version
# 查看变量
go env 
```

## vscode安装插件

VSCode Go —— 官方维护的插件（下个视频单独讲解）
Git 工具 —— GitLens
方便查看代码块 —— Bracket Pair Colorizer 2
代码标记 —— Bookmark
API 开发利器 —— Rest Client
API 开发利器 —— Thunder Client

    www.boredapi.com/
    www.boredapi.com/api/activity/

Prettier (全栈开发) —— 开启 format on save 功能
自动关闭标签 —— Auto Close Tag
自动更新标签名称 —— Auto Rename Tag
HTML 简写 —— Emmet

## Go Modules依赖包查找机制

下载的第三方依赖存储在GOPATH/pkg/mod下

go install生成的可执行文件存储在GOPATH/bin下

依赖包查找顺序

- 工作目录
- GOPATH/pkg/mod
- GOROOT/src

# 常用指令

## go help

查看文档

```bash
go help build
```

## go build

对源码和依赖的文件进行打包，生成可执行文件。

```bash
go build <源文件>
```

```bash
go build -o <可执行文件名> <源文件>
```

## go install

编译并安装包或依赖，安装到GOPATH/bin目录下

```bash
go install <源文件>
```

## go get

```bash
go get <第三方包>
```

把依赖库添加到module中，如果本机之前从未安装过则先下载并安装。

会在GOPATH/pkg/mod目录下是生成包目录，同时在GOPATH/bin目录下生成可执行文件。

## go mod tidy

引用项目需要的依赖增加到go.mod文件。

去掉go.mod文件中项目不需要的依赖。

## go run

编译并运行程序

## go test

执行测试代码

## go tool

执行go自带的工具

go tool pprof对cpu、内存和协程进行监控

go tool trace跟踪协程的执行情况

## go vet

检查代码中的静态错误

# 变量

## 变量类型

| 类型     | go变量类型                                                      | fmt输出  |
| -------- | --------------------------------------------------------------- | -------- |
| 整数     | int int8 int16 int32 int64 uint uint8<br />uint16 uint32 uint64 | %d       |
| 浮点数   | float32 float64                                                 | %f %e %g |
| 复数     | complex128 complex64                                            | %v       |
| 布尔型态 | bool                                                            | %t       |
| 指针     | uintprt                                                         | %d       |
| 引用     | map slice channel                                               | %v       |
| 字节     | byte                                                            | %d       |
| 任意符   | rune                                                            | %d       |
| 字符串   | string                                                          | %s       |
| 错误     | error                                                           | %v       |

## 变量初始化

函数内部的变量（非全局变量）可以通过:=声明并初始化

```bash
a:=3
```

## 匿名变量

```bash
_=2+4
```

匿名变量不占用命名空间，不会分配内存，因此可以重复使用。

# 常量

```bash
const (
  PI=3.14
  E=2.71
)
const (
  a=100
  b   # 100 和上一行相同
  c   # 100 和上一行相同
)
```

# iota

`iota`是go语言的常量计数器，只能在常量的表达式中使用。

 使用 `iota`时只需要记住以下两点

1.`iota`在 `const`关键字出现时将被重置为0。

2.`const`中每新增一行常量声明将使 `iota`计数一次(iota可理解为 `const`语句块中的行索引)。

```go
const(
  a=iota # 0
  b      # 1
  c      # 2
  d      # 3
)
const(
  a=iota # 0
  b      # 1
  _      # 2
  d      # 3
)
const(
  a=iota # 0
  b=30  
  c=iota # 2
  d      # 3
)
const(
  _=iota         # iota=0
  KB=1<<(10*iota)# iota=1
  MB=1<<(10*iota)# iota=2
  BG=1<<(10*iota)# iota=3
  TB=1<<(10*iota)# iota=4
)
const(
  a,b=iota+1,iota+2 # 1,2 iota=0
  c,d               # 2,3 iota=1
  e,f               # 3,4 iota=2
)
```

# 字面量

没有出现变量名，直接出现值。基础类型的字面量相当于常量。

# 变量作用域

```go
var (
  A=3   // 全局变量，大写字母开头，所有地方都可以访问，跨package访问需要带上package名称
  b=4   // 全局变量，小写字母开头，本package内部可以访问  
)
func fool(){
  b := 5 // 局部变量，仅本函数内部可以访问，内部声明的变量可以跟外部变量冲突，以内部为准
  {
    b := 6 // 仅{}圈定作用域可以访问，可以跟外部变量有冲突
  }
}
```

全局变量必须通过var/const声明，不能使用var:=

# 注释

## godoc

godoc可以为项目代码导出网页版注释文档

需要先安装go get golang.org/x/tools/cmd/godoc

启动http:godoc -http=:6060

启动浏览器访问http://127.0.0.1:6060/pkg/go-course/entrance_class

# 数据类型

## 基础数据类型

| 类型          | 长度(字节) | 默认值 | 说明                         |
| ------------- | ---------- | ------ | ---------------------------- |
| bool          | 1          | false  |                              |
| byte          | 1          | 0      | uint8，取值范围[0,255]       |
| rune          | 4          | 0      | Unicode Code Point，int32    |
| int，uint     | 4或8       | 0      | 32位或64位，取决于操作系统   |
| int8，uint8   | 1          | 0      | -128～127，0～255            |
| int16，uint16 | 2          | 0      | -32768～32767，0～65535      |
| int32，uint32 | 4          | 0      | -21亿～21亿，0～4亿          |
| int64，uint64 | 8          | 0      |                              |
| float32       | 4          | 0.0    |                              |
| float64       | 8          | 0.0    |                              |
| complex64     | 8          |        |                              |
| complex128    | 16         |        |                              |
| uinptr        | 4或8       |        | 以存储指针uint32或uint64整数 |

## 复合数据类型

| 类型      | 默认值 | 说明        |
| --------- | ------ | ----------- |
| array     |        | 值类型      |
| struct    |        | 值类型      |
| string    | ""     | UTF-8字符串 |
| slice     | nil    | 引用类型    |
| map       | nil    | 引用类型    |
| channel   | nil    | 引用类型    |
| interface | nil    | 接口        |
| function  | nil    | 函数        |

# 指针

```go
func scope() {
	var a int
	var pointer unsafe.Pointer = unsafe.Pointer(&a)
	var p uintptr = uintptr(pointer)
	var ptr *int = &a
	fmt.Printf("pointer %p,p %d %x,ptr %p\n",pointer,p,p,ptr)
}
```

结果

```go
pointer 0xc00001a0f0,p 824633827568 c00001a0f0,ptr 0xc00001a0f0
```

# 异常

```go
	e := errors.New("divide by zero")
	fmt.Printf("%v\n",e)
	fmt.Printf("%+v\n",e)
	fmt.Printf("%#v\n",e)
```

结果

```go
divide by zero
divide by zero
&errors.errorString{s:"divide by zero"}
```
