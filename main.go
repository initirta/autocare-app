package main

import "fmt"

func main() {
	var kendaraan DaftarKendaraan
	var servis DaftarServis
	var nKendaraan int = 0
	var nServis int = 0

	var pilihan int

	for {
		fmt.Println("1. Kelola Data Kendaraan")
		fmt.Println("2. Catat Riwayat Servis")
		fmt.Println("3. Edit/Ubah Data")
		fmt.Println("4. Cari Data Kendaraan/Servis")
		fmt.Println("5. Tampilkan Laporan & Statistik")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu (1-6): ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			for {
				var subPilihan int
				fmt.Println("\nMENU KELOLA KENDARAAN")
				fmt.Println("1. Tambah kendaraan")
				fmt.Println("2. Lihat semua kendaraan")
				fmt.Println("3. Kembali ke menu utama")
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
		} else if pilihan == 3 {
			fmt.Println("\n[INFO] Masuk ke Menu Edit Data...")
		} else if pilihan == 4 {
			fmt.Println("\n[INFO] Masuk ke Menu Cari Data...")
		} else if pilihan == 5 {
			fmt.Println("\n[INFO] Masuk ke Menu Statistik...")
		} else if pilihan == 6 {
			fmt.Println("\nTerima kasih telah menggunakan AutoCare, sampai jumpa!")
			break
		} else {
			fmt.Println("\nPilihan tidak valid!")
		}
		fmt.Println()
	}
}
