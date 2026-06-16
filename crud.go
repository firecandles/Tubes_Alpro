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

func sequentialsearchnama(namacari string) {
	ketemu := false
	fmt.Println("\n=== Hasil Pencarian (Sequential Search) ===")
	for i := 0; i < jumlahperangkat; i++ {
		if daftarperangkat[i].nama == namacari {
			fmt.Printf("Ditemukan di No %d:\n", i+1)
			fmt.Println("  Ruangan:", daftarperangkat[i].ruangan)
			fmt.Println("  Nama   :", daftarperangkat[i].nama)
			fmt.Println("  Watt   :", daftarperangkat[i].watt, "W")
			fmt.Println("  Jam    :", daftarperangkat[i].jampakai, "jam/hari")			
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Perangkat tidak ditemukan.")
	}
}

func sequientialsearchruangan(ruangancari string) {
	ketemu := false
	fmt.Println("\n=== Perangkat di Ruangan:", ruangancari, "===")
	for i := 0; i < jumlahperangkat; i++ {
		if daftarperangkat[i].ruangan == ruangancari {
			fmt.Printf("No %d: %s (%.1f W, %.1f jam)\n",
				i+1,
				daftarperangkat[i].nama,
				daftarperangkat[i].watt,
				daftarperangkat[i].jampakai)
				ketemu = true
		}
	}
	if !ketemu{
		fmt.Println("Tidak ada perangkat di ruangan tersebut.")
	}
}

func binarysearchnama(namacari string) {
	if jumlahperangkat == 0 {
		fmt.Println("Data kosong!")
		return
	}

	var salinan [100]perangkat
	for i := 0; i < jumlahperangkat; i++ {
		salinan[i] = daftarperangkat[i]
	}

	for i := 1; i < jumlahperangkat; i++ {
		kunci := salinan[i]
		j := i - 1
		for j >= 0 && salinan[j].nama > kunci.nama {
			salinan[j+1] = salinan[j]
			j--
		}
		salinan[j+1] = kunci
	}

	kiri := 0
	kanan := jumlahperangkat - 1
	ketemu := false 

	fmt.Println("\n=== Hasil Binary Search ===")
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		namatengah := salinan[tengah].nama
		namacaritarget := namacari

		if namatengah == namacaritarget {
			fmt.Println("Perangkat ditemukan!")
			fmt.Println("  Ruangan:", salinan[tengah].ruangan)
			fmt.Println("  Nama   :", salinan[tengah].nama)
			fmt.Println("  Watt   :", salinan[tengah].watt, "W")
			fmt.Println("  Jam    :", salinan[tengah].jampakai, "jam/hari")		
			ketemu = true
			break
		} else if namacaritarget < namatengah {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if !ketemu {
		fmt.Println("Perangkat tidak ditemukan.")
	}
}
