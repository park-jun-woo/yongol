//ff:func feature=gen-ffhash type=test control=sequence
//ff:what test: TestInjectIsIdempotent — 두번 호출해도 출력이 동일 (idempotency)

package ffhash

import (
	"bytes"
	"testing"
)

func TestInjectIsIdempotent(t *testing.T) {
	src := []byte("/" + "/ff:func feature=demo type=handler control=sequence\n" +
		"/" + "/ff:what Demo — demo function\n" +
		"package demo\n\n" +
		"func Demo(x int) int {\n\treturn x + 1\n}\n")
	once := InjectCheckedLine(src)
	twice := InjectCheckedLine(once)
	if !bytes.Equal(once, twice) {
		t.Errorf("InjectCheckedLine not idempotent:\nfirst:\n%s\nsecond:\n%s", once, twice)
	}
}
