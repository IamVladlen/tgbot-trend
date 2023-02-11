package repository

import (
	"github.com/IamVladlen/trend-bot/internal/entity"
	"github.com/IamVladlen/trend-bot/pkg/mongodb"
)

const (
	_mongoRequestTimeout = 5

	_chatCollection = "chat"
)

type Repository struct {
	Chat
}

type Chat interface {
	ChangeCountry(chat entity.Chat) error
	GetCountry(id int) (string, error)
}

func New(mg *mongodb.DB) *Repository {
	return &Repository{
		newCountryRepo(mg),
	}
}
