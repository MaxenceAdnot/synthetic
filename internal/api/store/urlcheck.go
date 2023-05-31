package store

import (
	"github.com/maxenceadnot/synthetic/pkg/api/utils"

	"github.com/google/uuid"
)

type URLCheckError string

func (e URLCheckError) Error() string {
	return string(e)
}

const (
	ErrInvalidURL        = URLCheckError("invalid url")
	ErrInvalidMethod     = URLCheckError("invalid method")
	ErrInvalidHTTPStatus = URLCheckError("invalid expected status")
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
		return nil, ErrInvalidURL
	}

	if !utils.IsHTTPMethod(method) {
		return nil, ErrInvalidMethod
	}

	if !utils.IsHTTPStatusCode(expectedStatus) {
		return nil, ErrInvalidHTTPStatus
	}

	return &URLCheck{
		ID:             uuid.New().ID(),
		AccountID:      accountID,
		URL:            url,
		Method:         method,
		ExpectedStatus: expectedStatus,
	}, nil
}
