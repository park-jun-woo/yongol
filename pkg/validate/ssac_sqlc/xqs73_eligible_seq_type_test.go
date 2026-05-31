//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"testing"
)

func TestXqs73EligibleSeqType(t *testing.T) {
	for _, ty := range []string{"get", "post", "put"} {
		if !xqs73EligibleSeqType(ty) {
			t.Errorf("%q should be eligible", ty)
		}
	}
	for _, ty := range []string{"delete", "call", "empty"} {
		if xqs73EligibleSeqType(ty) {
			t.Errorf("%q should not be eligible", ty)
		}
	}
}
