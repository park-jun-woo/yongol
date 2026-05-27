//ff:func feature=cli type=test control=sequence
//ff:what formatLocation test — file:line prefix 생성 검증

package main

import (
	"testing"
)

func TestFormatLocation(t *testing.T) {
	t.Run("FileAndLine", func(t *testing.T) {
		got := formatLocation("api.yaml", 42)
		if got != "api.yaml:42: " {
			t.Errorf("expected 'api.yaml:42: ', got '%s'", got)
		}
	})
	t.Run("FileOnly", func(t *testing.T) {
		got := formatLocation("api.yaml", 0)
		if got != "api.yaml: " {
			t.Errorf("expected 'api.yaml: ', got '%s'", got)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		got := formatLocation("", 0)
		if got != "" {
			t.Errorf("expected empty, got '%s'", got)
		}
	})
}
