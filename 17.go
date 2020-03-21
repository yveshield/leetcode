package main

func letterCombinations(digits string) []string {
	letters := make(map[byte][]string)
	letters['2'] = []string{"a", "b", "c"}
	letters['3'] = []string{"d", "e", "f"}
	letters['4'] = []string{"g", "h", "i"}
	letters['5'] = []string{"j", "k", "l"}
	letters['6'] = []string{"m", "n", "o"}
	letters['7'] = []string{"p", "q", "r", "s"}
	letters['8'] = []string{"t", "u", "v"}
	letters['9'] = []string{"w", "x", "y", "z"}

	res := []string{}

	if digits == "" {
		return res
	}

	res = append(res, "")

	for i := 0; i < len(digits); i++ {
		nums := letters[digits[i]]
		pre := res
		res = []string{}
		for _, num := range nums {
			for _, v := range pre {
				res = append(res, v+num)
			}
		}
	}

	return res
}
