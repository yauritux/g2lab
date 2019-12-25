package storage

import (
	"testing"
	"os"
	ent "g2lab.co/ecommerce/entity"
)

func TestWriteToJSONFileNoDataToWrite(t *testing.T) {
	err := WriteToJSONFile(nil)
	if err == nil {
		t.Error("Expected to got an error")
	}
}

func TestWriteToJSONFile(t *testing.T) {
	item := ent.NewItem("001", "Nike Shoes")
	item.Stock = 750
	item.Price = 950000.0
	items, _ := ent.AppendItem(item)
	err := WriteToJSONFile(items)
	if err != nil {
		t.Error("Expected no error")
	}
}

func TestReadFromUnexistingJSONFile(t *testing.T) {
	os.Remove("data.json")
	_, err := ReadFromJSONFile()
	if err == nil {
		t.Error("Expected to got an error due to unexisting file")
	}
}

func TestReadFromJSONFile(t *testing.T) {
	os.Remove("data.json")
	item := ent.NewItem("001", "Nike Shoes")
	item.Stock = 750
	item.Price = 950000.0
	items, _ := ent.AppendItem(item)
	WriteToJSONFile(items)
	items, err := ReadFromJSONFile()
	if err != nil {
		t.Errorf("Expected no error, but got an error of %v", err)
	}
	if len(items) == 0 {
		t.Errorf("Expected total item of %d, got %d", 1, len(items))
	}
	if items["001"].Name != "Nike Shoes" {
		t.Errorf("Expected item name of %s, got %s", "Nike Shoes", items["001"].Name)
	}
}