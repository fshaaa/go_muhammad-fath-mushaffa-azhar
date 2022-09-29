# Summary 18: ORM and Code Structure (MVC)
## ORM 
ORM atau kependekan dari Object Relational Mapping.

ORM merupakan sebuah teknik yang merubah suatu table menjadi sebuah object yang nantinya mudah untuk digunakan. 
Object yang dibuat memiliki property yang sama dengan field — field yang ada pada table tersebut.

Salah satu toolnya yaitu gorm.io

## GROM
Install gorm dan mysql di Go:
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## Code Structure (MVC)
Struktur Kode menggunakan MVC agar pembacaan kode lebih mudah

Kepanjangan MVC adalah Model, View, Controlling, yang mana nantinya akan memecah beberapa fungsi 
yang terkait dengan program di sebuah folder. Sehingga memudahkan untuk membaca sebuah kode.