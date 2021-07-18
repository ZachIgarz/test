package entities

//PurchaseResumeRequest ...
type PurchaseResumeRequest struct {
	realDate string
	days     string
}

//NewPurchaseResumeRequest ...
func NewPurchaseResumeRequest(realDate string, days string) *PurchaseResumeRequest {
	return &PurchaseResumeRequest{
		realDate: realDate,
		days:     days,
	}
}

//RealDate ...
func (purchase *PurchaseResumeRequest) RealDate() string {
	return purchase.realDate
}

func (purchase *PurchaseResumeRequest) Days() string {
	return purchase.days
}
