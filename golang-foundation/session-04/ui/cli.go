package ui

import (
	"bufio"
	"fmt"
	"strconv"

	ent "g2lab.co/ecommerce/entity"
	json "g2lab.co/ecommerce/storage"
)

func ShowProducts() {
	data, err := json.ReadFromJSONFile()
	if err != nil {
		fmt.Println("--- No Products to be displayed ---")
		return
	}

	ent.SetItems(data)

	if len(ent.GetItems()) == 0 {
		fmt.Println("--- No Products to be displayed ---")
		return
	}

	for _, v := range ent.GetItems() {
		fmt.Printf("|%5s\t|%-30s\t|%5d\t|%15.2f\n", v.ID, v.Name, v.Stock, v.Price)
	}
}

func AddProduct(scanner *bufio.Scanner) error {
	fmt.Print("\nItem code:")
	scanner.Scan()
	id := scanner.Text()
	fmt.Print("\nItem name: ")
	scanner.Scan()
	name := scanner.Text()

	item := ent.NewItem(id, name)

	fmt.Print("\nItem stock: ")
	scanner.Scan()
	stock, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return err
	}
	item.Stock = stock

	fmt.Print("\nItem price: ")
	scanner.Scan()
	price, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		return err
	}
	item.Price = price

	_, err = ent.AppendItem(item)
	if err != nil {
		return err
	}

	if err = WriteToJSONFile(); err != nil {
		fmt.Printf("Failed to write data into json file.Error=%s", err.Error())
		return err
	}
	
	fmt.Printf("Successfully added %s into the system.\n\n", item.Name)
	return nil
}

func WriteToJSONFile() error {
	return json.WriteToJSONFile(ent.GetItems())
}