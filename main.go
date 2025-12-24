package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ariboss89/koda-b5-go/internals/payments"
	"github.com/ariboss89/koda-b5-go/internals/people"
	"github.com/ariboss89/koda-b5-go/internals/profile"
	"github.com/ariboss89/koda-b5-go/internals/shapes"
	"github.com/ariboss89/koda-b5-go/internals/slice"
)

type Address struct {
	city, province, country string
}

func main() {
	ShowMenu()

	//testing pointer
	// var address1 = Address{"batam", "kepri", "indonesia"}

	// var address2 = &address1

	// address2.city = "Jepang"

	// fmt.Print(address1.city)
	// fmt.Print(address2.city)

	// file, err := os.Open("/src/welcome.txt")

	// defer fmt.Printf("Menutup file")
	// defer file.Close()

	// if file != nil {
	// 	read, _ := (io.ReadAll(file))
	// 	if read != nil {
	// 		println(string(read))
	// 	} else {
	// 		println("Terjadi error")

	// 	}
	// } else {
	// 	panicable(err)
	// }

	//fmt.Printf("")

	//people.GreetUs("ari", "12", "111")

}

// func panicable(error) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("print continue ...")

// 		}
// 	}()
// 	panic("Terjadi kepanikan")
// }

func ShowMenu() {
	var choice int
	var radius int
	var height int

	//scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Masukkan inputan anda: ")
	// scanner.Scan()
	// input := scanner.Text()

	for {
		fmt.Println("\nWELCOME USER")
		fmt.Println("Please choose the menu belows to run function as per as your needs :")
		fmt.Println("1. Calculate Circumference And Area of Circle")
		fmt.Println("2. Print Triangle")
		fmt.Println("3. Show Slice")
		fmt.Println("4. Show Profile")
		fmt.Println("5. Payment")
		fmt.Println("0. Exit")

		fmt.Println("Select your choice (1,2,3,4,5):")
		fmt.Scanf("%d", &choice)
		// scanner.Scan()
		// choice := scanner.Text()

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
		case choice == 5:
			fmt.Println("================================================")
			fmt.Println("You have choosen menu number 5 :")
			people.Greet("ari")
			fmt.Println("================================================")
		case choice == 6:
			fmt.Println("================================================")
			fmt.Println("You have choosen menu number 6 :")
			(ShowPayment())
			fmt.Println("================================================")
		default:
			fmt.Println("==========================================================================")
			fmt.Printf("You have choosen menu number %d and the menu is not available !!\n", choice)
			fmt.Println("Goodbye !!")
			fmt.Println("==========================================================================")
		}

		if choice > 8 || choice <= 0 {
			break
		}

	}
}

func ShowPayment() {
	var total int
	var choice int
	var arrTotal []int
	//var answer string
	scanner := bufio.NewScanner(os.Stdin)

	for {

		//defer payments.PayWithDcelup(payDcelup, 000)
		userService := &payments.UserService{}
		bankService := &payments.PayWithBankService{}
		qrisService := &payments.PayWithQrisService{}
		dCelupService := &payments.PayWithDcelupService{}

		for {

			fmt.Printf("Apakah anda ingin menambahkan barang (y/n) ? ")
			scanner.Scan()
			input := scanner.Text()

			if input == "y" {
				fmt.Println("\nWelcome User")
				fmt.Printf("Please input sum of payment:")
				fmt.Scanf("%d", &total)
				arrTotal = append(arrTotal, total)

				if total < 0 {
					fmt.Println("Error")
					break
				}

			} else {
				fmt.Printf("\n=======================================:\n")
				fmt.Println("1. Pay via Bank")
				fmt.Println("2. Pay via Qris")
				fmt.Println("3. Pay via Dcelup")
				fmt.Println("4. Exit")
				fmt.Printf("Please choose your payment method:")
				fmt.Scanf("%d", &choice)
				fmt.Printf("\n=======================================:\n")
				// scanner.Scan()
				// input := scanner.Text()
				//var arrTotal []int

				switch true {
				case choice == 1:
					userService.RegisterPay(bankService, arrTotal)
					os.Exit(0)
				case choice == 2:
					userService.RegisterPay(qrisService, arrTotal)
					os.Exit(0)
				case choice == 3:
					userService.RegisterPay(dCelupService, arrTotal)
				case choice == 4:
					userService.RegisterPay(dCelupService, arrTotal)
					os.Exit(0)
				default:
					fmt.Println("Error")
				}
			}

		}

		// 	if answer == "n" || answer == "N" {
		// 		fmt.Printf("Please input sum of payment:")
		// 		fmt.Scanf("%d", &total)

		// 	} else {

		// 		// for {
		// 		// 	fmt.Printf("Apakah anda ingin melakukan pembayaran kembali (y/n) ? ")
		// 		// 	scanner.Scan()
		// 		// 	input := scanner.Text()

		// 		// 	if input == "y" || input == "Y" {
		// 		// 		fmt.Printf("Masukkan total pembayaran: ")
		// 		// 		fmt.Scanf("%d", &total)
		// 		// 		payments.PayWithAripay(pay, total)
		// 		// 	} else {
		// 		// 		// ari := payments.PayWithAripay(pay, total)
		// 		// 		// fmt.Print(ari)
		// 		// 	}

		// 		// }

		// 	}
		// }

	}
}
