package util

import "memory/src/core/conditions"

// Verify if a char is a digit
func IsDigit(v string) bool {
	var numbers []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, item := range numbers {
		if item == v {
			return true
		}
	}
	return false
}

// A valid number includes digits a single point followed by others digits. The point is optional.
func IsNumber(v string) bool {
	var integer bool = true

	for _, char := range v {
		if string(char) == "." && integer {
			integer = false
		} else if string(char) == "." && !integer {
			return false
		} else if !IsDigit(string(char)) {
			return false
		}
	}
	return true
}

// true / false / condition (which returns true/false)
func IsBoolean(v string) bool {
	if v == "true" || v == "false" {
		return true
	}
	return conditions.IsCondition(v)
}