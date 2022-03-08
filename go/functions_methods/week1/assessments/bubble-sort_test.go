package main

import "testing"

func TestSwap_ShouldSwapSequencialElementsFromSlice(test *testing.T) {
	slice := []int{1, 2, 3}
	Swap(slice, 0)

	if slice[0] != 2 || slice[1] != 1 {
		test.Error("Element weren't swapped")
	}
}

func TestSwap_ShouldNotSwap_When_IndexIsFromLastElement(test *testing.T) {
	slice := []int{1, 2, 3}
	Swap(slice, 2)

	if slice[2] != 3 {
		test.Error("Element swapped")
	}
}

func TestBubbleSort_ShouldSortSlice(test *testing.T) {
	slice := []int{3, 2, 1, 9, 20, 13}

	BubbleSort(slice)

	if slice[0] != 1 || slice[1] != 2 || slice[2] != 3 {
		test.Error("Slice not sorted")
	}
}
