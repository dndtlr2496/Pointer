# Pointer
 Pointer Using Example with go

 ## 함수에서 인자를 변수의 함숫값으로 가져오면 함수내에서 전역변수를 바꿀 수 있다.
 '''
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
'''

'''
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
'''