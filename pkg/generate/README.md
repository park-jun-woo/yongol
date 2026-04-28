# pkg/generate

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

`yongol generate <specs> <arts>` 의 코드 생성 최상위 orchestrator. backend / frontend / hurl / OPA Rego 4 가지 산출 경로를 분기 호출. ERROR 또는 WARNING 이 하나라도 있으면 generate 호출 자체가 거절됨 (CLI 게이트).

## 공개 함수

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Generate` | `(fs *yongol.Fullstack, artifactsDir string, backend BackendType, frontend FrontendType, opts ...Option) error` | 진입점. runBackend → runFrontend → hurl_mirror → copy OPA Rego 순. |
| `runBackend` | `(fs, artifactsDir, backend) error` | `BackendType == GoGin` → `gogin.Generate` |
| `runFrontend` | `(fs, frontend) error` | `FrontendType == React` → `react.Generate` |
| `copyOPARego` | `(fs, artifactsDir) error` | `specs/policy/*.rego` → `arts/backend/policy/*.rego` |
| `runMigration` | `(fs, artifactsDir) error` | DDL diff → `arts/db/migrations/NNNN_*.up.sql` (+ down stub) |

## 공개 타입

| 타입 | 설명 |
|---|---|
| `BackendType` | `string` — `GoGin = "go-gin"` (현재 1종) |
| `FrontendType` | `string` — `React = "react"` (현재 1종) |
| `Option` / `Config` | `apply_generate_options.go` — generate 옵션 (`--clean`, dry-run 등) |

## 산출물 맵

```
arts/
├── backend/   ← gogin.Generate (cmd/main.go, internal/{api,db,auth,middleware,model,statemachine,service,...}, policy/*.rego, go.mod)
├── frontend/  ← react.Generate (향후)
├── db/migrations/  ← migration.Generate (NNNN_<desc>.up.sql + down stub)
└── tests/     ← hurl_mirror.Generate (smoke.hurl)
```

## 하위 패키지

| 패키지 | 역할 |
|---|---|
| `gogin/` | Go+Gin 백엔드 orchestrator (주력) |
| `react/` | React 프론트엔드 (향후) |
| `nestjs/` | NestJS (계획 단계) |
| `migration/` | DDL diff 마이그레이션 emit |
| `hurl_mirror/` | specs/tests → arts/tests 복사 |
| `prepared/` | 사전 준비 산출물 |
| `splitter/` | 큰 산출물 분할 |
| `ffhash/`, `filefunc/` | filefunc 어노테이션 / 해시 메타 주입 |

## 외부 도구 의존

| 도구 | 호출 위치 | 용도 |
|---|---|---|
| `oapi-codegen` v2.6+ | gogin [1] | OpenAPI → types + StrictServerInterface |
| `sqlc` v1.30+ | gogin [3] | DDL + queries → Go model + query methods |
| `gofmt`, `go mod tidy` | gogin [10][11] | 포매팅, 의존성 정리 |

## generate 게이트 (`cmd/yongol/generate_cmd.go`)

1. `ParseAll` → diagnostic 0 확인. 2. `validate.Validate` → ERROR 0 확인. 3. WARNING 0 확인. 셋 다 통과해야 `Generate` 호출. 부분 산출물은 남을 수 있음 (각 단계 error 시 즉시 중단 + `fmt.Errorf("<step>: %w", err)`).
