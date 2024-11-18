package main

import(
	"testing"
	"unicode"
)

func HasSymbol(str string)bool{
	for _, char := range str{
		if unicode.IsSymbol(char) || unicode.IsPunct(char){
			return true
		}
	}
	return false
}

func HasNumber(str string)bool{
	for _, char := range str{
		if unicode.IsDigit(char){
			return true
		}
	}
	return false
}

func HasUpperCase(str string)bool{
	for _, char := range str{
		if unicode.IsUpper(char){
			return true
		}
	}
	return false
}

func TestGeneratePassword(t *testing.T){
	testCases := []struct{
		name string
		config PasswordConfig
		expectedLength int
		expectSymbol bool
		expectNumber bool
		expectUpper bool
	}{
		{"Default", PasswordConfig{length: 10, IncludeSymbols: false, IncludeNumbers: false, IncludeUpper: false}, 10, false, false, false},
		{"Has symbols", PasswordConfig{length: 10, IncludeSymbols: true, IncludeNumbers: false, IncludeUpper: false}, 10, true, false, false},
		{"Has numbers", PasswordConfig{length: 10, IncludeSymbols: false, IncludeNumbers: true, IncludeUpper: false}, 10, false, true, false},
		{"Has uppercase", PasswordConfig{length: 10, IncludeSymbols: false, IncludeNumbers: false, IncludeUpper: true}, 10, false, false, true},

	}

	for _, tcase := range testCases{
		t.Run(tcase.name, func(t *testing.T){
			password := tcase.config.GeneratePassword()
				// Check password length
				if len(password) != tcase.expectedLength {
					t.Errorf("expected length %d, got %d", tcase.expectedLength, len(password))
				}

					// Check if symbols are included 
				if tcase.expectSymbol != HasSymbol(password) {
					t.Errorf("expected symbol presence: %v", tcase.expectSymbol)
				}

				// Check if numbers are included 
				if tcase.expectNumber != HasNumber(password) {
					t.Errorf("expected number presence: %v", tcase.expectNumber)
				}

				// Check if uppercase letters are included
				if tcase.expectUpper != HasUpperCase(password){
					t.Errorf("expected uppercase presence: %v", tcase.expectUpper)
				}

			})
	}
}