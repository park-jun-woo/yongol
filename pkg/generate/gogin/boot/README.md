# pkg/generate/gogin/boot

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

`arts/<project>/backend/cmd/main.go` 를 생성하는 Block Builder. 1 블록 = 1 파일 (`block_*.go`). orchestrator 가 활성 블록만 수집해 import dedup + 본문 조립.

> 상위: [`pkg/generate/gogin/README.md`](../README.md) [9].

## 공개 함수

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Generate` | `(fs *yongol.Fullstack, artifactsDir, modulePath string) error` | 진입점. `collectActiveBlocks` → `deduplicateImports` → `assembleBody` → `writeMainGo` |
| `collectActiveBlocks` | `(fs) []MainBlock` | `MainBlock.Active` 로 필터 |
| `deduplicateImports` | `(blocks, modulePath) []string` | 블록별 import 합산 후 중복 제거 |
| `assembleBody` | `(blocks) string` | main 본문 조립 (블록 순서 고정) |
| `writeMainGo` | `(artifactsDir, imports, body) error` | `cmd/main.go` 작성 + gofmt |
| `writeEnvHelperFiles` | `(artifactsDir) error` | top-level helper (envInt/envDuration/...) 별도 파일 emit |
| `baseCandidateBlocks` | `(fs) []MainBlock` | 후보 블록 정렬 목록 |
| `block_*` (블록 빌더 다수) | `(fs) MainBlock` | 아래 표 — 각 블록 1 파일 |

## 공개 타입

| 타입 | 설명 |
|---|---|
| `MainBlock` | `{Name string; Active func(*Fullstack) bool; Imports []string; Lines []string}` |

## 블록 목록 (조립 순서)

| # | 블록 | 활성 조건 | ssac/pkg import |
|---|---|---|---|
| 0 | `logger-init` | 항상 | — (slog) |
| 0.5 | `env-helpers` | 항상 | — |
| 1 | `db-init` | 항상 | — (pgxpool + stdlib bridge) |
| 2 | `jwt-secret` | `manifest.backend.auth` 존재 | — |
| 2.5 | `auth-init` | auth 통합 | — |
| 3 | `authz-init` | SSaC `@auth` 사용 | `ssac/pkg/authz` |
| 4 | `queue-init` | `manifest.queue.backend` | `ssac/pkg/queue` |
| 5 | `session-init` | `manifest.session.backend` | `ssac/pkg/session` |
| 6 | `cache-init` | `manifest.cache.backend` | `ssac/pkg/cache` |
| 7 | `file-init` | `manifest.file.backend` | `ssac/pkg/file` |
| 7.3 | `csrf` | csrf 활성 | — |
| 7.5 | `request-id`, `request-validator`, `error-envelope`, `body-limit` | manifest 옵션 | — |
| 7.7 | `otel-init`, `prometheus`, `security-headers` | manifest 활성 | otel/prometheus 패키지 |
| 8 | `server` | 항상 | — (`&Server{Queries: ...}`) |
| 9 | `middleware` | `bearerAuth` 포함 | (자체 middleware 패키지) |
| 10 | `router` | 항상 | gin |
| 10.3 | `cors` | `manifest.backend.cors.enabled` | gin-contrib/cors |
| 10.5 | `health` | 항상 (DDL 있을 때 `pool.Ping`) | — |
| 10.7 | `register-handlers` | 항상 | api.RegisterHandlers |
| 11 | `gin-run` | 항상 | http.Server + SIGINT/SIGTERM graceful shutdown |

## 산출물 환경변수 (생성된 main.go 가 읽음)

| 변수 | 기본 | 효과 |
|---|---|---|
| `LOG_LEVEL` | `INFO` | DEBUG/INFO/WARN/ERROR |
| `LOG_FORMAT` | `JSON` | TEXT 가독형 |
| `DB_MAX_OPEN_CONNS` / `DB_MAX_IDLE_CONNS` / `DB_CONN_MAX_LIFETIME` | 25 / 5 / 5m | pgxpool 설정 |
| `CORS_ALLOW_ORIGINS` / `_METHODS` / `_CREDENTIALS` | manifest 값 | manifest 덮어쓰기 |
| `DATABASE_URL`, `JWT_SECRET`, `FILE_ROOT`, `S3_BUCKET`, OTEL/Prometheus 변수 | — | 블록별 |

## 입력

| 입력 | 소스 | 사용 |
|---|---|---|
| `fs.Manifest` | manifest.yaml | auth/session/cache/queue/file/middleware 판단 |
| `fs.ServiceFuncs` | SSaC | @auth/@call/@publish/@subscribe 사용 여부 |
| `fs.ParsedPolicies` | Rego | @ownership mapping 리터럴 (authz-init) |
| `fs.StateDiagrams` | states/ | state machine init |
| `modulePath` | `manifest.backend.module` | import 경로 |
| `artifactsDir` | CLI 인자 | 출력 경로 |

## 런타임 라이브러리

생성된 main.go 는 `github.com/park-jun-woo/ssac/pkg/<auth|authz|cache|file|mail|queue|session|crypto|storage|text|image>` 를 import. yongol 자체는 빌드 도구이므로 런타임에 import 되지 않음.

## import 경로 규칙

| 카테고리 | 경로 |
|---|---|
| ssac 런타임 | `github.com/park-jun-woo/ssac/pkg/<pkg>` |
| 프로젝트 내부 | `<manifest.backend.module>/internal/<pkg>` |
| 표준 / 외부 | 그대로 (`gin`, `pgx/v5/pgxpool`, `stdlib`, …) |
