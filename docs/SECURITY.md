# yongol — Security Defaults for Generated Handlers

Summary of the security defaults in the Go+Gin backend produced by `yongol generate`. Review before any production deployment.

## Safe Defaults Today (OK)

| Area | Default |
|------|--------|
| Authorization | OPA Rego `default allow := false` (deny-by-default) |
| Authentication | JWT HS256; bootstrap fails if `JWT_SECRET` length < 32 |
| Algorithm verification | `verify_token.go` checks `*jwt.SigningMethodHMAC` (partial algorithm-confusion defense) |
| Input validation | kin-openapi runtime validator enforces body/query/path schemas -> 400 on failure |
| SQL injection | All queries use sqlc prepared statements |
| Passwords | bcrypt (`golang.org/x/crypto/bcrypt`); on login failure, compare against a dummy hash to mitigate timing oracles |
| Sensitive-data logging | `redact.DefaultKeys` auto-masks slog attributes (`password / token / api_key / authorization / ssn / credit_card / cvv`). DDL `-- @sensitive` columns added automatically |
| Error responses | Internal errors return a `500` with empty body — no stack traces / SQL errors leaked |
| 401 vs 404 | Login returns `401 Invalid credentials` for both user-not-found and wrong-password (blocks existence oracle) |
| CORS | `allow_origins=["*"]` + `allow_credentials=true` is an ERROR at validate (CORS-01) |
| Panic | `gin.Default()` Recovery returns 500; stack traces written only to server logs |

## Gaps (to be improved)

| Area | Current | Risk |
|------|------|------|
| JWT refresh rotation | None — 7-day refresh issuance only | On theft, full 7-day window is valid |
| Request body size limit | None | Memory DoS from oversized requests |
| Rate limiting | None | Login brute force unguarded |
| HTTP security headers | HSTS / nosniff / X-Frame-Options / CSP unset | MIME sniffing, clickjacking |
| Error request_id | None | Harder error tracing |
| Secret fallback | `issue_token.go` has `secret == "" { secret = "secret" }` | Weak secret if main.go bypassed |

## User Checklist (Before Production)

1. **`JWT_SECRET`** — at least 32 random characters (`openssl rand -base64 48`).
2. **`OPA_POLICY_PATH`** — verify at deploy time (bootstrap also checks it).
3. **`CORS_ALLOW_ORIGINS`** — production URLs explicitly; never leave `http://localhost:3000`.
4. **`DATABASE_URL`** — require SSL (`?sslmode=require`).
5. **Reverse proxy (nginx/ALB)** — TLS termination and `Strict-Transport-Security`. yongol does not attach HSTS.
6. **Rate limiting** — proxy (`nginx limit_req`) or WAF; cap `Login` / `PasswordReset` at ~5 req/min/IP.
7. **Request body limit** — configure at the proxy (`client_max_body_size 1m`).
8. **Log sinks** — ingest slog JSON as-is. redact replaces sensitive fields with `***`. Add fields via `buildSensitiveKeys` or `-- @sensitive` in DDL.
9. **OPA policy** — deny-by-default is safe; when adding endpoints, verify allow rules are not missing.
10. **DB user** — dedicated least-privilege account; no `CREATE TABLE`.

## Related Commands

- `yongol validate <specs>` — validates CORS wildcard/credentials combo, JWT secret length, etc.
- `yongol generate <specs> <arts>` — emits code with the defaults above.

## Reporting Security Issues

Issues: `github.com/park-jun-woo/yongol`. CVE-level vulnerabilities: private report to `madosaja@gmail.com`.
