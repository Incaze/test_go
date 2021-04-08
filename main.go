package main

import (
	"log"
	"os"
	"test_go/internal"
	"test_go/internal/address"
)

const testFilepath = "address.xml"

func main() {
	xmlFile, err := os.Open(testFilepath)
	if err != nil {
		log.Printf("failed to open xml file: %s", testFilepath)
		return
	}

	itemsInfo, cityInfo, err := address.Process(xmlFile)
	defer xmlFile.Close()

	if err != nil {
		log.Println(err.Error())
		return
	}

	internal.WriteData(itemsInfo, cityInfo)
}
