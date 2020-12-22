package main

import (
	"fmt"
)

func main() {
	var a int
	var b string

	a = 10
	b = "some string"

	message := ""

	message = fmt.Sprintf("a = %d and b = %s", a, b)

	fmt.Println(message)
}
