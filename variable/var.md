
** golang变量 **

- 变量名在变量类型之前
- 使用var来申明一个变量

`
    var a int
    var s string
`

** 打印变量 **

- fmt来打印变量, Printf 是格式化打印
`
    package main

    import "fmt"

    func main(){
    	var a int
    	a = 10
    	fmt.Printf("Print value:%d\n",a)
    }
`

** 设置变量初值 **

`
    func variableInitialValue(){
    	var a int = 10
    	var s string = "test"
    	fmt.Println(a,s)
    }
`