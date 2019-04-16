package main

import "fmt"

//单if判断
func ifconditional() {
	var a int = 10
	if a < 100 {
		fmt.Printf("a 小于 100\n")
	}
	fmt.Printf("a = %d\n", a)
}

//if else 判断
func ifelseconditional() {
	var a int = 10
	if a < 5 {
		fmt.Printf("a 小于 5\n")
	} else {
		fmt.Printf("a 大于 5 \n")
	}
}

// if 多条件判断
func go_if() {
	a := 6
	if a < 10 && a > 5 {
		fmt.Println("5 > Value < 10 ")
	} else if a < 5 {
		fmt.Println("Value < 5")
	} else {
		fmt.Println("value Error")
	}

}

func main() {
	ifconditional()
	ifelseconditional()
	go_if()
}
