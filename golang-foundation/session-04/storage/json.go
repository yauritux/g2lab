package storage

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	ent "g2lab.co/ecommerce/entity"
)

func WriteToJSONFile(items map[string]*ent.Item) error {
	if items == nil {
		return errors.New("Failed to writing into file. No data to write!")
	}
	b, _ := json.Marshal(items)
	if err := ioutil.WriteFile("data.json", b, 0644); err != nil {
		return err
	}
	return nil
}

func ReadFromJSONFile() (map[string]*ent.Item, error) {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		return nil, err
	}
	items := ent.GetItems()
	if err = json.Unmarshal(data, &items); err != nil {
		return nil, err
	}
	return items, nil
}
