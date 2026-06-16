package main

import "fmt"

func tambahperangkat(ruangan, nama  string, watt, jampakai float64) {
	if jumlahperangkat >= 100 {
		fmt.Println("[ERROR] Data penuh! Maksimal 100 perangkat.")
		return
	}
	daftarperangkat[jumlahperangkat].ruangan = ruangan
	daftarperangkat[jumlahperangkat].nama = nama
	daftarperangkat[jumlahperangkat].watt = watt
	daftarperangkat[jumlahperangkat].jampakai = jampakai
	jumlahperangkat++
	fmt.Println("[OK] Perangkat berhasil ditambahkan!")
}

func tampilsemuaperangkat() {
	if jumlahperangkat == 0 {
		fmt.Println("Belum ada data perangkat yang tersimpan.")
		return
	}
	fmt.Println("\n============================================================")
	fmt.Printf("%-4s | %-20s | %-8s | %-8s | %-12s | %-10s\n",
		"No", "Nama Perangkat", "Watt", "Jam/Hari", "Ruangan", "Konsumsi(Wh)")
}