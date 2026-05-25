package main

import "fmt"

func tambahKendaraan(daftar *DaftarKendaraan, n *int) {
	if *n >= NMAX {
		fmt.Println("Memori aplikasi sudah penuh!")
		return
	}

	fmt.Println("-- INPUT DATA KENDARAAN BARU --")

	fmt.Print("Masukkan Plat Nomor: ")
	fmt.Scan(&daftar[*n].PlatNomor)

	fmt.Print("Masukkan Nama Pemilik: ")
	fmt.Scan(&daftar[*n].NamaPemilik)

	fmt.Print("Masukkan Jenis Kendaraan (Mobil/Truk/Motor): ")
	fmt.Scan(&daftar[*n].JenisKendaraan)

	*n++
	fmt.Println("[INFO] Data kendaraan berhasil ditambahkan!")
}

func showKendaraan(daftar DaftarKendaraan, n int) {
	if n == 0 {
		fmt.Println("Belum ada data kendaraan yang terdaftar.")
		return
	}

	fmt.Println("-- DAFTAR KENDARAAN --")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. Plat: %s | Pemilik: %s | Jenis Kendaraan: %s\n", i+1, daftar[i].PlatNomor, daftar[i].NamaPemilik, daftar[i].JenisKendaraan)
	}
}

// sequential search, cek plat yang terdaftar
func cekPlat(daftar DaftarKendaraan, n int, plat string) bool {
	for i := 0; i < n; i++ {
		if daftar[i].PlatNomor == plat {
			return true
		}
	}
	return false
}

// catat transaksi servis
func tambahServis(daftarS *DaftarServis, nS *int, daftarK DaftarKendaraan, nK int) {
	if *nS >= NMAX {
		fmt.Println("Memori penuh!")
		return
	}

	fmt.Println("CATAT RIWAYAT SERVIS")

	var inputPlat string

	fmt.Print("Plat nomor kendaraan: ")
	fmt.Scan(&inputPlat)

	if !cekPlat(daftarK, nK, inputPlat) {
		fmt.Println("Plat belum tersedia, silakan daftar terlebih dahulu di men pertama.")
		return
	}
	daftarS[*nS].PlatNomor = inputPlat

	daftarS[*nS].IDServis = *nS + 1
	fmt.Printf("ID servis: %d", daftarS[*nS].IDServis)

	fmt.Print("Tanggal (DD-MM-YYYY): ")
	fmt.Scan(&daftarS[*nS].Tanggal)

	fmt.Print("Jenis tindakan: ")
	fmt.Scan(&daftarS[*nS].JenisTindakan)

	fmt.Print("Biaya: ")
	fmt.Scan(&daftarS[*nS].Biaya)
	*nS++

	fmt.Print("Data berhasil dicatat.")
}

func showServis(daftar DaftarServis, n int) {
	if n == 0 {
		fmt.Print("Belum ada riwayat servis yang tercatat.")
		return
	}

	fmt.Println("RIWAYAT SERVIS")

	for i := 0; i < n; i++ {
		fmt.Printf("%d. ID servis: %d | Tanggal: %s | Jenis tindakan: %s | Biaya: %d\n", i+1, daftar[i].IDServis, daftar[i].Tanggal, daftar[i].JenisTindakan, daftar[i].Biaya)
	}
}
