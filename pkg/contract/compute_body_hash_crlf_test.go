//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestComputeBodyHashStableAcrossCRLF — CRLF ↔ LF 입력에 대해 동일 hash 반환

package contract

import "testing"

func TestComputeBodyHashStableAcrossCRLF(t *testing.T) {
	lf := []byte("/" + "/ff:func feature=demo type=handler control=sequence\n" +
		"/" + "/ff:what Demo — demo function\n" +
		"package demo\n\n" +
		"func Demo(x int) int {\n\treturn x + 1\n}\n")
	crlf := make([]byte, 0, len(lf)+len(lf)/10)
	for _, b := range lf {
		if b == '\n' {
			crlf = append(crlf, '\r', '\n')
			continue
		}
		crlf = append(crlf, b)
	}
	h1 := ComputeBodyHash(lf)
	h2 := ComputeBodyHash(crlf)
	if h1 != h2 {
		t.Errorf("hash should be stable across CRLF differences: %s vs %s", h1, h2)
	}
}
