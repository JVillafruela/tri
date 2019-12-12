package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Item struct {
	Text     string
	Priority int
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}

}

// save items slide #219
func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("SaveItems : %s %s \n", filename, string(b))
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, nil
	}
	return items, nil
}
