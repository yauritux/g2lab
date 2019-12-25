package entity

import (
	"errors"
	"fmt"
)

type Item struct {
	ID    string `json:"item_id"`
	Name  string `json:"item_name"`
	Stock int `json:"item_stock"`
	Price float64 `json:"item_price"`
}

var items map[string]*Item

func init() {
	items = make(map[string]*Item)
}

func NewItem(id, name string) *Item {
	return &Item{
		ID:   id,
		Name: name,
	}
}

func AppendItem(item *Item) (map[string]*Item, error) {
	if _, found := items[item.ID]; found {
		return items, errors.New(fmt.Sprintf("Item with ID %s already exists", item.ID))
	}
	items[item.ID] = item
	return items, nil
}

func SetItems(data map[string]*Item) {
	items = data
}

func GetItems() map[string]*Item {
	return items
}
