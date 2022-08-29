# Summary 5 : Basic Programming
## Golang (Go Language)
Go merupakan bahasa pemrograman yang dibuat oleh Google pada tahun 2007, dan dipublikasikan pada tahun 2009. Sekarang ini Go merupakan salah satu bahasa pemrograman yang populer, dikarenakna memiliki kelebihan dari sisi kecepatan, stabilitas, dan keandalan.

## Go Module 
Pada dasarnya pengembangan project Go di `GOPATH`, untuk menginisialisasi project dengan cara
```
go mod init <project-name>
```
pada umunya nama project sama dengan nama folder workspace yang akan digunakan

## Variabel
Variabel adalah suatu tanda yang digunakan sebagai penampung suatu nilai. Variabel memiliki tipe data, alamat, dan nilai. 
Tipe data pada bahasa Go sendiri dibedakan menjadi 3 yaitu String, Numeric(int, float, complex), Boolean.
Adapun cara menginsialisasi variabel yaitu dengan cara
```
var <variable-name> <data-type> = <value>
```
atau bisa tanpa menggunakan `value`, sehingga
```
var <variable-name> <data-type>
```
dan masih banyak lagi. 

Adapun cara ringkasnya yaitu:
```
<variable-name> := <value>
```
pada inisialisasi diatas, tipe data variabel bergantung pada valuenya.

## Control Structure
Pada bahasa Go ini sama dengan kebanyakan bahasa pemrograman yang lain, yang mana saat melakukan *branching* menggunakan `if-else` atau `switch`, dan saat malakukan looping menggunakan `for`.

Cara Insialisasi `if-else` yaitu
```
if <condition> {

} else if <condition> {

} else {

}
```
Note: Kurung kurawal harus diletakkan sebaris dengan `if, else if, else`, dan `else-if dan else` tidak boleh beda baris dengan kurung kurawal sebelumnya.

Cara inisiaisasi `switch` yaitu
```
switch <condition>{
  case value:

  case value:

  default:

}
```

Cara inisialisai `for` yaitu
```
for <init>;<condition>;<past>{

} 
```
Note: bisa ommited salah satu ketiganya atau ketiga-tiganya dengan syarat conditon dan past-nya berada di dalam for agat tidak terjadi unlimited looping.

## Referensi
- https://alta.id/courses/become-a-master-of-golang-programme/16691/ 