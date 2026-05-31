//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestOpsReferencePath_Empty
package ssac

import (
	"testing"
)

func TestOpsReferencePath_Empty(t *testing.T) {
	if opsReferencePath(nil) {
		t.Error("nil ops should return false")
	}
}
