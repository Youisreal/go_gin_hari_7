# Go Hari 7

> **Go Hari 8** - Proyek pembelajaran Go (Golang) membuat user-api dan todo-api.

## 📖 Tentang Proyek

Proyek ini adalah hasil eksplorasi dan pembelajaran bahasa pemrograman Go dengan membuat REST API lengkap (user-api dan todo-api) menggunakan PostgreSQL dan GORM.

## 🚀 Fitur

- **Go Modules**: Menggunakan manajemen dependensi modern
- **Clean Code**: Struktur kode yang mudah dipelihara
- **Testing**: Unit tests dengan framework bawaan Go
- **Performance**: Optimasi performa dengan profiling tools
- **Database Integration**: Koneksi PostgreSQL menggunakan GORM
- **RESTful API**: CRUD lengkap (Create, Read, Update, Delete) untuk resource User dan Todo
- **AutoMigrate**: Tabel `users` dan `todos` dibuat otomatis saat server dijalankan
- **Pagination**: Endpoint list mendukung query `page` dan `limit`

## 🛠️ Prerequisites

- **Go** >= 1.21 (pastikan sudah terinstall, proyek ini dikembangkan dengan Go 1.27)
- **Git** (untuk version control)
- **VS Code** (opsional, dengan ekstensi Go)
- **DATABASE_URL** (variabel lingkungan wajib untuk koneksi database)

## 🧰 Tech Stack

