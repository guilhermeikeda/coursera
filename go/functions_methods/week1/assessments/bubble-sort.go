package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	slice := ReadUserInput()

	BubbleSort(slice)

	fmt.Println("Result:")
	print(slice)
}

func ReadUserInput() []int {
	var slice []int
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter an integer or press enter to sort")
		text, err := reader.ReadString('\n')

		if err != nil {
			return slice
		}

		inputValue, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil {
			break
		}

		slice = append(slice, inputValue)
		if len(slice) == 10 {
			break
		}
	}

	return slice
}

func BubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func Swap(slice []int, i int) {
	if isLastIndex(slice, i) {
		return
	}

	temp := slice[i]
	slice[i] = slice[i+1]
	slice[i+1] = temp
}

func isLastIndex(slice []int, i int) bool {
	return len(slice)-1 == i
}

func print(slice []int) {
	for _, v := range slice {
		fmt.Printf("%d ", v)
	}
}
