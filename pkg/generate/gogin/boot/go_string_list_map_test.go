//ff:func feature=gen-gogin type=test control=sequence
//ff:what goStringListMap — map[string][]string 를 Go 리터럴로 렌더 (deterministic)
package boot

import (
	"strings"
	"testing"
)

func TestGoStringListMap(t *testing.T) {
	if got := goStringListMap(nil); got != `map[string][]string{}` {
		t.Errorf("nil map = %q, want empty literal", got)
	}
	if got := goStringListMap(map[string][]string{}); got != `map[string][]string{}` {
		t.Errorf("empty map = %q, want empty literal", got)
	}

	out := goStringListMap(map[string][]string{
		"b": {"y"},
		"a": {"x", "z"},
	})
	if !strings.HasPrefix(out, "map[string][]string{") {
		t.Fatalf("missing literal prefix: %q", out)
	}
	// Keys are emitted in sorted order for deterministic codegen.
	ai := strings.Index(out, `"a":`)
	bi := strings.Index(out, `"b":`)
	if ai < 0 || bi < 0 || ai > bi {
		t.Fatalf("keys not sorted: %q", out)
	}
	if !strings.Contains(out, `[]string{"x", "z"}`) {
		t.Errorf("value for a not rendered: %q", out)
	}
}
