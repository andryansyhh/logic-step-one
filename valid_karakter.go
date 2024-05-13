package main

import (
	"fmt"
	"strings"
)

func IsValid() bool {

	var input string
	fmt.Println("Masukkan string untuk divalidasi:")
	fmt.Scanln(&input)

	// Pemetaan untuk karakter pembuka dan penutup
	mapping := map[rune]rune{
		'{': '}',
		'[': ']',
		'<': '>',
	}

	stack := []rune{} // stack untuk melacak karakter pembuka

	for _, char := range input {
		// Jika itu karakter pembuka, tambahkan ke stack
		if strings.ContainsAny(string(char), "{[<") {
			stack = append(stack, char)
		} else if strings.ContainsAny(string(char), "}>]") {
			// Jika itu karakter penutup
			// Pastikan stack tidak kosong dan karakter penutup cocok dengan karakter pembuka teratas di stack
			if len(stack) == 0 || mapping[stack[len(stack)-1]] != char {
				return false
			}
			// Pop karakter pembuka yang sesuai dari stack
			stack = stack[:len(stack)-1]
		}
	}

	// Pastikan stack kosong setelah semua karakter diproses
	return len(stack) == 0
}
