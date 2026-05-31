//ff:func feature=cli type=test control=sequence
//ff:what printLevelDiags output format test — 4 file/line prefix cases + level filtering
package main

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

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
