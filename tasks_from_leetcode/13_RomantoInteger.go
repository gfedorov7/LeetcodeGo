package tasks_from_leetcode

func romanToInt(s string) int {
	alphabet := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	i := len(s) - 1
	var result, lastSymbol int
	var found bool
	for i >= 0 {
		if i-1 >= 0 {
			lastSymbol, found = alphabet[s[i-1:i+1]]
			if found {
				i -= 2
				result += lastSymbol
				continue
			}
		}
		lastSymbol, _ = alphabet[s[i:i+1]]
		result += lastSymbol
		i--
	}

	return result
}
