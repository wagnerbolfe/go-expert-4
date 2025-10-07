package main

import (
	"net/http"

	"github.com/wagnerbolfe/goexpert/API/configs"
	"github.com/wagnerbolfe/goexpert/API/internal/entity"
	"github.com/wagnerbolfe/goexpert/API/internal/infra/database"
	"github.com/wagnerbolfe/goexpert/API/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
}
