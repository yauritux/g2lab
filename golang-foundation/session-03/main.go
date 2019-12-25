package main

import (
	"bufio"
	"fmt"
	"os"
	_ "strconv"
	ent "g2lab/ecommerce/entity"
)

func showMenu() {
	fmt.Println("Program Menu:")
	fmt.Println("1 Product Entry")
	fmt.Println("2 Browse Product")
	fmt.Println("3 Exit")
	fmt.Println()
	fmt.Println("Please enter your selected menu number [e.g.: '1', '2'], or '3' to exit the program.")
	fmt.Print("> ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var opt string
menu:
	for {
		showMenu()
		scanner.Scan()
		opt = scanner.Text()
		switch opt {
		case "1":
			if err := ent.AddProduct(scanner); err != nil {
				fmt.Printf("Cannot add product. Error=%s\n", err.Error())
			}
			fmt.Println("Press <enter> to continue...")
			fmt.Println()
			scanner.Scan()
		case "2":
			ent.ShowProducts()
			fmt.Println("Press <enter> to continue...")
			fmt.Println()
			scanner.Scan()
		case "3":
			fmt.Println("Thanks for your visit!")
			fmt.Println()
			break menu
		default:
			fmt.Printf("Sorry, %s is not listed in the menu option. Try again!\n\n", opt)
			fmt.Println("Press <enter> to continue...")
			fmt.Println()
			scanner.Scan()
		}
	}
}
