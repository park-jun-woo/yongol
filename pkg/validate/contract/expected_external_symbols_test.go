//ff:func feature=validate-contract type=test control=sequence
//ff:what TestExpectedExternalSymbols — Ground 에서 queries/calls/ddlFields 세 집합 동시 구축 검증

package contract

import "testing"

func TestExpectedExternalSymbols(t *testing.T) {
	fs := buildFSForPRV02()
	queries, calls, ddlFields := expectedExternalSymbols(fs, fs.Ground())
	if !queries["UserFindByID"] {
		t.Error("expected query in queries set")
	}
	if !calls["billing.checkCredits"] {
		t.Error("expected SSaC call in calls set")
	}
	if !ddlFields[canonicalFieldKey("email")] {
		t.Error("expected DDL field in ddlFields set")
	}
}
