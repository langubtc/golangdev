package main

import "fmt"

func loopFunc() {

	for a := 0; a < 20; a++ {
		fmt.Printf("当前循环值为%d\n", a)
	}

}

func loopVariable() {
	var a int = 10
	var b int = 20
	for a < b {
		a++
		fmt.Printf("a 的值为%d\n", a)
	}

}

func loopList() {
	numberList := [6]int{1, 2, 3, 4, 5, 6}
	for i, x := range numberList {
		fmt.Printf("当前值为%d,索引为%d\n", x, i)
	}

}

func main() {
	loopList()
	loopVariable()
	loopFunc()

}
