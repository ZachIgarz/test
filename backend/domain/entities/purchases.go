package entities

import (
	"time"
)

/*Purchases estructura de las compras*/
type Purchases struct {
	ClientID int64     `json:"clientId"`
	Phone    int64     `json:"phone,omitempty"`
	Nombre   string    `json:"nombre"`
	Compro   bool      `json:"compro"`
	Tdc      string    `json:"tdc,omitempty"`
	Monto    float64   `json:"monto,omitempty"`
	Date     time.Time `json:"date"`
}
