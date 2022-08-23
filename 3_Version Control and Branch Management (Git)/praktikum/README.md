# Penjelasan Tugas Praktikum Sec 3
### Membuat Repository
Link: https://github.com/fshaaa/3_git

### Implementasi Branching
saya disini menggunakan branch master, development, featureA, featureB, dengan menggunakan 
```
git checkout -b development
git checkout -b featureA
git checkout development
git checkout -b featureB
```

### Implementasi Intruksi git push, push, stash, merge
#### `git push`
setiap melakukan commit, saya melakukan push ke akun GitHub sesuai dengan kondisinya
```
git push origin <nama_branch>
```

#### `git stash` 
dilakukan ketika ingin menyimpan suatu perubahan tetapi tidak ingin kita commit karena belum yakin pada perubahan tersebut. 
![Git Stash](https://github.com/fshaaa/go_muhammad-fath-mushaffa-azhar/blob/master/3_Version%20Control%20and%20Branch%20Management%20(Git)/screenshot/git%20stash.png)
Tapi karena tidak digunakan maka tidak saya apply atau pop.

#### `git merge`
dilakukan saat branch yang sudah diperbarui diagbungkan dengan root branchnya
```
git checkout master
git merge --no-ff development
```
![Git Merge](https://github.com/fshaaa/go_muhammad-fath-mushaffa-azhar/blob/master/3_Version%20Control%20and%20Branch%20Management%20(Git)/screenshot/git%20merge%20w%20master.png)
Pada merge kali ini tidak terjadi conflict karena perubahan tidak terjadi pada line yang terdapat sama.

#### `git pull`
dilakukan ketika ingin mengambil suatu perubahan dari Github menuju repo lokal.
```
git pull origin master
```
![Git Pull](https://github.com/fshaaa/go_muhammad-fath-mushaffa-azhar/blob/master/3_Version%20Control%20and%20Branch%20Management%20(Git)/screenshot/git%20pull.png)
Pada kali ini saya mengambil sebuah commit dan file README.md dari Github ke repo lokal.

### Implementasi Merge Conflict dan Merge no fast forward
Merge Conflict terjadi apabila di 2 branch yang akan di gabungkan terdapat perbedaan kode program di line yang sama
contohnya saat saya merge development yang terlebih dahulu dimerge dengan featureA lalu akan dimerge lagi dengan featureB.

Untuk perintahnya yaitu 
```
git checkout development
git merge --no-ff featureB
```
![Merge Conflict](https://github.com/fshaaa/go_muhammad-fath-mushaffa-azhar/blob/master/3_Version%20Control%20and%20Branch%20Management%20(Git)/screenshot/git%20merge%20(w%20merge%20conflict).png)
Pada gambar tersebut dapat terlihat ada merge conflict yang terdapat pada line 9.

Untuk menyelesaikan Merge Conflict ada beberapa cara yaitu dipilih salah satu atau dipilih keduannya.
![Fixing Merge Conflict](https://github.com/fshaaa/go_muhammad-fath-mushaffa-azhar/blob/master/3_Version%20Control%20and%20Branch%20Management%20(Git)/screenshot/fixing%20git%20merge.png)
Penyelesaian tersebut adalah dipilih keduanya.

Pada perintah diatas ditulis `--no-ff` yang mana artinya merge dengan no fast forward.
![No Fast Forward](https://github.com/fshaaa/go_muhammad-fath-mushaffa-azhar/blob/master/3_Version%20Control%20and%20Branch%20Management%20(Git)/screenshot/graph%20in%20github.png)
![No Fasth Forward in VSCode](https://github.com/fshaaa/go_muhammad-fath-mushaffa-azhar/blob/master/3_Version%20Control%20and%20Branch%20Management%20(Git)/screenshot/graph%20in%20vscode.png)
yang mana perintah tersebut akan terlihat berbeda ketika kita melihat graphnya, yaitu ketika di-*merge* graphnya akan tetap berbelok
dan mengarah pada root branchnya, tidak menjadi satu seperti ketika menggunakan fast forward.

## Terimakasih
