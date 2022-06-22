package main

import (
	"fmt"
	"math/big"
	"sync"
)

// threeOnce: 很精妙，源码
// 1. var struct: 变量结构体
// 2. 将需要单例初始化的结构体"Float"和"sync.Once"放在一起
var threeOnce struct {
	sync.Once
	v *big.Float
}

func three() *big.Float {
	threeOnce.Do(func() {
		threeOnce.v = big.NewFloat(3.0)
	})
	return threeOnce.v
}

func main() {

	three()
	fmt.Println(threeOnce.v)

}
