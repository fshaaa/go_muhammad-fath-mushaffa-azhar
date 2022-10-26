# Summary 22: Dockers
## Docker
Docker adalah sebuah sistem operasi (OS) virtual
yang mengirimkan perangkat lunak dalam bentuk kontainer.

## Container VS Virtual Machine
| Container             | Virtual Machine                         |
|-----------------------|-----------------------------------------|
| Ringan                | Berat                                   |
| Performa maksimum     | Performa Terbatas                       |
| Virtual pada level OS | Virtual pada level Hardware             |
| Proses terisolasi     | Hardware terisolasi sehingga lebih aman |

## Docker Volume
Docker Volume adalah sebuah storage atau tempat penyimpanan data di container.
Dengan adanya docker volume ini data-data tidak akan terhapus ketika container
dimatikan.

## Docker Network
Secara default, container yang dibangun itu terisolasi dengan container lainnya.
Oleh sebab itu diperlukan penghubung antar container sehingga antar container 
dapat berhubungan, maka diperlukannya sebuah Docker Network.