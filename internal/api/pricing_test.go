package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductListingsExposeUnitPriceAndCurrency(t *testing.T) {
	server := newTestServer()
	for _, path := range []string{
		"/api/v1/products/featured",
		"/api/v1/products/featured?search=origen",
		"/api/v1/products/featured?category=coffee",
	} {
		t.Run(path, func(t *testing.T) {
			response := httptest.NewRecorder()
			server.ServeHTTP(response, httptest.NewRequest(http.MethodGet, path, nil))
			if response.Code != http.StatusOK {
				t.Fatalf("status = %d", response.Code)
			}
			products := decodeJSON(t, response)["data"].([]any)
			if len(products) == 0 {
				t.Fatal("expected products")
			}
			for _, item := range products {
				product := item.(map[string]any)
				if product["currency"] != "COP" {
					t.Fatalf("expected explicit COP currency: %v", product)
				}
				price, ok := product["price"].(float64)
				if !ok || price <= 0 {
					t.Fatalf("expected numeric unit price: %v", product)
				}
				if product["id"] == "prd_001" && price != 42900 {
					t.Fatalf("unit price changed: %v", product)
				}
			}
		})
	}
}
