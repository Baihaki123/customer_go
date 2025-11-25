package main

import (
	"log"
	stdhttp "net/http"

	"github.com/Baihaki123/customer-service/internal/config"
	delivery "github.com/Baihaki123/customer-service/internal/delivery/http"
	"github.com/Baihaki123/customer-service/internal/repository"
	"github.com/Baihaki123/customer-service/internal/usecase"

	"github.com/gorilla/mux"
)

func main() {
	db := config.ConnectDB()

	customerRepo := repository.NewCustomerRepository(db)
	customerUC := usecase.NewCustomerUsecase(customerRepo)

	r := mux.NewRouter()
	delivery.NewCustomerHandler(r, customerUC)

	log.Println("Server running at :8000")
	if err := stdhttp.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
