package users

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Handler menyimpan dependensi yang dibutuhkan semua handler user
type Handler struct {
	DB *gorm.DB
}

// respondError mengirim pesan error dengan format seragam {"error": "..."}
func respondError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

// validationMessage mengubah error binding menjadi pesan yang ramah
func validationMessage(err error) string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) && len(ve) > 0 {
		fe := ve[0]
		switch fe.Tag() {
		case "required":
			return fe.Field() + " wajib diisi"
		case "min":
			return fe.Field() + " minimal " + fe.Param() + " karakter"
		}
		return fe.Field() + " tidak valid"
	}
	return "Format JSON tidak valid"
}

// parseID mengambil path param :id dan mengubahnya menjadi angka
func parseID(c *gin.Context) (string, bool) {
	id := c.Param("id")
	if id == "" {
		respondError(c, http.StatusBadRequest, "ID harus tepat")
		return "", false
	}
	return id, true
}

func (h *Handler) List(c *gin.Context) {
	query := h.DB.WithContext(c.Request.Context()).Model(&User{})
	// Pagination: nilai tidak valid diganti nilai bawaan
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if limit < 1 || limit > 100 {
		limit = 10
	}
	// Session(&gorm.Session{}) membuat salinan aman agar query yang sama
	// bisa dipakai dua kali (Count lalu Find) tanpa saling memengaruhi
	var total int64
	if err := query.Session(&gorm.Session{}).Count(&total).Error; err != nil {
		log.Println("List count:", err)
		respondError(c, http.StatusInternalServerError, "Gagal mengambil data")
		return
	}
	users := []User{} // slice kosong agar hasilnya [] (bukan null)
	err := query.Order("id").Limit(limit).Offset((page - 1) * limit).Find(&users).Error
	if err != nil {
		log.Println("List find:", err)
		respondError(c, http.StatusInternalServerError, "Gagal mengambil data")
		return
	}
	c.Header("X-Total-Count", strconv.FormatInt(total, 10))
	c.JSON(http.StatusOK, users)
}

// GetByID - GET /todos/:id
func (h *Handler) GetByID(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var u User
	err := h.DB.WithContext(c.Request.Context()).First(&u, "id=?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		respondError(c, http.StatusNotFound, "User tidak ditemukan")
		return
	}
	if err != nil {
		log.Println("GetByID:", err)
		respondError(c, http.StatusInternalServerError, "Gagal mengambil data")
		return
	}
	c.JSON(http.StatusOK, u)
}

// Create - POST /todos
func (h *Handler) Create(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, http.StatusBadRequest, validationMessage(err))
		return
	}
	now := time.Now()
	// ID diisi otomatis oleh GORM setelah INSERT berhasil
	u := User{Fullname: req.Fullname, Email: req.Email, Password: req.Password, Created_at: now, Updated_at: now}
	if err := h.DB.WithContext(c.Request.Context()).Create(&u).Error; err != nil {
		log.Println("Create:", err)
		respondError(c, http.StatusInternalServerError, "Gagal menyimpan data")
		return
	}
	c.JSON(http.StatusCreated, u)
}

// Update - PUT /todos/:id
func (h *Handler) Update(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondError(c, http.StatusBadRequest, validationMessage(err))
		return
	}
	db := h.DB.WithContext(c.Request.Context())
	var u User
	if err := db.First(&u, "id=?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			respondError(c, http.StatusNotFound, "User tidak ditemukan")
			return
		}
		log.Println("Update find:", err)
		respondError(c, http.StatusInternalServerError, "Gagal mengubah data")
		return
	}
	now := time.Now()
	// Pakai map agar is_done = false tetap ikut diubah
	// (Updates dengan struct mengabaikan field bernilai nol)
	err := db.Model(&u).Updates(map[string]any{
		"fullname":   req.Fullname,
		"email":      req.Email,
		"password":   req.Password,
		"created_at": req.Created_at,
		"updated_at": now,
	}).Error
	if err != nil {
		log.Println("Update:", err)
		respondError(c, http.StatusInternalServerError, "Gagal mengubah data")
		return
	}
	c.JSON(http.StatusOK, u)
}

// Delete - DELETE /todos/:id
func (h *Handler) Delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	result := h.DB.WithContext(c.Request.Context()).Delete(&User{}, "id=?", id)
	if result.Error != nil {
		log.Println("Delete:", result.Error)
		respondError(c, http.StatusInternalServerError, "Gagal menghapus data")
		return
	}
	// RowsAffected = jumlah baris yang terhapus; 0 berarti ID tidak ada
	if result.RowsAffected == 0 {
		respondError(c, http.StatusNotFound, "User tidak ditemukan")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus"})
}
