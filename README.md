# Go Hari 8

> **Go Hari 8** - Proyek pembelajaran Go (Golang) yang dibangun dengan semangati harian praktis.

## 📖 Tentang Proyek

Proyek ini adalah hasil eksplorasi dan pembelajaran bahasa pemrograman Go melalui tantangan harian selama 7 hari. Setiap hari akan membawa kita merasakan kemajuan progresi dalam menguasai Go dari dasar hingga menulis kode yang efisien dan bersih.

## 🚀 Fitur

- **Go Modules**: Menggunakan manajemen dependensi modern
- **Clean Code**: Struktur kode yang mudah dipelihara
- **Testing**: Unit tests dengan framework bawaan Go
- **Performance**: Optimasi performa dengan profiling tools
- **Database Integration**: Koneksi PostgreSQL menggunakan GORM

## 🛠️ Prerequisites

- **Go** >= 1.21 (pastikan sudah terinstall)
- **Git** (untuk version control)
- **VS Code** (opsional, dengan ekstensi Go)
- **DATABASE_URL** (variabel lingkungan wajib untuk koneksi database)

## 📦 Instalasi

Ikuti langkah-langkah berikut untuk menjalankan proyek ini:

```bash
# Clone repository

git clone https://github.com/username/go-hari-7.git
cd go-hari-7

# Set DATABASE_URL (wajib)

export DATABASE_URL="postgres://user:password@localhost:5432/dbname?sslmode=disable"

# Download dependensi

go mod download

# Run program

go run main.go

## 🔧 Konfigurasi Database

Proyek ini menggunakan **PostgreSQL** dengan koneksi melalui variabel lingkungan `DATABASE_URL`. Pastikan variabel ini diatur sebelum menjalankan aplikasi:

export DATABASE_URL="postgres://username:password@localhost:5432/go-hari-7?sslmode=disable"

Contoh format `DATABASE_URL`:

postgres://user:password@host:port/database?sslmode=disable

## 📊 Project Structure

go-hari-7/
├── .gitignore # File untuk mengatur file yang tidak di-track git
├── README.md # Dokumentasi proyek
├── go.mod # Module definition & dependencies
├── go.sum # Dependency checksums
├── main.go # Entry point aplikasi
├── internal/ # Package internal
├── database/ # Package database
│ └── database.go # Fungsi Connect() yang memerlukan DATABASE_URL
└── router/ # Package router (Gin)
├── router.go # Setup router
├── handler.go # Handler endpoint
└── setup.go # Setup server

## 🎯 Tujuan Pembelajaran

1. Memahami dasar-dasar Go: tipe data, fungsi, struct
2. Menguasai konservasi: panic, recover, channels
3. Membuat API REST dengan Gin
4. Implementasi database dengan GORM
5. Testing dan benchmarking
6. Deployment dengan Docker (jika applicable)

## 🤝 Contributing

Terima kasih sudah menggunakan proyek ini! Jika Anda memiliki saran atau kontribusi, silakan buka issue atau pull request.

## 📝 License

Proyek ini dibuat untuk tujuan edukasi dan pembelajaran. Silakan gunakan dan modifikasi sesuai kebutuhan Anda.

<div align="center">
  <p>Made with ❤️ using <a href="https://go.dev/">Go Lang</a></p>
</div>
```
