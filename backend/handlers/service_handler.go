package handlers

import (
	"database/sql"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"backend/models"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	DB *sql.DB
}

func NewServiceHandler(db *sql.DB) *ServiceHandler {
	return &ServiceHandler{DB: db}
}

// saveUploadedFile menyimpan file gambar ke direktori uploads dan mengembalikan URL publiknya
func (h *ServiceHandler) saveUploadedFile(c *gin.Context, file *multipart.FileHeader) (string, error) {
	if err := os.MkdirAll("uploads", 0755); err != nil {
		return "", err
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext == "" {
		ext = ".jpg"
	}
	newFilename := fmt.Sprintf("%d_%d%s", time.Now().UnixNano(), file.Size, ext)
	dst := filepath.Join("uploads", newFilename)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		return "", err
	}

	// Buat URL yang dapat diakses publik oleh browser
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	baseURL := fmt.Sprintf("%s://%s", scheme, c.Request.Host)
	fileURL := fmt.Sprintf("%s/uploads/%s", baseURL, newFilename)
	return fileURL, nil
}

// GET /api/services
func (h *ServiceHandler) GetServices(c *gin.Context) {
	rows, err := h.DB.Query("SELECT id, name, COALESCE(description, ''), COALESCE(icon, ''), image_url, created_at FROM services ORDER BY id ASC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	services := make([]models.Service, 0)
	for rows.Next() {
		var s models.Service
		if err := rows.Scan(&s.ID, &s.Name, &s.Description, &s.Icon, &s.ImageURL, &s.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		services = append(services, s)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, services)
}

// GET /api/services/:id
func (h *ServiceHandler) GetServiceByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Layanan tidak ditemukan"})
		return
	}

	var s models.Service
	query := "SELECT id, name, COALESCE(description, ''), COALESCE(icon, ''), image_url, created_at FROM services WHERE id = ?"
	err = h.DB.QueryRow(query, id).Scan(&s.ID, &s.Name, &s.Description, &s.Icon, &s.ImageURL, &s.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Layanan tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, s)
}

// POST /api/services
// Mendukung format JSON biasa (image_url berupa URL teks) maupun multipart/form-data (upload file gambar atau URL)
func (h *ServiceHandler) CreateService(c *gin.Context) {
	var name, description, icon string
	var imageURL *string

	contentType := c.ContentType()

	if strings.Contains(contentType, "multipart/form-data") {
		name = c.PostForm("name")
		description = c.PostForm("description")
		icon = c.PostForm("icon")

		if imgText := c.PostForm("image_url"); imgText != "" {
			imageURL = &imgText
		}

		// Periksa jika ada file gambar yang diunggah
		file, err := c.FormFile("image")
		if err == nil && file != nil {
			uploadedURL, saveErr := h.saveUploadedFile(c, file)
			if saveErr != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file gambar"})
				return
			}
			imageURL = &uploadedURL
		}

		if strings.TrimSpace(name) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Field 'name' wajib diisi"})
			return
		}
	} else {
		// Default JSON body
		var input models.CreateServiceInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Field 'name' wajib diisi"})
			return
		}
		name = input.Name
		description = input.Description
		icon = input.Icon
		imageURL = input.ImageURL
	}

	query := "INSERT INTO services (name, description, icon, image_url) VALUES (?, ?, ?, ?) RETURNING id, name, COALESCE(description, ''), COALESCE(icon, ''), image_url, created_at"
	var created models.Service
	err := h.DB.QueryRow(query, name, description, icon, imageURL).Scan(
		&created.ID,
		&created.Name,
		&created.Description,
		&created.Icon,
		&created.ImageURL,
		&created.CreatedAt,
	)

	if err != nil {
		res, execErr := h.DB.Exec("INSERT INTO services (name, description, icon, image_url) VALUES (?, ?, ?, ?)", name, description, icon, imageURL)
		if execErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": execErr.Error()})
			return
		}
		lastID, _ := res.LastInsertId()
		_ = h.DB.QueryRow("SELECT id, name, COALESCE(description, ''), COALESCE(icon, ''), image_url, created_at FROM services WHERE id = ?", lastID).Scan(
			&created.ID,
			&created.Name,
			&created.Description,
			&created.Icon,
			&created.ImageURL,
			&created.CreatedAt,
		)
	}

	c.JSON(http.StatusCreated, created)
}

// PUT /api/services/:id
// Mendukung format JSON biasa maupun multipart/form-data
func (h *ServiceHandler) UpdateService(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Layanan tidak ditemukan"})
		return
	}

	// Cek apakah service ada
	var current models.Service
	checkQuery := "SELECT id, name, COALESCE(description, ''), COALESCE(icon, ''), image_url, created_at FROM services WHERE id = ?"
	err = h.DB.QueryRow(checkQuery, id).Scan(&current.ID, &current.Name, &current.Description, &current.Icon, &current.ImageURL, &current.CreatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Layanan tidak ditemukan"})
		return
	}

	var name, description, icon string
	var imageURL *string = current.ImageURL

	contentType := c.ContentType()

	if strings.Contains(contentType, "multipart/form-data") {
		name = c.PostForm("name")
		description = c.PostForm("description")
		icon = c.PostForm("icon")

		if imgText, exists := c.GetPostForm("image_url"); exists {
			imageURL = &imgText
		}

		// Periksa jika ada file gambar yang diunggah
		file, err := c.FormFile("image")
		if err == nil && file != nil {
			uploadedURL, saveErr := h.saveUploadedFile(c, file)
			if saveErr != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file gambar"})
				return
			}
			imageURL = &uploadedURL
		}

		if strings.TrimSpace(name) == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Field 'name' wajib diisi"})
			return
		}
	} else {
		var input models.UpdateServiceInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Field 'name' wajib diisi"})
			return
		}
		name = input.Name
		description = input.Description
		icon = input.Icon
		imageURL = input.ImageURL
	}

	query := "UPDATE services SET name = ?, description = ?, icon = ?, image_url = ? WHERE id = ?"
	_, err = h.DB.Exec(query, name, description, icon, imageURL, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updated models.Service
	err = h.DB.QueryRow("SELECT id, name, COALESCE(description, ''), COALESCE(icon, ''), image_url, created_at FROM services WHERE id = ?", id).Scan(
		&updated.ID,
		&updated.Name,
		&updated.Description,
		&updated.Icon,
		&updated.ImageURL,
		&updated.CreatedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// DELETE /api/services/:id
func (h *ServiceHandler) DeleteService(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Layanan tidak ditemukan"})
		return
	}

	res, err := h.DB.Exec("DELETE FROM services WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Layanan tidak ditemukan"})
		return
	}

	c.Status(http.StatusNoContent)
}
