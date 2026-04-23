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

// authInitSameSiteCommentLines — SameSite comment preamble.
var authInitSameSiteCommentLines = []string{
	`// Phase020 — SameSite string → http.SameSite enum. Values outside`,
	`// {Lax, Strict, None} fall back to Lax which is the OWASP-recommended`,
	`// default for same-site SaaS.`,
	`var sameSite http.SameSite`,
}

// authInitSameSiteSwitchLines — http.SameSite enum resolution.
var authInitSameSiteSwitchLines = []string{
	`case "Strict":`,
	`	sameSite = http.SameSiteStrictMode`,
	`case "None":`,
	`	sameSite = http.SameSiteNoneMode`,
	`default:`,
	`	sameSite = http.SameSiteLaxMode`,
	`}`,
}

// authInitDDLBootstrapLines — idempotent refresh_tokens DDL bootstrap.
var authInitDDLBootstrapLines = []string{
	`// Phase002 — bootstrap refresh_tokens schema (idempotent). Kept in`,
	`// main.go so a fresh DB is usable without running a separate`,
	`// migration tool; real deployments should instead run the DDL via`,
	`// their migration pipeline and drop this block.`,
	`if _, err := conn.ExecContext(ctx, auth.RefreshTokensDDL); err != nil {`,
	`	slog.Error("refresh_tokens DDL", "err", err)`,
	`	os.Exit(1)`,
	`}`,
}

// authInitStoreInjectionLines — RefreshStore 주입 + srv 필드 대입.
var authInitStoreInjectionLines = []string{
	`// Phase004/Phase009 — inject the RefreshStore into the Server so SSaC`,
	`// handlers that call auth.RefreshToken / auth.RefreshRotate / auth.Logout`,
	`// can reach it via server.RefreshStore without threading the DB handle`,
	`// through every handler signature. Phase009 moved the auth-refresh`,
	`// route onto the canonical openapi + SSaC path, so this block does`,
	`// not mount any gin route — it only wires store + config.`,
	`srv.RefreshStore = refreshStore`,
}
