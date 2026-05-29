//ff:func feature=gen-ffhash type=test control=sequence
//ff:what test: TestInjectUpdatesExistingLine — 기존 //ff:checked 라인의 hash 값을 갱신 (중복 생성 금지)

package ffhash

import (
	"bytes"
	"strings"
	"testing"
)

func TestInjectUpdatesExistingLine(t *testing.T) {
	stale := []byte("/" + "/ff:func feature=demo type=handler control=sequence\n" +
		"/" + "/ff:what Demo — demo function\n" +
		"/" + "/ff:checked llm=yongol-gen hash=deadbeef\n" +
		"package demo\n\n" +
		"func Demo(x int) int {\n\treturn x + 1\n}\n")
	out := InjectCheckedLine(stale)
	if bytes.Contains(out, []byte("hash=deadbeef")) {
		t.Error("expected stale hash to be replaced")
	}
	if strings.Count(string(out), "/"+"/ff:checked") != 1 {
		t.Error("expected exactly one //ff:checked line")
	}
}
