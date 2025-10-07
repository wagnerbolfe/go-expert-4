package database

import "github.com/wagnerbolfe/goexpert/API/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emaiId string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
