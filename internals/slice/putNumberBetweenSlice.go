package slice

import "fmt"

func PutNumberBetweenSlice() {
	var numbers = []int{50, 75, 66, 20, 32, 90}

	numbers = append(
		numbers[:3],
		append([]int{88}, numbers[3:]...)...,
	)

	for i, v := range numbers {
		fmt.Printf("Index %d: %d\n", i, v)
	}

}
