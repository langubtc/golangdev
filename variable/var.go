package main

import "fmt"

//在函数外不能使用:=，只能使用var开头定义变量
var (
	username = "xiaobai"
	age      = 5
	password = "test"
)

//设置变量初始值
func variableInitialValue() {
	var a int = 10
	var s string = "test"
	fmt.Println(a, s)
}

//设置变量
func variableSet() {
	var a int
	a = 10
	fmt.Printf("Print value:%d\n", a)
}

//设置变量多值
func variableMultiVule() {
	var a, b, c int
	a = 1
	b = 2
	c = 3

	var q, w, e string = "test", "v", "d"

	fmt.Println(q, w, e)
	fmt.Println(a, b, c)

}

//变量定义简单写法 使用:= 定义
func variableSorter() {
	name := "xiaobai"
	fmt.Println(name)
}

func main() {
	variableMultiVule()
	variableSet()
	variableInitialValue()
	variableSorter()

	fmt.Println(username, password, age)
}
