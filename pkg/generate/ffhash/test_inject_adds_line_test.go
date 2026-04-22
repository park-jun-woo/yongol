//ff:func feature=gen-ffhash type=test control=sequence
//ff:what test: TestInjectAddsCheckedLine — 최초 호출에서 //ff:checked 라인을 올바른 위치에 삽입

package ffhash

import (
	"bytes"
	"testing"
)

func TestInjectAddsCheckedLine(t *testing.T) {
	src := []byte("/" + "/ff:func feature=demo type=handler control=sequence\n" +
		"/" + "/ff:what Demo — demo function\n" +
		"package demo\n\n" +
		"func Demo(x int) int {\n\treturn x + 1\n}\n")
	out := InjectCheckedLine(src)
	if !bytes.Contains(out, []byte("/"+"/ff:checked llm=yongol-gen hash=")) {
		t.Fatalf("expected //ff:checked line in output, got:\n%s", out)
	}
	idxWhat := bytes.Index(out, []byte("/"+"/ff:what"))
	idxChecked := bytes.Index(out, []byte("/"+"/ff:checked"))
	idxPkg := bytes.Index(out, []byte("package demo"))
	if !(idxWhat < idxChecked && idxChecked < idxPkg) {
		t.Errorf("expected checked line between //ff:what and package, got what=%d checked=%d pkg=%d", idxWhat, idxChecked, idxPkg)
	}
}
