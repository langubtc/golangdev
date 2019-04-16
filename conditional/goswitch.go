package main

import "fmt"

func switchSelect() {
	var d_type = "mem"
	switch {
	case d_type == "disk":
		fmt.Println("Select disk")
	case d_type == "mem":
		fmt.Println("Select mem")
	case d_type == "cpu":
		fmt.Println("Select Cpu")
	default:
		fmt.Println("Success!")
	}
}

//函数传参数a,b int类型，op为操作 string类型，最后int是函数返回值类型
//在switch 中也可以直接带上变量进行判断，就不需要在case中在重复写条件判断
func switchCalculator(a, b int, op string) int {
	var result int

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)

	}
	return result
}

//switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。
func main() {
	switchSelect()

	//计算器
	returnResult := switchCalculator(2, 3, "*")
	fmt.Println(returnResult)

}
