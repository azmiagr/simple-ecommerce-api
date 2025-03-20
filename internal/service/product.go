package service

import (
	"golang-ecommerce/entity"
	"golang-ecommerce/internal/repository"
	"golang-ecommerce/model"
)

type IProductService interface {
	CreateProduct(param *model.AddProductRequest) (*entity.Product, error)
	GetAllProducts(page int) ([]*model.GetAllProducts, error)
	GetProductByName(productName string, page int) ([]*model.SearchProduct, error)
}

type ProductService struct {
	ProductRepository repository.IProductRepository
}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (ps *ProductService) CreateProduct(param *model.AddProductRequest) (*entity.Product, error) {
	product := &entity.Product{
		ProductName:        param.ProductName,
		ProductDescription: param.ProductDescription,
		Price:              param.Price,
		Stock:              param.Stock,
		StoreID:            param.StoreID,
	}

	product, err := ps.ProductRepository.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductService) GetAllProducts(page int) ([]*model.GetAllProducts, error) {
	limit := 6
	offset := (page - 1) * limit

	products, err := ps.ProductRepository.GetAllProducts(limit, offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (ps *ProductService) GetProductByName(productName string, page int) ([]*model.SearchProduct, error) {
	limit := 6
	offset := (page - 1) * limit
	product, err := ps.ProductRepository.GetProductByName(productName, limit, offset)
	if err != nil {
		return nil, err
	}

	return product, nil
}
