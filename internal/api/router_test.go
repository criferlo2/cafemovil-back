package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"cafemovil-back/internal/api"
	"cafemovil-back/internal/catalog"
	"cafemovil-back/internal/user"
)

func newTestServer() http.Handler {
	userService := user.NewService(user.NewMemoryRepository())
	catalogService := catalog.NewService(catalog.NewSeedRepository())
	return api.NewRouter(userService, catalogService)
}

func decodeJSON(t *testing.T, response *httptest.ResponseRecorder) map[string]any {
	t.Helper()
	var body map[string]any
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("response body must be valid JSON: %v; body=%s", err, response.Body.String())
	}
	return body
}

func TestRegisterCreatesUserSessionWithNameAndPhone(t *testing.T) {
	server := newTestServer()
	requestBody := []byte(`{"name":"Ana Cafetera","phone":"3001234567"}`)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d with body %s", response.Code, response.Body.String())
	}
	body := decodeJSON(t, response)
	if body["token"] == "" {
		t.Fatalf("expected session token in response: %#v", body)
	}
	userBody, ok := body["user"].(map[string]any)
	if !ok {
		t.Fatalf("expected user object in response: %#v", body)
	}
	if userBody["name"] != "Ana Cafetera" || userBody["phone"] != "3001234567" || userBody["guest"] != false {
		t.Fatalf("unexpected registered user: %#v", userBody)
	}
}

func TestRegisterRejectsMissingRequiredFields(t *testing.T) {
	server := newTestServer()
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewReader([]byte(`{"name":"Ana"}`)))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d with body %s", response.Code, response.Body.String())
	}
	body := decodeJSON(t, response)
	if body["error"] == "" {
		t.Fatalf("expected validation error response: %#v", body)
	}
}

func TestGuestLoginCreatesGuestSession(t *testing.T) {
	server := newTestServer()
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/guest", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d with body %s", response.Code, response.Body.String())
	}
	body := decodeJSON(t, response)
	userBody, ok := body["user"].(map[string]any)
	if !ok {
		t.Fatalf("expected user object in response: %#v", body)
	}
	if userBody["guest"] != true || userBody["name"] != "Guest" || body["token"] == "" {
		t.Fatalf("unexpected guest session: %#v", body)
	}
}

func TestCategoriesReturnsCoffeeFirst(t *testing.T) {
	server := newTestServer()
	request := httptest.NewRequest(http.MethodGet, "/api/v1/categories", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", response.Code, response.Body.String())
	}
	var body struct {
		Data []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"data"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("response body must be valid JSON: %v", err)
	}
	if len(body.Data) < 5 {
		t.Fatalf("expected at least five categories, got %#v", body.Data)
	}
	if body.Data[0].Name != "Café" {
		t.Fatalf("expected Café as first category, got %#v", body.Data[0])
	}
}

func TestFeaturedProductsIncludeRequiredFieldsAndSupportFilters(t *testing.T) {
	server := newTestServer()
	request := httptest.NewRequest(http.MethodGet, "/api/v1/products/featured?category=coffee&brand=devocion&search=origen", nil)
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d with body %s", response.Code, response.Body.String())
	}
	var body struct {
		Data []struct {
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Brand        string  `json:"brand"`
			CategoryID   string  `json:"categoryId"`
			CategoryName string  `json:"categoryName"`
			Presentation string  `json:"presentation"`
			Price        float64 `json:"price"`
			Featured     bool    `json:"featured"`
		} `json:"data"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("response body must be valid JSON: %v", err)
	}
	if len(body.Data) == 0 {
		t.Fatalf("expected filtered products, got none")
	}
	for _, product := range body.Data {
		if product.ID == "" || product.Name == "" || product.Brand == "" || product.CategoryID == "" || product.CategoryName == "" || product.Presentation == "" || product.Price <= 0 {
			t.Fatalf("product is missing required v1 fields: %#v", product)
		}
		if product.CategoryID != "coffee" || product.Brand != "Devoción" || !product.Featured {
			t.Fatalf("product does not respect filters/featured rule: %#v", product)
		}
	}
}
