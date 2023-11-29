## Echo + GORM Clean Architecture Boilerplate

Projek boilerplate yang saya miliki adalah kerangka awal untuk mengembangkan aplikasi berbasis Go yang kuat dan efisien dengan teknologi-teknologi kunci berikut:

- Echo: Framework web yang ringan dan cepat untuk membangun aplikasi web dan RESTful API dengan Go. Echo memiliki banyak fitur bawaan yang memungkinkan Anda untuk dengan cepat membangun layanan web yang handal dan efisien.

- GORM: ORM (Object-Relational Mapping) yang kuat untuk Go yang memungkinkan Anda - berinteraksi dengan database dengan mudah. GORM memudahkan pemodelan data Anda dalam kode Go, sehingga Anda dapat mengakses database Anda dengan nyaman dan aman.

- GORM Driver MySQL: Driver khusus MySQL untuk GORM. Ini memungkinkan Anda berkomunikasi dengan database MySQL secara langsung dari aplikasi Go Anda menggunakan GORM.

Dengan menggunakan boilerplate ini, Anda dapat memulai proyek Go Anda dengan cepat, mengintegrasikan basis data dengan mudah menggunakan GORM, dan menjalankan server Anda dengan dukungan hot reload untuk pengembangan yang lebih efisien. Ini adalah dasar yang kuat untuk membangun berbagai jenis aplikasi web dan layanan RESTful dengan Go.

### Table of Content
- [Installation](#installation)
- [Environment](#environment)
- [Documentation](#documentation)
- [Contributors](#contributors)


### Installation

- Clone a repository
```bash
git clone https://github.com/hafidznaufl/clean-architecture-boilerplate.git && cd clean-architecture-boilerplate
```

- Install all dependencies
```bash
go mod tidy
```
### Environment

File `.env.example` adalah file contoh konfigurasi yang digunakan dalam proyek ini. File ini berisi daftar variabel lingkungan yang harus diatur dalam file `.env` yang sesungguhnya untuk menjalankan proyek dengan benar. Silakan salin file ini sebagai referensi untuk mengatur variabel lingkungan yang sesuai.


Berikut adalah daftar variabel lingkungan yang diperlukan dalam file `.env`:

1. **DB_USER**: Nama pengguna database.
2. **DB_PASS**: Kata sandi pengguna database.
3. **DB_HOST**: Host database.
4. **DB_PORT**: Port database.
5. **DB_NAME**: Nama database yang digunakan.
6. **JWT_SECRET**: Kunci JWT Bearer.
7. **OPEN_AI_TOKEN**: Token OpenAI.
8. **FIREBASE_URL**: Link base url firebase.

#### How to use `.env.example`

- Duplikat file `.env.example` sebagai `.env` dan membuatnya secara otomatis apabila belum tersedia

```bash
cp -n .env .env
```
- Isi nilai variabel lingkungan pada `.env` dengan lingkungan yg anda miliki

### Documentation

Dalam proyek ini, kami menggunakan beberapa teknologi kunci untuk membangun layanan web yang kuat dan efisien. Berikut adalah daftar link ke dokumentasi resmi dan repository GitHub untuk masing-masing teknologi tersebut:

- **Echo**
  - [Dokumentasi Resmi Echo](https://echo.labstack.com/)
  - [Echo GitHub Repository](https://github.com/labstack/echo)

- **GORM**
  - [Dokumentasi Resmi GORM](https://gorm.io/docs/)
  - [GORM GitHub Repository](https://github.com/go-gorm/gorm)

- **GORM Driver untuk MySQL**
  - [Dokumentasi GORM tentang Menghubungkan ke Database MySQL](https://gorm.io/docs/connecting_to_the_database.html#MySQL)

Silakan klik tautan-tautan di atas untuk mengakses dokumentasi resmi dan repository GitHub dari masing-masing teknologi. Dokumentasi ini akan memberikan informasi lebih lanjut, petunjuk penggunaan, dan referensi yang dibutuhkan untuk bekerja dengan teknologi-teknologi tersebut dalam proyek Anda.

## Contributors 

[![Avatar Hafidz Naufal](https://avatars.githubusercontent.com/hafidznaufl?s=50)](https://github.com/hafidznaufl)
