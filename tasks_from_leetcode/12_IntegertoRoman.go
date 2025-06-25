package tasks_from_leetcode

func romanAddValueByRoman(result *string, lastSymbols int, alphabet *map[int]string) {
	var count int
	multiplicities := []int{1000, 500, 100, 50, 10, 5, 1}
	tempResult := ""
	for _, multiplicity := range multiplicities {
		if lastSymbols > multiplicity {
			count = lastSymbols / multiplicity
			for i := 0; i < count; i++ {
				value, _ := (*alphabet)[multiplicity]
				tempResult += value
			}
			lastSymbols -= multiplicity * count
		}
	}
	*result = tempResult + *result

}

func intToRoman(num int) string {
	alphabet := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
		4:    "IV",
		9:    "IX",
		40:   "XL",
		90:   "XC",
		400:  "CD",
		900:  "CM",
	}

	result := ""
	multiplier := 1
	for num > 0 {
		lastSymbol := num % 10 * multiplier
		num /= 10
		multiplier *= 10

		if value, found := alphabet[lastSymbol]; found {
			result = value + result
			continue
		}

		romanAddValueByRoman(&result, lastSymbol, &alphabet)
	}

	return result
}
