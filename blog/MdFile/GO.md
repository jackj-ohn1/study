### 改变GO参数来提高下载包的速度
`go env -w GOPROXY=https://goproxy.io,direct`

`go env -w GO111MODULE=on`

- GO 里面的一些细节
~~~go
func Hello(hello ...string){
  //fmt.Printf(hello)
}
// 传入可以是Hello("hello","ok"),里面的形参hello是一个切片，但传入实参的方法更贴近
~~~

### GO命令

#### 1.go build 命令

- 我们可以在 `go build` 的后面提供多个文件名，`go build` 会编译这些源码，输出可执行文件(.exe)

~~~go
go build file1.go file2.go

// 执行这个可执行文件
./file.exe
~~~

- 额外参数

| 附加参数 | 备  注                                      |
| -------- | ------------------------------------------- |
| -v       | 编译时显示包名                              |
| -p n     | 开启并发编译，默认情况下该值为 CPU 逻辑核数 |
| -a       | 强制重新构建                                |
| -n       | 打印编译时会用到的所有命令，但不真正执行    |
| -x       | 打印编译时会用到的所有命令                  |
| -race    | 开启竞态检测                                |

#### 2.go clean命令

- 清除编译文件

> 包含的文件
>
> ```go
> _obj/            旧的object目录，由Makefiles遗留
> _test/           旧的test目录，由Makefiles遗留
> _testmain.go     旧的gotest文件，由Makefiles遗留
> test.out         旧的test记录，由Makefiles遗留
> build.out        旧的test记录，由Makefiles遗留
> *.[568ao]        object文件，由Makefiles遗留
> 
> DIR(.exe)        由go build产生
> DIR.test(.exe)   由go test -c产生
> MAINFILE(.exe)   由go build MAINFILE.go产生
> *.so             由 SWIG 产生
> ```

