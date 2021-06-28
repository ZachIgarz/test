package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ZachIgarz/golangIpCom/application"
	"github.com/ZachIgarz/golangIpCom/domain/ports"
	"github.com/ZachIgarz/golangIpCom/infrastructure/restclients"

	"github.com/ZachIgarz/golangIpCom/infrastructure/controllers/get"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	routes()
}

func routes() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	//Permisos a cualquiera
	handler := cors.AllowAll().Handler(router)

	router.HandleFunc("/resumen/{clave}", get.NewPurchaseResume(getPurchasesUseCase()).Init).Methods("GET")

	//Escucha el puerto para ver las peticiones
	//Agrega el puerto a la url
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
func getPurchasesUseCase() application.PurchasesUseCase {
	return application.NewPurchasesApplication(getPurchasesClient())
}

func getPurchasesClient() ports.PurchasesClient {
	return restclients.PurchaseRestClient{}
}
