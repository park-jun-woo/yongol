//ff:func feature=stml-gen type=test-helper control=sequence
//ff:what 코드에 substring이 포함되어 있는지 검증하는 테스트 헬퍼
package stml

import ("strings"; "testing")

func assertContains(t *testing.T, code, substr string) {
	t.Helper()
	if !strings.Contains(code, substr) { t.Errorf("generated code does not contain %q\n--- code ---\n%s", substr, code) }
}
