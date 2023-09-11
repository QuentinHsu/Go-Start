package main

import "fmt"

func fool(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

func main() {
	c := fool("abc", 3456)
	fmt.Println("c = ", c)
}
