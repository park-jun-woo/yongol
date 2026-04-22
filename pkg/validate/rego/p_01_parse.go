//ff:func feature=validate type=rule control=iteration dimension=1 topic=rego-structural
//ff:what P-1 — detects Rego file parse errors by re-parsing

package rego

import (
	"path/filepath"

	regoparser "github.com/park-jun-woo/yongol/pkg/parser/rego"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// p01Parse re-parses policy/ to surface rego parse diagnostics.
func p01Parse(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if len(fs.ParsedPolicies) > 0 {
		return nil
	}
	_, diags := regoparser.ParsePolicies(filepath.Join(fs.SpecsDir, "policy"))
	for i := range diags {
		diags[i].Phase = diagnostic.PhaseValidate
		diags[i].Message = "[P-1] " + diags[i].Message
		diags[i].Advice = "Write Rego syntax according to the official OPA documentation"
	}
	return diags
}
