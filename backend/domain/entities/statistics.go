package entities

import "math"

/*Statistics estructura de las stadisticas*/
type Statistics struct {
	Total         float64            `json:"total"`
	ComprasPorTDC map[string]float64 `json:"comprasPorTDC"`
	NoCompraron   int64              `json:"noCompraron"`
	CompraMasAlta float64            `json:"compraMasAlta"`
}

/*GetTotalPurchases total of purchases*/
func (statistics *Statistics) GetTotalPurchases(purchaseList []Purchases) {
	var total float64
	for i := 0; i < len(purchaseList); i++ {
		if purchaseList[i].Compro == true {
			total = total + purchaseList[i].Monto
		}
	}
	statistics.Total = statistics.Total + roundFloat(total)

}

//WithoutPurchases  Without Purchases
func (statistics *Statistics) WithoutPurchases(purchaseList []Purchases) {
	var total int64
	for i := 0; i < len(purchaseList); i++ {
		if purchaseList[i].Compro == false {
			total++
		}
	}
	statistics.NoCompraron = statistics.NoCompraron + total
}

//HighestPurchases the HIGHEST!!!!
func (statistics *Statistics) HighestPurchases(purchaseList []Purchases) {
	var total float64 = 0
	for i := 0; i < len(purchaseList); i++ {
		if purchaseList[i].Monto > total {
			total = purchaseList[i].Monto
		}
	}

	statistics.CompraMasAlta = statistics.CompraMasAlta + roundFloat(total)
}

//PurchasesByCreditCards return map of puchases for every credit card
func (statistics *Statistics) PurchasesByCreditCards(purchaseList []Purchases) {
	mapPurchases := make(map[string]float64)

	for i := 0; i < len(purchaseList); i++ {
		key := purchaseList[i].Tdc
		if purchaseList[i].Compro == true {
			value := mapPurchases[key] + purchaseList[i].Monto
			mapPurchases[key] = roundFloat(value)
		}
	}

	for key, value := range mapPurchases {
		statistics.ComprasPorTDC[key] = statistics.ComprasPorTDC[key] + value
	}
}

func roundFloat(purchaseList float64) float64 {
	return math.Round(purchaseList*100) / 100
}
