//ff:func feature=gen-react type=generator control=sequence
//ff:what runOpenAPITypescript — openapi-typescript 를 spawn 하여 src/types/api.d.ts 생성

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
