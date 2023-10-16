package model_question3

import (
	"errors"
	"strings"
	"tes_coding_siesta/init/question3"
)

func Question3(request question3.PhoneNumber) (response question3.Question3, err error) {

	if len(request.Number) < 11 {
		return response, errors.New("invalid number")
	}

	request.Number = strings.ReplaceAll(request.Number, " ", "")

	if string(request.Number[0]) == "0" {
		response.Format = getCountryCode(request.Country)+request.Number[1:]
	} else {
		response.Format = request.Number[1:]
	}

	return
}

func getCountryCode(country string) string {
	countryCodes := map[string]string{
		"indonesia": "62",
		"usa":       "1",
		"japan":     "81",
		"singapore": "60",
	}

	country = strings.ToLower(country)

	if code, found := countryCodes[country]; found {
		return code
	}
	return "Country code not found"
}
