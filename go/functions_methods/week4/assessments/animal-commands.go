package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	name string
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	name string
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	var animals []Animal

	for {
		fmt.Print("> ")

		var userInput UserInput
		userInput.Read()

		switch userInput.command {
		case "newanimal":
			createAnimal(&animals, userInput)
		case "query":
			queryAnimal(animals, userInput)
		default:
			fmt.Println("Invalid command")
		}
	}
}

func createAnimal(animals *[]Animal, userInput UserInput) {
	var newAnimal Animal
	switch userInput.request {
	case "cow":
		newAnimal = Cow{name: userInput.animalName}
	case "bird":
		newAnimal = Bird{name: userInput.animalName}
	case "snake":
		newAnimal = Snake{name: userInput.animalName}
	default:
		newAnimal = nil
	}

	if newAnimal == nil {
		fmt.Println("Requested animal does not exists!")
	} else {
		fmt.Println("Created it!")
		*animals = append(*animals, newAnimal)
	}
}

func queryAnimal(animals []Animal, userInput UserInput) {
	var animalFound Animal
	for _, animal := range animals {
		switch animal.(type) {
		case Cow:
			if animal.(Cow).name == userInput.animalName {
				animalFound = animal.(Cow)
			}
		case Bird:
			if animal.(Bird).name == userInput.animalName {
				animalFound = animal.(Bird)
			}
		case Snake:
			if animal.(Snake).name == userInput.animalName {
				animalFound = animal.(Snake)
			}
		}
	}

	if animalFound == nil {
		fmt.Printf("Animal with name %s not found \n", userInput.animalName)
		return
	}

	requestInformation := map[string]func(){"eat": animalFound.Eat, "speak": animalFound.Speak, "move": animalFound.Move}
	action, found := requestInformation[userInput.request]
	if !found {
		fmt.Println(fmt.Sprintf("%s does not %s", userInput.animalName, userInput.request))
		return
	}
	action()
	// switch userInput.request {
	// case "eat":
	// 	animalFound.Eat()
	// case "speak":
	// 	animalFound.Speak()
	// case "move":
	// 	animalFound.Move()
	// }
}

type UserInput struct {
	command    string
	animalName string
	request    string
}

func (userInput *UserInput) Read() {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	inputValue := strings.Split(text, " ")
	if len(inputValue) == 3 {
		userInput.command = strings.TrimSpace(strings.ToLower(inputValue[0]))
		userInput.animalName = strings.TrimSpace(inputValue[1])
		userInput.request = strings.TrimSpace(inputValue[2])
	}
}
