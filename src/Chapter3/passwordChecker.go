package main

import "unicode"

func passwordChecker(pw string) bool {
	if len(pw) < 8 || len(pw) > 15 {
		return false
	}
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false
	for _, ch := range pw {
		switch {
		case unicode.IsUpper(ch):
			if !hasUpper {
				hasUpper = true
			}
		case unicode.IsLower(ch):
			if !hasLower {
				hasLower = true
			}
		case unicode.IsDigit(ch):
			if !hasNumber {
				hasNumber = true
			}
		case unicode.IsSymbol(ch):
			if !hasSymbol {
				hasSymbol = true
			}
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}
