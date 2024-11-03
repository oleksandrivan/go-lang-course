package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow

type Cow struct {
	name string
}

func (cow Cow) Eat(){
	fmt.Printf("The food of '%s' is 'grass'", cow.name)
}

func (cow Cow) Move(){
	fmt.Printf("The locomotion of '%s' is 'walk'", cow.name)
}
func (cow Cow) Speak(){
	fmt.Printf("The sound of '%s' is 'moo'", cow.name)
}

// Bird

type Bird struct {
	name string
}

func (bird Bird) Eat(){
	fmt.Printf("The food of '%s' is 'grass'", bird.name)
}

func (bird Bird) Move(){
	fmt.Printf("The locomotion of '%s' is 'walk'", bird.name)
}
func (bird Bird) Speak(){
	fmt.Printf("The sound of '%s' is 'moo'", bird.name)
}

// Snake

type Snake struct {
	name string
}

func (snake Snake) Eat(){
	fmt.Printf("The food of '%s' is 'grass'", snake.name)
}

func (snake Snake) Move(){
	fmt.Printf("The locomotion of '%s' is 'walk'", snake.name)
}
func (snake Snake) Speak(){
	fmt.Printf("The sound of '%s' is 'moo'", snake.name)
}



func performAction(animal Animal, action string) {
	switch {
	case strings.EqualFold(action, "eat"):
		animal.Eat()
	case strings.EqualFold(action, "move"):
		animal.Move()
	case strings.EqualFold(action, "speak"):
		animal.Speak()
	}
}

func main() {

	animals := make(map[string]Animal)

	for {
		fmt.Printf("\n> ")
		reader := bufio.NewReader(os.Stdin)
		inputString, err := reader.ReadString('\n')
		fields := strings.Fields(inputString)
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.EqualFold("newanimal", fields[0]) {
			animalName := fields[1]
			animalType := fields[2]
			newAnimal := createAnimal(animalName, animalType)
			fmt.Printf("Created it!")
			animals[animalName] = newAnimal
		} else if strings.EqualFold("query", fields[0]) {
			animal, exists := animals[fields[1]]
			if exists {
				performAction(animal, fields[2])
			} else {
				fmt.Printf("Animal '%s' not found!", fields[1])
			}
		} else {
			fmt.Printf("Command '%v' is not valid", fields)
		}
	}
}

func createAnimal(animalName, animalType string) Animal {
	if strings.EqualFold("cow", animalType) {
		return Cow{name: animalName}
	} else if strings.EqualFold("bird", animalType) {
		return Bird{name: animalName}
	} else if strings.EqualFold("snake", animalType) {
		return Snake{name: animalName}
	}
	return nil
}
