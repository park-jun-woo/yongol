# pkg/validate/manifest

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

`manifest.yaml` 자체 정합성 검증 — 메타 (`C-*`), CORS, observability, security headers, auth 모드/TTL/secret, builtin backend (cache/session/queue/auth refresh) 의 canonical DDL+sqlc 강제.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙 — 메타 (C-*)

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| C-2 | `c02ApiVersion` | `apiVersion` 이 `yongol/v1` 이어야 함 (ERROR) | IF-ELSE | ✓ |
| C-3 | `c03Kind` | `kind` 가 `Project` 이어야 함 (ERROR) | IF-ELSE | ✓ |
| C-4 | `c04MetadataName` | `metadata.name` 비어있지 않아야 함 (ERROR) | IF-ELSE | ✓ |
| C-5 | `c05BackendModule` | `backend.module` 비어있지 않아야 함 (ERROR) | IF-ELSE | ✓ |

## 검증 규칙 — CORS / Observability / Security headers

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| CORS-01 | `cors01WildcardCredentials` | `allow_origins=["*"]` + `allow_credentials=true` 동시 금지 (ERROR) | IF-ELSE | ✓ |
| OBS-001 | `obs01MetricsPath` | `metrics.path` 가 `/` 로 시작 (ERROR) | IF-ELSE | ✓ |
| OBS-002 | `obs02MetricsPathNotOpenAPI` | `metrics.path` 와 OpenAPI path 충돌 금지 (ERROR) | IF-ELSE | ✓ |
| OBS-003 | `obs03TracingExporter` | `tracing.exporter` 는 `otlp`/`stdout`/`noop` 중 하나 (ERROR) | IF-ELSE | ✓ |
| OBS-004 | `obs04TracingSampleRate` | `tracing.sample_rate` 가 `[0.0, 1.0]` 범위 (ERROR) | IF-ELSE | ✓ |
| SEC-301 | `sec301CspPermissive` | CSP `default-src` 가 `*`/`'unsafe-eval'` 포함 (WARNING) | IF-ELSE | ✓ |
| SEC-302 | `sec302HstsShort` | HSTS `max_age < 15552000` (180일) (WARNING, `0` 은 명시적 비활성화로 면제) | IF-ELSE | ✓ |

## 검증 규칙 — Auth

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| SEC-201 | `sec201CookieWithoutCsrf` | `auth.mode=cookie\|hybrid` + `csrf.enabled=false` 금지 (ERROR) | IF-ELSE | ✓ |
| SEC-401 | `sec401JwtSecretEnvRequired` | `backend.auth.secret` 리터럴 금지, `secret_env` 만 허용 (ERROR) | IF-ELSE | ✓ |
| SEC-402 | `sec402AccessTtlUpperBound` | `backend.auth.access_token_ttl` 30분 이하 (WARNING, OWASP 권장) | IF-ELSE | ✓ |
| SEC-403 | `sec403AuthModeEnum` | `backend.auth.mode` 는 `cookie`/`bearer`/`hybrid` 중 하나 (ERROR) | IF-ELSE | ✓ |

## 검증 규칙 — Builtin backend canonical DDL+sqlc 강제 (XN*-90)

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XNA-90 | `xna90RefreshRequiresSqlc` | `backend.auth` 구성 시 `refresh_tokens` DDL + `RefreshTokenInsert/FindByHash/Revoke/RevokeAll` sqlc 쿼리 강제 (ERROR) | IF-ELSE | ✓ |
| XNC-90 | `xnc90CacheBackendRequiresSqlc` | `cache.backend=postgres` 시 `fullend_cache` DDL + `CacheSet/CacheGet/CacheDelete` sqlc 쿼리 강제 (ERROR) | IF-ELSE | ✓ |
| XNQ-90 | `xnq90QueueBackendRequiresSqlc` | `queue.backend=postgres` 시 `fullend_queue` DDL + `QueuePublish/QueuePoll/QueueAck` sqlc 쿼리 강제 (ERROR) | IF-ELSE | ✓ |
| XNS-90 | `xns90SessionBackendRequiresSqlc` | `session.backend=postgres` 시 `fullend_sessions` DDL + `SessionSet/SessionGet/SessionDelete` sqlc 쿼리 강제 (ERROR) | IF-ELSE | ✓ |

`XN*-90` advice 는 `ssac/pkg/<infra>/interface.yaml` 의 `canonical_ddl` + `canonical_queries` 를 인라인으로 렌더한다 (`canonicalAdvice`).

## Defeater

없음.
