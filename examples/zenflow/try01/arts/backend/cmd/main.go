//ff:func feature=main type=command control=sequence
//ff:what main — 애플리케이션 엔트리포인트 (DB/JWT/authz/queue/cache/session/file/router/gin 초기화)
//ff:checked llm=yongol-gen hash=299155b2
package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/park-jun-woo/ssac/pkg/auth"
	"github.com/park-jun-woo/zenflow/internal/api"
	"github.com/park-jun-woo/zenflow/internal/db"
	"github.com/park-jun-woo/zenflow/internal/middleware"
	"github.com/park-jun-woo/zenflow/internal/service"
	"log/slog"
	"net/http"
	"os"
	"time"
	_ "github.com/lib/pq"
)

func main() {
	logLevel := parseLogLevel(os.Getenv("LOG_LEVEL"))
	sensitiveKeys := buildSensitiveKeys([]string{"password_hash"})
	handler := newSlogHandler(logLevel, sensitiveKeys)
	slog.SetDefault(slog.New(handler))

	ctx, cancelBootstrap := context.WithCancel(context.Background())
	defer cancelBootstrap()
	slog.Info("connecting to database")
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("db init", "err", err)
		os.Exit(1)
	}
	defer conn.Close()
	conn.SetMaxOpenConns(envInt("DB_MAX_OPEN_CONNS", 25))
	conn.SetMaxIdleConns(envInt("DB_MAX_IDLE_CONNS", 5))
	conn.SetConnMaxLifetime(envDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute))
	queries := db.New(conn)
	slog.Info("database connected", "max_open", conn.Stats().MaxOpenConnections)

	if v := os.Getenv("JWT_SECRET"); v == "" {
		slog.Error("JWT_SECRET is required")
		os.Exit(1)
	} else if len(v) < 32 {
		slog.Error("JWT_SECRET must be at least 32 characters", "length", len(v))
		os.Exit(1)
	}

	initAuthz(conn)

	srv := &service.Server{
		DB: conn,
		Queries: queries,
	}

	r := gin.Default()

	ridTrustUpstream := envBool("BACKEND_ERROR_REQUEST_ID_TRUST_UPSTREAM", true)
	ridHeader := envString("BACKEND_ERROR_REQUEST_ID_HEADER", "X-Request-Id")
	r.Use(middleware.RequestID(ridTrustUpstream, ridHeader))

	middleware.ExposeInternalError = envBool("BACKEND_ERROR_EXPOSE_INTERNAL_ERROR", false)
	r.Use(middleware.ErrorEnvelopeMiddleware())

	corsCfg := cors.Config{
		AllowMethods:     envStringList("CORS_ALLOW_METHODS", []string{"GET", "POST", "PUT", "PATCH", "DELETE"}),
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: envBool("CORS_ALLOW_CREDENTIALS", true),
		MaxAge:           time.Duration(3600000000000),
	}
	corsCfg.AllowOrigins = envStringList("CORS_ALLOW_ORIGINS", []string{"http://localhost:3000"})
	r.Use(cors.New(corsCfg))

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
	r.GET("/ready", readyHandlerWithDB(conn))

	publicOps := map[string]bool{
		"Login": true,
		"Signup": true,
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
	// Phase002 — bootstrap refresh_tokens schema (idempotent). Kept in
	// main.go so a fresh DB is usable without running a separate
	// migration tool; real deployments should instead run the DDL via
	// their migration pipeline and drop this block.
	if _, err := conn.ExecContext(ctx, auth.RefreshTokensDDL); err != nil {
		slog.Error("refresh_tokens DDL", "err", err)
		os.Exit(1)
	}
	refreshStore := &auth.RefreshStore{DB: conn, DetectReuseLogoutAll: false}
	// Phase004/Phase009 — inject the RefreshStore into the Server so SSaC
	// handlers that call auth.RefreshToken / auth.RefreshRotate / auth.Logout
	// can reach it via server.RefreshStore without threading the DB handle
	// through every handler signature. Phase009 moved the auth-refresh
	// route onto the canonical openapi + SSaC path, so this block does
	// not mount any gin route — it only wires store + config.
	srv.RefreshStore = refreshStore

	headerLimit := envInt64("BACKEND_HTTP_HEADER_LIMIT", 1048576)
	runServerWithGracefulShutdown(r, cancelBootstrap, int(headerLimit))
}
