package structure

import (
	"encoding/xml"
)

const RootElement = "root"

type Root struct {
	XMLName xml.Name `xml:"root"`
	Items   []Item   `xml:"item"`
}
