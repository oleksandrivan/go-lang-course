package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	name, eat, speak, move string
}

func (animal Animal) Eat() {
	fmt.Printf("Animal %s eats %s", animal.name, animal.eat)
}

func (animal Animal) Move() {
	fmt.Printf("Animal %s moves %s", animal.name, animal.move)
}

func (animal Animal) Speak() {
	fmt.Printf("Animal %s speaks %s", animal.name, animal.speak)
}

func (animal *Animal) performAction(action string) {
	switch {
	case strings.EqualFold(action, "eat"):
		animal.Eat()
	case strings.EqualFold(action, "move"):
		animal.Move()
	case strings.EqualFold(action, "speak"):
		animal.Speak()
	}
}

var (
	cow   = Animal{name: "cow", eat: "grass", move: "walk", speak: "moo"}
	bird  = Animal{name: "bird", eat: "worms", move: "fly", speak: "peep"}
	snake = Animal{name: "snake", eat: "mice", move: "slither", speak: "hiss"}
)

func main() {

	var animal, action string

	for {
		fmt.Printf("\n> ")
		_, err := fmt.Scanln(&animal, &action)
		if err == nil {
			inputAnimal := getAnimal(animal)
			inputAnimal.performAction(action)
		} else {
			break
		}
	}
}

func getAnimal(s string) *Animal {
	switch {
	case strings.EqualFold(s, "cow"):
		return &cow
	case strings.EqualFold(s, "bird"):
		return &bird
	case strings.EqualFold(s, "snake"):
		return &snake
	default:
		return nil
	}
}
