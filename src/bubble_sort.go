package main

func bubble_sort(arr *[]int) {
	if len(*arr) > 0 {
		for i := 0; i < len(*arr)-1; i++ {
			for j := i + 1; j < len(*arr); j++ {
				if (*arr)[i] > (*arr)[j] {
					swap(&(*arr)[i], &(*arr)[j])
				}
			}
		}
	}
}
