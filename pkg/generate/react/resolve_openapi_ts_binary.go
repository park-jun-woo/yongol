//ff:func feature=gen-react type=loader control=iteration dimension=1
//ff:what resolveOpenapiTsBinary — openapi-typescript 바이너리 해결 (env → PATH → npx)

package react

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

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
