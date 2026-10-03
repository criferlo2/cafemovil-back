package catalog

import "strings"

type SeedRepository struct {
	categories []Category
	products   []Product
}

func NewSeedRepository() *SeedRepository {
	categories := []Category{
		{ID: "coffee", Name: "Café"},
		{ID: "accessories", Name: "Accesorios"},
		{ID: "machines", Name: "Máquinas para hacer café"},
		{ID: "manuals", Name: "Manuales"},
		{ID: "books", Name: "Libros"},
	}

	return &SeedRepository{
		categories: categories,
		products: []Product{
			{ID: "prd_001", Name: "Origen Sierra Nevada", Brand: "Devoción", CategoryID: "coffee", CategoryName: "Café", Presentation: "Libra 500g", Price: 42900, Featured: true},
			{ID: "prd_002", Name: "Café Especial Huila", Brand: "Amor Perfecto", CategoryID: "coffee", CategoryName: "Café", Presentation: "Kilogramo", Price: 79900, Featured: true},
			{ID: "prd_003", Name: "Sachets Cold Brew", Brand: "Juan Valdez", CategoryID: "coffee", CategoryName: "Café", Presentation: "Caja x 10 sachets", Price: 29900, Featured: true},
			{ID: "prd_004", Name: "Prensa Francesa Vidrio", Brand: "Bialetti", CategoryID: "accessories", CategoryName: "Accesorios", Presentation: "600ml", Price: 89900, Featured: true},
			{ID: "prd_005", Name: "Cafetera Moka Express", Brand: "Bialetti", CategoryID: "machines", CategoryName: "Máquinas para hacer café", Presentation: "6 tazas", Price: 159900, Featured: true},
			{ID: "prd_006", Name: "Manual de Catación Básica", Brand: "Café Lab", CategoryID: "manuals", CategoryName: "Manuales", Presentation: "PDF", Price: 19900, Featured: true},
			{ID: "prd_007", Name: "Historias del Café Colombiano", Brand: "Editorial Andina", CategoryID: "books", CategoryName: "Libros", Presentation: "Tapa blanda", Price: 54900, Featured: true},
		},
	}
}

func (repository *SeedRepository) Categories() []Category {
	return append([]Category(nil), repository.categories...)
}

func (repository *SeedRepository) FeaturedProducts(filter ProductFilter) []Product {
	products := make([]Product, 0)
	for _, product := range repository.products {
		if !product.Featured {
			continue
		}
		if filter.CategoryID != "" && product.CategoryID != filter.CategoryID {
			continue
		}
		if filter.Brand != "" && normalize(product.Brand) != normalize(filter.Brand) {
			continue
		}
		if filter.Search != "" && !containsFold(product.Name, filter.Search) && !containsFold(product.Brand, filter.Search) && !containsFold(product.Presentation, filter.Search) {
			continue
		}
		products = append(products, product)
	}
	return products
}

func containsFold(value string, search string) bool {
	return strings.Contains(normalize(value), normalize(search))
}

func normalize(value string) string {
	replacer := strings.NewReplacer(
		"á", "a", "é", "e", "í", "i", "ó", "o", "ú", "u", "ü", "u", "ñ", "n",
		"Á", "a", "É", "e", "Í", "i", "Ó", "o", "Ú", "u", "Ü", "u", "Ñ", "n",
	)
	return replacer.Replace(strings.ToLower(value))
}
