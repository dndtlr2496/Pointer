package main

import "fmt"

func main() {
	a := 1

	A(&a)
	fmt.Println(a)
}

func A(c *int) {
	fmt.Println(*c)
	*c++
}
