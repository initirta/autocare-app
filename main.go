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
		fmt.Println("1. Kelola Data Kendaraan")
		fmt.Println("2. Kelola Data Pemilik")
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
				fmt.Println("\nMENU KELOLA KENDARAAN")
				fmt.Println("1. Tambah kendaraan")
				fmt.Println("2. Lihat semua kendaraan")
				fmt.Println("3. Kembali ke menu utama")
				fmt.Print("Pilih menu (1-3): ")
				fmt.Scan(&subPilihan)

				if subPilihan == 1 {
					tambahKendaraan(&kendaraan, &nKendaraan)
					break
				} else if subPilihan == 2 {
					showKendaraan(kendaraan, nKendaraan)
					break
				} else if subPilihan == 3 {
					break
				} else {
					fmt.Println("Pilihan tidak valid!")
				}
			}
		} else if pilihan == 2 {
			for {
				for {
				var subPilihan int
				fmt.Println("\nMENU KELOLA PEMILIK")
				fmt.Println("1. Tambah pemilik")
				fmt.Println("2. Lihat semua pemilik")
				fmt.Println("3. Kembali ke menu utama")
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
					break
				} else if subPilihan == 2 {
					showServis(servis, nServis)
					break
				} else if subPilihan == 3 {
					break
				} else {
					fmt.Print("Pilihan tidak valid!")
				}
			}
		} else if pilihan == 4 {
			fmt.Println("\n[INFO] Masuk ke Menu edit Data...")
		} else if pilihan == 5 {
			fmt.Println("\n[INFO] Masuk ke Menu cari Data...")
		} else if pilihan == 6 {
			fmt.Println("\n[INFO] Masuk ke Menu Statistik...")
		} else if pilihan == 7 {
			fmt.Println("\nTerima kasih telah menggunakan AutoCare, sampai jumpa!")
			break
		}else {
			fmt.Println("Pilihan tidak valid!")
		}
		fmt.Println()
	}
}
