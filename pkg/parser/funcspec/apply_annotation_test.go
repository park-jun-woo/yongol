//ff:func feature=funcspec type=test control=sequence
//ff:what applyAnnotation / parseCommentGroup / countFuncAnnotations — @func/@error/@description
package funcspec

import (
	"testing"
)

func TestApplyAnnotation(t *testing.T) {
	t.Run("func sets name and line", func(t *testing.T) {
		var s FuncSpec
		applyAnnotation("@func hashPassword", 7, &s)
		if s.Name != "hashPassword" || s.Line != 7 {
			t.Errorf("spec = %+v", s)
		}
	})
	t.Run("error sets status", func(t *testing.T) {
		var s FuncSpec
		applyAnnotation("@error 422", 1, &s)
		if s.ErrStatus != 422 {
			t.Errorf("ErrStatus = %d, want 422", s.ErrStatus)
		}
	})
	t.Run("error non-numeric ignored", func(t *testing.T) {
		var s FuncSpec
		applyAnnotation("@error notanumber", 1, &s)
		if s.ErrStatus != 0 {
			t.Errorf("ErrStatus = %d, want 0", s.ErrStatus)
		}
	})
	t.Run("description", func(t *testing.T) {
		var s FuncSpec
		applyAnnotation("@description hashes a password", 1, &s)
		if s.Description != "hashes a password" {
			t.Errorf("Description = %q", s.Description)
		}
	})
	t.Run("unknown ignored", func(t *testing.T) {
		var s FuncSpec
		applyAnnotation("@unknown x", 1, &s)
		if s.Name != "" || s.Description != "" || s.ErrStatus != 0 {
			t.Errorf("spec should be empty: %+v", s)
		}
	})
}
