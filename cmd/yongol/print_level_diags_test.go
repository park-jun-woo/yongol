//ff:func feature=cli type=test control=iteration dimension=1
//ff:what printLevelDiags output format test — 4 file/line prefix cases + level filtering
package main

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func TestPrintLevelDiagsFileLine(t *testing.T) {
	cases := []struct {
		name string
		diag diagnostic.Diagnostic
		want string
	}{
		{
			name: "file and line",
			diag: diagnostic.Diagnostic{
				File: "service/gig/create.ssac", Line: 42,
				Level: diagnostic.LevelError, Message: "X-99: bad",
			},
			want: "  - service/gig/create.ssac:42: X-99: bad\n",
		},
		{
			name: "file only",
			diag: diagnostic.Diagnostic{
				File: "api/openapi.yaml", Line: 0,
				Level: diagnostic.LevelError, Message: "X-99: bad",
			},
			want: "  - api/openapi.yaml: X-99: bad\n",
		},
		{
			name: "no location",
			diag: diagnostic.Diagnostic{
				File: "", Line: 0,
				Level: diagnostic.LevelError, Message: "X-99: bad",
			},
			want: "  - X-99: bad\n",
		},
		{
			name: "with advice",
			diag: diagnostic.Diagnostic{
				File: "db/users.sql", Line: 5,
				Level: diagnostic.LevelError, Message: "D-2: missing NOT NULL → Advice: Add NOT NULL constraint to the column",
			},
			want: "  - db/users.sql:5: D-2: missing NOT NULL\n      ↳ Advice: Add NOT NULL constraint to the column\n",
		},
	}
	for _, tc := range cases {
		runPrintLevelDiagsCase(t, tc.name, tc.diag, tc.want)
	}
}

// TestPrintLevelDiagsLevelFilter exercises the level-mismatch (continue) branch:
// diagnostics whose Level differs from the requested level are skipped entirely,
// so no "[step]" header or "- ..." line is emitted.
func TestPrintLevelDiagsLevelFilter(t *testing.T) {
	t.Run("mismatched level skipped", func(t *testing.T) {
		var buf bytes.Buffer
		step := validate.StepResult{
			Name: "filter-step",
			Diagnostics: []diagnostic.Diagnostic{
				{File: "db/x.sql", Line: 1, Level: diagnostic.LevelWarning, Message: "W-1: warn only"},
			},
		}
		// Request errors, but only a warning is present → fully skipped.
		printLevelDiags(&buf, step, diagnostic.LevelError)
		if got := buf.String(); got != "" {
			t.Fatalf("expected empty output when no diag matches level, got: %q", got)
		}
	})
	t.Run("matching level after a mismatch", func(t *testing.T) {
		var buf bytes.Buffer
		step := validate.StepResult{
			Name: "mixed-step",
			Diagnostics: []diagnostic.Diagnostic{
				{File: "db/x.sql", Line: 1, Level: diagnostic.LevelWarning, Message: "W-1: warn"},
				{File: "db/y.sql", Line: 2, Level: diagnostic.LevelError, Message: "E-1: err"},
			},
		}
		printLevelDiags(&buf, step, diagnostic.LevelError)
		got := buf.String()
		if got == "" {
			t.Fatal("expected output for the matching error diagnostic")
		}
		// The warning must be skipped; only the error line should appear.
		if want := "E-1: err"; !bytes.Contains(buf.Bytes(), []byte(want)) {
			t.Fatalf("expected output to contain %q, got: %q", want, got)
		}
		if bytes.Contains(buf.Bytes(), []byte("W-1: warn")) {
			t.Fatalf("warning diagnostic should have been skipped, got: %q", got)
		}
	})
}
