package ports

import (
	domainEntities "github.com/ZachIgarz/test-api-rest/domain/entities"
	"github.com/ZachIgarz/test-api-rest/infrastructure/entities"
)

//PurchasesClient ..
type PurchasesClient interface {
	Get(purchaseResumeRequest entities.PurchaseResumeRequest) (purchaseList [][]domainEntities.Purchases, err error)
}
