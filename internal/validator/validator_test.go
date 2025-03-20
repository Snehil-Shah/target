package validator

import (
	"testing"
)

func FuzzValidateEmail(f *testing.F) {
	// Add seed corpus
	seeds := []string{
		"test@example.com",
		"user.name@domain.com",
		"invalid-email",
		"no@dots",
		"@nodomain.com",
		"noatsign.com",
		"user@multiple@ats.com",
		"verylonglocal_part_that_exceeds_the_arbitrary_limit_we_set_in_our_validator_function_to_demonstrate_fuzzing@example.com",
		"special!chars@example.com",
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
	// Add seed corpus
	seeds := []string{
		"123px",
		"42.0 meters",
		"-17.5kg",
		"3.14159 radians",
		"invalid",
		"42",
		"1e10 light-years",
		"1e20 parsecs",
		".",
		"-",
		"---",
		"...",
		"42.0", // Will cause panic
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		// The fuzzer will detect and report any panics automatically
		_, _, _ = ParseNumeric(input)
	})
}
