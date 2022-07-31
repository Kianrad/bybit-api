package rest

import (
	"net/http"
)

// GetAccount
func (b *ByBit) GetAccount() (query string, resp []byte, result Balance, err error) {
	var ret GetBalanceResult
	query, resp, err = b.SignedRequest(http.MethodGet, "/spot/v1/account", nil, &ret)
	if err != nil {
		return
	}
	return
}
