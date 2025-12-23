package main

import (
	"fmt"

	"github.com/ariboss89/koda-b5-go/internals/profile"
	"github.com/ariboss89/koda-b5-go/internals/shapes"
	"github.com/ariboss89/koda-b5-go/internals/slice"
)

func main() {
	ShowMenu()
}

func ShowMenu() {
	var choice int
	var radius int
	var height int

	for {
		fmt.Println("\nWELCOME USER")
		fmt.Println("Please choose the menu belows to run function as per as your needs :")
		fmt.Println("1. Calculate Circumference And Area of Circle")
		fmt.Println("2. Print Triangle")
		fmt.Println("3. Show Slice")
		fmt.Println("4. Show Profile")
		fmt.Println("0. Exit")

		fmt.Println("Select your choice (1,2,3,4) : ")
		fmt.Scanf("%d", &choice)

		switch true {
		case choice == 1:
			fmt.Println("================================================")
			fmt.Println("You have choosen menu number 1 :")
			fmt.Println("Input circle radius : ")
			fmt.Scanf("%d", &radius)
			fmt.Println(shapes.PrintCircumferenceAndAreaOfCircle(radius))
			fmt.Println("================================================")
		case choice == 2:
			fmt.Println("================================================")
			fmt.Println("You have choosen menu number 2 :")
			fmt.Println("Input height of triangle : ")
			fmt.Scanf("%d", &height)
			(shapes.PrintRightTriangle(height))
			fmt.Println("\n================================================")
		case choice == 3:
			fmt.Println("================================================")
			fmt.Println("You have choosen menu number 3 :")
			(slice.PutNumberBetweenSlice())
			fmt.Println("================================================")
		case choice == 4:
			fmt.Println("================================================")
			fmt.Println("You have choosen menu number 4 :")
			fmt.Println(profile.GetProfile())
			fmt.Println("================================================")
		default:
			fmt.Println("==========================================================================")
			fmt.Printf("You have choosen menu number %d and the menu is not available !!\n", choice)
			fmt.Println("Goodbye !!")
			fmt.Println("==========================================================================")
		}

		if choice > 4 {
			break
		}

	}
}
