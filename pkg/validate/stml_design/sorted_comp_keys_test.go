//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestStmlDesignHelpers — unit tests for the pure stml_design helper functions
package stml_design

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

func TestSortedCompKeys(t *testing.T) {
	m := map[string]design.ComponentToken{"z": {}, "a": {}}
	if got := sortedCompKeys(m); !reflect.DeepEqual(got, []string{"a", "z"}) {
		t.Errorf("sortedCompKeys = %v", got)
	}
}
