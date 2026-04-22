//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWarnExceeds_Violations — Q1 위반 소스에서 WarnExceeds가 WARN 반환

package qcheck

import "testing"

func TestWarnExceeds_Violations(t *testing.T) {
	warns := WarnExceeds("deep.go", deepSrc, DefaultLimits())
	if len(warns) == 0 {
		t.Fatalf("want Q1 warn on deep src")
	}
}
