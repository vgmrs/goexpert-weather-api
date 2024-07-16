package validator

import (
	"regexp"
)

func IsValidCEP(cep string) bool {
	return regexp.MustCompile(`^\d{5}-?\d{3}$`).MatchString(cep)
}
