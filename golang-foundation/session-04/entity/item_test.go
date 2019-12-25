package entity

import "testing"

func TestNewItem(t *testing.T) {
	item := NewItem("001", "Nike Shoes")
	if item.ID != "001" {
		t.Errorf("Expected item ID %s, got %s", "001", item.ID)
	}
	if item.Name != "Nike Shoes" {
		t.Errorf("Expected item name %s, got %s", "Nike Shoes", item.Name)
	}
}

func TestSetAndGetItemStock(t *testing.T) {
	item := NewItem("001", "Nike Shoes")
	item.Stock = 700
	if item.Stock != 700 {
		t.Errorf("Expected item stock %d, got %d", 700, item.Stock)
	}
}

func TestSetAndGetItemPrice(t *testing.T) {
	item := NewItem("001", "Nike Shoes")
	item.Price = 750000.0
	if item.Price != 750000.0 {
		t.Errorf("Expected item price %f, got %f", 750000.0, item.Price)
	}
}

func TestAppendItem(t *testing.T) {
	item := NewItem("001", "Nike Shoes")
	item.Stock = 200
	item.Price = 750000.0
	items, err := AppendItem(item)
	if err != nil {
		t.Errorf("Expected no error, but got an error of %v", err)
	}
	if _, found := items["001"]; !found {
		t.Errorf("Expected %s to be exist inside the items", item.ID)
	}
}

func TestAppendItemExistingID(t *testing.T) {
	item1 := NewItem("001", "Nike Shoes")
	item1.Stock = 500
	item1.Price = 900000.0
	items, err := AppendItem(item1)
	item2 := NewItem("001", "Programming T-Shirt")
	item2.Stock = 1300
	item2.Price = 250000.0
	items, err = AppendItem(item2)
	if err == nil {
		t.Error("Expected an error, but no errors is thrown")
	}
	if len(items) != 1 {
		t.Errorf("Expected total items of %d, but got %d", 1, len(items))
	}
}

func TestGetItems(t *testing.T) {
	items := GetItems()
	if items == nil {
		t.Error("Expected items not nil")
	}
	item := NewItem("001", "Nike Shoes")
	item.Stock = 500
	item.Price = 750000.0
	AppendItem(item)
	items = GetItems()
	if len(items) != 1 {
		t.Errorf("Expected total items of %d, got %d", 1, len(items))
	}
}