package catalog

import "strings"

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Brand        string  `json:"brand"`
	CategoryID   string  `json:"categoryId"`
	CategoryName string  `json:"categoryName"`
	Presentation string  `json:"presentation"`
	Price        float64 `json:"price"`
	Featured     bool    `json:"featured"`
}

type ProductFilter struct {
	CategoryID string
	Brand      string
	Search     string
}

type Repository interface {
	Categories() []Category
	FeaturedProducts(filter ProductFilter) []Product
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service *Service) Categories() []Category {
	return service.repository.Categories()
}

func (service *Service) FeaturedProducts(filter ProductFilter) []Product {
	filter.CategoryID = strings.TrimSpace(filter.CategoryID)
	filter.Brand = strings.TrimSpace(filter.Brand)
	filter.Search = strings.TrimSpace(filter.Search)
	return service.repository.FeaturedProducts(filter)
}
