package main

import "fmt"

func main() {
	var kendaraan DaftarKendaraan
	var servis DaftarServis
	var pemilik DaftarPemilik
	var nKendaraan int = 0
	var nServis int = 0
	var nPemilik int = 0

	var pilihan int

	for {
		fmt.Println("1. Kelola Data Pemilik")
		fmt.Println("2. Kelola data kendaraan")
		fmt.Println("3. Catat Riwayat Servis")
		fmt.Println("4. Edit/Ubah Data")
		fmt.Println("5. Cari Data Kendaraan/Servis")
		fmt.Println("6. Tampilkan Laporan & Statistik")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu (1-7): ")
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
				fmt.Println("3. Kembali ke menu utama")
				fmt.Print("Pilih menu (1-3): ")
				fmt.Scan(&subPilihan)

				if subPilihan == 1 {
					tambahKendaraan(&kendaraan, &nKendaraan, pemilik, nPemilik)
				} else if subPilihan == 2 {
					showKendaraan(kendaraan, nKendaraan, pemilik, nPemilik)
				} else if subPilihan == 3 {
					break
				} else {
					fmt.Println("Pilihan tidak valid!")
				}
			}
		} else if pilihan == 3 {
			for {
				var subPilihan int
				fmt.Println("\n MENU CATAT RIWAYAT SERVIS")
				fmt.Println("1. Catat servis baru")
				fmt.Println("2. Lihat riwayat servis")
				fmt.Println("3. Kembali ke menu utama")
				fmt.Scan(&subPilihan)

				if subPilihan == 1 {
					tambahServis(&servis, &nServis, kendaraan, nKendaraan)
				} else if subPilihan == 2 {
					showServis(servis, nServis)
				} else if subPilihan == 3 {
					break
				} else {
					fmt.Print("Pilihan tidak valid!")
				}
			}
		} else if pilihan == 4 {
			for {
				var subPilihan int
				fmt.Println("\n>> MENU EDIT/UBAH DATA <<")
				fmt.Println("1. Edit Data Pemilik")
				fmt.Println("2. Edit Data Kendaraan")
				fmt.Println("3. Hapus Data Pemilik")
				fmt.Println("4. Hapus Data Kendaraan")
				fmt.Println("5. Kembali ke Menu Utama")
				fmt.Print("Pilih menu (1-3): ")
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
		} else if pilihan == 5 {
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
		} else if pilihan == 6 {
			showStatistik(servis, nServis)
		} else if pilihan == 7 {
			fmt.Println("")
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
		fmt.Println()
	}
}
