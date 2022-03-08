package main

import "fmt"

func main() {
	slice := []int{}

	fmt.Print("Enter up to ten integers")
	for i := 0; i < 10; i++ {
		var input int
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}

		slice = append(slice, input)
	}

	BubbleSort(slice)
	print(slice)
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
