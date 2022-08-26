package rest

import (
	"net/http"
)

// GetAccount
func (b *ByBit) GetAccount() (query string, resp []byte, result Balance, err error) {
	var ret GetBalanceResult
	params := map[string]interface{}{}
	query, resp, err = b.SignedRequest(http.MethodGet, "/spot/v1/account", params, &ret)
	if err != nil {
		return
	}
	return
}

func (b *ByBit) CreateSpotOrder(symbol string,
	side string,
	orderType string,
	qty float64,
	timeInForce string) (query string, resp []byte, result Balance, err error) {
	var ret GetBalanceResult
	params := map[string]interface{}{}
	query, resp, err = b.SignedRequest(http.MethodGet, "/spot/v1/order", params, &ret)
	if err != nil {
		return
	}
	return
}
