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
		fmt.Println("============================================================")
	for i := 0; i < jumlahperangkat; i++ {
		konsumsi := daftarperangkat[i].watt * daftarperangkat[i].jampakai
		fmt.Printf("%-4d | %-20s | %-8.1s | %-8.1f | %-12f | %-10.1f\n",
			i+1,
			daftarperangkat[i].ruangan,
			daftarperangkat[i].nama,
			daftarperangkat[i].watt,
			daftarperangkat[i].jampakai,
			konsumsi)
	}
	fmt.Println("============================================================")
}

func ubahperangkat(index int, nama string, watt float64, jamPakai float64, ruangan string) {
	if index < 0 || index >= jumlahperangkat {
		fmt.Println("[ERROR] Nomor perangkat tidak valid!")
		return
	}
	daftarperangkat[index].ruangan = ruangan
	daftarperangkat[index].nama = nama
	daftarperangkat[index].watt = watt
	daftarperangkat[index].jampakai = jamPakai
	fmt.Println("[OK] Data perangkat berhasil diubah!")
}

func hapusPerangkat(index int) {
	if index < 0 || index >= jumlahperangkat {
		fmt.Println("[ERROR] Nomor perangkat tidak valid!")
		return
	}
	namaHapus := daftarperangkat[index].nama
	for i := index; i < jumlahperangkat-1; i++ {
		daftarperangkat[i] = daftarperangkat[i+1]
	}
	jumlahperangkat--
	fmt.Println("[OK] Perangkat", namaHapus, "berhasil dihapus!")
}
