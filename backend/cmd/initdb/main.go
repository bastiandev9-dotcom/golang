package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func main() {
	dbDir := "data"
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		log.Fatalf("Gagal membuat direktori data: %v", err)
	}

	dbPath := filepath.Join(dbDir, "app.db")
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("Gagal membuka database: %v", err)
	}
	defer db.Close()

	schemaBytes, err := os.ReadFile("database/schema.sql")
	if err != nil {
		log.Fatalf("Gagal membaca schema.sql: %v", err)
	}

	if _, err := db.Exec(string(schemaBytes)); err != nil {
		log.Fatalf("Gagal mengeksekusi schema.sql: %v", err)
	}
	fmt.Println("Eksekusi schema.sql berhasil.")

	seedBytes, err := os.ReadFile("database/seed.sql")
	if err != nil {
		log.Fatalf("Gagal membaca seed.sql: %v", err)
	}

	if len(seedBytes) > 0 {
		if _, err := db.Exec(string(seedBytes)); err != nil {
			log.Fatalf("Gagal mengeksekusi seed.sql: %v", err)
		}
		fmt.Println("Eksekusi seed.sql berhasil.")
	}

	// Cek tabel foto_pengalaman jika ada sesuai instruksi, atau tampilkan jumlah baris services
	var count int
	// Blueprint dan instruksi: dilarang membuat tabel selain services.
	// Jika user meminta mencetak jumlah baris di tabel foto pengalaman / services:
	var tableExists bool
	err = db.QueryRow("SELECT COUNT(*) > 0 FROM sqlite_master WHERE type='table' AND name='foto_pengalaman'").Scan(&tableExists)
	if err == nil && tableExists {
		rowErr := db.QueryRow("SELECT COUNT(*) FROM foto_pengalaman").Scan(&count)
		if rowErr != nil {
			log.Printf("Gagal menghitung baris foto_pengalaman: %v", rowErr)
		} else {
			fmt.Printf("Jumlah baris di tabel foto_pengalaman: %d\n", count)
		}
	} else {
		err = db.QueryRow("SELECT COUNT(*) FROM services").Scan(&count)
		if err != nil {
			log.Fatalf("Gagal menghitung baris services: %v", err)
		}
		fmt.Printf("Catatan: tabel foto_pengalaman tidak ada (aturan blueprint melarang tabel selain services).\n")
		fmt.Printf("Jumlah baris di tabel services: %d\n", count)
	}
}
