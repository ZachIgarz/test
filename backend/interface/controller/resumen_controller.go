package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ZachIgarz/test-api-rest/infrastructure/entities"
)

type purchaseResume struct {
	purchaseInteractor interactor.Purchase
}

type PurchaseController interface {
	Resumen(c Context) error
	//CreateUser(c Context) error
}

func NewPurchaseController(purchaseInteractor interactor.Purchase) PurchaseController {
	return &purchaseResume{
		purchaseInteractor: purchaseInteractor,
	}
}

func (purchaseResume *purchaseResume) Resumen(c Context) error {

	date := c.Get("date")
	days := c.Get("days")

	//TODO: refactor this
	return c.JSON(http.StatusOK, nil)
}

/*Init ...*/
func (purchaseResume *PurchaseResume) Init(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	date := strings.Split(path, "/")
	realDate := date[2]
	dias := r.URL.Query().Get("dias")

	purchaseResumeRequest := entities.NewPurchaseResumeRequest(realDate, dias)

	statistics, error := purchaseResume.purchasesUseCase.Handler(*purchaseResumeRequest)

	if error != nil {
		http.Error(w, "an error has occurred trying to get the statistics ", http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(statistics)

}
