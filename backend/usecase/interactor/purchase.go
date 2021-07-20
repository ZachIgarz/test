package interactor

type purchaseInteractor struct {
}

type Purchase interface {
	Purchases(date string, days uint) ([]string, error)
}

func NewPurchaseInteractor() *purchaseInteractor {
	return &purchaseInteractor{}
}
