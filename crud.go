package main

import "fmt"

type Perangkat struct {
	nama     string
	watt     float64
	jamPakai float64 
	ruangan  string
}

var daftarPerangkat [100]Perangkat
var jumlahPerangkat int = 0

func tambahPerangkat(ruangan, nama  string, watt, jamPakai float64) {
	if jumlahPerangkat >= 100 {
		fmt.Println("[ERROR] Data penuh! Maksimal 100 perangkat.")
		return
	}
	daftarPerangkat[jumlahPerangkat].ruangan = ruangan
	daftarPerangkat[jumlahPerangkat].nama = nama
	daftarPerangkat[jumlahPerangkat].watt = watt
	daftarPerangkat[jumlahPerangkat].jamPakai = jamPakai
	jumlahPerangkat++
	fmt.Println("Perangkat berhasil ditambahkan!")
}

func tampilSemuaPerangkat() {
	if jumlahPerangkat == 0 {
		fmt.Println("Belum ada data perangkat yang tersimpan.")
		return
	}
	fmt.Println("\n============================================================")
	fmt.Printf("%-4s | %-12s | %-20s | %-8s | %-8s | %-10s\n",
		"No", "Ruangan", "Nama Perangkat", "Watt", "Jam/Hari", "Konsumsi(Wh)")
		fmt.Println("============================================================")
	for i := 0; i < jumlahPerangkat; i++ {
		konsumsi := daftarPerangkat[i].watt * daftarPerangkat[i].jamPakai
		fmt.Printf("%-4d | %-12s | %-20s | %-8.1f | %-8.1f | %-10.1f\n",
			i+1,
			daftarPerangkat[i].ruangan,
			daftarPerangkat[i].nama,
			daftarPerangkat[i].watt,
			daftarPerangkat[i].jamPakai,
			konsumsi)
	}
	fmt.Println("============================================================")
}

func ubahPerangkat(index int,ruangan, nama string, watt, jamPakai float64) {
	if index < 0 || index >= jumlahPerangkat {
		fmt.Println("[ERROR] Nomor perangkat tidak valid!")
		return
	}
	daftarPerangkat[index].ruangan = ruangan
	daftarPerangkat[index].nama = nama
	daftarPerangkat[index].watt = watt
	daftarPerangkat[index].jamPakai = jamPakai
	fmt.Println("Data perangkat berhasil diubah!")
}

func hapusPerangkat(index int) {
	if index < 0 || index >= jumlahPerangkat {
		fmt.Println("[ERROR] Nomor perangkat tidak valid!")
		return
	}
	namaHapus := daftarPerangkat[index].nama
	for i := index; i < jumlahPerangkat-1; i++ {
		daftarPerangkat[i] = daftarPerangkat[i+1]
	}
	jumlahPerangkat--
	fmt.Println("Perangkat", namaHapus, "berhasil dihapus!")
}

