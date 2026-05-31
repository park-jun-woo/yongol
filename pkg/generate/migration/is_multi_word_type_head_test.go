//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestIsMultiWordTypeHead — CHARACTER/TIMESTAMP/TIME/DOUBLE 만 true
package migration

import (
	"testing"
)

func TestIsMultiWordTypeHead(t *testing.T) {
	for _, s := range []string{"CHARACTER", "TIMESTAMP", "TIME", "DOUBLE"} {
		if !isMultiWordTypeHead(s) {
			t.Errorf("isMultiWordTypeHead(%q) = false, want true", s)
		}
	}
	for _, s := range []string{"INTEGER", "TEXT", "character"} {
		if isMultiWordTypeHead(s) {
			t.Errorf("isMultiWordTypeHead(%q) = true, want false", s)
		}
	}
}
