package main

type kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}

type mobil struct {
	kendaraan
}

func (m *mobil) berjalan() {
	m.tambahKecepatan(10)
}

func (m *mobil) tambahKecepatan(kecepatanBaru int) {
	m.kecepatanPerJam = m.kecepatanPerJam + kecepatanBaru
}

func main() {
	mobilCepat := mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	mobilLamban := mobil{}
	mobilLamban.berjalan()
}
