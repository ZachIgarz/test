package application

import (
	domainEntities "github.com/ZachIgarz/golangIpCom/domain/entities"
	"github.com/ZachIgarz/golangIpCom/domain/ports"
	"github.com/ZachIgarz/golangIpCom/infrastructure/entities"
)

//PurchasesUseCase ...
type PurchasesUseCase interface {
	Handler(purchaseResumeRequest entities.PurchaseResumeRequest) (domainEntities.Statistics, error)
}

//PurchasesApplication ...
type PurchasesApplication struct {
	purchaseClient ports.PurchasesClient
}

//NewPurchasesApplication ....
func NewPurchasesApplication(purchaseClient ports.PurchasesClient) *PurchasesApplication {
	return &PurchasesApplication{
		purchaseClient: purchaseClient,
	}

}

/*Handler manejador de la peticion */
func (purchasesApplication *PurchasesApplication) Handler(purchaseResumeRequest entities.PurchaseResumeRequest) (statistics domainEntities.Statistics, err error) {

	purchaseList, error := purchasesApplication.purchaseClient.Get(purchaseResumeRequest)

	if error != nil {
		err := error

		return statistics, err
	}

	statistics.ComprasPorTDC = make(map[string]float64)
	for _, purchase := range purchaseList {
		obtainStatistics(&statistics, purchase)
	}

	return statistics, error
}

/*obtainStatistics logica de las estadisticas de la consulta del api*/
func obtainStatistics(statistics *domainEntities.Statistics, purchaseList []domainEntities.Purchases) {
	statistics.GetTotalPurchases(purchaseList)
	statistics.WithoutPurchases(purchaseList)
	statistics.HighestPurchases(purchaseList)
	statistics.PurchasesByCreditCards(purchaseList)
}
