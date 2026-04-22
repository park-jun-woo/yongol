//ff:func feature=cli type=test-helper control=sequence
//ff:what printLevelDiags single-case runner helper — buf/step/assert bundle
package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// runPrintLevelDiagsCase executes a single printLevelDiags test case.
// Extracted from TestPrintLevelDiagsFileLine range body to satisfy Q4.
func runPrintLevelDiagsCase(t *testing.T, name string, diag diagnostic.Diagnostic, want string) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		var buf bytes.Buffer
		step := validate.StepResult{
			Name:        "test-step",
			Diagnostics: []diagnostic.Diagnostic{diag},
		}
		printLevelDiags(&buf, step, diagnostic.LevelError)
		got := buf.String()
		if !strings.Contains(got, want) {
			t.Fatalf("expected output to contain:\n%q\ngot:\n%q", want, got)
		}
	})
}
