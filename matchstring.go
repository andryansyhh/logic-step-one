package main

import (
	"fmt"
	s "strings"
)

func MatchStrings() string {
	var N int
	fmt.Println("Masukkan jumlah string:")
	fmt.Scanln(&N)

	// Meminta input string dari pengguna
	fmt.Println("Masukkan string:")
	strings := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&strings[i])
	}

	matched := make(map[string]bool)
	var result []int

	for i := 0; i < N; i++ {
		if matched[strings[i]] {
			continue
		}
		// loop untuk membandingkan dengan string lain
		for j := i + 1; j < N; j++ {
			if s.EqualFold(strings[i], strings[j]) {
				result = append(result, i+1)
				result = append(result, j+1)
				matched[strings[i]] = true
				matched[strings[j]] = true
				break
			}
		}
		if len(result) > 0 {
			break
		}
	}

	if len(result) == 0 {
		return "false"
	}

	// format hasil menjadi string
	output := ""
	for i := 0; i < len(result); i++ {
		output += fmt.Sprintf("%d", result[i])
		if i < len(result)-1 {
			output += " "
		}
	}

	return output
}
