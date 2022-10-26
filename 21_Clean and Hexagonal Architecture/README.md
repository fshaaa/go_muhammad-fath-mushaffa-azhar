# Summary 21: Clean and Hexagonal Architecture
## Clean Architecture
Dalam konsep clean architecture ini, setiap komponen yang ada bersifat
independen tidak bergantung pada library eksternal yang spesifik.

### The Goal of Clean Architecture
Tujuan menggunakan clean Architecture adalah supaya code yang sudah dibuat
lebih modular, scallabel, dan maintainable
- Modular -> dapat dengan mudah mengganti dipendensi satu ke dipendensi lain
- Scallabel -> dapat dengan mudah menambahkan feature baru dan lain sebagainya
- Maintainable -> deapat dengan mudah memperbaiki issue

### The Constraint before designed Clean Architecture
- Independent of Frameworks
- Testable
- Independent of UI
- Independent of Database
- Independent of any External

### CA Layer
- Entity Layer (opsional) -> business object yang menpresentatifkan konsep application 
- Usecase - Domain layer -> terdiri dari business logic
- Controller - Presentation layer -> menampilkan data ke client dan meng-handle user interactions
- Drivers - Data layer -> mengelola data aplikasi

### Clean Architecture VS Domain Driven Design (DDD)
- Clean Architecture is a sofware architecture
- Domain Driven Design (DDD) is sofware design technique