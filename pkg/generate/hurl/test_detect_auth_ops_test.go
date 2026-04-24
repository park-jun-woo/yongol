//ff:func feature=gen-hurl type=test control=iteration dimension=1
//ff:what TestDetectAuthOps — SSaC shape 기반 signup/login 감지 table-driven 검증 (BUG-023 회귀 방지)

package hurl

import (
	"testing"
)

// TestDetectAuthOps exercises the shape-detection matrix described in
// plans/gen/hurl02/Phase003-AuthOpShapeDetection.md. Naming variance
// (Register / Signup / Join / EnrollStudent / SignIn) must not affect
// classification — only the SSaC body shape does.
func TestDetectAuthOps(t *testing.T) {
	cases := detectAuthOpsFixtures()
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			runDetectAuthOpsCase(t, tc)
		})
	}
}
