package main

type perangkat struct {
	nama, ruangan  string
	watt, jampakai float64
}

var daftarperangkat [100]perangkat
var jumlahperangkat int = 0