| Layer             | Teknologi                                                             |
| ----------------- | --------------------------------------------------------------------- |
| **Language**      | Go 1.27                                                               |
| **Web Framework** | [Gin](https://github.com/gin-gonic/gin)                               |
| **ORM**           | [GORM](https://gorm.io) + driver PostgreSQL                           |
| **Validation**    | [go-playground/validator](https://github.com/go-playground/validator) |
| **Config**        | [godotenv](https://github.com/joho/godotenv) (baca file `.env`)       |
| **UUID**          | Package standar `uuid` (RFC 9562)                                     |
| **Database**      | PostgreSQL (lokal / Supabase)                                         |

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
```

## 🔧 Konfigurasi Database

Proyek ini menggunakan **PostgreSQL** dengan koneksi melalui variabel lingkungan `DATABASE_URL`. Pastikan variabel ini diatur sebelum menjalankan aplikasi:

```bash
export DATABASE_URL="postgres://username:password@localhost:5432/go-hari-7?sslmode=disable"
```

Contoh format `DATABASE_URL`:

```bash
postgres://user:password@host:port/database?sslmode=disable
```

> **Catatan:** Proyek ini juga dapat dijalankan dengan **Supabase**. Ambil connection string dari dashboard Supabase pada bagian _Project Settings → Database → Connection string_ (pilih **URI**), lalu tempelkan nilainya ke `DATABASE_URL`. File `.env` juga didukung dan akan dibaca otomatis oleh `godotenv`. Pastikan file `.env` tidak ikut ter-_commit_ ke repository.

## 📊 Project Structure

```
go-hari-7/
├── .gitignore           # File untuk mengatur file yang tidak di-track git
├── README.md            # Dokumentasi proyek
├── go.mod               # Module definition & dependencies
├── go.sum               # Dependency checksums
├── main.go              # Entry point aplikasi (koneksi DB + AutoMigrate + server)
├── database/            # Package database
│   └── database.go        # Fungsi Connect() yang memerlukan DATABASE_URL
├── router/              # Package router (Gin)
│   └── router.go        # SetupRouter() mendaftarkan semua route di /api/v1
├── todo/                # Package todo
│   ├── model.go         # Model Todo + request DTO
│   ├── handler.go       # Handler CRUD todo
│   └── routes.go        # Registrasi route /todos
└── users/               # Package users
    ├── model.go         # Model User + request DTO
    ├── handler.go       # Handler CRUD user
    └── routes.go        # Registrasi route /users
```

### Arsitektur Layer

Setiap resource (user & todo) mengikuti pola **Model → Handler → Routes**:

| Layer                      | Tanggung Jawab                                                                        |
| -------------------------- | ------------------------------------------------------------------------------------- |
| **Model** (`model.go`)     | Definisi struct tabel GORM dan struct request (DTO) beserta aturan validasi `binding` |
| **Handler** (`handler.go`) | Menangani request HTTP, validasi input, query database, dan pengiriman response JSON  |
| **Routes** (`routes.go`)   | Mendaftarkan endpoint pada `gin.RouterGroup`                                          |

`router.SetupRouter(db)` menggabungkan seluruh route di bawah prefix `/api/v1`, sedangkan `main.go` menangani koneksi database, `AutoMigrate`, dan menjalankan server.

## 📡 API Documentation

Semua endpoint berada di bawah base URL:

```
http://localhost:8080/api/v1
```

### 👤 User API — `/api/v1/users`

Model `User` pada tabel `users`:

| Field        | Tipe   | Keterangan                                              |
| ------------ | ------ | ------------------------------------------------------- |
| `id`         | UUID   | Primary key, di-generate otomatis (`gen_random_uuid()`) |
| `full_name`  | string | Wajib diisi, minimal 3 karakter                         |
| `email`      | string | Wajib diisi, format email valid, unik                   |
| `password`   | string | Wajib diisi                                             |
| `created_at` | time   | Waktu pembuatan data                                    |
| `update_at`  | time   | Waktu perubahan data terakhir                           |

#### Endpoint

| Method   | Endpoint            | Deskripsi                         | Status Sukses |
| -------- | ------------------- | --------------------------------- | ------------- |
| `GET`    | `/api/v1/users`     | List semua user dengan pagination | `200 OK`      |
| `GET`    | `/api/v1/users/:id` | Detail user berdasarkan ID        | `200 OK`      |
| `POST`   | `/api/v1/users`     | Membuat user baru                 | `201 Created` |
| `PUT`    | `/api/v1/users/:id` | Update user berdasarkan ID        | `200 OK`      |
| `DELETE` | `/api/v1/users/:id` | Hapus user berdasarkan ID         | `200 OK`      |

#### Request Body

```json
{
  "full_name": "Budi Santoso",
  "email": "budi@example.com",
  "password": "rahasia123"
}
```

> **Catatan:** Saat melakukan `PUT`, field `created_at` juga ikut dikirim dari body. Jika field tersebut tidak disertakan, nilainya akan menjadi kosong. Sertakan `created_at` dengan nilai ISO-8601 (contoh `2025-01-01T10:00:00Z`) agar waktu pembuatan tetap dipertahankan.

#### Response

```json
{
  "id": "6f1c9a2e-3b4d-4c7a-8e9f-0a1b2c3d4e5f",
  "full_name": "Budi Santoso",
  "email": "budi@example.com",
  "password": "rahasia123",
  "created_at": "2025-01-01T10:00:00Z",
  "update_at": "2025-01-01T10:00:00Z"
}
```

#### Contoh Penggunaan (curl)

```bash
# List user (page 1, 10 data per halaman)
curl -X GET "http://localhost:8080/api/v1/users?page=1&limit=10"

# Detail user
curl -X GET "http://localhost:8080/api/v1/users/6f1c9a2e-3b4d-4c7a-8e9f-0a1b2c3d4e5f"

# Create user
curl -X POST "http://localhost:8080/api/v1/users" \
  -H "Content-Type: application/json" \
  -d '{"full_name":"Budi Santoso","email":"budi@example.com","password":"rahasia123"}'

# Update user
curl -X PUT "http://localhost:8080/api/v1/users/6f1c9a2e-3b4d-4c7a-8e9f-0a1b2c3d4e5f" \
  -H "Content-Type: application/json" \
  -d '{"full_name":"Budi Santoso Update","email":"budi@example.com","password":"rahasia456"}'

# Delete user
curl -X DELETE "http://localhost:8080/api/v1/users/6f1c9a2e-3b4d-4c7a-8e9f-0a1b2c3d4e5f"
```

### ✅ Todo API — `/api/v1/todos`

Model `Todo` pada tabel `todos`:

| Field     | Tipe   | Keterangan                      |
| --------- | ------ | ------------------------------- |
| `id`      | int    | Primary key auto increment      |
| `title`   | string | Wajib diisi, minimal 3 karakter |
| `is_done` | bool   | Status selesai, default `false` |

#### Endpoint

| Method   | Endpoint            | Deskripsi                                   | Status Sukses |
| -------- | ------------------- | ------------------------------------------- | ------------- |
| `GET`    | `/api/v1/todos`     | List semua todo, bisa difilter `?done=true` | `200 OK`      |
| `GET`    | `/api/v1/todos/:id` | Detail todo berdasarkan ID                  | `200 OK`      |
| `POST`   | `/api/v1/todos`     | Membuat todo baru                           | `201 Created` |
| `PUT`    | `/api/v1/todos/:id` | Update todo berdasarkan ID                  | `200 OK`      |
| `DELETE` | `/api/v1/todos/:id` | Hapus todo berdasarkan ID                   | `200 OK`      |

#### Request Body

```json
{
  "title": "Belajar Go dan Gin",
  "is_done": false
}
```

#### Response

```json
{
  "id": 1,
  "title": "Belajar Go dan Gin",
  "is_done": false
}
```

#### Contoh Penggunaan (curl)

```bash
# List todo dengan filter status
curl -X GET "http://localhost:8080/api/v1/todos?done=false&page=1&limit=10"

# Detail todo
curl -X GET "http://localhost:8080/api/v1/todos/1"

# Create todo
curl -X POST "http://localhost:8080/api/v1/todos" \
  -H "Content-Type: application/json" \
  -d '{"title":"Belajar Go dan Gin","is_done":false}'

# Update todo
curl -X PUT "http://localhost:8080/api/v1/todos/1" \
  -H "Content-Type: application/json" \
  -d '{"title":"Belajar Go dan Gin","is_done":true}'

# Delete todo
curl -X DELETE "http://localhost:8080/api/v1/todos/1"
```

### 📄 Pagination

Endpoint `GET` pada collection mengembalikan array data dan header `X-Total-Count` berisi jumlah total baris.

| Query Param | Default | Keterangan                                                                 |
| ----------- | ------- | -------------------------------------------------------------------------- |
| `page`      | `1`     | Nomor halaman, nilai < 1 akan diubah ke `1`                                |
| `limit`     | `10`    | Jumlah data per halaman, nilai di luar rentang `1-100` akan diubah ke `10` |

Contoh response header:

```
X-Total-Count: 42
```

### ⚠️ Format Error

Semua error dikirim dalam format JSON yang seragam:

```json
{
  "error": "Email tidak valid"
}
```

| Status Code                 | Keterangan                                                   |
| --------------------------- | ------------------------------------------------------------ |
| `400 Bad Request`           | Body JSON tidak valid atau gagal melewati validasi `binding` |
| `404 Not Found`             | Data dengan ID tersebut tidak ditemukan                      |
| `500 Internal Server Error` | Terjadi kesalahan pada sisi server / query database          |

Contoh pesan validasi: `Fullname wajib diisi`, `Fullname minimal 3 karakter`, `Email tidak valid`, `Format JSON tidak valid`.

Contoh response sukses pada endpoint `DELETE`:

```json
{
  "message": "User berhasil dihapus"
}
```

## 🎯 Tujuan Pembelajaran

1. Memahami dasar-dasar Go: tipe data, fungsi, struct
2. Menguasai konservasi: panic, recover, channels
3. Membuat API REST dengan Gin
4. Implementasi database dengan GORM
5. Membangun **user-api** dan **todo-api** dengan layer model, handler, dan routes
6. Validasi input menggunakan `binding` dari `go-playground/validator`
7. Menambahkan pagination, filter, dan auto-migration tabel
8. Testing dan benchmarking
9. Deployment dengan Docker (jika applicable)

## 🤝 Contributing

Terima kasih sudah menggunakan proyek ini! Jika Anda memiliki saran atau kontribusi, silakan buka issue atau pull request.

## 📝 License

Proyek ini dibuat untuk tujuan edukasi dan pembelajaran. Silakan gunakan dan modifikasi sesuai kebutuhan Anda.

<div align="center">
  <p>Made with ❤️ using <a href="https://go.dev/">Go Lang</a></p>
</div>
