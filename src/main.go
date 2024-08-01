package main

import "fmt"

func changeValue(x *int) {
	*x = *x + *x
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	arr := []int{}
	arr = append(arr, 8, 6, 9, 2, 4, 10, 11, 23, 0, 2, 1)
	fmt.Println(arr)
	bubble_sort(&arr)
	fmt.Println(arr)
}
