# 📘 Customer API

Customer API adalah layanan backend yang menyediakan fitur manajemen data customer menggunakan arsitektur **Clean Architecture**.  
API ini mendukung operasi CRUD (Create, Read, Update, Delete) serta pagination.

---

## 🚀 Teknologi
- Go 1.22+
- GORM
- MySQL / PostgreSQL
- Gorilla Mux
- Clean Architecture Pattern

---

## 📂 Fitur Utama
- ✔ Menampilkan daftar customer (pagination)
- ✔ Menambah customer baru
- ✔ Melihat detail customer
- ✔ Mengupdate customer
- ✔ Menghapus customer

---

## 📁 Struktur Proyek

internal/
│── config/
│ └── db.go
│── domain/
│ └── customer.go
│── repository/
│ └── customer_repository.go
│── usecase/
│ └── customer_usecase.go
│── delivery/
│ └── http/
│ └── customer_handler.go
cmd/
│── server/
│ └── main.go

## 🗄 Model Customer

Field JSON yang digunakan pada API:

```json
{
  "nationality_id": 1,
  "cst_name": "Nama Customer",
  "cst_dob": "1990-01-20",
  "cst_phoneNum": "08123456789",
  "cst_email": "email@example.com"
}

📡 Endpoint API

Base URL:

http://localhost:8000

1️⃣ GET /customers

Menampilkan daftar customer (pagination).

Query Parameters
Parameter	Tipe	Default	Deskripsi
page	int	1	Halaman
limit	int	10	Jumlah data per halaman
Contoh CURL
curl -X GET "http://localhost:8000/customers?page=1&limit=5"

2️⃣ POST /customers

Menambah customer baru.

Body (JSON)
{
  "nationality_id": 1,
  "cst_name": "Budi Setiawan",
  "cst_dob": "1995-01-20",
  "cst_phoneNum": "08123456789",
  "cst_email": "budi@example.com"
}

Contoh CURL
curl -X POST http://localhost:8000/customers \
  -H "Content-Type: application/json" \
  -d '{
    "nationality_id": 1,
    "cst_name": "Budi Setiawan",
    "cst_dob": "1995-01-20",
    "cst_phoneNum": "08123456789",
    "cst_email": "budi@example.com"
  }'

3️⃣ GET /customers/{id}

Menampilkan detail customer berdasarkan ID.

Contoh CURL
curl -X GET http://localhost:8000/customers/1

4️⃣ PUT /customers/{id}

Mengupdate data customer.

Body (JSON)
{
  "nationality_id": 2,
  "cst_name": "Budi Update",
  "cst_dob": "1995-01-20",
  "cst_phoneNum": "0888888888",
  "cst_email": "budi.update@example.com"
}

Contoh CURL
curl -X PUT http://localhost:8000/customers/1 \
  -H "Content-Type: application/json" \
  -d '{
    "nationality_id": 2,
    "cst_name": "Budi Update",
    "cst_dob": "1995-01-20",
    "cst_phoneNum": "0888888888",
    "cst_email": "budi.update@example.com"
  }'

5️⃣ DELETE /customers/{id}

Menghapus customer berdasarkan ID.

Contoh CURL
curl -X DELETE http://localhost:8000/customers/1

▶ Cara Menjalankan Server
go mod tidy
go run ./cmd/server


Server berjalan di:

http://localhost:8000
