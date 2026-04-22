//ff:func feature=contract type=test control=sequence
//ff:what test: TestComputeBodyHashChangesOnBodyChange — 본문 1 byte 변경에도 hash 가 달라져야 함

package contract

import "testing"

func TestComputeBodyHashChangesOnBodyChange(t *testing.T) {
	prefix := "/" + "/ff:func feature=demo type=handler control=sequence\n" +
		"/" + "/ff:what Demo — demo function\n" +
		"package demo\n\n"
	base := []byte(prefix + "func Demo(x int) int {\n\treturn x + 1\n}\n")
	mutated := []byte(prefix + "func Demo(x int) int {\n\treturn x + 2\n}\n")
	h1 := ComputeBodyHash(base)
	h2 := ComputeBodyHash(mutated)
	if h1 == h2 {
		t.Errorf("hash should differ when body changes, got %s == %s", h1, h2)
	}
}
