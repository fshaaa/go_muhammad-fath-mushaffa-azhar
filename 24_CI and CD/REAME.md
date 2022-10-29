# Summary 24: Continuous Intergration and Continuous Deployment
## Continuous Intergration (CI)
continuous Intergration adalah sebuah proses automasi yang dilakukan untuk 
mengintergasikan berbagai kode dari berbagai sumber potensial untuk memebangun
atau mengujinya

## The Cycle CI
Check in Changes -> Fetch Changes -> Build -> Test -> Fail or Succeed -> 
Notify Success or Fail

## Continuous Deployment (CD)
Continuous Deployment adalah sebuah proses terus menerus yang mana semua
proses dari test sampai production sudah fully automated tanpa Human Intervention.

## The Cycle CD
Unit test -> Platform test -> Deliver to staging -> Application Acceptance Test -> 
Deploy to production -> Post deploy tests

## Continuous Delivery vs Continuous Deployment 
Continuous Delivery kadang tidak sepenuhnya bekerja otomatis masih terdapat human intervation
pada proses deployment ke production.

Deployment memang masih manual tapi itu digunakan saat approving saja.