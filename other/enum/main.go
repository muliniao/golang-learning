package main

import "fmt"

type Season int32

const (
	Spring Season = iota + 1
	Summer
	Autumn
	Winter
)

var (
	SeasonName = map[Season]string{
		Spring: "Spring",
		Summer: "Summer",
		Autumn: "Autumn",
		Winter: "Winter",
	}

	SeasonValue = map[string]Season{
		"Spring": Spring,
		"Summer": Summer,
		"Autumn": Autumn,
		"Winter": Winter,
	}
)

// 返回实例名称
func (s Season) name() string {
	return SeasonName[s]
}

func (s Season) ordinal() int {
	return 0
}

func (s Season) compareTo() int {
	return 0
}

func (s Season) valueOf() {

	return
}

func main() {

	a := Winter
	fmt.Println(a)
	fmt.Println(a.name())

}
