# 3 Version Control and Branch Management (Git)

Tugas Praktikum: https://github.com/fshaaa/3_git

## Version Control System 
Version Control System merupakan sebuah sistem yang membantu untuk men-*tracking*sebuah perubahan 
pada berkas atau file-file dari waktu ke waktu.

### Jenis-jenis Version Control System

#### Single User
Single User ini sebuah jenis VCS yang mencatat perubahannya di lokal repository 
atau sesuai dengan namanya yaitu untuk satu user saja.
Contoh: SCCS, RCS

#### Centralized 
Sesuai namanya, pada jenis ini sebuah repository akan terpusat ke sebuah server yang nantinya para developer
mengerjakan sebuah proyek yang tersimpan di dalam server tersebut.
Contoh: CVS, Perforce, Subvesrion, Microsoft Team Foundation Server

#### Distributed
Jenis ini merupakan yang sering digunakan oleh para developer karena pada jenis ini developer mempunyai repo 
di lokal dan di cloud(server), para developer ini dapat mengambil atau memasukkan suatu perubahan pada repo yang ada
di cloud(server).
Contoh: Git, Mercurial, Bazaar

#### Layanan Cloud
Layanan Cloud merupakan sebuah server yang digunakan untuk menyimpan berbagai macam berkas(kode program, dll). 
Pada saat ini sudah banyak layanan cloud yang tersedia dan tidak menutup kemungkinan lagi beberapa tahun kedepan
akan bertambah atau peningkatan layanan cloud yang sudah tersedia saat ini.
Layanan Cloud yang populer saat ini yaitu GitHub, GitLab, BicBucket, dll. 

Pada sesi kali ini akan menggunakan layanan cloud dari Github

## Git dan GitHub
Git merupakan salah satu version control system yang populer yang digunakan oleh para developer
untuk mengembangkan sofware secara bersama-sama. 

Github adalah penyimpanan cloud yang sudah terintegrasi dengan version control system.

### Installing and Setting
Silahkan mengunjungi situs https://git-scm.com/downloads lalu download sesuai dengan sistem operasinya masing-masing.

Setelah itu di-*setting* dengan memasukkan nama dan email, dengan cara
```
git config --global user.name "your-name"
git config --global user.email "your-email"
```

Untuk melakukan inisialisasi Repository ada 2 cara yaitu membuat repo baru dan *cloning* suatu repo.
Membuat repo baru dengan cara:
```
git init
git remote add <remote_name> <link_repo>
git push <remote_name> <branch_name>
```
atau *cloning* sebuah repo:
```
git clone <url_repo>
```

### Saving Changes
Terdapat 3 tempat yang digunakan yaitu Working Directory, Staging Area, Repository
#### Working Directory
Tempat yang digunakan oleh para developer untuk menulis atau memodifikasi suatu kode program. 

Apabila sudah selesai, dapat memasukkan hasil tersebut ke Staging Area, dengan cara:
```
git add <file_name>
```
atau memasukkan semua perubahan file
```
git add .
```
#### Staging Area
Menambahkan suatu file atau kode program yang selanjutnya akan di-commit.
Agar perubahan itu dapat disimpan dan ter-*tracking* maka dilakukan commit 
```
git commit -m "your-message"
```
maka commit ini akan membawa perubahan fila ke repository
#### Repository
Sesudah melakukan commit maka peruban tersebut sudah masuk ke repository lokal, namun 
belum masuk ke repository server, untuk memesukkan commit tersebut dilakukan push
```
git push <remote-name> <branch-name>
```
maka push akan mnegirim perubahan ke repo server (github)

Selain itu ada pula stashing yang mana ada apabila kita melakukan 
```
git stash
```
maka akan menyimpan perubahan sementara teteapi tidak dicommit
dan dapat mengambil dengan cara 
'''
git stash apply
'''
### Inspecting Repository
Dengan melakukan commit, dapat melihat perubahan apa saja yang sudah terjadi di file, 
untuk mengetahui alur commit dapat dilakukan
```
git log --online
```
dan nanti akan muncul kode commit dan pesan yang sudah dbuat saat melakukan commit.

Apabila ingin kembali pada commit sebelumnya dapat melakukan
```
git reset <commit_code> --soft
```
atau mengembalikan semua commit dengan menghapus semua commit setelahnya
```
git reset <commit_code> --hard
```

### Syncing
Sinkronisasi antara repo lokal dan repo di github sangatlah penting,
oleh karena itu setiap ada perubahan sebaiknya kita lakukan push ke repo github
```
git push <remotea_name> <branch_name>
```
Apabila ternyata di repo gtihub sudah ada yang orang lain yang commit dan ingin mengambil perubahan 
tersebut ke repo lokal maka bisa melakukan 
```
git pull <remote_name> <branch_name>
```
atau dengan mengambil seluruh commit di repo Github
```
git fetch
```
Juga dapat melakukan Pull Request di github untuk melakukan perubahan yang mengharuskan merge dengan branch yang lain.

### Branches
Pada dasarnya apabila sudah melakukan git init, repo berada di branch master/main.
Tetapi pada branch ini, jangan pernah dilakukan perubahan apabila belum mencapai final. 

Maka dari itu dibutuhkan branch lain untuk melakukan perkembangan suatu program, ini dapat dilakukan dengan cara
```
git branch <branch_name>
```
setelah itu pindah ke branch lain dengan cara
```
git checkout <branch_name>
```
atau dapat melakuakan membuat dan pindah sekaligus dengan cara
```
git checkout -b <branch_name>
```

Untuk melihat daftar branch yang tersedia yaitu
```
git branch --list
```
atau melihat semua yang termasuk di repo di Github dengan cara
```
git branch -a
```

Nah untuk menggabungkan suatu branch ke branch yang lain, ini dengan fungsi merge
```
git merge <dest_branch_name>
```

Setelah di-*merge*, branch dapat dihapus
```
git branch -D <branch_name>
```

## Referensi
- https://docs.google.com/presentation/d/1Y1pqmfT-p5ximnWJ4jrXrHcbrlBYnDdeoRLlvJ8dHpE/edit#slide=id.g6513da4353_0_0
- https://www.hostinger.co.id/tutorial/cara-menggunakan-github-perintah-dasar-github
- https://git-scm.com/book/id/v2/Memulai-Tentang-Version-Control

