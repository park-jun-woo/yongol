//ff:func feature=gen-react type=generator control=sequence
//ff:what openapi-typescript 를 spawn 하여 src/types/api.d.ts 생성. 실패 시 fail-fast 에러.

package react

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// runOpenAPITypescript spawns `npx openapi-typescript <spec> -o <dest>`.
// The binary is resolved from YONGOL_OPENAPI_TS_PROJECT_DIR/node_modules,
// then PATH. Missing installation yields a precise, actionable error —
// silent skip would leave api.ts importing non-existent types and confuse
// downstream consumers.
func runOpenAPITypescript(specPath, destPath string) error {
	if err := os.MkdirAll(filepath.Dir(destPath), 0o755); err != nil {
		return fmt.Errorf("create types dir: %w", err)
	}

	bin, env, err := resolveOpenapiTsBinary()
	if err != nil {
		// Emit a stub with a clear reason instead of crashing the whole
		// generate pipeline — the stub still makes the import compile
		// loosely (any) so developers can run `npm install` and retry.
		writeOpenapiTsStub(destPath, err)
		return fmt.Errorf("openapi-typescript not available: %w", err)
	}

	cmd := exec.Command(bin[0], append(bin[1:], specPath, "-o", destPath)...)
	cmd.Env = append(os.Environ(), env...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("openapi-typescript failed: %v: %s", err, strings.TrimSpace(stderr.String()))
	}
	return nil
}

// resolveOpenapiTsBinary returns (argv, env) for invoking openapi-typescript.
// Order:
//
//  1. YONGOL_OPENAPI_TS_PROJECT_DIR/node_modules/.bin/openapi-typescript
//  2. YONGOL_SWC_PROJECT_DIR/node_modules/.bin/openapi-typescript  (reuse)
//  3. `openapi-typescript` on PATH
//  4. `npx --yes openapi-typescript` — deterministic if npx is installed.
func resolveOpenapiTsBinary() ([]string, []string, error) {
	candidates := []string{
		os.Getenv("YONGOL_OPENAPI_TS_PROJECT_DIR"),
		os.Getenv("YONGOL_SWC_PROJECT_DIR"),
	}
	for _, dir := range candidates {
		if dir == "" {
			continue
		}
		local := filepath.Join(dir, "node_modules", ".bin", "openapi-typescript")
		if _, err := os.Stat(local); err == nil {
			return []string{local}, nil, nil
		}
	}
	if p, err := exec.LookPath("openapi-typescript"); err == nil {
		return []string{p}, nil, nil
	}
	if _, err := exec.LookPath("npx"); err == nil {
		return []string{"npx", "--yes", "openapi-typescript"}, nil, nil
	}
	return nil, nil, fmt.Errorf("install Node.js (>=18) and run `npm install --save-dev openapi-typescript` (or ensure `npx` is on PATH)")
}

// writeOpenapiTsStub writes a minimal types/api.d.ts that makes TypeScript
// compile (barely) while surfacing the failure reason in a comment. This
// is strictly a fallback so AI iteration loops aren't blocked by a missing
// node_modules install; the generate command still returns a non-nil error
// so automation surfaces the issue.
func writeOpenapiTsStub(destPath string, reason error) {
	content := "// openapi-typescript could not run: " + reason.Error() + "\n" +
		"// Install it in your frontend project: npm install --save-dev openapi-typescript\n" +
		"export type paths = Record<string, any>\n"
	_ = os.WriteFile(destPath, []byte(content), 0o644)
}
