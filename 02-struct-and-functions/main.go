package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	a := 3
	b := 9

	result := Sum(a, b)

	fmt.Println(fmt.Sprintf("the result of summing %d and %d is %d", a, b, result))

	jarekRomaniuk := &person{
		Name: "John Doe",
		Age:  31,
	}
	fmt.Println(jarekRomaniuk.SayHello())
	jarekRomaniuk.GetOld()
	fmt.Println(jarekRomaniuk.SayHello())

}

func Sum(a int, b int) int {
	return a + b
}

func (p *person) SayHello() string {
	msg := fmt.Sprintf("Hi! My name is %s and I'm %d years old.", p.Name, p.Age)
	return msg
}

func (p *person) GetOld() {
	p.Age++
}
