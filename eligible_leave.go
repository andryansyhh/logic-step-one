package main

import (
	"fmt"
	"time"
)

func checkLeaveEligibility() (bool, string) {
	var holidayCount int
	var joinDateStr, leaveDateStr string
	var leaveDuration int

	fmt.Print("Masukkan jumlah cuti bersama: ")
	fmt.Scanln(&holidayCount)

	fmt.Print("Masukkan tanggal join karyawan (format: yyyy-mm-dd): ")
	fmt.Scanln(&joinDateStr)

	fmt.Print("Masukkan tanggal rencana cuti (format: yyyy-mm-dd): ")
	fmt.Scanln(&leaveDateStr)

	fmt.Print("Masukkan durasi cuti (hari): ")
	fmt.Scanln(&leaveDuration)

	joinDate, _ := time.Parse("2006-01-02", joinDateStr)
	leaveDate, _ := time.Parse("2006-01-02", leaveDateStr)

	// Hitung jumlah hari dari tanggal join karyawan sampai dengan akhir tahun
	firstYearDays := int((time.Date(joinDate.Year(), 12, 31, 0, 0, 0, 0, time.UTC).Sub(joinDate)).Hours() / 24)

	// Hitung jumlah cuti pribadi untuk karyawan baru di tahun pertama
	firstYearLeave := int((float64(firstYearDays) / 365) * float64(holidayCount))

	// Tentukan apakah karyawan baru masih dalam periode 180 hari pertama
	if leaveDate.Sub(joinDate).Hours()/24 < 180 {
		return false, "Karena belum 180 hari sejak tanggal join karyawan"
	}

	// Tentukan jumlah cuti pribadi yang bisa diambil
	availableLeave := firstYearLeave
	if leaveDuration > availableLeave {
		return false, "Karena hanya boleh mengambil " + fmt.Sprintf("%d", availableLeave) + " hari cuti"
	}

	// Jika durasi cuti lebih dari 3 hari, cek apakah berturut-turut
	if leaveDuration > 3 {
		return false, "Karena cuti pribadi maksimal 3 hari berturut-turut"
	}

	return true, ""
}
