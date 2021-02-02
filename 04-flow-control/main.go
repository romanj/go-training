package main

import (
	"fmt"
)

func main() {
	line := []string{"Ana", "Billy", "Joe"}

	if len(line)%2 != 0 {
		fmt.Println("it's an odd line")
	} else {
		fmt.Println("it's an even line")
	}

	switch len(line) {
	case 1:
		fmt.Println("one persone in line")
		break

	case 2:
		fmt.Println("two persons in line")

	case 3:
		fmt.Println("three persons in line")

	default:
		fmt.Println("none of cases ware a match")
	}

	for i := 0; i < len(line); i++ {
		fmt.Println(fmt.Sprintf("current person in line by using a regular for loop is %s", line[i]))
	}

	for i, value := range line {
		fmt.Println(fmt.Sprintf("range current index is %d and current value is %s", i, value))
	}

	i := 0
	for i < 5 {
		fmt.Println(fmt.Sprintf("current index using for loop with only the condition is %d", i))
		i++
	}

	i = 0
	for {
		fmt.Println(fmt.Sprintf("current index usign for loop without any condition is %d", i))
		i++
		if i == 5 {
			break
		}
	}
}
