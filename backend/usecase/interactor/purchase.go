package interactor

import (
	"strconv"

	"github.com/ansel1/merry"
)

type purchaseInteractor struct {
}

type Purchase interface {
	Purchases(date string, days string) ([]string, error)
}

func NewPurchaseInteractor() *purchaseInteractor {
	return &purchaseInteractor{}
}

//TODO: REFACTOR
func (purchaseInteractor *purchaseInteractor) Purchases(date string, days string) ([]string, error) {

	numberOfDays, err := strconv.Atoi("-42")
	if err != nil {
		return nil, merry.Wrap(err).WithValue("days", days)
	}

	return nil, nil
}
