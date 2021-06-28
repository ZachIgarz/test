package restclients

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	domainEntities "github.com/ZachIgarz/golangIpCom/domain/entities"
	"github.com/ZachIgarz/golangIpCom/infrastructure/entities"
)

const purchaseURL = "https://apirecruit-gjvkhl2c6a-uc.a.run.app/compras/"

//PurchaseRestClient ...
type PurchaseRestClient struct {
}

/*Get consume api y retorna respuesta*/
func (purchaseRestClient PurchaseRestClient) Get(purchaseResumeRequest entities.PurchaseResumeRequest) (purchaseList [][]domainEntities.Purchases, err error) {
	purchaseRestClientURL := (purchaseURL) + (purchaseResumeRequest.RealDate())

	if len(purchaseResumeRequest.Days()) > 0 {
		date := purchaseResumeRequest.RealDate()

		layoutISO := "2006-01-02"
		t, _ := time.Parse(layoutISO, date)
		l, _ := strconv.Atoi(purchaseResumeRequest.Days())
		var dateNext time.Time = t

		for i := 0; i < l; i++ {

			purchase := []domainEntities.Purchases{}

			dateNext = dateNext.AddDate(0, 0, 1)
			realDate := dateNext.String()[0:10]

			purchaseRestClientURL = (purchaseURL) + (realDate)
			response, err := http.Get(purchaseRestClientURL)
			if err != nil {

				return purchaseList, err
			}

			respuesta, _ := ioutil.ReadAll(response.Body)
			dec := json.NewDecoder(bytes.NewReader(respuesta))
			dec.Decode(&purchase)

			purchaseList = append(purchaseList, purchase)

		}

		return
	}
	purchase := []domainEntities.Purchases{}
	response, err := http.Get(purchaseRestClientURL)
	if err != nil {

		return purchaseList, err
	}

	respuesta, _ := ioutil.ReadAll(response.Body)

	dec := json.NewDecoder(bytes.NewReader(respuesta))
	dec.Decode(&purchase)

	purchaseList = append(purchaseList, purchase)

	//DEFER
	defer response.Body.Close()

	return purchaseList, err
}
