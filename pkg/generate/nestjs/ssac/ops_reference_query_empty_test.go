//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestOpsReferenceQuery_Empty
package ssac

import (
	"testing"
)

func TestOpsReferenceQuery_Empty(t *testing.T) {
	if opsReferenceQuery(nil) {
		t.Error("nil ops should return false")
	}
}
