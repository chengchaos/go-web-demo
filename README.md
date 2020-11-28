# go-web-demo
A web project with go languange implement

## Setup

### Unix

1, Install go

2, Set up go module 

### Windows

1, Install go

2, Set up go module

[https://goproxy.cn/](https://goproxy.cn/)

```bat
> go env -w  GO111MODULE=on
> go env -w  GOPROXY=https://goproxy.cn
> go mod init github.com/chengchaos/go-web-demo
go: creating new go.mod: module github.com/chengchaos/go-web-demo

```

## Notes

Go 语言提供了用于创建 Web 服务器的标准库。使得创建一个服务器的步骤非常简单，只要调用 `ListenAndServe` 并传入网络地址以及负责处理请求的处理器（Handler）作为参数就可以了。

如果网络地址参数为空，那么服务器使用默认的 80 端口进行网络连接；

如果处理器的参数为 `nil` 那么服务器使用默认的多路复用器 `DefaultServeMux` 

```go
package main

import "net/http"

// 最简单的服务器
func main() {
    http.ListenAndServe("", nil)
}
```

我们除了可以通过 `ListenAndServe` 的参数对服务器的网络的地址和处理器进行配置之外，还可以通过 `Server` 结构对服务器进行更详细的配置。

```go
package main

import "net/http"

func main() {
    server := http.Server{
        Addr: "127.0.0.1:8080",
        Handler: nil,
    }
    server.ListenAndServe()
}
```

下面的代码展示了 Server 结构所有可选的配置项：

```go
type Server struct {
    Addr string
    Handler Handler
    ReadTimeout time.Duration
    WriteTimeout time.Duration
    MaxhandlerBytes int
    TLSConfig *tls.Config
    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
    ConnState func(net.Conn, ConnState)
    ErrorLog *log.Logger
}
```

### 处理器

在 Go 语言中，一个处理器就是一个拥有 `ServeHTTP` 方法的借口，这个 `ServerHTTP` 方法需要接受两个参数：第一个参数是一个 `ResponseWriter` 借口，第二个参数是一个指向 `Request` 结构的指针。

```go
ServeHTTP(http.ResponseWriter, *http.Request)
```

DefaultServeMax 既是 ServeMux 结构的实例，也是 Handler 结构的实例，因此 DefaultServeMux 不仅是一个多路复用器，它还是一个处理器。

不过 DefaultServeMux 处理器和其他一般的处理器不同，它唯一要做的就是根据请求的 URL 将请求重定向到不同的处理器。

### 处理器函数

处理器函数就是与处理器有相同行为的函数：这些函数与 ServeHTTP 方法有相同的签名。

```go
package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello")
}

func main() {
    server := http.Server {
        Addr : "127.0.0.1:8080",
    }
    http.HandleFunc("/hello", hello)

    server.ListenAndServe()
}
```

Go 语言有用一种 HandlerFunc 函数类型，它可以把一个带有正确签名的函数 f 转换成一个带有方法 f 的 Handler。

处理器函数只是创建处理器的一种便利方法。

http.HandleFunc 函数的源代码：

```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}

```

ServeMux.HandleFunc 的方法定义：

```go
func(mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    mxu.Handle(pattern, HandlerFunc(handler))
}
```

### 串联

```go
package main

import (
    "fmt"
    "net/http"
    "reflect"
    "runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
        fmt.Println("Handler function called - " + name)
        h(w, r)
    }
}

func main() {
    server := http.Server{
        Addr : "127.0.0.1:8080",
    }

    http.HandleFunc("/hello", log(hello))
    server.ListenAndServe()
}
```

### httprouter

[https://github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)


```bash
go get github.com/julienschmidt/httprouter

```


