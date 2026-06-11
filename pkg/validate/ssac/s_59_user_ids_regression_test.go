//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what TestS59_UserIDsRegression — BUG-099 회귀: sqlc UserIDS 정본 / GoPascal UserIds FAIL+did-you-mean 확인

package ssac

import (
	"strings"
	"testing"
)

// TestS59_UserIDsRegression covers the BUG-099 case: the canonical sqlc field
// is "UserIDS"; the ToGoPascal spelling "UserIds" must fail with a did-you-mean
// pointing at "UserIDS".
func TestS59_UserIDsRegression(t *testing.T) {
	schema := map[string][]string{
		"SSaC.var.Notify.group": {"ID", "UserIDS"},
	}

	t.Run("sqlc UserIDS passes", func(t *testing.T) {
		fs := callFS("Notify", "group", "UserIDS", schema)
		assertNoDiag(t, s59DottedField(fs), "[S-59]")
	})

	t.Run("GoPascal UserIds fails with did-you-mean UserIDS", func(t *testing.T) {
		fs := callFS("Notify", "group", "UserIds", schema)
		diags := s59DottedField(fs)
		assertDiag(t, diags, "[S-59]")
		if len(diags) != 1 {
			t.Fatalf("want 1 diag, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, `did you mean "UserIDS"?`) {
			t.Errorf("Message missing did-you-mean UserIDS: %q", diags[0].Message)
		}
	})
}
