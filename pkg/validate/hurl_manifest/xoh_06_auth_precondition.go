//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-manifest
//ff:what XOH-06 — 보호 구간 operation 호출 전에 인증 스텝이 선행되어야 함

package hurl_manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh06AuthPrecondition enforces XOH-06 (WARNING): when the manifest
// declares an auth middleware, any hurl call to a protected endpoint
// must be preceded — within the same file — by either an auth-issuing
// step or a request carrying the auth credential. Heuristics:
//
//   - An "auth step" is an entry whose path contains `/auth/` or
//     `/login` / `/signin` (these are the conventional prefixes) and
//     whose response is 2xx. Users who deviate from convention carry a
//     manual Authorization / Cookie header on every protected call,
//     which the rule also accepts.
//   - A "protected endpoint" is anything that is NOT an auth step.
//     This biases toward WARNINGs on public endpoints that omit an
//     Authorization header, but only when the project has any auth
//     middleware at all. Users can quiet the noise by setting
//     backend.auth to null.
func xoh06AuthPrecondition(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return nil
	}
	byFile := groupByFile(fs.HurlEntries)
	var diags []diagnostic.Diagnostic
	for _, entries := range byFile {
		diags = append(diags, checkFileAuth(entries)...)
	}
	return diags
}

// checkFileAuth walks a single file's entries and emits a WARNING for
// every protected call that precedes any auth step and lacks its own
// auth header.
func checkFileAuth(entries []hurl.HurlEntry) []diagnostic.Diagnostic {
	authIssued := false
	var diags []diagnostic.Diagnostic
	for _, e := range entries {
		if isAuthPath(e.Path) {
			if is2xx(e.StatusCode) {
				authIssued = true
			}
			continue
		}
		if authIssued || carriesAuthHeader(e.Headers) {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    e.File,
			Line:    e.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[XOH-06] " + e.Method + " " + e.Path + " called without a prior auth step",
			Advice:  "Add a preceding login request, reuse a captured token via Authorization: Bearer {{token}}, or present the session cookie",
		})
	}
	return diags
}

// isAuthPath matches the path conventions yongol recognises as auth
// endpoints. Matching the path (not the operationId) lets the rule work
// even when the OpenAPI document was not parsed.
func isAuthPath(p string) bool {
	low := strings.ToLower(p)
	if strings.Contains(low, "/auth/") {
		return true
	}
	for _, suffix := range []string{"/login", "/signin", "/register", "/signup"} {
		if strings.HasSuffix(low, suffix) || strings.Contains(low, suffix+"/") {
			return true
		}
	}
	return false
}

// is2xx reports whether a hurl-asserted status code is 2xx. Empty means
// "not asserted" and is treated as success (hurl default behaviour).
func is2xx(code string) bool {
	if code == "" {
		return true
	}
	return len(code) == 3 && code[0] == '2'
}

// carriesAuthHeader returns true when at least one request header looks
// like it transports auth state — Authorization, Cookie, or the
// __Host-access_token / refresh_token cookies yongol sets by default.
func carriesAuthHeader(headers []hurl.HurlHeader) bool {
	for _, h := range headers {
		name := strings.ToLower(h.Name)
		if name == "authorization" || name == "cookie" {
			return true
		}
	}
	return false
}

// groupByFile bins entries by file so each hurl file is reasoned about
// independently. Hurl does not share captures across files, so auth
// state does not cross this boundary either.
func groupByFile(entries []hurl.HurlEntry) map[string][]hurl.HurlEntry {
	out := map[string][]hurl.HurlEntry{}
	for _, e := range entries {
		out[e.File] = append(out[e.File], e)
	}
	return out
}
