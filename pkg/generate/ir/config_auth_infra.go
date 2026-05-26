//ff:type feature=gen-ir type=model
//ff:what AuthInfraConfig -- auth 인프라 어댑터 설정 (refresh_tokens DDL + RefreshStore)

package ir

// AuthInfraConfig holds the resolved auth infrastructure configuration.
// Active when manifest.backend.auth is declared.
type AuthInfraConfig struct {
	// Mode is one of "bearer", "cookie", or "hybrid".
	Mode string

	// SecretEnv is the environment variable name for the JWT secret.
	SecretEnv string

	// AccessTokenTTL is the access token lifetime (e.g. "15m").
	AccessTokenTTL string

	// RefreshTokenTTL is the refresh token lifetime (e.g. "168h").
	RefreshTokenTTL string
}
