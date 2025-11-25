package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Baihaki123/customer-service/internal/domain"
	"github.com/Baihaki123/customer-service/internal/usecase"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	UC usecase.CustomerUsecase
}

func NewCustomerHandler(r *mux.Router, uc usecase.CustomerUsecase) {
	h := &CustomerHandler{UC: uc}
	r.HandleFunc("/customers", h.Create).Methods("POST")
	r.HandleFunc("/customers", h.FindAll).Methods("GET")
	r.HandleFunc("/customers/{id}", h.FindByID).Methods("GET")
	r.HandleFunc("/customers/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/customers/{id}", h.Delete).Methods("DELETE")
}

func (h *CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c domain.Customer
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.UC.Create(&c); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func (h *CustomerHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	customers, err := h.UC.GetAll()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h *CustomerHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	customer, err := h.UC.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func (h *CustomerHandler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	var c domain.Customer
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.ID = id // pastikan ID update sesuai URL

	if err := h.UC.Update(&c); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func (h *CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.UC.Delete(id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// implement FindAll, FindByID, Update, Delete seperti contoh yang saya kirim sebelumnya
