package main

import "fmt"

func tambahKendaraan(daftarK *DaftarKendaraan, nK *int, daftarP DaftarPemilik, nP int) {
	if *nK >= NMAX {
		fmt.Println("Memori aplikasi sudah penuh!")
		return
	}

	fmt.Println("INPUT DATA KENDARAAN BARU")

	var inputID string
	fmt.Print("Masukkan ID pemilik kendaraan: ")
	fmt.Scan(&inputID)

	indeksPemilik := cariPemilik(daftarP, nP, inputID)
	if indeksPemilik == -1 {
		fmt.Println("ID pemilik belum terdaftar!")
		return
	}
	daftarK[*nK].IDPemilik = inputID
	fmt.Print("Masukkan Plat Nomor: ")
	fmt.Scan(&daftarK[*nK].PlatNomor)

	fmt.Print("Masukkan Jenis Kendaraan (Mobil/Motor): ")
	fmt.Scan(&daftarK[*nK].JenisKendaraan)

	fmt.Print("Masukkan Tahun Produksi: ")
	fmt.Scan(&daftarK[*nK].TahunProduksi)

	*nK++
	fmt.Println("Data kendaraan berhasil ditambahkan!")
}

func showKendaraan(daftarK DaftarKendaraan, nK int, daftarP DaftarPemilik, nP int) {
	if nK == 0 {
		fmt.Println("Belum ada data kendaraan yang terdaftar.")
		return
	}

	fmt.Println("-- DAFTAR KENDARAAN --")
	for i := 0; i < nK; i++ {
		indeksPemilik := cariPemilik(daftarP, nP, daftarK[i].IDPemilik)
		fmt.Printf("Data ke-%d:\n", i+1)
		fmt.Printf("Plat Nomor: %s\n", daftarK[i].PlatNomor)
		fmt.Printf("Jenis Kendaraan: %s\n", daftarK[i].JenisKendaraan)
		fmt.Printf("Tahun Produksi: %d\n", daftarK[i].TahunProduksi)

		if indeksPemilik != -1 {
			fmt.Printf("Nama pemilik: %s (ID: %s)\n", daftarP[indeksPemilik].Nama, daftarP[indeksPemilik].IDPemilik)
		} else {
			fmt.Println("Data tidak diketahui/ditemukan")
		}
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
func tambahServis(daftarS *DaftarServis, nS *int, daftarK DaftarKendaraan, nK int, daftarKategori DaftarKategori, nKategori int) {
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

	fmt.Println("\nDAFTAR KATEGORI")
	showKategori(daftarKategori, nKategori)
	var kode string
	fmt.Print("Pilih kode kategori: ")
	fmt.Scan(&kode)

	if cariKategori(daftarKategori, nKategori, kode) == -1 {
		fmt.Println("Kode kategori tidak ditemukan!")
		return
	}
	daftarS[*nS].kodeKategori = kode

	daftarS[*nS].IDServis = *nS + 1
	fmt.Println("ID servis: ", daftarS[*nS].IDServis)

	fmt.Println("Masukkan Tanggal Servis")

	fmt.Print("Tanggal (DD): ")
	fmt.Scan(&daftarS[*nS].Tanggal.Hari)

	fmt.Print("Bulan (MM): ")
	fmt.Scan(&daftarS[*nS].Tanggal.Bulan)

	fmt.Print("Tahun (YYYY): ")
	fmt.Scan(&daftarS[*nS].Tanggal.Tahun)

	fmt.Print("Jenis tindakan: ")
	fmt.Scan(&daftarS[*nS].JenisTindakan)

	fmt.Print("Biaya: ")
	fmt.Scan(&daftarS[*nS].Biaya)
	*nS++

	fmt.Print("Data berhasil dicatat.")
}

func showServis(daftar DaftarServis, n int, daftarKategori DaftarKategori, nKategori int) {
	if n == 0 {
		fmt.Print("Belum ada riwayat servis yang tercatat.")
		return
	}

	fmt.Println("RIWAYAT SERVIS")

	indeks := cariKategori(daftarKategori, nKategori, daftar[0].JenisTindakan)
	for i := 0; i < n; i++ {
		fmt.Printf("Nota ke-%d: \n", i+1)
		fmt.Printf("ID Servis: %d\n", daftar[i].IDServis)
		fmt.Printf("Plat Nomor: %s\n", daftar[i].PlatNomor)
		if indeks != -1 {
			fmt.Printf("kode kategori: %s\n", daftarKategori[indeks].kodeKategori)
		}
		fmt.Printf("Tanggal: %02d-%02d-%d\n", daftar[i].Tanggal.Hari, daftar[i].Tanggal.Bulan, daftar[i].Tanggal.Tahun)
		fmt.Printf("Tindakan: %s\n", daftar[i].JenisTindakan)
		fmt.Printf("Biaya: Rp. %d\n", daftar[i].Biaya)
	}
}

// sequential search, cek ID pemilik yang terdaftar
func cariPemilik(daftar DaftarPemilik, n int, id string) int {
	for i := 0; i < n; i++ {
		if daftar[i].IDPemilik == id {
			return i
		}
	}
	return -1
}

func tambahPemilik(daftar *DaftarPemilik, n *int) {
	if *n >= NMAX {
		fmt.Println("Data pemilik sudah penuh!")
		return
	}
	fmt.Println("INPUT DATA PEMILIK")

	autoID := fmt.Sprintf("%03d", *n+1)
	daftar[*n].IDPemilik = autoID
	fmt.Println("ID pemilik: ", autoID)

	fmt.Print("Nama pemilik: ")
	fmt.Scan(&daftar[*n].Nama)

	fmt.Print("Alamat: ")
	fmt.Scan(&daftar[*n].Alamat)

	fmt.Print("Nomor telepon: ")
	fmt.Scan(&daftar[*n].NoTelp)

	*n++
	fmt.Println("Data pemilik berhasil ditambahkan.")
}

func showPemilik(daftar DaftarPemilik, n int) {
	if n == 0 {
		fmt.Println("Belum ada data pemilik yang terdaftar.")
		return
	}

	fmt.Println("-- DAFTAR PEMILIK --")
	for i := 0; i < n; i++ {
		fmt.Printf("Pemilik ke-%d: \n", i+1)
		fmt.Printf("ID Pemilik: %s\n", daftar[i].IDPemilik)
		fmt.Printf("Nama: %s\n", daftar[i].Nama)
		fmt.Printf("Alamat: %s\n", daftar[i].Alamat)
		fmt.Printf("Nomor Telepon: %s\n", daftar[i].NoTelp)
	}
}

func editPemilik(daftar *DaftarPemilik, n int) {
	if n == 0 {
		fmt.Println("Belum ada data pemilik.")
		return
	}
	fmt.Println("\nEDIT DATA PEMILIK")

	var cariID string
	fmt.Print("Masukkan ID pemilik: ")
	fmt.Scan(&cariID)

	indeks := cariPemilik(*daftar, n, cariID)

	if indeks == -1 {
		fmt.Println("ID pemilik tidak ditemukan.")
		return
	}

	fmt.Println("\nData Lama Ditemukan")
	fmt.Printf("Nama: %s\n", daftar[indeks].Nama)
	fmt.Printf("Alamat: %s\n", daftar[indeks].Alamat)
	fmt.Printf("Nomor Telp: %s\n", daftar[indeks].NoTelp)

	var namaBaru, alamatBaru, noTelpBaru string
	fmt.Println("Masukkan Data Baru")

	fmt.Print("Nama: ")
	fmt.Scan(&namaBaru)

	fmt.Print("Alamat: ")
	fmt.Scan(&alamatBaru)

	fmt.Print("No. Telp: ")
	fmt.Scan(&noTelpBaru)

	var konfirmasi string
	fmt.Print("\nApakah yakin ingin menyimpan perubahan? (y/n) ")
	fmt.Scan(&konfirmasi)

	if konfirmasi == "y" || konfirmasi == "Y" {
		daftar[indeks].Nama = namaBaru
		daftar[indeks].Alamat = alamatBaru
		daftar[indeks].NoTelp = noTelpBaru
		fmt.Println("Data pemilik berhasil diperbarui.")
	} else {
		fmt.Println("Perubahan dibatalkan.")
	}
}

func editKendaraan(daftarK *DaftarKendaraan, nK int, daftarP DaftarPemilik, nP int) {
	if nK == 0 {
		fmt.Println("Belum ada data kendaraan yang bisa diubah.")
		return
	}

	fmt.Println("\nEDIT DATA KENDAAAN")
	var cariPlat string
	fmt.Print("Masukkan plat nomor kendaraan: ")
	fmt.Scan(&cariPlat)

	indeksK := -1

	for i := 0; i < nK; i++ {
		if daftarK[i].PlatNomor == cariPlat {
			indeksK = i
			break
		}
	}

	if indeksK == -1 {
		fmt.Println("Plat nomor tidak ditemukan")
		return
	}

	fmt.Println("Data lama ditemukan:")
	fmt.Printf("Plat Nomor: %s\n", daftarK[indeksK].PlatNomor)
	fmt.Printf("Jenis Tindakan: %s\n", daftarK[indeksK].JenisKendaraan)
	fmt.Printf("Tahun Produksi: %d\n", daftarK[indeksK].TahunProduksi)
	fmt.Printf("ID Pemilik: %s\n", daftarK[indeksK].IDPemilik)

	var jenisBaru, idBaru string
	var tahunBaru int

	fmt.Println("Masukkan Data Baru")

	fmt.Print("Jenis Kendaraan (Mobil/Truk/Motor): ")
	fmt.Scan(&jenisBaru)

	fmt.Print("Tahun Produksi: ")
	fmt.Scan(&tahunBaru)

	for {
		fmt.Print("Masukkan ID pemilik baru untuk kendaraan ini: ")
		fmt.Scan(&idBaru)

		if cariPemilik(daftarP, nP, idBaru) != -1 {
			break
		}
		fmt.Println("ID Pemilik tidak terdaftar!")
	}

	var konfirmasi string
	fmt.Print("\nApakah yakin ingin menyimpan perubahan? (y/n): ")
	fmt.Scan(&konfirmasi)

	if konfirmasi == "y" || konfirmasi == "Y" {
		daftarK[indeksK].JenisKendaraan = jenisBaru
		daftarK[indeksK].TahunProduksi = tahunBaru
		daftarK[indeksK].IDPemilik = idBaru
		fmt.Println("Data kendaraan berhasil diperbarui.")
	} else {
		fmt.Println("Perubahan dibatalkan.")
	}
}

// sequential search cari kendaraan berdasar plat
func cariKendaraan(daftarK DaftarKendaraan, nK int, findPlat string) int {
	for i := 0; i < nK; i++ {
		if daftarK[i].PlatNomor == findPlat {
			return i
		}
	}
	return -1
}

func detailCariKendaraan(daftarK DaftarKendaraan, nK int, daftarP DaftarPemilik, nP int) {
	if nK == 0 {
		fmt.Println("\nBelum ada data kendaraan di sistem.")
		return
	}

	var plat string
	fmt.Print("Masukkan Plat Nomor yang dicari: ")
	fmt.Scan(&plat)

	indeksK := cariKendaraan(daftarK, nK, plat)

	if indeksK == -1 {
		fmt.Println("Data kendaraan dengan plat tersebut tidak ditemukan.")
	} else {
		fmt.Printf("Plat nomor: %s\n", daftarK[indeksK].PlatNomor)
		fmt.Printf("Jenis kendaraan: %s\n", daftarK[indeksK].JenisKendaraan)
		fmt.Printf("Tahun produksi: %d\n", daftarK[indeksK].TahunProduksi)
	}

	indeksP := cariPemilik(daftarP, nP, daftarK[indeksK].IDPemilik)
	if indeksP != -1 {
		fmt.Printf("Nama pemilik: %s (ID: %s)\n", daftarP[indeksP].Nama, daftarP[indeksP].IDPemilik)
		fmt.Printf("No. telp: %s\n", daftarP[indeksP].NoTelp)
	}

}

// binary
func cariServis(daftarS DaftarServis, nS int, findID int) int {
	low := 0
	high := nS - 1

	for low <= high {
		mid := (low + high) / 2

		if daftarS[mid].IDServis == findID {
			return mid
		} else if daftarS[mid].IDServis < findID {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func detailCariServis(daftarS DaftarServis, nS int) {
	if nS == 0 {
		fmt.Println("\nBelum ada riwayat servis di sistem.")
		return
	}

	var findID int
	fmt.Print("Masukkan ID servis yang dicari: ")
	fmt.Scan(&findID)

	indeksS := cariServis(daftarS, nS, findID)

	if indeksS == -1 {
		fmt.Println("Nota servis dengan ID tersebut tidak ditemukan.")
	} else {
		fmt.Printf("ID Servis: %d\n", daftarS[indeksS].IDServis)
		fmt.Printf("Plat Nomor: %s\n", daftarS[indeksS].PlatNomor)
		fmt.Printf("Tanggal: %02d-%02d-%d\n", daftarS[indeksS].Tanggal.Hari, daftarS[indeksS].Tanggal.Bulan, daftarS[indeksS].Tanggal.Tahun)
		fmt.Printf("Tindakan: %s\n", daftarS[indeksS].JenisTindakan)
		fmt.Printf("Biaya: Rp %d\n", daftarS[indeksS].Biaya)
	}
}

func showStatistik(daftarS DaftarServis, nS int, daftarKategori DaftarKategori, nKategori int) {
	if nS == 0 {
		fmt.Println("Belum ada data riwayat servis.")
		return
	}

	var findMonth int
	fmt.Println("Masukkan bulan yang ingin dilihat (1-12): ")
	fmt.Scan(&findMonth)

	if findMonth < 1 || findMonth > 12 {
		fmt.Println("Input tidak valid!")
		return
	}

	var listKategori [NMAX]string
	var platTerhitung [NMAX]string

	nData := 0
	nPlat := 0
	totalPendapatan := 0
	jumlahKendaraan := 0

	for i := 0; i < nS; i++ {
		if daftarS[i].Tanggal.Bulan == findMonth {
			listKategori[nData] = daftarS[i].kodeKategori
			nData++

			totalPendapatan += daftarS[i].Biaya

			adaPlat := false
			for j := 0; j < nPlat; j++ {
				if platTerhitung[j] == daftarS[i].PlatNomor{
					adaPlat = true
				}
			}
			if !adaPlat {
				platTerhitung[nPlat] = daftarS[i].PlatNomor
				nPlat++
				jumlahKendaraan++
			}
		}
	}

	if nData == 0 {
		fmt.Printf("Tidak ada aktivitas servis pada bulan ke-%d.\n", findMonth)
		return
	}

	jumlahMaksimum := 0
	for i := 0; i < nData; i++ {
		hitung := 0
		for j := 0; j < nData; j++ {
			if listKategori[i] == listKategori[j] {
				hitung++
			}
		}
		if hitung > jumlahMaksimum {
			jumlahMaksimum = hitung
		}
	}

	fmt.Printf("LAPORAN STATISTIK BULAN %d\n", findMonth)
	fmt.Printf("Total Pendapatan: Rp %d\n", totalPendapatan)
	fmt.Printf("Total Unit Diservis: %d kendaraan\n", jumlahKendaraan)
	fmt.Printf("Frekuensi Kerusakan: Terbanyak muncul %d kali\n", jumlahMaksimum)
	fmt.Println("Kategori Kerusakan Paling Sering Keluar: ")

	for i := 0; i < nData; i++ {
		hitung := 0
		for j := 0; j < nData; j++ {
			if listKategori[i] == listKategori[j] {
				hitung++
			}
		}
		if hitung == jumlahMaksimum {
			sudahDicetak := false
			for k := 0; k < i; k++ {
				if listKategori[i] == listKategori[k] {
					sudahDicetak = true
					break
				}
			}
			if !sudahDicetak {
				idx := cariKategori(daftarKategori, nKategori, listKategori[i])
				if idx != -1 {
					fmt.Printf("%s\n (%d kali)\n", daftarKategori[idx].namaKerusakan, hitung)
				}
			}
		}
	}
}

func hapusPemilik(daftar *DaftarPemilik, n *int, daftarK DaftarKendaraan, nK int) {
	if *n == 0 {
		fmt.Print("Belum ada data pemilik untuk dihapus.")
		return
	}

	fmt.Println("\nHAPUS DATA PEMILIK")
	var findID string
	fmt.Print("Masukkan ID pemilik yang ingin dihapus: ")
	fmt.Scan(&findID)

	indeks := cariPemilik(*daftar, *n, findID)
	if indeks == -1 {
		fmt.Println("ID pemilik tidak ditemukan.")
		return
	}

	for i := 0; i < nK; i++ {
		if daftarK[i].IDPemilik == findID {
			fmt.Print("Gagal hapus, kendaraan masih terdaftar di sistem!")
			return
		}
	}
	var konfirmasi string
	fmt.Printf("Apakah Anda yakin ingin menghapus pemilik %s? (y/n): ", daftar[indeks].Nama)
	fmt.Scan(&konfirmasi)

	if konfirmasi == "y" || konfirmasi == "Y" {
		for i := indeks; i < *n-1; i++ {
			daftar[i] = daftar[i+1]
		}
		*n--
		fmt.Print("Data pemilih berhasil dihapus.")
	} else {
		fmt.Print("Hapus data dibatalkan.")
	}
}

func hapusKendaraan(daftarK *DaftarKendaraan, nK *int) {
	if *nK == 0 {
		fmt.Println("Belum ada data kendaraan untuk dihapus.")
		return
	}

	fmt.Println("\nHAPUS DATA KENDARAAN")
	var findPlat string
	fmt.Print("Masukkan plat nomor kendaraan yang ingin dihapus: ")
	fmt.Scan(&findPlat)

	indeksK := cariKendaraan(*daftarK, *nK, findPlat)
	if indeksK == -1 {
		fmt.Println("Plat nomor tidak ditemukan.")
		return
	}

	var konfirmasi string
	fmt.Printf("Apakah Anda yakin ingin menghapus kendaraan plat %s? (y/n): ", findPlat)
	fmt.Scan(&konfirmasi)

	if konfirmasi == "Y" || konfirmasi == "y" {
		for i := indeksK; i < *nK-1; i++ {
			daftarK[i] = daftarK[i+1]
		}
		*nK--
		fmt.Print("Data kendaraan berhasil dihapus")
	} else {
		fmt.Print("Hapus data dibatalkan.")
	}
}

func tambahKategori(daftar *DaftarKategori, n *int) {
	if *n >= NMAX {
		fmt.Println("Data kategori sudah penuh!")
		return
	}
	fmt.Println("INPUT DATA KATEGORI KERUSAKAN")

	autoID := fmt.Sprintf("%03d", *n+1)
	daftar[*n].kodeKategori = autoID
	fmt.Println("Kode kategori: ", autoID)

	fmt.Print("Nama kerusakan: ")
	fmt.Scan(&daftar[*n].namaKerusakan)
	*n++
	fmt.Println("Data kategori kerusakan berhasil ditambahkan.")
}

func showKategori(daftar DaftarKategori, n int) {
	if n == 0 {
		fmt.Println("Belum ada data kategori kerusakan yang terdaftar.")
		return
	}
	fmt.Println("-- DAFTAR KATEGORI KERUSAKAN --")
	for i := 0; i < n; i++ {
		fmt.Printf("Kode Kategori: %s, Nama Kerusakan: %s\n", daftar[i].kodeKategori, daftar[i].namaKerusakan)
	}

}

func cariKategori(daftar DaftarKategori, n int, kode string) int {
	for i := 0; i < n; i++ {
		if daftar[i].kodeKategori == kode {
			return i
		}
	}
	return -1
}