![](http://c.biancheng.net/uploads/allimg/200116/4-200116102KJ01.gif)

- 额外参数
  - -i 清除关联的安装的包和可运行文件，也就是通过`go install`安装的文件；
  - -n 把需要执行的清除命令打印出来，但是不执行，这样就可以很容易的知道底层是如何运行的；
  - -r 循环的清除在 import 中引入的包；
  - -x 打印出来执行的详细命令，其实就是 -n 打印的执行版本；
  - -cache 删除所有`go build`命令的缓存
  - -testcache 删除当前包所有的测试结果

#### 3.go run 命令

- `go run`命令会编译源码，并且直接执行源码的 main() 函数，**不会在当前目录留下可执行文件**。相当于先`go build`再运行`可执行文件`

#### 4.gofmt命令

- `gofmt `是一个 cli 程序，会优先读取标准输入，如果传入了文件路径的话，会格式化这个文件，如果传入一个目录，会格式化目录中所有 .go 文件，如果不传参数，会格式化当前目录下的所有 .go 文件

~~~go
[]T{T{}, T{}}  s[a:len(s)]  for x, _ = range v {...} 
命令:gofmt -s ./main.go
// 格式 gofmt 参数 文件名
// 简化后
[]T{{}, {}}    s[a:]        for x = range v {...}
~~~

| 标记名称    | 标记描述                                                     |
| ----------- | ------------------------------------------------------------ |
| -l          | 仅把那些不符合格式化规范的、需要被命令程序改写的源码文件的绝对路径打印到标准输出。而不是把改写后的全部内容都打印到标准输出。 |
| -w          | 把改写后的内容直接写入到文件中，而不是作为结果打印到标准输出。 |
| -r          | 添加形如“a[b:len(a)] -> a[b:]”的重写规则。如果我们需要自定义某些额外的格式化规则，就需要用到它。 |
| -s          | 简化文件中的代码。                                           |
| -d          | 只把改写前后内容的对比信息作为结果打印到标准输出。而不是把改写后的全部内容都打印到标准输出。 命令程序将使用 diff 命令对内容进行比对。在 Windows 操作系统下可能没有 diff 命令，需要另行安装。 |
| -e          | 打印所有的语法错误到标准输出。如果不使用此标记，则只会打印每行的第 1 个错误且只打印前 10 个错误。 |
| -comments   | 是否保留源码文件中的注释。在默认情况下，此标记会被隐式的使用，并且值为 true。 |
| -tabwidth   | 此标记用于设置代码中缩进所使用的空格数量，默认值为 8。要使此标记生效，需要使用“-tabs”标记并把值设置为 false。 |
| -tabs       | 是否使用 tab（'\t'）来代替空格表示缩进。在默认情况下，此标记会被隐式的使用，并且值为 true。 |
| -cpuprofile | 是否开启 CPU 使用情况记录，并将记录内容保存在此标记值所指的文件中。 |



#### 5.go mod命令

- 什么是go.mod?
  - Go.mod是Golang1.11版本新引入的官方包管理工具用于解决之前没有地方记录依赖包具体版本的问题，方便依赖包的管理。
  - 官方定义:`Modules是相关Go包的集合，是源代码交换和版本控制的单元。go命令直接支持使用Modules，包括记录和解析对其他模块的依赖性。Modules替换旧的基于GOPATH的方法，来指定使用哪些源文件。`

~~~go
// 生成一个go.mod文件，之后的包的管理都是通过这个文件管理。
//子目录里是不需要init的，所有的子目录里的依赖都会组织在根目录的go.mod文件里
go mod init taskname

// 增加缺少的包，删除无用的包
go mod tidy

~~~



- 常用的`go mod`命令如下表所示：

| 命令            | 作用                                           |
| --------------- | ---------------------------------------------- |
| go mod download | 下载依赖包到本地（默认为 GOPATH/pkg/mod 目录） |
| go mod edit     | 编辑 go.mod 文件                               |
| go mod graph    | 打印模块依赖图                                 |
| go mod init     | 初始化当前文件夹，创建 go.mod 文件             |
| go mod tidy     | 增加缺少的包，删除无用的包                     |
| go mod vendor   | 将依赖复制到 vendor 目录下                     |
| go mod verify   | 校验依赖                                       |
| go mod why      | 解释为什么需要依赖                             |

#### 6.go get 命令

- go get 命令可以借助代码管理工具通过远程拉取或更新代码包及其依赖包，并自动完成编译和安装。整个过程就像安装一个 App 一样简单。这个命令可以动态获取远程代码包

> 这个命令在内部实际上分成了两步操作：第一步是下载源码包，第二步是执行 go install。下载源码包的 go 工具会自动根据不同的域名调用不同的源码工具
>
> - BitBucket (Mercurial Git)
> - GitHub (Git)
> - Google Code Project Hosting (Git, Mercurial, Subversion

`go get 包名  是直接把远程包下载到GOPATH里面，get成功后就可以直接使用这个包`

| 附加参数  | 备  注                                 |
| --------- | -------------------------------------- |
| -v        | 显示操作流程的日志及信息，方便检查错误 |
| -u        | 下载丢失的包，但不会更新已经存在的包   |
| -d        | 只下载，不安装                         |
| -insecure | 允许使用不安全的 HTTP 方式进行下载操作 |

#### 7.go install 命令

- go install 只是将编译的中间文件放在 GOPATH 的 pkg 目录下，以及固定地将编译结果放在 GOPATH 的 bin 目录下。这个命令在内部实际上分成了两步操作：第一步是生成结果文件（可执行文件或者 .a 包），第二步会把编译好的结果移到 $GOPATH/pkg 或者 $GOPATH/bin。

`使用 go install 来执行代码，参考下面的 shell：`

~~~go
export GOPATH=/home/davy/golangbook/code // 指定输出的GOPATH路径
go install chapter11/goinstall
~~~

#### 8.其他命令

- go version 查看go当前的版本
- go env 查看当前go的环境变量
- go list 列出当前全部安装的package