//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestSequenceTag(t *testing.T) {
	if got := sequenceTag("call"); got != "call" {
		t.Errorf("sequenceTag = %q, want call", got)
	}
}
