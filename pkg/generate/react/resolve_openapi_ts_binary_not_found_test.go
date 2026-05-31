//ff:func feature=gen-react type=test control=sequence
//ff:what resolveOpenapiTsBinary env→PATH→npx 해결 순서·부재 에러 검증
package react

import (
	"os/exec"
	"testing"
)

func TestResolveOpenapiTsBinaryNotFound(t *testing.T) {
	isolatePATH(t)

	// Guard: ensure neither binary is reachable in the isolated PATH.
	if _, err := exec.LookPath("npx"); err == nil {
		t.Skip("npx unexpectedly resolvable in isolated PATH")
	}

	argv, env, err := resolveOpenapiTsBinary()
	if err == nil {
		t.Fatalf("expected error when no binary available, got argv=%v env=%v", argv, env)
	}
}
