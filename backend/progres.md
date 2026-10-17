# Progres Backend & Database

Dokumen monitoring changelog dan progres database serta API backend.

## Changelog Tabel & Skema

### Inisialisasi Database (SQLite)
- **Driver**: `modernc.org/sqlite`
- **Path Database**: `backend/data/app.db`
- **File Skema**: [`backend/database/schema.sql`](file:///d:/workshop/backend/database/schema.sql)
- **File Seed**: [`backend/database/seed.sql`](file:///d:/workshop/backend/database/seed.sql)

### Tabel yang Dibuat
1. **`services`**:
   - `id`: INTEGER, PRIMARY KEY AUTOINCREMENT
   - `name`: TEXT, NOT NULL
   - `description`: TEXT, NULLABLE
   - `icon`: TEXT (nama icon Bootstrap Icons)
   - `image_url`: TEXT, NULLABLE (URL gambar layanan baik external link maupun path static server)
   - `created_at`: DATETIME, DEFAULT CURRENT_TIMESTAMP
   - **Jumlah Baris Seed**: 4 baris data layanan beserta `image_url`

---

## API Service CRUD & Penanganan Gambar

REST API dengan Gin telah disempurnakan:

- `GET /api/services` (200 OK - Mengembalikan array JSON services tanpa pembungkus)
- `GET /api/services/:id` (200 OK - Mengembalikan satu objek service; 404 jika tidak ditemukan)
- `POST /api/services` (201 Created - Mendukung JSON biasa dan `multipart/form-data`)
- `PUT /api/services/:id` (200 OK - Mendukung JSON biasa dan `multipart/form-data`)
- `DELETE /api/services/:id` (204 No Content - Menghapus data service; 404 jika tidak ditemukan)

### Penanganan `image_url` & Static Files:
1. **Dukungan URL Link & File Fisik**:
   - Menerima `image_url` berupa link web langsung (`http://...` atau `https://...`) melalui payload JSON maupun Form Data.
   - Menerima unggahan file gambar langsung (field `image`) via `multipart/form-data`. File disimpan di direktori `backend/uploads/` dengan nama unik.
2. **Static File Serving**:
   - Mengaktifkan route `r.Static("/uploads", "./uploads")` sehingga gambar yang diunggah dapat diakses publik melalui browser pada alamat `http://localhost:8080/uploads/<nama_file>`.
3. **Format Respons**:
   - Struktur JSON asli tetap utuh dan konsisten tanpa merusak endpoint CRUD yang sudah ada.
