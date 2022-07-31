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
