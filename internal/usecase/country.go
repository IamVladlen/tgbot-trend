package usecase

import (
	"fmt"
)

type CountryUC struct {
	repo CountryRepo
}

// ChangeCountry changes country of fetched trends in chat.
func (uc *CountryUC) ChangeCountry(id int, text string) error {
	country, err := uc.validateCountry(text)
	if err != nil {
		return fmt.Errorf("usecase - ChangeCountry: %w", err)
	}

	if err := uc.repo.ChangeCountry(id, country); err != nil {
		return fmt.Errorf("usecase - ChangeCountry: %w", err)
	}

	return nil
}

// TODO: Switch to map after increasing the number of countries

// validateCountry converts emoji to plain text and returns
// an error if there is no reference.
func (uc *CountryUC) validateCountry(text string) (string, error) {
	switch text {
	case "🇩🇪":
		return "DE", nil
	case "🇪🇸":
		return "ES", nil
	case "🇫🇷":
		return "FR", nil
	case "🇮🇹":
		return "IT", nil
	case "🇬🇧":
		return "GB", nil
	case "🇷🇺":
		return "RU", nil
	case "🇺🇦":
		return "UA", nil
	case "🇺🇸", "🇺🇲":
		return "US", nil
	default:
		return "", errInvalidCountry
	}
}

func newCountryUC(repo CountryRepo) *CountryUC {
	return &CountryUC{
		repo: repo,
	}
}
