package structure

import (
	"fmt"
)

const ItemElement = "item"

type Item struct {
	City   string `xml:"city,attr"`
	Street string `xml:"street,attr"`
	House  string `xml:"house,attr"`
	Floor  string `xml:"floor,attr"`
}

func (item *Item) GetRecord() string {
	return fmt.Sprintf("%s %s %s %s", item.City, item.Street, item.House, item.Floor)
}
