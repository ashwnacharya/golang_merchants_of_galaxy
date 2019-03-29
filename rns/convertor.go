package rns

import (
	"fmt"
)

// ConvertToInteger converts a Roman numeral to Integer
func ConvertToInteger(input string) (int, error) {

	var digits []int
	romanLookup := make(map[string]int)

	romanLookup["I"] = 1
	romanLookup["V"] = 5
	romanLookup["X"] = 10
	romanLookup["L"] = 50
	romanLookup["C"] = 100
	romanLookup["D"] = 500
	romanLookup["M"] = 1000


	for _, character := range input {

		found := false
		for _, valid := range "IVXLCDM" {
			if string(character) == string(valid) {
				digits = append(digits, romanLookup[string(character)])
				found = true
				break
			}
		}

		if found == false {
			return 0, fmt.Errorf("%c is not a valid roman digit", character)
		}
	}
	
	sum := 0
	length := len(digits)

	for index, digit := range digits {
		if index == length -1 {
			sum = sum + digit
			continue
		}

		if digit < digits[index+1] {
			if (digit == 1 && (digits[index+1] == 5 || digits[index+1] == 10)) || 
				(digit == 10 && (digits[index+1] == 50 || digits[index+1] == 100)) ||
				(digit == 100 && (digits[index+1] == 500 || digits[index+1] == 1000)) {
				sum = sum - digit
			} else {
				return 0, fmt.Errorf("%s is not a valid roman numeral", input)
			}
				
		} else {
			sum = sum + digit
		}
	}

	return sum, nil
}

// ConvertToRoman converts an integer to a roman numeral, represented as an array of strings.
func ConvertToRoman(number int) ([]string) {
	integerToRomanLookup := make(map[int][]string)

	integerToRomanLookup[1000] = []string{"M"}
	integerToRomanLookup[900]  = []string{"C", "M"}
	integerToRomanLookup[500]  = []string{"D"}
	integerToRomanLookup[400]  = []string{"C", "D"}
	integerToRomanLookup[100]  = []string{"C"}
	integerToRomanLookup[50]   = []string{"L"}
	integerToRomanLookup[40]   = []string{"X", "L"}
	integerToRomanLookup[10]   = []string{"X"}
	integerToRomanLookup[5]    = []string{"V"}
	integerToRomanLookup[4]    = []string{"I", "V"}
	integerToRomanLookup[1]  =   []string{"I"}


	digitKeys := []int{1000, 900, 500, 400, 100, 50, 40, 10, 5, 4, 1}
	digitIndex := 0

	var numeral []string

	for number > 0 {
				if number >= digitKeys[digitIndex] {
			numeral = append(numeral, integerToRomanLookup[digitKeys[digitIndex]]...)
			number = number - digitKeys[digitIndex]
		} else {
			digitIndex = digitIndex + 1
		}
	}

	return numeral
}