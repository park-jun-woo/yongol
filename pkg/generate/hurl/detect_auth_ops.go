//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what detectAuthOps — SSaC 본문 shape 로 signup/login 후보 감지 (이름 독립)

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// detectAuthOps scans every public OpenAPI operation that carries a
// `password` request-body field and classifies it by the shape of its
// SSaC implementation:
//
//   - has `@verify-password`                       → login
//   - has `@call auth.HashPassword(...)`           → signup
//   - has both (combined signup-with-auto-login)   → signup + WARNING
//
// Operations whose SSaC body matches neither pattern are skipped (with
// a WARNING) — they are auth-shape but not yongol's standard flow
// (e.g. magic-link, refresh-token endpoints).
//
// When multiple candidates exist for the same role, the first by
// operationId alphabetical order wins (deterministic) and a WARNING is
// emitted listing all candidates and the chosen one.
//
// Returns (nil, nil, nil) when OpenAPIDoc is absent. Callers treat a
// nil signup OR nil login as "skip that step" — the smoke will still
// emit whichever side was detected.
func detectAuthOps(fs *yongol.Fullstack) (signup, login *detectedAuthOp, warnings []string) {
	if fs == nil || fs.OpenAPIDoc == nil {
		return nil, nil, nil
	}

	funcsByOpID := indexServiceFuncsByOpID(fs.ServiceFuncs)

	var signupCands []detectedAuthOp
	var loginCands []detectedAuthOp

	for path, pathItem := range fs.OpenAPIDoc.Paths.Map() {
		if pathItem == nil {
			continue
		}
		classifyPathItemAuthOps(path, pathItem, funcsByOpID, &signupCands, &loginCands, &warnings)
	}

	signup = pickCandidate(signupCands, "signup", &warnings)
	login = pickCandidate(loginCands, "login", &warnings)
	return signup, login, warnings
}
