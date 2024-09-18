package product

import (
	"errors"
	"sistemaLogistica/internal/controllers"
	"sistemaLogistica/internal/entity"
)

type product struct {
	productRepository      controllers.ProductRepository
	totalProductRepository controllers.TotalProductRepository
}

func (p product) Create(info entity.Product) (entity.Product, error) {
	err := p.validations(info)
	if err != nil {
		return entity.Product{}, err
	}

	createdProduct, err := p.productRepository.Insert(info)
	if err != nil {
		return entity.Product{}, err
	}

	inform := entity.TotalProduct{Name: info.Item}

	_, err = p.totalProductRepository.Upsert(inform)
	if err != nil {
		return entity.Product{}, err
	}

	return createdProduct, nil
}

func (p product) validations(info entity.Product) error {
	if info.ClientId <= 0 {
		return errors.New("invalid id")
	}

	if info.Item == "" {
		return errors.New("invalid name")
	}

	if info.Quantity <= 0 {
		return errors.New("invalid quantity")
	}

	if info.Kind != "in" && info.Kind != "out" {
		return errors.New("invalid movement type, use 'in' or 'out'")
	}

	if info.Date.IsZero() {
		return errors.New("invalid date")
	}

	return nil
}

func (p product) Read(info entity.Product) (entity.Product, error) {

	product, err := p.productRepository.Get(info.ClientId)
	if err != nil {
		return entity.Product{}, err
	}

	inform := entity.TotalProduct{Name: info.Item}

	_, err = p.totalProductRepository.Get(inform.Name)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil

}

func (p product) GetMetrics(productName string) (entry int, out int, err error) {
	totalProduct, err := p.totalProductRepository.Get(productName)
	if err != nil {
		return
	}

	product, err := p.productRepository.GetByName(productName)
	if err != nil {
		return
	}

	return totalProduct.Total, totalProduct.Total - product.Quantity, err
}

func (p product) Update(info entity.Product) (entity.Product, error) {
	return p.productRepository.Update(info)
}

func (p product) Delete(info entity.Product) error {
	return p.productRepository.Delete(info)
}

func NewProductUseCase(productRepo controllers.ProductRepository, totalProductsRepo controllers.TotalProductRepository) controllers.ProductUsecase {
	return product{productRepository: productRepo, totalProductRepository: totalProductsRepo}
}
