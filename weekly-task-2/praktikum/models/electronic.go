package models

type Electronics struct {
	ID             string `json:"id" from:"id"`
	Nama           string `json:"nama" from:"nama"`
	Deskripsi      string `json:"deskripsi" from:"deskripsi"`
	JumlahStok     int    `json:"jumlahStok" from:"jumlahStok"`
	Harga          int    `json:"harga" from:"harga"`
	KategoriBarang int    `json:"kategoriBarang" from:"kategoriBarang"`
}
