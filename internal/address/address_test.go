package address_test

import (
	"os"
	"reflect"
	"test_go/internal/address"
	"test_go/internal/address/structure"
	"testing"
)

const (
	benchFilepath = "../../address.xml"
	testFilePath  = "../../addr.xml"
)

var itemsInfoExpected = map[string]int{
	"City Азов, Street Просека, улица, House 156, Floor 3":             1,
	"City Балаково, Street Барыши, местечко, House 67, Floor 2":        2,
	"City Барнаул, Street Дальняя улица, House 56, Floor 2":            2,
	"City Барнаул, Street Дальняя улица, House 56, Floor 3":            1,
	"City Братск, Street Большая Октябрьская улица, House 65, Floor 5": 1,
}

var cityInfoExpected = map[string]structure.CityInfo{
	"Азов":     {Floors: [5]int{0, 0, 1, 0, 0}},
	"Балаково": {Floors: [5]int{0, 2, 0, 0, 0}},
	"Барнаул":  {Floors: [5]int{0, 2, 1, 0, 0}},
	"Братск":   {Floors: [5]int{0, 0, 0, 0, 1}},
}

func TestProcess(t *testing.T) {
	xmlFile, err := os.Open(testFilePath)
	if err != nil {
		t.Fatalf("failed to open xml file: %s", testFilePath)
	}

	itemsInfoActual, cityInfoActual, err := address.Process(xmlFile)
	if err != nil {
		t.Fatal(err.Error())
	}

	eqItemsInfo := reflect.DeepEqual(itemsInfoExpected, itemsInfoActual)
	if !eqItemsInfo {
		t.Fatal("Not equals items info")
	}

	eqCityInfo := reflect.DeepEqual(cityInfoExpected, cityInfoActual)
	if !eqCityInfo {
		t.Fatal("Not equals city info")
	}

	defer xmlFile.Close()
}

func BenchmarkProcess(b *testing.B) {
	xmlFile, err := os.Open(benchFilepath)
	if err != nil {
		b.Fatalf("failed to open xml file: %s", benchFilepath)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, err = address.Process(xmlFile)
	}
	if err != nil {
		b.Fatal(err.Error())
	}

	defer xmlFile.Close()
}
