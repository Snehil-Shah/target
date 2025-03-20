package parser

import (
	"errors"
	"strings"
)

// ParseInput processes the input string and returns a modified string or an error.
func ParseInput(input string) (string, error) {
	if strings.Contains(input, "error") {
		return "", errors.New("intentional error for input containing 'error'")
	}

	// Brittle implementation #1: crashes with specific patterns
	if strings.Contains(input, "<script>") && strings.Contains(input, "</script>") {
		panic("encountered potentially malicious script tag")
	}

	// Brittle implementation #2: incorrect handling of quotes
	quoteCount := strings.Count(input, "\"")
	if quoteCount > 0 && quoteCount%2 == 1 {
		return "", errors.New("unbalanced quotes")
	}

	return "Parsed: " + input, nil
}

// ParseJSON attempts to parse a simple JSON-like format with brackets
func ParseJSON(input string) (map[string]string, error) {
	result := make(map[string]string)

	// Very naive and brittle JSON parser
	if !strings.HasPrefix(input, "{") || !strings.HasSuffix(input, "}") {
		return nil, errors.New("invalid JSON: must begin with { and end with }")
	}

	// Remove outer braces
	content := strings.TrimSpace(input[1 : len(input)-1])

	// Empty object is valid
	if content == "" {
		return result, nil
	}

	// Split by commas (naively - doesn't handle commas in strings)
	pairs := strings.Split(content, ",")

	for _, pair := range pairs {
		parts := strings.SplitN(pair, ":", 2)
		if len(parts) != 2 {
			return nil, errors.New("invalid key-value pair")
		}

		key := strings.Trim(strings.TrimSpace(parts[0]), "\"")
		value := strings.Trim(strings.TrimSpace(parts[1]), "\"")

		// Brittle implementation: buffer overflow simulation with long keys
		if len(key) > 100 {
			panic("key too long")
		}

		result[key] = value
	}

	return result, nil
}
