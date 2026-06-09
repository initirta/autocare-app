package main

const NMAX int = 100

type Kendaraan struct {
	PlatNomor      string
	JenisKendaraan string
	TahunProduksi  int
	IDPemilik      string
}

type Pemilik struct {
	IDPemilik string
	Nama      string
	Alamat    string
	NoTelp    string
}

type Waktu struct {
	Hari  int
	Bulan int
	Tahun int
}

type Servis struct {
	IDServis      int
	PlatNomor     string
	Tanggal       Waktu
	JenisTindakan string
	Biaya         int
}

type DaftarKendaraan [NMAX]Kendaraan
type DaftarServis [NMAX]Servis
type DaftarPemilik [NMAX]Pemilik
