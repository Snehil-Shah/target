package parser

import (
	"testing"
)

func FuzzParseInput(f *testing.F) {
	// Add seed corpus
	seeds := []string{
		"hello world",
		"a simple test",
		"",
		"something with error in it",
		"quotes like \"this\"",
		"unbalanced \"quotes",
		"<script>alert('test')</script>",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, input string) {
		// This will catch panics
		result, err := ParseInput(input)

		// Basic validation
		if err == nil && !testing.Short() {
			if result == "" {
				t.Error("Empty result returned without error")
			}
		}
	})
}

func FuzzParseJSON(f *testing.F) {
	// Add seed corpus
	seeds := []string{
		"{}",
		"{\"key\":\"value\"}",
		"{key: value}",
		"{\"name\":\"test\",\"type\":\"fuzz\"}",
		"{invalid",
		"not json at all",
		"{\"very_long_key_that_repeats_many_times_very_long_key_that_repeats_many_times_very_long_key_that_repeats_many_times\":\"value\"}",
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
