//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what detectAuthOps — SSaC 본문 shape 로 signup/login 후보 감지 (이름 독립)

package hurl

import (
	"fmt"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// detectedAuthOp captures the minimum metadata smoke.hurl needs to emit
// an auth step: operationId + HTTP method + URL path. method is always
// "POST" under yongol's DSL (signup/login never use GET).
type detectedAuthOp struct {
	OpID   string
	Path   string
	Method string
}

// authRole is the role tag stored in scenarioCtx.authOpIDs for O(1)
// "is this op an auth op?" lookup during downstream phase building.
type authRole string

const (
	authRoleSignup authRole = "signup"
	authRoleLogin  authRole = "login"
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
		for method, op := range pathItem.Operations() {
			if op == nil || op.OperationID == "" {
				continue
			}
			if !isPublicOp(op) {
				continue
			}
			if !hasPasswordField(op) {
				continue
			}
			fn := funcsByOpID[op.OperationID]
			isLogin := isLoginShape(fn)
			isSignup := isSignupShape(fn)
			switch {
			case isSignup && isLogin:
				signupCands = append(signupCands, detectedAuthOp{OpID: op.OperationID, Path: path, Method: method})
				warnings = append(warnings, fmt.Sprintf(
					"detect_auth_ops: %q declares both @verify-password and @call auth.HashPassword — treating as combined signup (auto-login)",
					op.OperationID))
			case isSignup:
				signupCands = append(signupCands, detectedAuthOp{OpID: op.OperationID, Path: path, Method: method})
				if fn != nil && !hasUserCreatePost(fn) {
					warnings = append(warnings, fmt.Sprintf(
						"detect_auth_ops: %q calls auth.HashPassword but no companion @post <Model>.Create({PasswordHash: ...}) was found",
						op.OperationID))
				}
			case isLogin:
				loginCands = append(loginCands, detectedAuthOp{OpID: op.OperationID, Path: path, Method: method})
			default:
				if fn != nil {
					warnings = append(warnings, fmt.Sprintf(
						"detect_auth_ops: %q looks auth-shaped (public + password field) but SSaC body matches neither signup nor login pattern — skipped",
						op.OperationID))
				}
			}
		}
	}

	signup = pickCandidate(signupCands, "signup", &warnings)
	login = pickCandidate(loginCands, "login", &warnings)
	return signup, login, warnings
}

// indexServiceFuncsByOpID builds a lookup from operationId (ServiceFunc.Name)
// to the parsed SSaC ServiceFunc. yongol's convention pins the .ssac
// function name to the OpenAPI operationId (e.g. `func Signup() {}` ↔
// operationId: Signup), so direct name match is safe.
func indexServiceFuncsByOpID(funcs []ssac.ServiceFunc) map[string]*ssac.ServiceFunc {
	out := make(map[string]*ssac.ServiceFunc, len(funcs))
	for i := range funcs {
		fn := &funcs[i]
		if fn.Name == "" {
			continue
		}
		out[fn.Name] = fn
	}
	return out
}

// pickCandidate returns the deterministic pick from a candidate list:
// alphabetical-first operationId. Emits a WARNING when more than one
// candidate exists (user should resolve ambiguity by renaming).
func pickCandidate(cands []detectedAuthOp, role string, warnings *[]string) *detectedAuthOp {
	if len(cands) == 0 {
		return nil
	}
	sort.SliceStable(cands, func(i, j int) bool {
		return cands[i].OpID < cands[j].OpID
	})
	if len(cands) > 1 {
		names := make([]string, len(cands))
		for i, c := range cands {
			names[i] = c.OpID
		}
		*warnings = append(*warnings, fmt.Sprintf(
			"detect_auth_ops: multiple %s candidates %v — using %q",
			role, names, cands[0].OpID))
	}
	c := cands[0]
	return &c
}
