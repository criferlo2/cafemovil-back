package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"cafemovil-back/internal/catalog"
	"cafemovil-back/internal/user"
)

type Router struct {
	userService    *user.Service
	catalogService *catalog.Service
	mux            *http.ServeMux
}

func NewRouter(userService *user.Service, catalogService *catalog.Service) http.Handler {
	router := &Router{
		userService:    userService,
		catalogService: catalogService,
		mux:            http.NewServeMux(),
	}
	router.routes()
	return router
}

func (router *Router) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	router.mux.ServeHTTP(response, request)
}

func (router *Router) routes() {
	router.mux.HandleFunc("POST /api/v1/auth/register", router.register)
	router.mux.HandleFunc("POST /api/v1/auth/guest", router.guest)
	router.mux.HandleFunc("GET /api/v1/categories", router.categories)
	router.mux.HandleFunc("GET /api/v1/products/featured", router.featuredProducts)
	router.mux.HandleFunc("GET /health", router.health)
}

func (router *Router) register(response http.ResponseWriter, request *http.Request) {
	var payload struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}
	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		writeError(response, http.StatusBadRequest, "invalid JSON body")
		return
	}

	session, err := router.userService.Register(payload.Name, payload.Phone)
	if err != nil {
		if errors.Is(err, user.ErrInvalidUser) {
			writeError(response, http.StatusBadRequest, "name and phone are required")
			return
		}
		writeError(response, http.StatusInternalServerError, "could not create user")
		return
	}

	writeJSON(response, http.StatusCreated, session)
}

func (router *Router) guest(response http.ResponseWriter, request *http.Request) {
	writeJSON(response, http.StatusCreated, router.userService.Guest())
}

func (router *Router) categories(response http.ResponseWriter, request *http.Request) {
	writeJSON(response, http.StatusOK, map[string]any{"data": router.catalogService.Categories()})
}

func (router *Router) featuredProducts(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	products := router.catalogService.FeaturedProducts(catalog.ProductFilter{
		CategoryID: query.Get("category"),
		Brand:      query.Get("brand"),
		Search:     query.Get("search"),
	})
	writeJSON(response, http.StatusOK, map[string]any{"data": products})
}

func (router *Router) health(response http.ResponseWriter, request *http.Request) {
	writeJSON(response, http.StatusOK, map[string]string{"status": "ok"})
}

func writeJSON(response http.ResponseWriter, status int, body any) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	_ = json.NewEncoder(response).Encode(body)
}

func writeError(response http.ResponseWriter, status int, message string) {
	writeJSON(response, status, map[string]string{"error": message})
}