func sequentialSearchNama(namacari string) {
	ketemu := false
	fmt.Println("\n=== Hasil Pencarian (Sequential Search) ===")
	for i := 0; i < jumlahPerangkat; i++ {
		if daftarPerangkat[i].nama == namacari {
			fmt.Printf("Ditemukan di No %d:\n", i+1)
			fmt.Println("  Ruangan:", daftarPerangkat[i].ruangan)
			fmt.Println("  Nama   :", daftarPerangkat[i].nama)
			fmt.Println("  Watt   :", daftarPerangkat[i].watt, "W")
			fmt.Println("  Jam    :", daftarPerangkat[i].jamPakai, "jam/hari")			
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Perangkat tidak ditemukan.")
	}
}

func sequentialSearchRuangan(ruangancari string) {
	ketemu := false
	fmt.Println("\n=== Perangkat di Ruangan:", ruangancari, "===")
	for i := 0; i < jumlahPerangkat; i++ {
		if daftarPerangkat[i].ruangan == ruangancari {
			fmt.Printf("No %d: %s (%.1f W, %.1f jam)\n",
				i+1,
				daftarPerangkat[i].nama,
				daftarPerangkat[i].watt,
				daftarPerangkat[i].jamPakai)
				ketemu = true
		}
	}
	if !ketemu{
		fmt.Println("Tidak ada perangkat di ruangan tersebut.")
	}
}

func binarySearchNama(namacari string) {
	if jumlahPerangkat == 0 {
		fmt.Println("Data kosong!")
		return
	}

	var salinan [100]Perangkat
	for i := 0; i < jumlahPerangkat; i++ {
		salinan[i] = daftarPerangkat[i]
	}

	for i := 1; i < jumlahPerangkat; i++ {
		kunci := salinan[i]
		j := i - 1
		for j >= 0 && salinan[j].nama > kunci.nama {
			salinan[j+1] = salinan[j]
			j--
		}
		salinan[j+1] = kunci
	}

	kiri := 0
	kanan := jumlahPerangkat - 1
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
			fmt.Println("  Jam    :", salinan[tengah].jamPakai, "jam/hari")		
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
func selectionSortKonsumsi() {
	if jumlahPerangkat == 0 {
		fmt.Println("Data kosong!")
		return
	}
	for i := 0; i < jumlahPerangkat-1; i++ {
		indexMax := i
		for j := i + 1; j < jumlahPerangkat; j++ {
			konsumsiJ := daftarPerangkat[j].watt * daftarPerangkat[j].jamPakai
			konsumsiMax := daftarPerangkat[indexMax].watt * daftarPerangkat[indexMax].jamPakai
			if konsumsiJ > konsumsiMax {
				indexMax = j
			}
		}
		if indexMax != i {
			temp := daftarPerangkat[i]
			daftarPerangkat[i] = daftarPerangkat[indexMax]
			daftarPerangkat[indexMax] = temp
		}
	}
	fmt.Println("Urutan data berdasarkan konsumsi energi tertinggi ke terendah")
	tampilSemuaPerangkat()
}

func insertionSortNama() {
	if jumlahPerangkat == 0 {
		fmt.Println("Data kosong!")
		return
	}
	for i := 1; i < jumlahPerangkat; i++ {
		kunci := daftarPerangkat[i]
		j := i - 1
		for j >= 0 && daftarPerangkat[j].nama > kunci.nama {
			daftarPerangkat[j+1] = daftarPerangkat[j]
			j--
		}
		daftarPerangkat[j+1] = kunci
	}
	fmt.Println("Urutan data berdasarkan nama (A-Z)!")
	tampilSemuaPerangkat()
}

func tampilStatistik() {
	if jumlahPerangkat == 0 {
		fmt.Println("Belum ada data perangkat.")
		return
	}

	totalWh := 0.0
	indexTerborosWatt := 0
	indexTerborosKonsumsi := 0

	for i := 0; i < jumlahPerangkat; i++ {
		konsumsi := daftarPerangkat[i].watt * daftarPerangkat[i].jamPakai
		totalWh = totalWh + konsumsi

		if daftarPerangkat[i].watt > daftarPerangkat[indexTerborosWatt].watt {
			indexTerborosWatt = i
		}

		konsumsiMax := daftarPerangkat[indexTerborosKonsumsi].watt * daftarPerangkat[indexTerborosKonsumsi].jamPakai
		if konsumsi > konsumsiMax {
			indexTerborosKonsumsi = i
		}
	}

	totalKwh := totalWh / 1000.0

	fmt.Println("\n============================================")
	fmt.Println("        STATISTIK KONSUMSI LISTRIK")
	fmt.Println("============================================")
	fmt.Println("Total perangkat         :", jumlahPerangkat, "unit")
	fmt.Printf("Total konsumsi harian   : %.2f Wh (%.4f kWh)\n", totalWh, totalKwh)
	fmt.Println("--------------------------------------------")
	fmt.Println("Perangkat paling boros (watt tertinggi):")
	fmt.Printf("  -> %s (%.1f W)\n", daftarPerangkat[indexTerborosWatt].nama, daftarPerangkat[indexTerborosWatt].watt)
	fmt.Println("Perangkat paling boros (konsumsi harian):")
	konsumsiTerboros := daftarPerangkat[indexTerborosKonsumsi].watt * daftarPerangkat[indexTerborosKonsumsi].jamPakai
	fmt.Printf("  -> %s (%.1f Wh/hari)\n", daftarPerangkat[indexTerborosKonsumsi].nama, konsumsiTerboros)
	fmt.Println("============================================")
}

func tampilMenu() {
	fmt.Println("\n+------------------------------------------+")
	fmt.Println("|   POWERLOG - Pencatatan Konsumsi Listrik |")
	fmt.Println("+------------------------------------------+")
	fmt.Println("1. Tambah Perangkat                      ")
	fmt.Println("2. Tampilkan Semua Perangkat             ")
	fmt.Println("3. Ubah Data Perangkat                   ")
	fmt.Println("4. Hapus Perangkat                       ")
	fmt.Println("5. Cari Perangkat                        ")
	fmt.Println("6. Urutkan Data                          ")
	fmt.Println("7. Statistik Konsumsi                    ")
	fmt.Println("0. Keluar                                ")
	fmt.Print("Pilih menu: ")
}

func menuCari() {
	fmt.Println("\n--- Menu Pencarian ---")
	fmt.Println("1. Cari berdasarkan Nama (Sequential Search)")
	fmt.Println("2. Cari berdasarkan Ruangan (Sequential Search)")
	fmt.Println("3. Cari berdasarkan Nama (Binary Search)")
	fmt.Print("Pilih: ")
}

func menuUrut() {
	fmt.Println("\n--- Menu Pengurutan ---")
	fmt.Println("1. Urutkan berdasarkan Konsumsi Energi Tertinggi (Selection Sort)")
	fmt.Println("2. Urutkan berdasarkan Nama A-Z (Insertion Sort)")
	fmt.Print("Pilih: ")
}

func main() {
	var pilihan int

	fmt.Println("==========================================")
	fmt.Println("  Selamat datang di aplikasi POWERLOG!")
	fmt.Println("==========================================")

	for {
		tampilMenu()
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			var nama, ruangan string
			var watt, jam float64
			fmt.Print("Nama ruangan    : ")
			fmt.Scan(&ruangan)
			fmt.Print("Nama perangkat  : ")
			fmt.Scan(&nama)
			fmt.Print("Daya (watt)     : ")
			fmt.Scan(&watt)
			fmt.Print("Jam pakai/hari  : ")
			fmt.Scan(&jam)
			tambahPerangkat(ruangan, nama, watt, jam)

		} else if pilihan == 2 {
			tampilSemuaPerangkat()

		} else if pilihan == 3 {
			tampilSemuaPerangkat()
			if jumlahPerangkat > 0 {
				var no int
				var nama, ruangan string
				var watt, jam float64
				fmt.Print("Masukkan nomor yang mau diubah: ")
				fmt.Scan(&no)
				fmt.Print("Ruangan baru    : ")
				fmt.Scan(&ruangan)
				fmt.Print("Nama baru       : ")
				fmt.Scan(&nama)
				fmt.Print("Daya baru (watt): ")
				fmt.Scan(&watt)
				fmt.Print("Jam baru/hari   : ")
				fmt.Scan(&jam)
				ubahPerangkat(no-1, ruangan, nama, watt, jam)
			}

		} else if pilihan == 4 {
			tampilSemuaPerangkat()
			if jumlahPerangkat > 0 {
				var no int
				fmt.Print("Masukkan nomor yang mau dihapus: ")
				fmt.Scan(&no)
				hapusPerangkat(no - 1)
			}

		} else if pilihan == 5 {
			menuCari()
			var subPilihan int
			fmt.Scan(&subPilihan)
			if subPilihan == 1 {
				var nama string
				fmt.Print("Nama yang dicari: ")
				fmt.Scan(&nama)
				sequentialSearchNama(nama)
			} else if subPilihan == 2 {
				var ruangan string
				fmt.Print("Nama ruangan: ")
				fmt.Scan(&ruangan)
				sequentialSearchRuangan(ruangan)
			} else if subPilihan == 3 {
				var nama string
				fmt.Print("Nama yang dicari (harus persis): ")
				fmt.Scan(&nama)
				binarySearchNama(nama)
			} else {
				fmt.Println("Pilihan tidak valid")
			}

		} else if pilihan == 6 {
			menuUrut()
			var subPilihan int
			fmt.Scan(&subPilihan)
			if subPilihan == 1 {
				selectionSortKonsumsi()
			} else if subPilihan == 2 {
				insertionSortNama()
			} else {
				fmt.Println("Pilihan tidak valid")
			}

		} else if pilihan == 7 {
			tampilStatistik()

		} else if pilihan == 0 {
			fmt.Println("\nTerima kasih sudah menggunakan PowerLog!")
			return

		} else {
			fmt.Println("[ERROR] Pilihan tidak valid, silakan coba lagi.")
		}
	}
}