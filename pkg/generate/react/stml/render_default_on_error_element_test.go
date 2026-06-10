//ff:func feature=stml-gen type=test control=sequence
//ff:what renderDefaultOnErrorElement — 기본 에러 슬롯(role="alert" 조건부 렌더) JSX 검증
package stml

import "testing"

func TestRenderDefaultOnErrorElement(t *testing.T) {
	got := renderDefaultOnErrorElement("loginError", 2)
	want := `  {loginError && <p role="alert" className="text-sm text-destructive">{loginError}</p>}`
	if got != want {
		t.Errorf("default on-error element = %q, want %q", got, want)
	}
}
