package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Please enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	fmt.Print("Enter the item name: ")
	scanner.Scan()
	item := scanner.Text()
	fmt.Print("Your ordered amount: ")
	scanner.Scan()
	qty, _ := strconv.Atoi(scanner.Text())
	fmt.Print("Price per item: ")
	scanner.Scan()
	price, _ := strconv.Atoi(scanner.Text())
	fmt.Println("Welcome,", name)
	fmt.Printf("Your've ordered %d %v\n", qty, item)
	fmt.Printf("Total price: %d\n", price*qty)
}
