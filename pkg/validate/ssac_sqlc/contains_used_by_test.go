//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestContainsUsedBy(t *testing.T) {
	usedBy := []string{"GET", "POST"}
	if !containsUsedBy(usedBy, "POST") {
		t.Error("expected POST present")
	}
	if containsUsedBy(usedBy, "DELETE") {
		t.Error("expected DELETE absent")
	}
	if containsUsedBy(nil, "GET") {
		t.Error("nil slice should not contain anything")
	}
}
