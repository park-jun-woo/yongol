//ff:func feature=contract type=test control=sequence
//ff:what test: TestComputeBodyHashStable — 같은 입력에 대해 같은 hash 반환 (8 hex chars)

package contract

import "testing"

func TestComputeBodyHashStable(t *testing.T) {
	src := []byte("/" + "/ff:func feature=demo type=handler control=sequence\n" +
		"/" + "/ff:what Demo — demo function\n" +
		"package demo\n\n" +
		"func Demo(x int) int {\n\treturn x + 1\n}\n")
	h1 := ComputeBodyHash(src)
	h2 := ComputeBodyHash(src)
	if h1 == "" {
		t.Fatal("expected non-empty hash")
	}
	if h1 != h2 {
		t.Errorf("hash not stable: %s vs %s", h1, h2)
	}
	if len(h1) != 8 {
		t.Errorf("expected 8 hex chars, got %d (%q)", len(h1), h1)
	}
}
