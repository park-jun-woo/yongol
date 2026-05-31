//ff:func feature=ssac-parse type=test control=sequence
//ff:what TestSSaCParseHelpers — unit tests for the pure ssac parser helper functions
package ssac

import (
	"testing"
)

func TestExtractQuoted(t *testing.T) {
	val, rest := extractQuoted(`"hello" world`)
	if val != "hello" || rest != "world" {
		t.Errorf("extractQuoted = (%q,%q), want (hello,world)", val, rest)
	}
	// No leading quote → ("", original-trimmed).
	v2, r2 := extractQuoted(`bare value`)
	if v2 != "" || r2 != "bare value" {
		t.Errorf("no-quote = (%q,%q)", v2, r2)
	}
	// Unterminated quote → ("", original).
	v3, r3 := extractQuoted(`"unterminated`)
	if v3 != "" || r3 != `"unterminated` {
		t.Errorf("unterminated = (%q,%q)", v3, r3)
	}
}
