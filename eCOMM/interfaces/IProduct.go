package interfaces

import "github.com/ArchiMrin/Project_GOLang/eCOMM/entities"

type IProduct interface {
	Insert(product *entities.Product) (string, error)
	GetProducts() ([]*entities.Product, error)
	GetProductById(id string) (*entities.Product, error)
}
