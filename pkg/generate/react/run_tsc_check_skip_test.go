//ff:func feature=gen-react type=test control=sequence
//ff:what RunTscCheck — node_modules 미설치 시 게이트 graceful skip(nil) 검증 (BUG-137 Phase041)

package react

import "testing"

func TestRunTscCheck_SkipsWithoutNodeModules(t *testing.T) {
	// No node_modules → graceful skip (warn, not fail), regardless of tsc.
	artifacts := writeFrontendFixture(t, "const n: number = 'x'\n", false)
	if err := RunTscCheck(artifacts); err != nil {
		t.Fatalf("expected skip (nil) without node_modules, got: %v", err)
	}
}
