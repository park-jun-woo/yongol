//ff:func feature=tsx-parser type=test control=sequence
//ff:what Parse — empty_page.tsx 는 Calls/FormFields/Imports 모두 비어있음

package tsx

import "testing"

func TestParse_EmptyPage(t *testing.T) {
	if !swcAvailable(t) {
		return
	}
	got, err := Parse("testdata/empty_page.tsx")
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(got.Calls) != 0 || len(got.FormFields) != 0 || len(got.Imports) != 0 {
		t.Errorf("empty page should have nothing, got %+v", got)
	}
}
