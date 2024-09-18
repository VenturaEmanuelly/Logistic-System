package controllers

import "sistemaLogistica/internal/entity"


type ProductRepository interface {
	Insert(info entity.Product) (entity.Product, error)
	Get(info int) (entity.Product, error)
	GetByName(name string) (entity.Product, error)
	Update(info entity.Product) (entity.Product, error)
	Delete(info entity.Product) error
}

type ProductHandler interface {
	Create(entity.Product) (entity.Product, error)
	Read(entity.Product) (entity.Product, error)
	Update(entity.Product) (entity.Product, error)
	Delete(entity.Product) error
}

type ProductUsecase interface {
	Create(info entity.Product) (entity.Product, error)
	GetMetrics(productName string) (entry int, out int, err error)
	Read(info entity.Product) (entity.Product, error)
	Update(info entity.Product) (entity.Product, error)
	Delete(info entity.Product) error
}

type TotalProductRepository interface {
	Upsert(entity.TotalProduct) (entity.TotalProduct, error)
	Get(string) (entity.TotalProduct, error)
}

type RepoSql interface {
	QueryRow(query string, any []any, dest ...any) error
}
