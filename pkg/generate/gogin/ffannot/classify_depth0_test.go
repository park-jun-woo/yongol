//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestClassifyDepth0 — depth-0 라인 for/switch/none 분류 검증
package ffannot

import (
	"testing"
)

func TestClassifyDepth0(t *testing.T) {
	cases := map[string]string{
		"for i := 0; i < n; i++ {": ControlIteration,
		"for{":                     ControlIteration,
		"for(":                     ControlIteration,
		"for":                      ControlIteration,
		"switch x {":               ControlSelection,
		"switch{":                  ControlSelection,
		"switch":                   ControlSelection,
		"if x {":                   "",
		"return nil":               "",
		"format := 1":              "", // "for" must be a prefix word, not substring
	}
	for line, want := range cases {
		if got := classifyDepth0(line); got != want {
			t.Errorf("classifyDepth0(%q) = %q, want %q", line, got, want)
		}
	}
}
