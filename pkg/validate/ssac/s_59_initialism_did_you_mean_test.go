//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what TestS59_InitialismDidYouMean — S-59: sqlc 표기 PASS / GoPascal 표기 FAIL+did-you-mean / 오타 FAIL+제안없음 (BUG-123)

package ssac

import (
	"strings"
	"testing"
)

func TestS59_InitialismDidYouMean(t *testing.T) {
	// Schema holds the canonical sqlc field spelling.
	schema := map[string][]string{
		"SSaC.var.Export.site": {"ID", "QueueExportRepoUrl"},
	}

	t.Run("sqlc spelling passes", func(t *testing.T) {
		fs := callFS("Export", "site", "QueueExportRepoUrl", schema)
		assertNoDiag(t, s59DottedField(fs), "[S-59]")
	})

	t.Run("GoPascal spelling fails with did-you-mean", func(t *testing.T) {
		fs := callFS("Export", "site", "QueueExportRepoURL", schema)
		diags := s59DottedField(fs)
		assertDiag(t, diags, "[S-59]")
		if len(diags) != 1 {
			t.Fatalf("want 1 diag, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, `did you mean "QueueExportRepoUrl"?`) {
			t.Errorf("Message missing did-you-mean suggestion: %q", diags[0].Message)
		}
		if !strings.Contains(diags[0].Advice, "QueueExportRepoUrl") {
			t.Errorf("Advice missing suggestion: %q", diags[0].Advice)
		}
	})

	t.Run("pure typo fails without suggestion", func(t *testing.T) {
		fs := callFS("Export", "site", "Bogus", schema)
		diags := s59DottedField(fs)
		assertDiag(t, diags, "[S-59]")
		if len(diags) != 1 {
			t.Fatalf("want 1 diag, got %d", len(diags))
		}
		if strings.Contains(diags[0].Message, "did you mean") {
			t.Errorf("non-near-match should not suggest, got: %q", diags[0].Message)
		}
	})
}
