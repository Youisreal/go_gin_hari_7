package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go_hari_7/todo"
)

// SetupRouter membuat router Gin lengkap dengan semua route.
// Dipakai oleh main.go dan oleh test.
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	todo.RegisterRoutes(v1, &todo.Handler{DB: db})
	return r
}
