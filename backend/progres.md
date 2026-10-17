# Progres Backend & Database

Dokumen monitoring changelog dan progres database backend.

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
   - `image_url`: TEXT, NULLABLE (URL gambar layanan)
   - `created_at`: DATETIME, DEFAULT CURRENT_TIMESTAMP
   - **Jumlah Baris Seed**: 4 baris data layanan beserta `image_url`

> [!NOTE]
> Sesuai instruksi dan aturan ketat, tidak ada tabel lain yang dibuat selain tabel `services`. Database telah di-reinit dan memuat 4 data seed lengkap dengan `image_url`.
