package internal

import (
	"fmt"
	"test_go/internal/address"
)

func WriteData(itemsInfo map[string]int, cityInfo map[string]address.CityInfo) {
	for key, val := range itemsInfo {
		if val == 1 {
			continue
		}

		fmt.Printf("Record: %s | Count duplicates: %d\n", key, val)
	}
	fmt.Println("---")
	for key, val := range cityInfo {
		fmt.Println("City: ", key)
		for cityKey, cityVal := range val.Floors {
			fmt.Printf("  Floor %d: %d\n", cityKey+1, cityVal)
		}
	}
	fmt.Println("---")
	fmt.Printf("Count records (Exclude duplicates): %d", len(itemsInfo))
}
