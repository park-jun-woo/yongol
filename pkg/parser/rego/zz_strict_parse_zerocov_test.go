//ff:func feature=policy type=test control=sequence
//ff:what TestStrictParse_ZeroCov — strictParse 정상/구문오류 분기 직접 호출

package rego

import "testing"

func TestStrictParse_ZeroCov(t *testing.T) {
	// valid RegoV1 module → nil.
	valid := "package authz\n\nimport rego.v1\n\ndefault allow := false\n\nallow if {\n\tinput.method == \"GET\"\n}\n"
	if diags := strictParse("policy.rego", valid); len(diags) != 0 {
		t.Errorf("valid module should yield no diagnostics, got %v", diags)
	}

	// syntax error → ast.Errors path → at least one R-1 diagnostic.
	broken := "package authz\n\nallow if {\n\tinput.method ==\n"
	if diags := strictParse("policy.rego", broken); len(diags) == 0 {
		t.Errorf("broken module should yield parse diagnostics")
	}
}
