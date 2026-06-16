package main

import "fmt"

type Perangkat struct {
	nama     string
	watt     float64
	jamPakai float64 
	ruangan  string
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
	fmt.Println("[OK] Data berhasil diurutkan berdasarkan konsumsi energi (tertinggi ke terendah)!")
	tampilSemuaPerangkat()
}