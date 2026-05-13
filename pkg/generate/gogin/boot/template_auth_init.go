package boot

// authInitHeaderLines — the static preamble rendered before dynamic values.
var authInitHeaderLines = []string{
	`// Phase003 — Configure ssac/pkg/auth. SecretEnv stores the env var NAME;`,
	`// IssueToken/RefreshToken/VerifyToken read os.Getenv(SecretEnv) on every`,
	`// call so secret rotation does not require re-Configure.`,
}

// authInitParseAccessTTL — time.ParseDuration for the access-token TTL.
// The second line carries the dynamic literal %q substitution.
var authInitParseAccessTTLLines = []string{
	`if err != nil {`,
	`	slog.Error("parse access_token_ttl", "err", err)`,
	`	os.Exit(1)`,
	`}`,
}

// authInitParseRefreshTTLLines — same shape as access-ttl for refresh TTL.
var authInitParseRefreshTTLLines = []string{
	`if err != nil {`,
	`	slog.Error("parse refresh_token_ttl", "err", err)`,
	`	os.Exit(1)`,
	`}`,
}

// authInitModeOverrideLines — BACKEND_AUTH_MODE env override comment + code.
var authInitModeOverrideLines = []string{
	`// Phase020 — BACKEND_AUTH_MODE env overrides the manifest default`,
	`// so the same binary can serve web (cookie) and mobile (bearer)`,
	`// deployments from a shared image.`,
}

// authInitModeSwitchLines — switch block selecting the final authMode value.
var authInitModeSwitchLines = []string{
	`if v := os.Getenv("BACKEND_AUTH_MODE"); v != "" {`,
	`	switch v {`,
	`	case "bearer", "cookie", "hybrid":`,
	`		authMode = v`,
	`	}`,
	`}`,
}

// authInitSameSiteLines — SameSite resolution via extracted helper.
var authInitSameSiteLines = []string{
	`// Phase020 — SameSite string → http.SameSite enum. Values outside`,
	`// {Lax, Strict, None} fall back to Lax which is the OWASP-recommended`,
	`// default for same-site SaaS.`,
}

// authInitSameSiteHelperFunc — parseSameSite top-level func emitted as a
// cmd/parse_same_site.go helper so main() body stays sequence-only.
var authInitSameSiteHelperFunc = `func parseSameSite(s string) http.SameSite {
	switch s {
	case "Strict":
		return http.SameSiteStrictMode
	case "None":
		return http.SameSiteNoneMode
	default:
		return http.SameSiteLaxMode
	}
}`

// authInitStoreInjectionLines — yongol-generated postgres RefreshStore 를
// package-level singleton 으로 설치 (ssac/purify Phase002).
//
// Before Phase002: main.go built a `&auth.RefreshStore{DB: conn}` struct
// literal and stored it on srv.RefreshStore; SSaC handlers referenced
// `server.RefreshStore` at every call site.
//
// After Phase002: ssac/pkg/auth exposes RefreshStore as an interface and
// auth.Init(store) sets the package-level default. Handlers now call
// auth.RefreshRotate(ctx, nil, token, ...) — ssac falls back to the
// singleton when store is nil, so no field threading is needed.
//
// The refresh_tokens DDL that used to live here (idempotent CREATE TABLE
// via auth.RefreshTokensDDL) is owned by the user's specs/db/ files now —
// Phase004 validate enforces that the refresh_tokens table is present
// when backend.auth is configured.
var authInitStoreInjectionLines = []string{
	`// Phase002 (ssac/purify) — install the yongol-generated postgres`,
	`// RefreshStore as the package-level auth singleton. Handlers that call`,
	`// auth.RefreshRotate / auth.Logout pass a nil store and let ssac fall`,
	`// back to this default.`,
	`auth.Init(infraauth.NewPostgres(queries))`,
}
