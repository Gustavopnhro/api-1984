package entity

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrInvalidId              = errors.New("invalid identifier")
	ErrNameIsRequired         = errors.New("name is required")
	ErrPriceMustBePositive    = errors.New("price must be positive")
	ErrQuantityMustBePositive = errors.New("quantity must be positive when")
)

type Item struct {
	ID       ID     `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Category string `json:"category"`
	gorm.Model
}

func CreateItem(name string, price int, quantity int, category string) (*Item, error) {
	item := Item{
		ID:       NewID(),
		Name:     name,
		Price:    price,
		Quantity: quantity,
		Category: category,
	}

	return &item, nil
}
