# Summary 11: Concurrent Programing
## GO Routine
goroutine merupakan salah satu bagian terpenting didalam concurrent programming di GO.

Dengan menggunakan goroutine, dapat mengeksekusi suatu program dengan menjalankan multi-core processor.
Maka dengan begitu program dapat dijalankan dengan kebih cepat.

## Channel dan Select 
Channel digunakan untuk menghubungkan goroutine satu dengan goroutine lain. Mekanismenya adalah mekakukan
serah-terima data melalui channel tersebut.

Select digunakan ketika menggunakan channel lebih dari satu, hal ini dapat memudahkan untuk mengontrol data
yang terhubung.

## Race Condition
Race Condition merupakan kondisi dimana lebih dari satu goroutine, mengakses data pada waktu yang bersamaan.
Ketika hal itu terjadi dapat menyebabkan nilai data menjadi kacau karena saling mengakses.

Untuk mencegah tersebut, ada `sync.Mutex` yang didalamnya terdapat `lock` dan `unlock` yang dapat mengatasi
race condition.
