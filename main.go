package main

import (
	"fmt"
	"log"
	"os"
	"test_go/internal/address"
)

const filepath = "address.xml"

func main() {
	xmlFile, err := os.Open(filepath)
	if err != nil {
		log.Printf("failed to open xml file: %s", filepath)
		return
	}

	itemsInfo, cityInfo, err := address.Process(xmlFile)
	defer xmlFile.Close()

	if err != nil {
		log.Println(err.Error())
		return
	}

	// some output
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
