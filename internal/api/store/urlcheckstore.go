package store

import "fmt"

type URLCheckStore interface {
	GetURLChecks() []URLCheck
	GetURLCheckByID(id uint32) (*URLCheck, error)
	GetURLChecksByAccount(account uint32) []URLCheck
	AddURLCheck(urlcheck URLCheck) error
	DeleteURLCheck(id uint32) error
}

type InMemoryURLCheckStore struct {
	urlChecks []URLCheck
}

func NewInMemoryURLCheckStore(urlchecks []URLCheck) *InMemoryURLCheckStore {
	return &InMemoryURLCheckStore{
		urlChecks: urlchecks,
	}
}

func (s *InMemoryURLCheckStore) GetURLChecks() []URLCheck {
	return s.urlChecks
}

func (s *InMemoryURLCheckStore) GetURLCheckByID(id uint32) (*URLCheck, error) {
	for _, urlCheck := range s.urlChecks {
		if urlCheck.ID == id {
			return &urlCheck, nil
		}
	}

	return nil, fmt.Errorf("url check not found")
}

func (s *InMemoryURLCheckStore) GetURLChecksByAccount(account uint32) []URLCheck {
	var urlChecksForAccount []URLCheck

	for _, urlCheck := range s.urlChecks {
		if urlCheck.AccountID == account {
			urlChecksForAccount = append(urlChecksForAccount, urlCheck)
		}
	}

	return urlChecksForAccount
}

func (s *InMemoryURLCheckStore) AddURLCheck(urlcheck URLCheck) error {
	s.urlChecks = append(s.urlChecks, urlcheck)

	return nil
}

func (s *InMemoryURLCheckStore) DeleteURLCheck(id uint32) error {
	for i, urlCheck := range s.urlChecks {
		if urlCheck.ID == id {
			s.urlChecks = append(s.urlChecks[:i], s.urlChecks[i+1:]...)

			return nil
		}
	}

	return fmt.Errorf("url check not found")
}
