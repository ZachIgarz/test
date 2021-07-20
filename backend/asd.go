package sas

import (
	"log"
	"net/http"
	"os"

	"github.com/ZachIgarz/test-api-rest/application"
	"github.com/ZachIgarz/test-api-rest/domain/ports"
	"github.com/ZachIgarz/test-api-rest/infrastructure/restclients"

	"github.com/ZachIgarz/test-api-rest/infrastructure/controllers/get"
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
