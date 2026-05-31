//ff:func feature=ssac-parse type=test control=sequence
//ff:what TestSSaCParseHelpers — unit tests for the pure ssac parser helper functions
package ssac

import (
	"testing"
)

func TestSplitTargetMessage(t *testing.T) {
	target, msg, rem := splitTargetMessage(`course "not found" 404`)
	if target != "course" || msg != "not found" || rem != "404" {
		t.Errorf("got (%q,%q,%q)", target, msg, rem)
	}
	// No quote → target only.
	t2, m2, r2 := splitTargetMessage(`bareTarget`)
	if t2 != "bareTarget" || m2 != "" || r2 != "" {
		t.Errorf("no-quote = (%q,%q,%q)", t2, m2, r2)
	}
}
