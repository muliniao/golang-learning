package main

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

func main() {

	tests := []struct {
		value interface{}
	}{
		{"1ABC_99/9"},
	}

	// 大写字母，数字和下划线
	//r := validation.Match(regexp.MustCompile("^[A-Z_][A-Z0-9_]*$")).Error("大写字母，数字和下划线, 不以下划线开头")

	//r := validation.Match(regexp.MustCompile("(?!A)")).Error("不以数字开头")

	// 以什么字符开头
	//r := validation.Match(regexp.MustCompile("^/.*$")).Error("不以数字开头")

	//r := validation.Match(regexp.MustCompile("(^[a-z0-9]+$)|(^[a-z0-9][-a-z0-9]+$)")).Error("only support lower case letters, digits and hyphen, cannot begin and end with hyphen")

	//r := validation.Match(regexp.MustCompile("(^[A-Z0-9a-z_]+$)|(^[A-Z0-9a-z_][-A-Z0-9a-z_]+$)|(^[A-Z0-9a-z_][.A-Z0-9a-z_]+$)")).Error("大写字母，数字和下划线, 不以下划线开头")
	r := validation.Match(regexp.MustCompile("^[A-Z0-9a-z_][A-Z0-9a-z_/.-]*$")).Error("大小写字母,数字,下划线,句号和连接号 不以下划线开头")

	err := r.Validate(tests[0].value)
	fmt.Println(err)
}
