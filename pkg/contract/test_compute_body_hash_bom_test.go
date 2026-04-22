//ff:func feature=contract type=test control=sequence
//ff:what test: TestComputeBodyHashStableAcrossBOM — 선두 UTF-8 BOM 유무에도 동일 hash 반환

package contract

import "testing"

func TestComputeBodyHashStableAcrossBOM(t *testing.T) {
	src := []byte("/" + "/ff:func feature=demo type=handler control=sequence\n" +
		"/" + "/ff:what Demo — demo function\n" +
		"package demo\n\n" +
		"func Demo(x int) int {\n\treturn x + 1\n}\n")
	withBOM := append([]byte{0xEF, 0xBB, 0xBF}, src...)
	h1 := ComputeBodyHash(src)
	h2 := ComputeBodyHash(withBOM)
	if h1 != h2 {
		t.Errorf("hash should be stable across BOM: %s vs %s", h1, h2)
	}
}
