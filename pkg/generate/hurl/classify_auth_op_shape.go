//ff:func feature=gen-hurl type=util control=selection
//ff:what classifyAuthOpShape — 단일 auth operation 의 shape 분류 + 후보/경고 기록

package hurl

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// classifyAuthOpShape inspects a single auth-shaped operation and records
// the appropriate candidate / warning. Kept as its own function so the
// per-operation case switch lives at depth 1.
func classifyAuthOpShape(
	path, method string,
	op *openapi3.Operation,
	fn *ssac.ServiceFunc,
	signupCands *[]detectedAuthOp,
	loginCands *[]detectedAuthOp,
	warnings *[]string,
) {
	isLogin := isLoginShape(fn)
	isSignup := isSignupShape(fn)
	switch {
	case isSignup && isLogin:
		*signupCands = append(*signupCands, detectedAuthOp{OpID: op.OperationID, Path: path, Method: method})
		*warnings = append(*warnings, fmt.Sprintf(
			"detect_auth_ops: %q declares both @verify-password and @call auth.HashPassword — treating as combined signup (auto-login)",
			op.OperationID))
	case isSignup:
		*signupCands = append(*signupCands, detectedAuthOp{OpID: op.OperationID, Path: path, Method: method})
		if fn != nil && !hasUserCreatePost(fn) {
			*warnings = append(*warnings, fmt.Sprintf(
				"detect_auth_ops: %q calls auth.HashPassword but no companion @post <Model>.Create({PasswordHash: ...}) was found",
				op.OperationID))
		}
	case isLogin:
		*loginCands = append(*loginCands, detectedAuthOp{OpID: op.OperationID, Path: path, Method: method})
	default:
		if fn != nil {
			*warnings = append(*warnings, fmt.Sprintf(
				"detect_auth_ops: %q looks auth-shaped (public + password field) but SSaC body matches neither signup nor login pattern — skipped",
				op.OperationID))
		}
	}
}
