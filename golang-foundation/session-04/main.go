package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var items = make(map[string]item, 5)

type customer struct {
	name     string
	email    string
	phone    string
	shipping shippingAddress
	address  defaultAddress
}

type shippingAddress struct {
	streetName, city, postCode string
}

type defaultAddress struct {
	streetName string
	city       string
	postCode   string
}

type item struct {
	id    string
	name  string
	stock int
	price int
}

type cart struct {
	customer   customer
	orderedQty int
	item       item
}

func showMenu() {
	fmt.Printf("\nAdmin Menu:\n")
	fmt.Println("1.1 Product Entry")
	fmt.Println("1.2 Edit Product")
	fmt.Println()
	fmt.Println("Customer Menu:")
	fmt.Println("2.1 Browse Product")
	fmt.Println("2.2 Add to Cart")
	fmt.Println("2.3 Show Cart Items")
	fmt.Println("2.4 Checkout")
	fmt.Println()
	fmt.Println("3. Exit")
	fmt.Println()
	fmt.Println("Please enter your selected menu number [e.g.: '1.1', '2.1', etc], or '3' to end the program.")
	fmt.Print("> ")
}

func productEntry(scanner *bufio.Scanner) {
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
}

func showProducts() {
	for _, v := range items {
		fmt.Println("Item Code:", v.id)
		fmt.Println("Item Name:", v.name)
		fmt.Println()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var selectedOption string
menu:
	for {
		showMenu()
		scanner.Scan()
		selectedOption = scanner.Text()
		switch selectedOption {
		case "1.1":
			productEntry(scanner)
		case "1.2":
			showProducts()
		case "3":
			fmt.Println("Thanks for your visit!")
			break menu
		}
	}
}
