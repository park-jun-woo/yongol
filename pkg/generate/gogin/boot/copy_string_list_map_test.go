//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what copyStringListMap — map[string][]string 의 deep copy (generator shared-state 보호)

package boot

import "testing"

func TestCopyStringListMap(t *testing.T) {
	if copyStringListMap(nil) != nil {
		t.Fatalf("nil input should return nil")
	}

	in := map[string][]string{"a": {"1", "2"}, "b": {"3"}}
	out := copyStringListMap(in)

	if len(out) != len(in) {
		t.Fatalf("len = %d, want %d", len(out), len(in))
	}
	if !equalStrings(out["a"], []string{"1", "2"}) || !equalStrings(out["b"], []string{"3"}) {
		t.Fatalf("values not copied faithfully: %v", out)
	}

	// Mutating the copy must not affect the source (deep copy guarantee).
	out["a"][0] = "mutated"
	if in["a"][0] != "1" {
		t.Errorf("source mutated through copy: %v", in["a"])
	}
}
