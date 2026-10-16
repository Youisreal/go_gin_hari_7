# Go Hari 7

> **Go Hari 7** - Proyek pembelajaran Go (Golang) yang dibangun dengan semangati harian praktis.

## 📖 Tentang Proyek

Proyek ini adalah hasil eksplorasi dan pembelajaran bahasa pemrograman Go melalui tantangan harian selama 7 hari. Setiap hari akan membawa kita merasakan kemajuan progresi dalam menguasai Go dari dasar hingga menulis kode yang efisien dan bersih.

## 🚀 Fitur

- **Go Modules**: Menggunakan manajemen dependensi modern
- **Clean Code**: Struktur kode yang mudah dipelihara
- **Testing**: Unit tests dengan framework bawaan Go
- **Performance**: Optimasi performa dengan profiling tools

## 🛠️ Prerequisites

- **Go** >= 1.21 (pastikan sudah terinstall)
- **Git** (untuk version control)
- **VS Code** (opsional, dengan ekstensi Go)

## 📦 Instalasi

Ikuti langkah-langkah berikut untuk menjalankan proyek ini:

```bash
# Clone repository
git clone https://github.com/username/go-hari-7.git
cd go-hari-7

# Download dependensi
go mod download

# Run program
go run main.go
```

## 🧪 Testing

Jalankan unit tests dengan perintah:

```bash
go test -v ./...
```

Untuk coverage:

```bash
go test -cover
```

## 📊 Project Structure

```
go-hari-7/
├── .gitignore          # File untuk mengatur file yang tidak di-track git
├── README.md           # Dokumentasi proyek
├── go.mod              # Module definition & dependencies
├── go.sum              # Dependency checksums
├── main.go             # Entry point aplikasi
└── internal/           # Package internal
```

## 🎯 Tujuan Pembelajaran

1. Memahami dasar-dasar Go: tipe data, fungsi, struct
2. Menguasai konservasi: panic, recover, channels
3. Membuat API REST dengan Gorilla Mux
4. Implementasi database dengan GORM
5. Testing dan benchmarking
6. Deployment dengan Docker (jika applicable)

## 🤝 Contributing

Terima kasih sudah menggunakan proyek ini! Jika Anda memiliki saran atau kontribusi, silakan buka issue atau pull request.

## 📝 License

Proyek ini dibuat untuk tujuan edukasi dan pembelajaran. Silakan gunakan dan modifikasikan sesuai kebutuhan Anda.

---

<div align="center">
  <p>Made with ❤️ using <a href="https://go.dev/">Go Lang</a></p>
</div>