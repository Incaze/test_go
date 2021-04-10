package structure_test

import (
	"test_go/internal/address/structure"
	"testing"
)

const (
	getRecordExpected                     = "City Saransk, Street Sovetskaya, House 35, Floor 5"
	getNumericFloorExpectedPositive       = 5
	getNumericFloorExpectedStringNegative = "provided non-numeric floor: foo"
	getNumericFloorExpectedLimitNegative  = "provided unsupported floor: 7"
	getNumericFloorProvidedLimitNegative  = "7"
	getNumericFloorProvidedStringNegative = "foo"
)

var getRecordActualItem = structure.Item{
	City:   "Saransk",
	Street: "Sovetskaya",
	House:  "35",
	Floor:  "5",
}

func TestItem_GetRecord(t *testing.T) {
	actual := getRecordActualItem.GetRecord()
	if actual != getRecordExpected {
		t.Fatalf("not equal: actual - %s; expected - %s", actual, getRecordExpected)
	}
}

func TestItem_GetNumericFloor_Positive(t *testing.T) {
	number, err := getRecordActualItem.GetNumericFloor()
	if err != nil {
		t.Fatal(err.Error())
	}
	if number != getNumericFloorExpectedPositive {
		t.Fatalf("not equal: actual - %d; expected - %d", number, getNumericFloorExpectedPositive)
	}
}

func TestItem_GetNumericFloor_StringNegative(t *testing.T) {
	test := structure.Item{
		Floor: getNumericFloorProvidedStringNegative,
	}
	_, err := test.GetNumericFloor()
	if err == nil {
		t.Fatal("Expected some error")
	}

	if err.Error() != getNumericFloorExpectedStringNegative {
		t.Fatalf("not equal: actual - %s; expected - %s", err, getNumericFloorExpectedStringNegative)
	}
}

func TestItem_GetNumericFloor_LimitNegative(t *testing.T) {
	test := structure.Item{
		Floor: getNumericFloorProvidedLimitNegative,
	}
	_, err := test.GetNumericFloor()
	if err == nil {
		t.Fatal("Expected some error")
	}

	if err.Error() != getNumericFloorExpectedLimitNegative {
		t.Fatalf("not equal: actual - %s; expected - %s", err, getNumericFloorExpectedLimitNegative)
	}
}

func BenchmarkItem_GetRecord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getRecordActualItem.GetRecord()
	}
}

func BenchmarkItem_GetNumericFloor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = getRecordActualItem.GetNumericFloor()
	}
}
