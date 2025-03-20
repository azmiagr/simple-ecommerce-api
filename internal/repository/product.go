package repository

import (
	"golang-ecommerce/entity"
	"golang-ecommerce/model"

	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(product *entity.Product) (*entity.Product, error)
	GetAllProducts(limit, offset int) ([]*model.GetAllProducts, error)
	GetProductByName(produtName string, limit, offset int) ([]*model.SearchProduct, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{db}
}

func (p *ProductRepository) CreateProduct(product *entity.Product) (*entity.Product, error) {
	err := p.db.Debug().Create(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductRepository) GetAllProducts(limit, offset int) ([]*model.GetAllProducts, error) {
	var products []*model.GetAllProducts

	err := p.db.Debug().Table("products").Select("products.product_name, products.price, stores.store_name").Joins("JOIN stores ON products.store_id = stores.store_id").Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductRepository) GetProductByName(produtName string, limit, offset int) ([]*model.SearchProduct, error) {
	var product []*model.SearchProduct
	searchQuery := "%" + produtName + "%"
	err := p.db.Debug().Table("products").Select("products.product_name, products.price, stores.store_name").Joins("JOIN stores ON products.store_id = stores.store_id").Where("products.product_name LIKE ?", searchQuery).Limit(limit).Offset(offset).Find(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}
