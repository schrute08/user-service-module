package utils

import (
	"fmt"
	"testing"
	"user-service-module/internal/errors"
)

func TestIsIDValid(t *testing.T) {
	tests := []struct {
		id     uint32
		result bool
	}{
		{0, false},
		{1, true},
		{12345, true},
	}

	for _, test := range tests {
		if got := IsIDValid(test.id); got != test.result {
			t.Errorf("IsIDValid(%d) = %v; want %v", test.id, got, test.result)
		}
	}
}

func TestGetInvalidIDs(t *testing.T) {
	tests := []struct {
		ids        []uint32
		invalidIDs []uint32
	}{
		{[]uint32{0, 1, 2, 0, 3}, []uint32{0, 0}},
		{[]uint32{1, 2, 3}, []uint32{}},
		{[]uint32{0, 0, 0}, []uint32{0, 0, 0}},
	}

	for _, test := range tests {
		got := GetInvalidIDs(test.ids)
		if len(got) != len(test.invalidIDs) {
			t.Errorf("GetInvalidIDs(%v) = %v; want %v", test.ids, got, test.invalidIDs)
		}
		for i := range got {
			if got[i] != test.invalidIDs[i] {
				t.Errorf("GetInvalidIDs(%v) = %v; want %v", test.ids, got, test.invalidIDs)
				break
			}
		}
	}
}

func TestIsValidPhone(t *testing.T) {
	tests := []struct {
		phone  string
		result bool
	}{
		{"1234567890", true},
		{"0123456789", false},
		{"123456789", false},
		{"12345678901", false},
		{"123-456-7890", false},
		{"", false},
	}

	for _, test := range tests {
		if got := isValidPhone(test.phone); got != test.result {
			t.Errorf("IsValidPhone(%s) = %v; want %v", test.phone, got, test.result)
		}
	}
}

func TestIsCityValid(t *testing.T) {
	tests := []struct {
		city   string
		result bool
	}{
		{"New York", true},
		{"Los-Angeles", true},
		{"San Francisco", true},
		{"123City", false},
		{"City!", false},
		{"", false},
	}

	for _, test := range tests {
		if got := isCityValid(test.city); got != test.result {
			t.Errorf("isCityValid(%s) = %v; want %v", test.city, got, test.result)
		}
	}
}

func TestValidateSearchRequest(t *testing.T) {
	tests := []struct {
		name        string
		city        string
		phone       string
		isValid     bool
		errContains string
	}{
		{
			name: "should validate when both city and phone are valid", 
			city: "New York", 
			phone: "1234567890", 
			isValid: true, 
			errContains: "",
		},
		{
			name: "should validate when both city and phone are valid 2", 
			city: "Los Angeles", 
			phone: "9876543210", 
			isValid: true, 
			errContains: "",
		},
		{
			name: "should validate when city is empty and phone is valid", 
			city: "", 
			phone: "1234567890", 
			isValid: true, 
			errContains: "",
		},
		{
			name: "should validate when city is valid and phone is empty",
			city: "New York",
			phone: "",
			isValid: true,
			errContains: "",
		},
		{
			name: "should invalidate when city is invalid and phone is valid", 
			city: "InvalidCity1", 
			phone: "1234567890", 
			isValid: false, 
			errContains: "city",
		},
		{
			name: "should invalidate when both city and phone are invalid", 
			city: "InvalidCity1", 
			phone: "0123456789", 
			isValid: false, 
			errContains: "city, phone",
		},
		{
			name: "should invalidate when both city and phone are empty", 
			city: "", 
			phone: "", 
			isValid: false, 
			errContains: "either city or phone must be provided",
		},
		{
			name: "should invalidate when city is invalid and phone is empty", 
			city: "InvalidCity1", 
			phone: "", 
			isValid: false, 
			errContains: "city",
		},
	}	

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			valid, err := ValidateSearchRequest(test.city, test.phone)
			if valid != test.isValid {
				t.Errorf("ValidateSearchRequest(%q, %q) valid = %v; want %v", test.city, test.phone, valid, test.isValid)
			}
			if err != nil && test.errContains != "" {
				expectedErr := fmt.Sprintf("%v: %v", errors.ErrInvalidFields, test.errContains)
				if err.Error() != expectedErr {
					t.Errorf("ValidateSearchRequest(%q, %q) err = %v; want %v", test.city, test.phone, err, expectedErr)
				}
			} else if err == nil && test.errContains != "" {
				t.Errorf("ValidateSearchRequest(%q, %q) expected error, got nil", test.city, test.phone)
			}
		})
	}
}
