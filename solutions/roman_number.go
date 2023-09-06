package solutions

import "strings"

var romToInt = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000}

var intToRom = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M"}

func RomanToInt(s string) int {
	totalValue := 0
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		symbol := string(s[i])
		currentValue := romToInt[symbol]
		if i > 0 {
			prev := peek(s, i)
			if prev < currentValue {
				currentValue -= prev
				i--
			}
		}
		totalValue += currentValue

	}
	return totalValue
}

func peek(s string, index int) int {
	return romToInt[string(s[index-1])]
}

func IntToRoman(num int) string {
	output := ""
	divider := 10
	result := 1

	for result > 0 {
		result = num / divider
		rem := num % divider
		place := divider / 10
		if rem!=0{
			rom := computeRoman(rem, place,result)
			output = rom + output
		}
		num = result*divider
		divider *= 10
	}
	return output
}

func computeRoman(rem, place,result int) string {
	// check the number itself
	if val, ok := intToRom[rem]; ok {
		return val
	}
	// check upper bound
	upper := rem + place
	romanNumber := intToRom[place]
	if val, ok := intToRom[upper]; ok {
		return romanNumber + val
	}

	// check lower
	for i:=1;i<4;i++{
		lower := rem-i*place
		if val, ok := intToRom[lower];ok{
			return val+strings.Repeat(romanNumber,i)
		}
	}
	return strings.Repeat(romanNumber, rem)
}
