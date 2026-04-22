//ff:func feature=cli type=reporter control=iteration dimension=1
//ff:what printSSOTSummary — tallies SSOT element counts from Fullstack and prints the Summary section

package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// printSSOTSummary writes an "SSOT Summary" block listing one line per
// source-of-truth kind with a human readable count. Nil parse results
// render as "0" so an incomplete specs directory still produces a
// meaningful dashboard.
func printSSOTSummary(w io.Writer, fs *yongol.Fullstack) {
	fmt.Fprintln(w, "SSOT Summary")
	endpoints := 0
	if fs.OpenAPIDoc != nil {
		for _, pi := range fs.OpenAPIDoc.Paths.Map() {
			endpoints += len(pi.Operations())
		}
	}
	policyRules := 0
	for _, p := range fs.ParsedPolicies {
		policyRules += len(p.Rules)
	}
	fmt.Fprintf(w, "  OpenAPI      %d endpoints\n", endpoints)
	fmt.Fprintf(w, "  DDL          %d tables\n", len(fs.DDLTables))
	fmt.Fprintf(w, "  SSaC         %d service functions\n", len(fs.ServiceFuncs))
	fmt.Fprintf(w, "  States       %d diagrams\n", len(fs.StateDiagrams))
	fmt.Fprintf(w, "  Policy       %d files (%d rules)\n", len(fs.ParsedPolicies), policyRules)
	fmt.Fprintf(w, "  Scenario     %d hurl files\n", len(fs.HurlFiles))
	fmt.Fprintf(w, "  Func         %d funcs\n", len(fs.ProjectFuncSpecs))
}
