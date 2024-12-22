package handler

import (
	"encoding/json"
	"github.com/Ddarli/app/shop/config"
	"github.com/Ddarli/app/shop/internal/service"
	"github.com/Ddarli/app/shop/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
)

type (
	Handler interface {
		getProducts(w http.ResponseWriter, r *http.Request)
		InitRouter()
	}

	httpHandler struct {
		router       *mux.Router
		service      service.Service
		tokenService service.TokenService
	}
)

func NewHttpHandler(service service.Service, tokenService service.TokenService) Handler {
	return &httpHandler{
		service:      service,
		router:       mux.NewRouter(),
		tokenService: tokenService,
	}
}

func (h *httpHandler) InitRouter() {
	authMiddleware := AuthMiddleware(config.Key)

	h.router.HandleFunc("/api/v1/products", func(w http.ResponseWriter, r *http.Request) {
		authMiddleware(http.HandlerFunc(h.getProducts)).ServeHTTP(w, r)
	}).Methods("GET")
	h.router.HandleFunc("/api/v1/order", func(w http.ResponseWriter, r *http.Request) {
		authMiddleware(http.HandlerFunc(h.handleOrder)).ServeHTTP(w, r)
	}).Methods("POST")
	http.Handle("/", h.router)

	h.router.HandleFunc("/api/v1/token", h.generateToken).Methods("GET")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func (h *httpHandler) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (h *httpHandler) handleOrder(w http.ResponseWriter, r *http.Request) {
	var request models.OrderRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	response, err := h.service.ProcessOrder(r.Context(), request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *httpHandler) generateToken(w http.ResponseWriter, r *http.Request) {
	var userID = "1"

	token, err := h.tokenService.GenerateAccessToken(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
