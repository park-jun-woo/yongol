//ff:func feature=cli type=test control=sequence
//ff:what parseModelFlag test — valid/invalid model flag 파싱 검증

package main

import (
	"testing"
)

func TestParseModelFlag(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		cases := []struct {
			input   string
			backend string
			model   string
		}{
			{"ollama:gemma4:e4b", "ollama", "gemma4:e4b"},
			{"xai:grok-3", "xai", "grok-3"},
			{"gemini:pro", "gemini", "pro"},
		}
		for _, tc := range cases {
			b, m, err := parseModelFlag(tc.input)
			if err != nil {
				t.Errorf("parseModelFlag(%q) error: %v", tc.input, err)
				continue
			}
			if b != tc.backend || m != tc.model {
				t.Errorf("parseModelFlag(%q) = (%q, %q), want (%q, %q)", tc.input, b, m, tc.backend, tc.model)
			}
		}
	})
	t.Run("NoColon", func(t *testing.T) {
		_, _, err := parseModelFlag("no-colon")
		if err == nil {
			t.Fatal("expected error for missing colon")
		}
	})
	t.Run("UnsupportedBackend", func(t *testing.T) {
		_, _, err := parseModelFlag("openai:gpt4")
		if err == nil {
			t.Fatal("expected error for unsupported backend")
		}
	})
	t.Run("EmptyModel", func(t *testing.T) {
		_, _, err := parseModelFlag("ollama:")
		if err == nil {
			t.Fatal("expected error for empty model")
		}
	})
}
