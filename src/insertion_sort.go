package main

func insertion_sort(arr *[]int) {
	if len(*arr) > 0 {
		for i := 0; i < len(*arr)-1; i++ {
			for j := i + 1; j > 0; j-- {
				if (*arr)[j] < (*arr)[j-1] {
					swap(&(*arr)[j], &(*arr)[j-1])
				}
			}
		}
	}
}
