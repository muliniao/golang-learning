package main

import "strings"

func main() {
	f := MakeSuffix(".jpg")
	f("winter")
	f("bird.jpg")
}

func MakeSuffix(suffix string) func(string) string {
	return func(name string) string {
		// 无后缀,则加上
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
