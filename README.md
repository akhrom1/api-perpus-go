# Golang API PERPUS

API Sederhana untuk manajement sistem pinjam buku di perpustakaan

## Authentication

API ini menggunakan **JWT**.

Gunakan kredensial berikut untuk mengakses endpoint:

- **Username:** `admin`
- **Password:** `admin123`
- **JWT_SECRET:** `supersecretkey1234567890`

---

## Dokumentasi 

Postman 
👉 https://documenter.getpostman.com/view/21073811/2sBXVhEWev

Swaggerhub
👉 https://app.swaggerhub.com/apis-docs/nonm/api-perpus-go/1.0.0

---

## Folder Tree

```text
.
├── config
│   ├── database.go
│   └── migrate.go
├── go.mod
├── go.sum
├── internal
│   ├── handlers
│   │   ├── auth_handler.go
│   │   ├── book_handler.go
│   │   ├── category_handler.go
│   │   ├── load_handler.go
│   │   └── member_handler.go
│   ├── middleware
│   │   └── jwt_auth.go
│   ├── models
│   │   ├── book.go
│   │   ├── category.go
│   │   ├── item_loan.go
│   │   ├── loan.go
│   │   └── member.go
│   ├── repositories
│   │   ├── book_repo.go
│   │   ├── category_repo.go
│   │   ├── load_repo.go
│   │   └── member_repo.go
│   └── services
│       ├── book_service.go
│       ├── category_sevice.go
│       ├── load_service.go
│       └── member_service.go
├── main.go
├── migrations
│   ├── 001_create_categories.sql
│   ├── 002_create_books.sql
│   ├── 003_create_members.sql
│   ├── 004_create_loans.sql
│   └── 005_create_loan_items.sql
└── routes
    └── api.go

```
---

## Daftar Route API

### Auth

| Method | Endpoint  | Deskripsi        |
| ------ | --------- | ---------------- |
| POST   | `/login`  | Login  |


### Categories

| Method | Endpoint                  | Deskripsi                         |
| ------ | ------------------------- | --------------------------------- |
| GET    | `/api/categories`         | Mengambil semua kategori          |
| POST   | `/api/categories`         | Menambahkan kategori baru         |
| PUT    | `/api/categories/:id`     | Mengubah kategori berdasarkan ID |
| DELETE | `/api/categories/:id`     | Menghapus kategori                |


### Books

| Method | Endpoint              | Deskripsi                      |
| ------ | --------------------- | ------------------------------ |
| GET    | `/api/books`          | Mengambil semua buku           |
| POST   | `/api/books`          | Menambahkan buku baru          |
| GET    | `/api/books/:id`      | Detail buku berdasarkan ID     |
| PUT    | `/api/books/:id`      | Mengubah data buku             |
| DELETE | `/api/books/:id`      | Menghapus buku                 |


### Members

| Method | Endpoint                 | Deskripsi                        |
| ------ | ------------------------ | -------------------------------- |
| GET    | `/api/members`           | Mengambil semua anggota          |
| POST   | `/api/members`           | Menambahkan anggota baru         |
| GET    | `/api/members/:id`       | Detail anggota berdasarkan ID    |
| PUT    | `/api/members/:id`       | Mengubah data anggota            |
| DELETE | `/api/members/:id`       | Menghapus anggota                |


### Loans

| Method | Endpoint                     | Deskripsi                                  |
| ------ | ---------------------------- | ------------------------------------------ |
| GET    | `/api/loans`                 | Mengambil semua data peminjaman            |
| GET    | `/api/loans/:id`             | Detail peminjaman berdasarkan ID           |
| POST   | `/api/loans`                 | Membuat transaksi peminjaman baru          |
| PUT    | `/api/loans/:id/return`      | Mengembalikan buku (menyelesaikan loan)    |

---


