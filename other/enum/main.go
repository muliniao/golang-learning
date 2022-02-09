package main

import "fmt"

type Season int

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)

// 返回实例名称
func (s Season) name() string {
	// 方案1
	//switch s {
	//case Spring:
	//	return "spring"
	//case Summer:
	//	return "summer"
	//case Autumn:
	//	return "autumn"
	//case Winter:
	//	return "winter"
	//default:
	//	return "unknown"
	//}

	// 方案2
	return [...]string{
		"spring",
		"summer",
		"autumn",
		"winter",
	}[s]
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
	fmt.Println(a.name())

}
