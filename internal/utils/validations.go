package utils

import (
	"fmt"
	"regexp"
	"user-service-module/internal/errors"
)

func IsIDValid(id uint32) bool {
	return id > 0
}

func GetInvalidIDs(ids []uint32) []uint32 {
	var invalidIDs []uint32
	for _, id := range ids {
		if !IsIDValid(id) {
			invalidIDs = append(invalidIDs, id)
		}
	}
	return invalidIDs
}

func IsValidPhone(phone string) bool {
	phoneRegex := `^[1-9]\d{9}$`
	return regexp.MustCompile(phoneRegex).MatchString(phone)
}

func isCityValid(city string) bool {
	cityRegex := `^[a-zA-Z]+(?:[\s-][a-zA-Z]+)*$`
	return regexp.MustCompile(cityRegex).MatchString(city)
}

func ValidateSearchRequest(city, phone string) (bool, error) {
	invalidFields := []string{}
	if !isCityValid(city) {
		invalidFields = append(invalidFields, "city")
	}
	if !IsValidPhone(phone) {
		invalidFields = append(invalidFields, "phone")
	}
	
	if len(invalidFields) > 0 {
		return false, fmt.Errorf("%w: %v", errors.ErrInvalidFields, invalidFields)
	}
	return true, nil	
}
