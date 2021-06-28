package get

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ZachIgarz/golangIpCom/application"
	"github.com/ZachIgarz/golangIpCom/infrastructure/entities"
)

/*PurchaseResume ... */
type PurchaseResume struct {
	purchasesUseCase application.PurchasesUseCase
}

/*NewPurchaseResume ...*/
func NewPurchaseResume(
	purchasesUseCase application.PurchasesUseCase) *PurchaseResume {
	return &PurchaseResume{
		purchasesUseCase: purchasesUseCase,
	}
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
