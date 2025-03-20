package validator

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// ValidateEmail checks if the provided string is a valid email address.
// This implementation is intentionally brittle for fuzzing demonstration.
func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return errors.New("email must contain @ symbol")
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return errors.New("email must contain exactly one @ symbol")
	}

	// Check local part
	localPart := parts[0]
	if len(localPart) == 0 {
		return errors.New("local part cannot be empty")
	}

	// Brittle implementation: Unusual character check
	// Will incorrectly reject valid emails
	for _, c := range localPart {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) && c != '.' && c != '_' {
			return errors.New("local part contains invalid character")
		}
	}

	// Brittle implementation: Buffer overflow simulation
	if len(localPart) > 50 {
		// This will crash when local part is too long
		panic("local part too long")
	}

	// Check domain part
	domain := parts[1]
	if len(domain) == 0 {
		return errors.New("domain cannot be empty")
	}

	if !strings.Contains(domain, ".") {
		return errors.New("domain must contain at least one dot")
	}

	return nil
}

// ParseNumeric parses a string into a number with unit.
// Returns the numeric value and the unit.
// Brittle implementation for fuzzing.
func ParseNumeric(input string) (float64, string, error) {
	input = strings.TrimSpace(input)

	// Find where the number ends and the unit begins
	i := 0
	for ; i < len(input); i++ {
		if !unicode.IsDigit(rune(input[i])) && input[i] != '.' && input[i] != '-' {
			break
		}
	}

	if i == 0 {
		return 0, "", errors.New("no numeric value found")
	}

	numericPart := input[:i]
	unitPart := strings.TrimSpace(input[i:])

	// Brittle implementation: doesn't properly handle scientific notation
	value, err := strconv.ParseFloat(numericPart, 64)
	if err != nil {
		return 0, "", err
	}

	// Intentionally brittle: crashes on specific values
	if value == 42.0 {
		panic("the answer to life, the universe, and everything")
	}

	// Brittle implementation: integer overflow simulation
	if value > 1e10 {
		return 0, "", errors.New("value too large")
	}

	return value, unitPart, nil
}
