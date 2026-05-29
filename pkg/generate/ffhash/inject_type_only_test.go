//ff:func feature=gen-ffhash type=test control=sequence
//ff:what test: TestInjectSkipsTypeOnlyFiles — func 이 없는 파일은 hash 계산 불가이므로 스킵

package ffhash

import (
	"bytes"
	"testing"
)

func TestInjectSkipsTypeOnlyFiles(t *testing.T) {
	src := []byte("/" + "/ff:type feature=api type=model\n" +
		"/" + "/ff:what Foo — demo model\n" +
		"package api\n\n" +
		"type Foo struct{ ID int64 }\n")
	out := InjectCheckedLine(src)
	if bytes.Contains(out, []byte("/"+"/ff:checked")) {
		t.Errorf("expected no injection for type-only file, got:\n%s", out)
	}
}
