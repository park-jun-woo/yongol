//ff:func feature=main type=command control=selection
//ff:what main — 애플리케이션 엔트리포인트 (DB/JWT/authz/queue/cache/session/file/router/gin 초기화)
//ff:checked llm=yongol-gen hash=fa19c806
package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/park-jun-woo/ssac/pkg/auth"
	"github.com/park-jun-woo/zenflow-try01/internal/api"
	"github.com/park-jun-woo/zenflow-try01/internal/db"
	"github.com/park-jun-woo/zenflow-try01/internal/middleware"
	"github.com/park-jun-woo/zenflow-try01/internal/service"
	"log/slog"
	"net/http"
	"os"
	"time"
	infraauth "github.com/park-jun-woo/zenflow-try01/internal/infra/auth"
)

func main() {
	logLevel := parseLogLevel(os.Getenv("LOG_LEVEL"))
	sensitiveKeys := buildSensitiveKeys([]string{"password_hash"})
	handler := newSlogHandler(logLevel, sensitiveKeys)
	slog.SetDefault(slog.New(handler))

	ctx, cancelBootstrap := context.WithCancel(context.Background())
	defer cancelBootstrap()
	slog.Info("connecting to database")
	poolCfg, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("db init: parse DATABASE_URL", "err", err)
		os.Exit(1)
	}
	poolCfg.MaxConns = int32(envInt("DB_MAX_OPEN_CONNS", 25))
	poolCfg.MinConns = int32(envInt("DB_MAX_IDLE_CONNS", 5))
	poolCfg.MaxConnLifetime = envDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute)
	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		slog.Error("db init", "err", err)
		os.Exit(1)
	}
	defer pool.Close()
	queries := db.New(pool)
	slog.Info("database connected", "max_conns", poolCfg.MaxConns)

	if v := os.Getenv("JWT_SECRET"); v == "" {
		slog.Error("JWT_SECRET is required")
		os.Exit(1)
	} else if len(v) < 32 {
		slog.Error("JWT_SECRET must be at least 32 characters", "length", len(v))
		os.Exit(1)
	}

	initAuthz(os.Getenv("OPA_POLICY_PATH"))

	srv := &service.Server{
		DB: pool,
		Queries: queries,
	}

	r := gin.Default()

	ridTrustUpstream := envBool("BACKEND_ERROR_REQUEST_ID_TRUST_UPSTREAM", true)
	ridHeader := envString("BACKEND_ERROR_REQUEST_ID_HEADER", "X-Request-Id")
	r.Use(middleware.RequestID(ridTrustUpstream, ridHeader))

	middleware.ExposeInternalError = envBool("BACKEND_ERROR_EXPOSE_INTERNAL_ERROR", false)
	r.Use(middleware.ErrorEnvelopeMiddleware())

	promEnabled := envBool("BACKEND_OBSERVABILITY_METRICS_ENABLED", true)
	promPath := envString("BACKEND_OBSERVABILITY_METRICS_PATH", "/metrics")
	if promEnabled {
		r.Use(middleware.PrometheusMiddleware())
		r.GET(promPath, middleware.PrometheusHandler())
	}

	shEnabled := envBool("BACKEND_SECURITY_HEADERS_ENABLED", true)
	shProfile := envString("BACKEND_SECURITY_HEADERS_PROFILE", "production")
	shHSTSMaxAge := envInt("BACKEND_SECURITY_HEADERS_HSTS_MAX_AGE", 31536000)
	shCSPReportOnly := envBool("BACKEND_SECURITY_HEADERS_CSP_REPORT_ONLY", false)
	if shEnabled {
		secHeadersCfg := middleware.SecurityHeadersConfig{
			Enabled:           true,
			Profile:           shProfile,
			HSTSMaxAge:        shHSTSMaxAge,
			HSTSIncludeSubs:   true,
			HSTSPreload:       false,
			CSPEnabled:        true,
			CSPReportOnly:     shCSPReportOnly,
			CSPDirectives:     map[string][]string{
			"base-uri": []string{"'self'"},
			"default-src": []string{"'self'"},
			"frame-ancestors": []string{"'none'"},
		},
			XFrameOptions:     "DENY",
			ReferrerPolicy:    "strict-origin-when-cross-origin",
			PermissionsPolicy: map[string][]string{
			"camera": []string{},
			"geolocation": []string{},
			"microphone": []string{},
		},
		}
		r.Use(middleware.SecurityHeadersMiddleware(secHeadersCfg))
	}

	bodyLimit := envInt64("BACKEND_HTTP_BODY_LIMIT", 1048576)
	multipartLimit := envInt64("BACKEND_HTTP_MULTIPART_LIMIT", 33554432)
	r.Use(middleware.BodyLimit(bodyLimit))
	r.Use(middleware.MultipartLimit(multipartLimit))

	validator, err := middleware.RequestValidator()
	if err != nil {
		slog.Error("bootstrap failed", "stage", "request-validator", "err", err)
		fmt.Fprintf(os.Stderr, "bootstrap failed: %v\n", err)
		os.Exit(1)
	}
	r.Use(validator)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.GET("/ready", readyHandlerWithDB(pool))

	publicOps := map[string]bool{
		"Login": true,
		"Register": true,
	}
	strictHandler := api.NewStrictHandler(srv, []api.StrictMiddlewareFunc{
		middleware.BearerAuthStrict(publicOps),
	})
	api.RegisterHandlers(r, strictHandler)

	// Phase003 — Configure ssac/pkg/auth. SecretEnv stores the env var NAME;
	// IssueToken/RefreshToken/VerifyToken read os.Getenv(SecretEnv) on every
	// call so secret rotation does not require re-Configure.
	accessTTL, err := time.ParseDuration("15m")
	if err != nil {
		slog.Error("parse access_token_ttl", "err", err)
		os.Exit(1)
	}
	refreshTTL, err := time.ParseDuration("168h")
	if err != nil {
		slog.Error("parse refresh_token_ttl", "err", err)
		os.Exit(1)
	}
	// Phase020 — BACKEND_AUTH_MODE env overrides the manifest default
	// so the same binary can serve web (cookie) and mobile (bearer)
	// deployments from a shared image.
	authMode := "bearer"
	if v := os.Getenv("BACKEND_AUTH_MODE"); v != "" {
		switch v {
		case "bearer", "cookie", "hybrid":
			authMode = v
		}
	}
	// Phase020 — SameSite string → http.SameSite enum. Values outside
	// {Lax, Strict, None} fall back to Lax which is the OWASP-recommended
	// default for same-site SaaS.
	var sameSite http.SameSite
	switch "Lax" {
	case "Strict":
		sameSite = http.SameSiteStrictMode
	case "None":
		sameSite = http.SameSiteNoneMode
	default:
		sameSite = http.SameSiteLaxMode
	}
	auth.Configure(auth.Config{
		SecretEnv:  "JWT_SECRET",
		AccessTTL:  accessTTL,
		RefreshTTL: refreshTTL,
		Mode:       authMode,
		CookieAttrs: auth.CookieAttrs{
			AccessName:  "__Host-access_token",
			RefreshName: "__Host-refresh_token",
			SameSite:    sameSite,
			AccessTTL:   accessTTL,
			RefreshTTL:  refreshTTL,
		},
	})
	// Phase002 (ssac/purify) — install the yongol-generated postgres
	// RefreshStore as the package-level auth singleton. Handlers that call
	// auth.RefreshRotate / auth.Logout pass a nil store and let ssac fall
	// back to this default.
	auth.Init(infraauth.NewPostgres(queries))

	headerLimit := envInt64("BACKEND_HTTP_HEADER_LIMIT", 1048576)
	runServerWithGracefulShutdown(r, cancelBootstrap, int(headerLimit))
}
