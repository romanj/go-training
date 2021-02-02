package main

import (
	"fmt"
)

func main() {
	a := 4
	b := a
	c := &a

	fmt.Println(fmt.Sprintf("the value of c is %v", c))

	printAll := func() {
		fmt.Println(fmt.Sprintf("a is %d, b is %d, c internal value is %d", a, b, *c))
	}

	printAll()
	a = 8
	fmt.Println(fmt.Sprintf("the value of c is %v", c))

	printAll()
}
