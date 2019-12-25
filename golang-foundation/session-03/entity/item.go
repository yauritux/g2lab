package entity

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
)

var items map[string]*Item

type Item struct {
	id    string
	name  string
	stock int
	price float64
}

func init() {
	items = make(map[string]*Item)
}

func NewItem(id, name string) (*Item, error) {
	if IsItemExist(id) {
		return nil, errors.New(fmt.Sprintf("Item with ID %s already exists", id))
	}
	return &Item{id: id, name: name}, nil
}

func IsItemExist(id string) bool {
	_, ok := items[id]
	return ok
}

func (i *Item) ID() string {
	return i.id
}

func (i *Item) Name() string {
	return i.name
}

func (i *Item) Stock() int {
	return i.stock
}

func (i *Item) SetStock(s int) *Item {
	i.stock = s
	return i
}

func (i *Item) Price() float64 {
	return i.price
}

func (i *Item) SetPrice(p float64) *Item {
	i.price = p
	return i
}

func ShowProducts() {
	if len(items) == 0 {
		fmt.Println("--- No Products to be displayed ---")
		return
	}

	for _, v := range items {
		fmt.Printf("|%5s\t|%-30s\t|%5d\t|%15.2f\n", v.id, v.name, v.stock, v.price)
	}
}

func AddProduct(scanner *bufio.Scanner) error {
	fmt.Print("\nItem code:")
	scanner.Scan()
	id := scanner.Text()
	fmt.Print("\nItem name: ")
	scanner.Scan()
	name := scanner.Text()

	item, err := NewItem(id, name)	
	if err != nil {
		return err
	}

	fmt.Print("\nItem stock: ")
	scanner.Scan()
	stock, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return err
	}
	item.stock = stock

	fmt.Print("\nItem price: ")
	scanner.Scan()
	price, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		return err
	}
	item.price = price

	items[item.id] = item
	fmt.Printf("Successfully added %s into the system.\n\n", item.name)
	return nil
}
