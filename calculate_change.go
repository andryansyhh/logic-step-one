package main

import (
	"fmt"
	"math"
)

func CalculateChange() (string, bool) {
	var totalBelanja, pembayaran float64

	fmt.Print("Total belanja seorang customer: Rp ")
	fmt.Scanln(&totalBelanja)

	fmt.Print("Pembeli membayar: Rp ")
	fmt.Scanln(&pembayaran)

	kembalian := pembayaran - totalBelanja

	// Jika uang yang dibayarkan kurang dari total belanja, kembalikan false
	if kembalian < 0 {
		return "False, kurang bayar", false
	}

	kembalianBulat := math.Floor(kembalian*100) / 100

	// Format kembalian dan pecahan uang
	result := fmt.Sprintf("Kembalian yang harus diberikan kasir: %.3f,\ndibulatkan menjadi %.3f\n", pembayaran-totalBelanja, kembalianBulat)
	result += "Pecahan uang:\n"

	// Pecahan uang yang tersedia
	pecahan := []float64{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}

	pecahanUang := make(map[float64]int)

	for _, nilai := range pecahan {
		jumlah := int(kembalianBulat / nilai)
		if jumlah > 0 {
			pecahanUang[nilai] = jumlah
			kembalianBulat -= float64(jumlah) * nilai
		}
	}

	for nilai, jumlah := range pecahanUang {
		if jumlah > 0 {
			if nilai >= 1000 {
				result += fmt.Sprintf("%d lembar %.0f\n", jumlah, nilai)
			} else {
				result += fmt.Sprintf("%d koin %.0f\n", jumlah, nilai)
			}
		}
	}

	return result, true
}
