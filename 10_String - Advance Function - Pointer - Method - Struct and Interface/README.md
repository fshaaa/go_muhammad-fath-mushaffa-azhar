# Summary 10 : String - Advance Function - Pointer - Method - Struct and Interface
## String 
String adalah sebuah tipe data yang digunakan untuk menyimpan nilai array of char (kumpulan karakter).

Di dalam Golang, terdapat beberapa fungsi string yaitu:
- Len
- Compare
- Contains
- Substring
- Replace
- Insert

## Advance Function
Di dalam Golang, mengadopsi beberapa fungsi yang unik, yaitu diantaranya
- Fungsi Variadic (parameternya tidak terbatas)
- Fungsi Closure (bisa disimpan di variabel)
- Fungsi Defer (dikerjakan setelah fungsi parent selesai)

## Pointer
Pointer adalah sebuah tipe data yang digunakan untuk menyimpan suatu alamat variabel lain.
Operator pointer:
- *(digunakan untuk mengakses nilai dari suatu variabel lain) 
- &(digunakan untuk mengakses dan mengembalikan alamat variabel lain)

## Method
Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya). Method bisa diakses lewat variabel objek. Keunggulan method dibanding fungsi biasa adalah memiliki akses ke property struct hingga level private.

## Struct and Interface
Struct adalah sebuah tipe data yang didalamnya terdapat beberapa tipe data yang lain(kumpulan tipe data yang lain).

Interface kumpulan definisi method yang tidak memiliki isi (hanya definisi saja), yang dibungkus dengan nama tertentu.

Interface merupakan tipe data. Nilai objek bertipe interface zero value-nya adalah `nil`