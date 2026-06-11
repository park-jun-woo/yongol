//ff:func feature=gen-react type=test control=sequence
//ff:what TestResType_201Body_TscInteg — Res<K> 가 201 본문을 채택하고 204 는 void 임을 tsc 로 증명 (BUG-128/Phase039)

package react

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestResType_201Body_TscInteg pins the BUG-128 regression at the type level:
// the Res<K> helper emitted by writeReqResTypes must adopt the 201 success
// body (so `.id` is accessible) while a 204/no-body op stays `void`. A pure Go
// string assertion cannot observe this — per-op body-vs-void is a TypeScript
// conditional-type computation — so this drives the real openapi-typescript →
// tsc toolchain over a fixture spec carrying a 201+body op and a 204 op.
//
// Skips (not fails) when openapi-typescript or tsc are unavailable, matching
// the other integ tests; the static string assertions in
// write_api_client_typed_sig_req_res_test.go remain the always-on guard.
func TestResType_201Body_TscInteg(t *testing.T) {
	dir := t.TempDir()

	spec := `openapi: 3.0.0
info: { title: phase039, version: "1" }
paths:
  /items:
    post:
      operationId: CreateItem
      responses:
        '201':
          description: created
          content:
            application/json:
              schema:
                type: object
                properties:
                  id: { type: integer }
                required: [id]
  /ping:
    post:
      operationId: Ping
      responses:
        '204':
          description: no content
`
	specPath := filepath.Join(dir, "spec.yaml")
	if err := os.WriteFile(specPath, []byte(spec), 0o644); err != nil {
		t.Fatal(err)
	}

	apiDTS := filepath.Join(dir, "api.d.ts")
	if err := runOpenAPITypescript(specPath, apiDTS); err != nil {
		t.Skipf("openapi-typescript unavailable: %v", err)
	}

	// probe.ts = the real Res<K>/Req<K> helpers (referencing only `operations`)
	// plus type-level assertions. `createId: number = (... Res<'CreateItem'>).id`
	// fails to compile with the old 200-only Res (void has no `.id`); `Ping`
	// (204) must remain assignable from `void`.
	var b strings.Builder
	writeReqResTypes(&b)
	probe := "import type { operations } from './api'\n\n" +
		b.String() +
		"const createId: number = (undefined as unknown as Res<'CreateItem'>).id\n" +
		"const pingRes: Res<'Ping'> = undefined as void\n" +
		"void createId; void pingRes;\n"
	if err := os.WriteFile(filepath.Join(dir, "probe.ts"), []byte(probe), 0o644); err != nil {
		t.Fatal(err)
	}

	argv := resolveTscArgv()
	if argv == nil {
		t.Skip("tsc unavailable (no node_modules/.bin/tsc and no npx)")
	}
	args := append(argv[1:], "--noEmit", "--strict", "--skipLibCheck", "probe.ts")
	cmd := exec.Command(argv[0], args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("tsc rejected the 201-body probe (Res<K> regression): %v\n%s", err, out)
	}
}
