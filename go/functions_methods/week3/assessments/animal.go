package main

import (
	"bufio"
	"errors"
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

func main() {
	animals := map[string]Animal{"cow": buildCow(), "bird": buildBird(), "snake": buildSnake()}

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print(">")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		var animalName string
		var action string

		input := strings.Split(text, " ")
		if len(input) == 2 {
			animalName = input[0]
			action = input[1]
		} else {
			fmt.Println("Invalid input. Try again")
			continue
		}

		animal, err := getByName(animals, animalName)
		if err != nil {
			fmt.Println(err)
			continue
		}

		animal.Execute(action)
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

func getByName(animals map[string]Animal, animalName string) (Animal, error) {
	for name, animal := range animals {
		if name == animalName {
			return animal, nil
		}
	}

	return Animal{}, errors.New(fmt.Sprintf("Animal %s não encontrado", animalName))
}
