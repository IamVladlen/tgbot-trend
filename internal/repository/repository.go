package repository

import (
	"github.com/IamVladlen/trend-bot/pkg/mongodb"
)

const (
	_mgdbRequestTimeout = 5

	_chatMgdbCollection = "chat"
)

type Repository struct {
	Country *countryRepo
}

// New creates repository instance.
func New(mg *mongodb.DB) *Repository {
	return &Repository{
		newCountryRepo(mg),
	}
}
