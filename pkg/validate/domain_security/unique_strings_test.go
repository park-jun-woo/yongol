//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestDomainSecurityHelpers — unit tests for the pure domain_security helper functions
package domain_security

import (
	"reflect"
	"testing"
)

func TestUniqueStrings(t *testing.T) {
	got := uniqueStrings([]string{"a", "b", "a", "c", "b"})
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("uniqueStrings = %v, want %v", got, want)
	}
	if got := uniqueStrings(nil); got != nil {
		t.Errorf("nil input → %v, want nil", got)
	}
}
