package main

import (
	"fmt"
	"mymodule/mypackage"
	"mymodule/utils"
)

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func add(a, b int) int {
	c := a + b
	return c
}

func main() {
	x := add(20, 10)
	if x > 25 {
		mypackage.HelloFromMyPackage()
	}
	x = factorial(2)

	if 10 == x {
		println("x=10")
	}

	if x > 50 {
		//do something
	}

	s := "ala ma kota"
	if len(s) > 20 {
		mypackage.HelloFromMyPackage2()
	}

	y := add(40, 10)
	if x >= y {
		fmt.Println("X>=Y")
	}

	utils.HelloFromUtils()
	//utils.hello()
}
