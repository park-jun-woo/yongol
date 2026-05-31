//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestStmlDesignHelpers — unit tests for the pure stml_design helper functions
package stml_design

import (
	"reflect"
	"testing"
)

func TestSortedKeys(t *testing.T) {
	m := map[string]string{"c": "1", "a": "2", "b": "3"}
	if got := sortedKeys(m); !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("sortedKeys = %v", got)
	}
	if got := sortedKeys(map[string]string{}); len(got) != 0 {
		t.Errorf("empty map should give empty slice, got %v", got)
	}
}
