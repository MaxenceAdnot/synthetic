package store

import (
	"fmt"

	"github.com/maxenceadnot/synthetic/pkg/api/utils"

	"github.com/google/uuid"
)

type URLCheck struct {
	ID             uint32 `json:"id"`
	AccountID      uint32 `json:"accountId"`
	URL            string `json:"url"`
	Method         string `json:"method"`
	ExpectedStatus int    `json:"expectedStatus"`
}

func NewURLCheck(accountID uint32, url, method string, expectedStatus int) (*URLCheck, error) {
	if !utils.IsURL(url) {
		return nil, fmt.Errorf("invalid url")
	}

	if !utils.IsHTTPMethod(method) {
		return nil, fmt.Errorf("invalid method")
	}

	if !utils.IsHTTPStatusCode(expectedStatus) {
		return nil, fmt.Errorf("invalid expected status")
	}

	return &URLCheck{
		ID:             uuid.New().ID(),
		AccountID:      accountID,
		URL:            url,
		Method:         method,
		ExpectedStatus: expectedStatus,
	}, nil
}
