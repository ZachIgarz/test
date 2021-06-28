package ports

import (
	domainEntities "github.com/ZachIgarz/golangIpCom/domain/entities"
	"github.com/ZachIgarz/golangIpCom/infrastructure/entities"
)

//PurchasesClient ..
type PurchasesClient interface {
	Get(purchaseResumeRequest entities.PurchaseResumeRequest) (purchaseList [][]domainEntities.Purchases, err error)
}
