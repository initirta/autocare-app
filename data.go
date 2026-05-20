package main

const NMAX int = 100

type Kendaraan struct {
	PlatNomor      string
	NamaPemilik    string
	JenisKendaraan string
}

type Servis struct {
	IDServis      int
	PlatNomor     string
	Tanggal       string
	JenisTindakan string
	Biaya         int
}

type DaftarKendaraan [NMAX]Kendaraan
type DaftarServis [NMAX]Servis
