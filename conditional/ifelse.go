package main

import "fmt"

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
	go_if()
}
