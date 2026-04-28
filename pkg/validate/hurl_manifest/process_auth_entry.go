//ff:func feature=validate type=rule control=sequence topic=hurl-manifest
//ff:what processAuthEntry — 한 hurl entry 에 대해 auth 상태 갱신 / WARNING append

package hurl_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// processAuthEntry decides whether entry e represents an auth-issuing
// step (updates authIssued) or a protected call (appends a WARNING when
// no prior auth exists). Returns the updated authIssued value for the
// caller's running state.
func processAuthEntry(e hurl.HurlEntry, authIssued bool, diags *[]diagnostic.Diagnostic) bool {
	if isAuthPath(e.Path) {
		if is2xx(e.StatusCode) {
			return true
		}
		return authIssued
	}
	if authIssued || carriesAuthHeader(e.Headers) {
		return authIssued
	}
	*diags = append(*diags, diagnostic.Diagnostic{
		File:    e.File,
		Line:    e.Line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: "[XOH-06] " + e.Method + " " + e.Path + " called without a prior auth step",
		Advice:  "Add a preceding login request, reuse a captured token via Authorization: Bearer {{token}}, or present the session cookie",
	})
	return authIssued
}
