//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestOpsReferenceBody_Empty
package ssac

import (
	"testing"
)

func TestOpsReferenceBody_Empty(t *testing.T) {
	if opsReferenceBody(nil) {
		t.Error("nil ops should return false")
	}
}
