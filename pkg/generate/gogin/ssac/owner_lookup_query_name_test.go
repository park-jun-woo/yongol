//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what ownerLookupQueryName 단위 테스트 (OwnerLookup<PascalResource> 생성)

package ssac

import "testing"

func TestOwnerLookupQueryName(t *testing.T) {
	cases := map[string]string{
		"workflow":      "OwnerLookupWorkflow",
		"execution_log": "OwnerLookupExecutionLog",
		"User":          "OwnerLookupUser",
	}
	for in, want := range cases {
		if got := ownerLookupQueryName(in); got != want {
			t.Errorf("ownerLookupQueryName(%q) = %q, want %q", in, got, want)
		}
	}
}
