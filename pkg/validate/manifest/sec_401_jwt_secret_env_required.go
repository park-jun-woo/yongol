//ff:func feature=validate type=rule control=sequence topic=manifest-auth
//ff:what SEC-401 — backend.auth.secret literal is forbidden; only secret_env is allowed

package manifest

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// sec401JWTSecretEnvRequired flags manifests that embed a literal JWT
// signing key under `backend.auth.secret`. Secrets must come from the
// process environment via `secret_env`; checking literals into the
// manifest leaks them into git history and prevents rotation. The rule
// is a hard error — yongol's generator has no fallback path for a
// literal secret since the Phase002 JWT Hardening removed it.
//
// Detection: re-parses the manifest.yaml as a yaml.Node and walks to
// backend.auth, then asserts the `secret` key is absent. `secret_env`
// is allowed and the intended field.
func sec401JWTSecretEnvRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.SpecsDir == "" {
		return nil
	}
	if fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return nil
	}

	path := filepath.Join(fs.SpecsDir, "manifest.yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	auth := pmanifest.FindAuthNode(data)
	if auth == nil || auth.Kind != yaml.MappingNode {
		return nil
	}

	// Walk the key list looking for a literal `secret:` assignment.
	for i := 0; i+1 < len(auth.Content); i += 2 {
		if auth.Content[i].Value == "secret" {
			return []diagnostic.Diagnostic{{
				File:    "manifest.yaml",
				Line:    auth.Content[i].Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[SEC-401] backend.auth.secret literal is forbidden — specify only an environment variable name in secret_env",
				Advice:  "Remove the secret field and use `secret_env: JWT_SECRET` instead",
			}}
		}
	}
	return nil
}
