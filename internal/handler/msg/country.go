package msg

import "fmt"

const (
	CallChangeCountry = `👇 Сlick on the country you are interested in. You can change this at any time.`
	ChangeCountryFail = `🤔 We this country in our database. Please select another one.`
)

func ChangeCountrySucc(country string) string {
	return fmt.Sprintf(`The country is set successfully. The new search region is %s`, country)
}
