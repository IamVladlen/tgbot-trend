package usecase

import (
	"fmt"
)

type CountryUC struct {
	repo CountryRepo
}

// ChangeCountry changes country of fetched trends in chat.
func (uc *CountryUC) ChangeCountry(id int, country string) error {
	if err := uc.repo.ChangeCountry(id, country); err != nil {
		return fmt.Errorf("usecase - ChangeCountry: %w", err)
	}

	return nil
}

func newCountryUC(repo CountryRepo) *CountryUC {
	return &CountryUC{
		repo: repo,
	}
}
