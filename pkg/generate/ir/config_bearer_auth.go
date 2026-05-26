//ff:type feature=gen-ir type=model
//ff:what BearerAuthConfig -- JWT/쿠키/하이브리드 인증 미들웨어 설정

package ir

// BearerAuthConfig holds the resolved auth middleware configuration.
type BearerAuthConfig struct {
	// Mode is one of "bearer", "cookie", or "hybrid".
	Mode string

	// SecretEnv is the environment variable name for the JWT secret.
	SecretEnv string

	// HasClaims is true when auth.claims is declared in the manifest.
	HasClaims bool
}
