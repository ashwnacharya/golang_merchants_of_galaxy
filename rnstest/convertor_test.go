package rnstest

import (
	"testing"
	"merchants_of_galaxy/rns"
)

func TestConvertToIntegerInvalidInput(t *testing.T) {
	_, err := rns.ConvertToInteger("NONSENSE")
	if err == nil {
		t.Error("Invalid input should throw an error")
	}
}

func TestConvertToInteger(t *testing.T) {
	number, err := rns.ConvertToInteger("XVIII")

	if err != nil {
		t.Error("XVIII is a valid input")
	} else if number != 18 {
		t.Errorf("XVIII is 18, not %d", number)
	}
}

func TestConvertToInteger2(t *testing.T) {
	number, err := rns.ConvertToInteger("MMMDCCCLXXXVIII")

	if err != nil {
		t.Error("XVIII is a valid input")
	} else if number != 3888 {
		t.Errorf("XVIII is 3888, not %d", number)
	}
}

func TestConvertToInteger3(t *testing.T) {
	number, err := rns.ConvertToInteger("XIV")

	if err != nil {
		t.Error("XVIII is a valid input")
	} else if number != 14 {
		t.Errorf("XIV is 14, not %d", number)
	}
}

func TestConvertToInteger4(t *testing.T) {
	number, err := rns.ConvertToInteger("CDXLIV")

	if err != nil {
		t.Error("CDXLIV is a valid input")
	} else if number != 444 {
		t.Errorf("XIV is 444, not %d", number)
	}
}

func TestConvertToInteger5(t *testing.T) {
	number, err := rns.ConvertToInteger("CMXLIV")

	if err != nil {
		t.Error("CMXLIV is a valid input")
	} else if number != 944 {
		t.Errorf("XIV is 944, not %d", number)
	}
}

func TestConvertToInteger6(t *testing.T) {
	_, err := rns.ConvertToInteger("VX")

	if err == nil {
		t.Error("VX is not a valid input")
	}
}

func TestConvertToRoman1(t *testing.T) {
	actual := rns.ConvertToRoman(1)
	expected := []string{"I"}

	if len(actual) != len(expected) {
		t.Error("Returned array is not of same length as expected array")
	} else {

		for index, value := range actual {
			if value != expected[index] {
				t.Errorf("Expected %s at position %d, got %s", expected[index], index, value)
			}
		}
	}
}

func TestConvertToRoman2(t *testing.T) {
	actual := rns.ConvertToRoman(3888)
	expected := []string{"M", "M", "M", "D", "C", "C", "C", "L", "X", "X", "X", "V", "I", "I", "I"}

	if len(actual) != len(expected) {
		t.Error("Returned array is not of same length as expected array")
	} else {

		for index, value := range actual {
			if value != expected[index] {
				t.Errorf("Expected %s at position %d, got %s", expected[index], index, value)
			}
		}
	}
}

func TestConvertToRoman3(t *testing.T) {
	actual := rns.ConvertToRoman(444)
	expected := []string{"C", "D", "X", "L", "I", "V"}

	if len(actual) != len(expected) {
		t.Error("Returned array is not of same length as expected array")
	} else {

		for index, value := range actual {
			if value != expected[index] {
				t.Errorf("Expected %s at position %d, got %s", expected[index], index, value)
			}
		}
	}
}