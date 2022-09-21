# Once(优先级：高)

## 实现机制
- 双检查机制
- 原子Load,判断done是否为0
- 加互斥锁，再次判断done是否为0
````go
type Once struct {
    done uint32
    m    Mutex
}

func (o *Once) Do(f func()) {
    if atomic.LoadUint32(&o.done) == 0 {
        o.doSlow(f)
    }
}

func (o *Once) doSlow(f func()) {
    o.m.Lock()
    defer o.m.Unlock()
    // 双检查
    if o.done == 0 {
        defer atomic.StoreUint32(&o.done, 1)
        f()
    }
}
````

## 易错场景
1. 死锁
````go
func main() {
    var once sync.Once
    once.Do(func() {
        once.Do(func() {
            fmt.Println("初始化")
        })
    })
}
````

2. 未初始化
````go
func main() {
    var once sync.Once
    var googleConn net.Conn // 到Google网站的一个连接

    once.Do(func() {
        // 建立到google.com的连接，有可能因为网络的原因，googleConn并没有建立成功，此时它的值为nil
        googleConn, _ = net.Dial("tcp", "google.com:80")
    })
    // 发送http请求
    googleConn.Write([]byte("GET / HTTP/1.1\r\nHost: google.com\r\n Accept: */*\r\n\r\n"))
    io.Copy(os.Stdout, googleConn)
}
````

## 错误案例




## 源码分析





## 总结


##