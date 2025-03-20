package validator

import (
	"testing"
)

func FuzzValidateEmail(f *testing.F) {
	// Add minimal seed corpus
	seeds := []string{
		"test@example.com",
		"",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, email string) {
		// The fuzzer will detect and report any panics automatically
		_ = ValidateEmail(email)
	})
}

func FuzzParseNumeric(f *testing.F) {
	// Add minimal seed corpus
	seeds := []string{
		"123px",
		"",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		// The fuzzer will detect and report any panics automatically
		_, _, _ = ParseNumeric(input)
	})
}
