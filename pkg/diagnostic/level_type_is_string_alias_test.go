//ff:func feature=orchestrator type=test control=sequence
//ff:what Level is string alias — string ↔ Level 왕복 변환 검증
package diagnostic_test

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// TestLevel_TypeIsStringAlias verifies that Level is convertible to and from string.
func TestLevel_TypeIsStringAlias(t *testing.T) {
	l := diagnostic.LevelError
	if string(l) != "ERROR" {
		t.Errorf("string(LevelError): want %q, got %q", "ERROR", string(l))
	}

	custom := diagnostic.Level("INFO")
	if string(custom) != "INFO" {
		t.Errorf("Level(\"INFO\"): round-trip failed, got %q", string(custom))
	}
}
