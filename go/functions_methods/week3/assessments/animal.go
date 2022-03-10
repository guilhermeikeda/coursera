package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() string {
	return animal.food
}

func (animal Animal) Speak() string {
	return animal.noise
}

func (animal Animal) Move() string {
	return animal.locomotion
}

func (animal Animal) Execute(action string) {
	action = strings.TrimSpace(strings.ToLower(action))

	switch action {
	case "eat":
		fmt.Println(animal.Eat())
	case "speak":
		fmt.Println(animal.Speak())
	case "move":
		fmt.Println(animal.Move())
	default:
		fmt.Printf("This animal doesn't %s\n", action)
	}
}

type UserInput struct {
	animalName string
	action     string
}

func (userInput *UserInput) Read() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	input := strings.Split(text, " ")
	if len(input) == 2 {
		userInput.animalName = input[0]
		userInput.action = input[1]
	}
}

func main() {
	animals := map[string]Animal{"cow": buildCow(), "bird": buildBird(), "snake": buildSnake()}

	for {
		fmt.Print(">")

		var userInput UserInput
		userInput.Read()

		animal, found := animals[userInput.animalName]
		if !found {
			fmt.Printf("Animal %s not found\n", userInput.animalName)
			continue
		}

		animal.Execute(userInput.action)
	}
}

func buildCow() Animal {
	return Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
}

func buildBird() Animal {
	return Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
}

func buildSnake() Animal {
	return Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hss",
	}
}
