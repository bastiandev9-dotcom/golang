## 1. Klien
[Nama usaha] adalah [jenis usaha]. Klien ingin mengelola daftar layanan sendiri lewat dashboard.

## 2. Struktur halaman
Landing page (/): Header, Hero, Tentang, Layanan, Portofolio, Kontak, Footer.
Bagian Layanan menampilkan kartu dari GET /api/services.

Dashboard (/dashboard):
- Tabel layanan dengan tombol tambah, ubah, hapus
- Form di modal Bootstrap, konfirmasi sebelum hapus
- Tanpa login

## 4. Data
Tabel: services
| Kolom       | Tipe     | Aturan                    |
| id          | INTEGER  | primary key, otomatis     |
| name        | TEXT     | wajib                     |
| description | TEXT     | boleh kosong              |
| icon        | TEXT     | nama ikon Bootstrap Icons |
| created_at  | DATETIME | otomatis                  |

## 5. API
| Method | Alamat            | Jawaban                  |
| GET    | /api/services     | 200, daftar layanan      |
| GET    | /api/services/:id | 200, satu layanan; 404   |
| POST   | /api/services     | 201, layanan baru; 400   |
| PUT    | /api/services/:id | 200, layanan diubah; 404 |
| DELETE | /api/services/:id | 204; 404                 |

Bentuk satu layanan:
{"id": 1, "name": "...", "description": "...", "icon": "globe", "created_at": "..."}
Daftar dikirim sebagai array JSON langsung, tanpa pembungkus. Port API: 8080.

## 6. Aturan teknis
- Frontend: Vue dan Bootstrap 5 di akar project
- Backend: Go dan Gin di folder backend, database SQLite dengan driver modernc.org/sqlite
- Pengaturan dibaca dari .env; .env tidak di-commit
- Dilarang: login, unggah gambar, tabel selain services

