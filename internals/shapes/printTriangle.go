package shapes

import "fmt"

func PrintRightTriangle(height int) {
	// for
	for x := 0; x <= height; x++ {
		fmt.Println("")
		for y := 0; y <= x; y++ {
			fmt.Printf("*")
		}
	}

	//while
	x := 0
	for x <= height {
		fmt.Println("")
		y := 0
		for y <= x {
			fmt.Printf("*")
			y++
		}
		x++
	}

	// for
	for x := range height + 1 {
		fmt.Println("")
		for y := range x {
			if y <= x {
				fmt.Printf("*")
			}
		}
	}
}
