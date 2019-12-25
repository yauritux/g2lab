package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var items = make(map[string]item, 5)

type item struct {
	id    string
	name  string
	stock int
	price int
}

func showMenu() {
	fmt.Println("Program Menu:")
	fmt.Println("1 Product Entry")
	fmt.Println("2 Browse Product")
	fmt.Println("3 Exit")
	fmt.Println()
	fmt.Println("Please enter your selected menu number [e.g.: '1', '2'], or '3' to exit the program.")
	fmt.Print("> ")
}

func addProduct(scanner *bufio.Scanner) {
	var item item
	fmt.Print("\nItem code:")
	scanner.Scan()
	item.id = scanner.Text()
	fmt.Print("\nItem name: ")
	scanner.Scan()
	item.name = scanner.Text()
	fmt.Print("\nItem stock: ")
	scanner.Scan()
	item.stock, _ = strconv.Atoi(scanner.Text())
	fmt.Print("\nItem price: ")
	scanner.Scan()
	item.price, _ = strconv.Atoi(scanner.Text())
	items[item.id] = item
	fmt.Printf("Successfully added %s into the system.\n\n", item.name)
}

func showProducts() {
	if len(items) == 0 {
		fmt.Println("--- No Products to be displayed ---")
		return
	}
	for _, v := range items {
		fmt.Println("Item Code:", v.id)
		fmt.Println("Item Name:", v.name)
		fmt.Println()
	}
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
			addProduct(scanner)
			fmt.Println("Press <enter> to continue...")
			fmt.Println()
			scanner.Scan()
		case "2":
			showProducts()
			fmt.Println("Press <enter> to continue...")
			fmt.Println()
			scanner.Scan()
		case "3":
			fmt.Println("Thanks for your visit!")
			break menu
		default:
			fmt.Printf("Sorry, %s is not listed in the menu option. Try again!\n\n", opt)			
		}
	}
}
