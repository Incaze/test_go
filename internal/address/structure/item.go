package structure

import (
	"fmt"
	"strconv"
)

const ItemElement = "item"

type Item struct {
	City   string `xml:"city,attr"`
	Street string `xml:"street,attr"`
	House  string `xml:"house,attr"`
	Floor  string `xml:"floor,attr"`
}

func (item *Item) GetRecord() string {
	return fmt.Sprintf("City %s, Street %s, House %s, Floor %s", item.City, item.Street, item.House, item.Floor)
}

func (item *Item) GetNumericFloor() (int, error) {
	floor, err := strconv.Atoi(item.Floor)
	if err != nil {
		return 0, fmt.Errorf("provided non-numeric floor: %s", item.Floor)
	}

	if floor > MaxFloor || floor < 1 {
		return 0, fmt.Errorf("provided unsupported floor: %d", floor)
	}

	return floor, nil
}
