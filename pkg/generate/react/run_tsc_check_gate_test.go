//ff:func feature=gen-react type=test control=sequence
//ff:what RunTscCheck — 정상 소스는 tsc 게이트 통과, 타입 깨진 소스(TS2322)는 실패 검증 (BUG-137 Phase041)

package react

import "testing"

func TestRunTscCheck_PassesAndFails(t *testing.T) {
	if resolveTscArgv() == nil {
		t.Skip("tsc unavailable (no node_modules/.bin/tsc and no npx)")
	}

	// Clean source type-checks → gate passes.
	okArtifacts := writeFrontendFixture(t, "const n: number = 1\nvoid n\n", true)
	if err := RunTscCheck(okArtifacts); err != nil {
		t.Fatalf("clean frontend should pass tsc gate, got: %v", err)
	}

	// Broken source (TS2322) → gate fails.
	badArtifacts := writeFrontendFixture(t, "const n: number = 'x'\nvoid n\n", true)
	if err := RunTscCheck(badArtifacts); err == nil {
		t.Fatal("type-broken frontend should fail tsc gate, got nil")
	}
}
