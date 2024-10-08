package utils

import (
	"fmt"
	"regexp"
	"strings"
	"user-service-module/internal/errors"
	pb "user-service-module/proto/user/userpb"
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

func isValidPhone(phone string) bool {
	phoneRegex := `^[1-9]\d{9}$`
	return regexp.MustCompile(phoneRegex).MatchString(phone)
}

func isCityValid(city string) bool {
	cityRegex := `^[a-zA-Z]+(?:[\s-][a-zA-Z]+)*$`
	return regexp.MustCompile(cityRegex).MatchString(city)
}

func ValidateSearchRequest(city, phone string, isMarried pb.MaritalStatus) (bool, error) {
	if city == "" && phone == "" {
		if isMarried == pb.MaritalStatus_UNKNOWN{
			return false, fmt.Errorf(
				"%w: %v", 
				errors.ErrInvalidFields, 
				"either city, phone or marital status must be provided",
			)
		} else {
			return true, nil
		}
		
	}

	var invalidFields []string
	if city != "" && !isCityValid(city) {
		invalidFields = append(invalidFields, "city")
	}
	if phone != "" && !isValidPhone(phone) {
		invalidFields = append(invalidFields, "phone")
	}

	if len(invalidFields) > 0 {
		// Join the invalid fields into a single string
		invalidFieldsStr := strings.Join(invalidFields, ", ")
		return false, fmt.Errorf("%w: %v", errors.ErrInvalidFields, invalidFieldsStr)
	}
	return true, nil
}
