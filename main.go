package main

import (
	"fmt"
	"go_hari_7/database"
	"go_hari_7/router"
	"go_hari_7/todo"
	"go_hari_7/users"
	"log"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Gagal terhubung ke database: ", err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	fmt.Println("✅ Terhubung ke database PostgreSQL")
	// AutoMigrate: membuat tabel todos jika belum ada, atau menambah kolom yang kurang
	if err := db.AutoMigrate(&todo.Todo{}, &users.User{}); err != nil {
		log.Fatal("Gagal migrasi database: ", err)
	}
	r := router.SetupRouter(db)
	fmt.Println("🚀 Server Gin berjalan di http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}

}
