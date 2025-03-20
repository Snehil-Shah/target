package parser

import (
	"testing"
)

func FuzzParseInput(f *testing.F) {
	// Add minimal seed corpus
	seeds := []string{
		"hello world",
		"",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		// This will catch panics
		_, _ = ParseInput(input)
		// Let the fuzzer discover issues on its own
	})
}

func FuzzParseJSON(f *testing.F) {
	// Add minimal seed corpus
	seeds := []string{
		"{}",
		"{\"key\":\"value\"}",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		// Call the function and catch any panics
		_, _ = ParseJSON(input)
		// The fuzzer will detect and report any panics automatically
	})
}
