# WaitGroup




## 易错场景
1. 计数器设置为负值



2. 等所有Add方法调用后再调用Wait



3. 错误重用WaitGroup



## 总结
1. 统一Add，再并发Done，最后Wait这种标准方式，来使用WaitGroup值
2. 不重用 WaitGroup。新建一个 WaitGroup 不会带来多大的资源开销，重用反而更容易出错
3. 保证所有的 Add 方法调用都在 Wait 之前
4. 不传递负数给 Add 方法，只通过 Done 来给计数值减 1
5. 不做多余的 Done 方法调用，保证 Add 的计数值和 Done 方法调用的数量是一样的
6. 不遗漏 Done 方法的调用，否则会导致 Wait hang 住无法返回
7. WaitGroup必须传地址,不能是引用类型
   - func f (i int, wg *sync.WaitGroup){}

## 遗留问题
1. 错误信息：如何处理不同协程的error(优先级高): 参考例子waitgroup_003


## 源码解析