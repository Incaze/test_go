package address

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"test_go/internal/address/structure"
)

func Process(xmlFile *os.File) (map[string]int, map[string]structure.CityInfo, error) {
	itemsInfo := make(map[string]int)
	cityInfo := make(map[string]structure.CityInfo)
	decoder := xml.NewDecoder(xmlFile)
	for {
		token, err := decoder.Token()
		if token == nil || err == io.EOF {
			break
		} else if err != nil {
			return nil, nil, fmt.Errorf("failed to decode xml file: %s", err.Error())
		}

		switch element := token.(type) {
		case xml.StartElement:
			if element.Name.Local == structure.ItemElement {
				var item structure.Item
				err = decoder.DecodeElement(&item, &element)
				if err != nil {
					return nil, nil, fmt.Errorf("failed to decode xml element: %s", err.Error())
				}

				floor, err := item.GetNumericFloor()
				if err != nil {
					return nil, nil, err
				}

				key := item.GetRecord()
				if _, ok := itemsInfo[key]; ok {
					itemsInfo[key]++
				} else {
					itemsInfo[key] = 1
				}

				if val, ok := cityInfo[item.City]; ok {
					val.Floors[floor-1]++
					cityInfo[item.City] = val
				} else {
					val.Floors[floor-1] = 1
					cityInfo[item.City] = val
				}
			}
		default:
			// to do nothing
		}
	}

	return itemsInfo, cityInfo, nil
}
