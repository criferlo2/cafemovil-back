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
			{ID: "prd_001", Name: "Origen Sierra Nevada", Brand: "Devoción", CategoryID: "coffee", CategoryName: "Café", Presentation: "Libra 500g", Price: new(float64(42900)), Currency: "COP", Featured: true},
			{ID: "prd_002", Name: "Café Especial Huila", Brand: "Amor Perfecto", CategoryID: "coffee", CategoryName: "Café", Presentation: "Kilogramo", Price: new(float64(79900)), Currency: "COP", Featured: true},
			{ID: "prd_003", Name: "Sachets Cold Brew", Brand: "Juan Valdez", CategoryID: "coffee", CategoryName: "Café", Presentation: "Caja x 10 sachets", Price: new(float64(29900)), Currency: "COP", Featured: true},
			{ID: "prd_004", Name: "Prensa Francesa Vidrio", Brand: "Bialetti", CategoryID: "accessories", CategoryName: "Accesorios", Presentation: "600ml", Price: new(float64(89900)), Currency: "COP", Featured: true},
			{ID: "prd_005", Name: "Cafetera Moka Express", Brand: "Bialetti", CategoryID: "machines", CategoryName: "Máquinas para hacer café", Presentation: "6 tazas", Price: new(float64(159900)), Currency: "COP", Featured: true},
			{ID: "prd_006", Name: "Manual de Catación Básica", Brand: "Café Lab", CategoryID: "manuals", CategoryName: "Manuales", Presentation: "PDF", Price: new(float64(19900)), Currency: "COP", Featured: true},
			{ID: "prd_007", Name: "Historias del Café Colombiano", Brand: "Editorial Andina", CategoryID: "books", CategoryName: "Libros", Presentation: "Tapa blanda", Price: new(float64(54900)), Currency: "COP", Featured: true},
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
