package msg

import "fmt"

const (
	CallChangeCountry       = `👇 Сlick on the country you are interested in. You can change this at any time.`
	ChangeCountryInputFail  = `🤔 We don't have this country in our database yet. Please select another one.`
	ChangeCountryServerFail = `🤔 Our server is temporarily down. Please wait for a while.`
)

func ChangeCountrySucc(country string) string {
	return fmt.Sprintf(`The country is set successfully. The new search region is %s`, country)
}
