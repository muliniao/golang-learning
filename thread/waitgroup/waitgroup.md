# WaitGroup(优先级：高)

## 实现机制
1. Add: 用来设置WaitGroup的计数值
2. Done: 用来将 WaitGroup 的计数值减 1，其实就是调用了 Add(-1)
3. Wait: 调用这个方法的 goroutine 会一直阻塞，直到 WaitGroup 的计数值变为 0

## 易错场景
1. 计数器设置为负值

- 调用Add时传入负数
````go
func main() {
    var wg sync.WaitGroup
    wg.Add(10)

    wg.Add(-10)//将-10作为参数调用Add，计数值被设置为0

    wg.Add(-1)//将-1作为参数调用Add，如果加上-1计数值就会变为负数。这是不对的，所以会触发panic
}
````

- 调用Done方法次数,超过WaitGroup的计数值
````go
func main() {
    var wg sync.WaitGroup
    wg.Add(1)

    wg.Done()

    wg.Done()
}
````

2. 等所有Add方法调用后再调用Wait
````go
func main() {
    var wg sync.WaitGroup
    go dosomething(100, &wg) // 启动第一个goroutine
    go dosomething(110, &wg) // 启动第二个goroutine
    go dosomething(120, &wg) // 启动第三个goroutine
    go dosomething(130, &wg) // 启动第四个goroutine

    wg.Wait() // 主goroutine等待完成
    fmt.Println("Done")
}

func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
    duration := millisecs * time.Millisecond
    time.Sleep(duration) // 故意sleep一段时间

    wg.Add(1)
    fmt.Println("后台执行, duration:", duration)
    wg.Done()
}
````

3. 错误重用WaitGroup
````go
func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        time.Sleep(time.Millisecond)
        wg.Done() // 计数器减1
        wg.Add(1) // 计数值加1
    }()
    wg.Wait() // 主goroutine等待，有可能和第7行并发执行
}
````

## 错误案例



## 总结
1. 统一Add，再并发Done，最后Wait这种标准方式，来使用WaitGroup值
2. 不重用 WaitGroup。新建一个 WaitGroup 不会带来多大的资源开销，重用反而更容易出错
3. 保证所有的 Add 方法调用都在 Wait 之前
4. 不传递负数给 Add 方法，只通过 Done 来给计数值减 1
5. 不做多余的 Done 方法调用，保证 Add 的计数值和 Done 方法调用的数量是一样的
6. 不遗漏 Done 方法的调用，否则会导致 Wait hang 住无法返回
7. WaitGroup必须传地址,不能是引用类型
   - func f (i int, wg *sync.WaitGroup){}
8. 使用vet进行辅助检查

## 遗留问题
1. 错误信息：如何处理不同协程的error(优先级高): 参考例子waitgroup_003
2. 在使用WaitGroup值实现一对多的 goroutine 协作流程时，怎样才能让分发子任务的 goroutine 获得各个子任务的具体执行结果？
3. 通常我们可以把 WaitGroup 的计数值，理解为等待要完成的 waiter 的数量。你可以试着扩展下 WaitGroup，来查询 WaitGroup 的当前的计数值吗？

## 源码解析



