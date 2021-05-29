package main

import "fmt"

type A struct {
	num int
}

func (t *A) Hello() {
	fmt.Println(t.num)
	t.num = 7
}

func main() {
	a := A{5}
	a.Hello()
	fmt.Println(a.num)
}
