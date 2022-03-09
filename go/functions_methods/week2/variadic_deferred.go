package main

import "fmt"

func main() {
	deferring()

	fmt.Println(maxValue(1, 2, 3))

	slice := []int{5, 23, 1}
	fmt.Println(maxValue(slice...))

}

func deferring() {
	defer fmt.Println("Bye") // Argumentos avaliados na declaração
	i := 1
	defer fmt.Println(i + 1)
	i++

	fmt.Printf("Hello %d \n", i)
}

func maxValue(values ...int) int {
	maxValue := -1
	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}
