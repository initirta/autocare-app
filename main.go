package main

import "fmt"

func main() {
	var kendaraan DaftarKendaraan
	var servis DaftarServis
	var pemilik DaftarPemilik
	var kategori DaftarKategori
	var nKendaraan int = 0
	var nServis int = 0
	var nPemilik int = 0
	var nKategori int = 0

	var pilihan int

	for {
		fmt.Println("1. Kelola Data Pemilik")
		fmt.Println("2. Kelola data kendaraan")
		fmt.Println("3. Kelola Kategori Service")
		fmt.Println("4. Catat Riwayat Servis")
		fmt.Println("5. Edit/Ubah Data")
		fmt.Println("6. Cari Data Kendaraan/Servis")
		fmt.Println("7. Tampilkan Laporan & Statistik")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu (1-8): ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			for {
				var subPilihan int
				fmt.Println("\nMENU KELOLA PEMILIK")
				fmt.Println("1. Tambah pemilik")
				fmt.Println("2. Lihat semua pemilik")
				fmt.Println("3. Kembali ke menu utama")
				fmt.Print("Pilih menu (1-3): ")
				fmt.Scan(&subPilihan)

				if subPilihan == 1 {
					tambahPemilik(&pemilik, &nPemilik)
				} else if subPilihan == 2 {
					showPemilik(pemilik, nPemilik)
				} else if subPilihan == 3 {
					break
				} else {
					fmt.Println("Pilihan tidak valid!")
				}
			}
		} else if pilihan == 2 {
			for {
				var subPilihan int
				fmt.Println("\nMENU KELOLA KENDARAAN")
				fmt.Println("1. Tambah kendaraan")
				fmt.Println("2. Lihat semua kendaraan")
				fmt.Println("3. Urutkan kendaraan berdasarkan tahun produksi")
				fmt.Println("4. Urutkan berdasarkan plat")
				fmt.Println("5. Kembali ke menu utama")
				fmt.Print("Pilih menu (1-5): ")
				fmt.Scan(&subPilihan)

				if subPilihan == 1 {
					tambahKendaraan(&kendaraan, &nKendaraan, pemilik, nPemilik)
				} else if subPilihan == 2 {
					showKendaraan(kendaraan, nKendaraan, pemilik, nPemilik)
				} else if subPilihan == 3 {
					selectionTahunKendaraan(&kendaraan, nKendaraan)
				} else if subPilihan == 4 {
					selectionPlat(&kendaraan, nKendaraan)
				} else if pilihan == 5 {
					break
				} else {
					fmt.Println("Pilihan tidak valid")
				}
			}
		} else if pilihan == 3 {
			for {
				var subPilihan int
				fmt.Println("\nMENU CATAT KATEGORI SERVICE")
				fmt.Println("1. Tambah kategori service")
				fmt.Println("2. Lihat semua kategori service")
				fmt.Println("3. Kembali ke menu utama")
				fmt.Print("Pilih menu (1-3): ")
				fmt.Scan(&subPilihan)

				if subPilihan == 1 {
					tambahKategori(&kategori, &nKategori)
				} else if subPilihan == 2 {
					showKategori(kategori, nKategori)
				} else if subPilihan == 3 {
					break
				} else {
					fmt.Println("Pilihan tidak valid!")
				}
			}
		} else if pilihan == 4 {
			for {
				var subPilihan int
				fmt.Println("\n MENU CATAT RIWAYAT SERVIS")
				fmt.Println("1. Catat servis baru")
				fmt.Println("2. Lihat riwayat servis")
				fmt.Println("3. Urutkan riwayat servis berdasarkan tanggal")
				fmt.Println("4. Kembali ke menu utama")
				fmt.Print("Pilih menu (1-4): ")
				fmt.Scan(&subPilihan)

				if subPilihan == 1 {
					tambahServis(&servis, &nServis, kendaraan, nKendaraan, kategori, nKategori)
				} else if subPilihan == 2 {
					showServis(servis, nServis, kategori, nKategori)
				} else if subPilihan == 3 {
					insertionTglServis(&servis, nServis)
				} else if pilihan == 4 {
					break
				} else {
					fmt.Println("Pilihan tidak valid!")
				}
			}
		} else if pilihan == 5 {
			for {
				var subPilihan int
				fmt.Println("\n>> MENU EDIT/UBAH DATA <<")
				fmt.Println("1. Edit Data Pemilik")
				fmt.Println("2. Edit Data Kendaraan")
				fmt.Println("3. Hapus Data Pemilik")
				fmt.Println("4. Hapus Data Kendaraan")
				fmt.Println("5. Kembali ke Menu Utama")
				fmt.Print("Pilih menu (1-5): ")
				fmt.Scan(&subPilihan)

				if subPilihan == 1 {
					editPemilik(&pemilik, nPemilik)
				} else if subPilihan == 2 {
					editKendaraan(&kendaraan, nKendaraan, pemilik, nPemilik)
				} else if subPilihan == 3 {
					hapusPemilik(&pemilik, &nPemilik, kendaraan, nKendaraan)
				} else if subPilihan == 4 {
					hapusKendaraan(&kendaraan, &nKendaraan)
				} else if subPilihan == 5 {
					break
				} else {
					fmt.Print("Pilihan tidak valid!")
				}
			}
		} else if pilihan == 6 {
			for {
				var subPilihan int
				fmt.Println("\nMENU CARI DATA")
				fmt.Println("1. Cari Data Kendaraan (by Plat)")
				fmt.Println("2. Cari Riwayat Servis (by ID Servis)")
				fmt.Println("3. Kembali ke Menu Utama")
				fmt.Print("Pilih menu (1-3): ")
				fmt.Scan(&subPilihan)

				if subPilihan == 1 {
					detailCariKendaraan(kendaraan, nKendaraan, pemilik, nPemilik)
				} else if subPilihan == 2 {
					detailCariServis(servis, nServis)
				} else if subPilihan == 3 {
					break
				} else {
					fmt.Println("Pilihan tidak valid!")
				}
			}
		} else if pilihan == 7 {
			showStatistik(servis, nServis, kategori, nKategori)
		} else if pilihan == 8 {
			fmt.Println("Terima kasih, program selesai.")
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
		fmt.Println()
	}
}
