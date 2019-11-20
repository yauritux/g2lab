package entity

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var Customers map[int]Customer

type Address struct {
	street, no, city, region string
}

type Customer struct {
	id                                int
	firstname, lastname, phone, email string
	shippingaddress                   Address
	defaultaddress                    Address
}

func init() {
	Customers = make(map[int]Customer)
}

func ViewCustomers() {
	if len(Customers) == 0 {
		fmt.Println("There's no products!")
	} else {
		fmt.Printf("|NO\t|FIRST NAME\t|LAST NAME\t|PHONE\t|EMAIL\t|DEFAULT ADDRESS\t|SHIPPING ADDRESS\t|\n")
		for _, v := range Customers {
			fmt.Printf("|%d.\t|%s\t|%s\t|%s\t|%s\t|%s\t|%s\t|\n", v.id, v.firstname, v.lastname, v.phone, v.email, v.defaultaddress.street+" "+v.defaultaddress.no+" "+v.defaultaddress.city+" "+v.defaultaddress.region, v.shippingaddress.street+" "+v.shippingaddress.no+" "+v.shippingaddress.city+" "+v.shippingaddress.region)
		}
	}
}

func CustomerMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Select Menu : ")
		fmt.Println("1. View Customers")
		fmt.Println("2. Add Customers")
		fmt.Println("3. Remove Customers")
		fmt.Println("4. Back")
		fmt.Printf("Input : ")
		scanner.Scan()
		sel, err := strconv.Atoi(scanner.Text())
		if err == nil {
			if sel != 1 && sel != 2 && sel != 3 && sel != 4 {
				fmt.Println("Wrong input!")
			} else if sel == 1 {
				viewCustomers()
			} else if sel == 2 {
				fmt.Printf("Input ID : ")
				scanner.Scan()
				id, _ := strconv.Atoi(scanner.Text())
				fmt.Printf("Input First Name : ")
				scanner.Scan()
				firstname := scanner.Text()
				fmt.Printf("Input Last Name : ")
				scanner.Scan()
				lastname := scanner.Text()
				fmt.Printf("Input Phone : ")
				scanner.Scan()
				phone := scanner.Text()
				fmt.Printf("Input Email : ")
				scanner.Scan()
				email := scanner.Text()
				fmt.Printf("Input Default Address Street : ")
				scanner.Scan()
				defstreet := scanner.Text()
				fmt.Printf("Input Default Address No : ")
				scanner.Scan()
				defno := scanner.Text()
				fmt.Printf("Input Default Address City : ")
				scanner.Scan()
				defcity := scanner.Text()
				fmt.Printf("Input Default Address Region : ")
				scanner.Scan()
				defregion := scanner.Text()

				var shipstreet, shipno, shipcity, shipregion, same string

				for {
					fmt.Printf("Is your shipping address same as default address? (Y/N) : ")
					scanner.Scan()
					same = scanner.Text()
					if same != "Y" && same != "N" {
						fmt.Println("Wrong input!")
					} else {
						break
					}
				}

				if same == "Y" {
					shipstreet = defstreet
					shipno = defno
					shipregion = defregion
					shipcity = defcity
				} else {
					fmt.Printf("Input Shipping Address Street : ")
					scanner.Scan()
					shipstreet = scanner.Text()
					fmt.Printf("Input Shipping Address No : ")
					scanner.Scan()
					shipno = scanner.Text()
					fmt.Printf("Input Shipping Address City : ")
					scanner.Scan()
					shipcity = scanner.Text()
					fmt.Printf("Input Shipping Address Region : ")
					scanner.Scan()
					shipregion = scanner.Text()
				}
				Customers[id] = Customer{id, firstname, lastname, phone, email, Address{defstreet, defno, defcity, defregion}, Address{shipstreet, shipno, shipcity, shipregion}}
				fmt.Println("Success")
			} else if sel == 3 {
				fmt.Printf("Input ID : ")
				scanner.Scan()
				id, _ := strconv.Atoi(scanner.Text())
				_, ok := Customers[id]
				if ok {
					delete(Customers, id)
					fmt.Println("Customer deleted")
				} else {
					fmt.Println("Customer not found!")
				}
			} else if sel == 4 {
				return
			}
		} else {
			fmt.Println("Wrong input!")
		}
	}
}
