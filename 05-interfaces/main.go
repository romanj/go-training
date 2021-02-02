package main

import (
	"fmt"
)

func main() {
	catto := &cat{
		Name: "Cat",
	}

	doggo := &dog{
		Name: "Dog",
	}

	animals := []speaker{catto, doggo}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

}

type speaker interface {
	Speak() string
}

type cat struct {
	Name string
}

type dog struct {
	Name string
}

func (c *cat) Speak() string {
	return fmt.Sprintf("cat %s says meoooow", c.Name)
}

func (c *dog) Speak() string {
	return fmt.Sprintf("dog %s says woof! woof!", c.Name)
}
