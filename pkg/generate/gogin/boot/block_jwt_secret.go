//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockJWTSecret — JWT secret 환경변수 읽기 블록

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// blockJWTSecret produces the JWT secret env var read. Active when
// manifest.backend.auth is configured.
//
// Phase003 — auth.VerifyToken reads os.Getenv(SecretEnv) on every call, so
// this block does not pass the secret value anywhere. It remains as a
// fail-fast check: empty or short secrets abort bootstrap before the server
// accepts traffic. The read value is intentionally discarded via _ after
// the length assertion.
func blockJWTSecret(fs *yongol.Fullstack) MainBlock {
	envVar := "JWT_SECRET"
	if fs.Manifest != nil && fs.Manifest.Backend.Auth != nil && fs.Manifest.Backend.Auth.SecretEnv != "" {
		envVar = fs.Manifest.Backend.Auth.SecretEnv
	}
	return MainBlock{
		Name:   "jwt-secret",
		Active: hasAuth,
		Lines: []string{
			`if v := os.Getenv("` + envVar + `"); v == "" {`,
			`	slog.Error("` + envVar + ` is required")`,
			`	os.Exit(1)`,
			`} else if len(v) < 32 {`,
			`	slog.Error("` + envVar + ` must be at least 32 characters", "length", len(v))`,
			`	os.Exit(1)`,
			`}`,
		},
	}
}
