//ff:func feature=gen-hurl type=test control=iteration dimension=1
//ff:what TestSmokeBodyUniqueEmail — Register/Login body의 email 필드에 {{newUuid}} placeholder 삽입 검증

package hurl

import (
	"strings"
	"testing"
)

// TestSmokeBodyUniqueEmail pins BUG-015 Phase003: every email field in
// smoke auth bodies must carry a unique-per-run placeholder so repeated
// runs against the same DB each Register with a fresh unique address
// (otherwise the DDL unique constraint trips on the 2nd run). The
// Register step seeds `smoke_email=smoke-{{newUuid}}@example.com` via
// [Options] and both Register + Login bodies reference {{smoke_email}}
// so the pair shares the SAME address within a single run (otherwise
// Login would hit a 401 against an email Register never saw).
func TestSmokeBodyUniqueEmail(t *testing.T) {
	fs := newSmokeFullstack(newAuthOnlyOpenAPI())
	steps := buildScenarioOrder(fs)

	authOps := map[string]bool{"Register": true, "Login": true}
	sawRegisterOption := false
	checked := 0
	for _, s := range steps {
		if !authOps[s.OperationID] {
			continue
		}
		checked++
		if !s.HasBody {
			t.Errorf("%s step missing body", s.OperationID)
			continue
		}
		if !strings.Contains(s.BodyJSON, "{{smoke_email}}") {
			t.Errorf("%s body missing {{smoke_email}} placeholder: %s", s.OperationID, s.BodyJSON)
		}
		if s.OperationID == "Register" && hasSmokeEmailOption(s.Options) {
			sawRegisterOption = true
		}
	}
	if checked == 0 {
		t.Fatalf("no auth steps were inspected; got opIDs=%v", stepOpIDs(steps))
	}
	if !sawRegisterOption {
		t.Errorf("Register step must seed smoke_email via [Options] containing {{newUuid}}; got no matching option")
	}
}
