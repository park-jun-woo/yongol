//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestStmlDesignHelpers — unit tests for the pure stml_design helper functions
package stml_design

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

func TestSortedTypoKeys(t *testing.T) {
	m := map[string]design.TypographyToken{"y": {}, "b": {}}
	if got := sortedTypoKeys(m); !reflect.DeepEqual(got, []string{"b", "y"}) {
		t.Errorf("sortedTypoKeys = %v", got)
	}
}
