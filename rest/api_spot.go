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
	price float64,
	timeInForce string,
	clientOID string,
) (query string, resp []byte, result Balance, err error) {
	var ret GetBalanceResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["side"] = side
	params["type"] = orderType
	if price > 0 {
		params["price"] = price
	}
	params["timeInForce"] = timeInForce
	params["qty"] = qty
	params["orderLinkId"] = clientOID
	query, resp, err = b.SignedRequest(http.MethodGet, "/spot/v1/order", params, &ret)
	if err != nil {
		return
	}
	return
}
