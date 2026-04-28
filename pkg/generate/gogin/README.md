# pkg/generate/gogin

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

Go + Gin 백엔드 코드 생성 orchestrator. 외부 도구(oapi-codegen, sqlc) 호출 + 하위 패키지 generator 조율 + 자체 glue (func copy, go.mod) 책임.

> 상위: [`pkg/generate/README.md`](../README.md). `pkg/generate/run_backend.go` 에서 `BackendType == GoGin` 일 때 호출.

## 공개 함수

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Generate` | `(fs *yongol.Fullstack, artifactsDir string, opts ...) error` | 파이프라인 진입점 ([1]~[11] 순차 실행) |
| `generateOpenAPIGoGin` | `(fs, artifactsDir) error` | [1] `oapi-codegen -generate types,strict-server,gin` |
| `generateSQLCGoGin` | `(fs, artifactsDir) error` | [3] `sqlc generate -f sqlc.yaml` (cmd.Dir = specs/db/) |
| `copyFuncSpecs` | `(fs, artifactsDir) error` | [8] `specs/func/<pkg>/*.go` → `arts/backend/internal/<pkg>/` 복사 |
| `generateGoMod` | `(fs, artifactsDir) error` | [11] go.mod 작성 + `go mod tidy` 실행 |
| `resolveGoModDeps` | `(fs) []dep` | go.mod 의 require 목록 결정 |
| `runGo` / `runGoModTidy` | `(...) error` | `go` / `go mod tidy` 실행 wrapper |
| `injectFFChecked` | `(file, hash)` | `//yg:checked` 해시 주입 |
| `cleanStaleFiles` | `(arts) error` | yongol 관리 파일 중 더이상 필요 없는 항목 제거 |

## 파이프라인 (실행 순서)

| # | 단계 | 활성 조건 | 위임 |
|---|---|---|---|
| [1] | oapi-codegen | 항상 | 자체 (`generate_openapi_gogin.go`) |
| [2] | postgresql 동적 SQL | 항상 | `pkg/generate/postgresql/` |
| [3] | sqlc generate | 항상 | 자체 (`generate_sqlc_gogin.go`) |
| [4] | auth (JWT+middleware+CurrentUser) | `manifest.backend.auth` 존재 | `gogin/auth/` |
| [5] | middleware (보강) | `bearerAuth` 등 | `gogin/middleware/` |
| [6] | model (CurrentUser 보강) | claims 존재 | `gogin/auth/` 와 통합 |
| [7] | statemachine | `len(fs.StateDiagrams) > 0` | `gogin/state/` |
| [8] | func copy | `specs/func/` 존재 | 자체 (`copy_func_specs.go`) |
| [9] | boot (cmd/main.go) | 항상 | `gogin/boot/` |
| [10] | handler (StrictServerInterface) | `len(fs.ServiceFuncs) > 0` | `gogin/ssac/` |
| [11] | go.mod | 항상 | 자체 (`generate_go_mod.go`) |

## 산출물 맵

```
arts/backend/
├── cmd/main.go                          ← [9]
├── internal/
│   ├── api/server.gen.go                ← [1] oapi-codegen
│   ├── db/{db,models,*.sql}.go          ← [3] sqlc (정적 + [2] 동적)
│   ├── auth/{issue,verify,refresh,reexport}.go  ← [4]
│   ├── middleware/bearerauth.go         ← [4][5]
│   ├── model/current_user.go            ← [4]
│   ├── statemachine/<id>.go             ← [7]
│   ├── <custom-func-pkg>/*.go           ← [8]
│   └── service/{server.go, <feature>/<func>.go}  ← [10]
└── go.mod                                ← [11]
```

## 외부 도구

| 도구 | 버전 | 단계 |
|---|---|---|
| `oapi-codegen` | v2.6+ | [1] |
| `sqlc` | v1.30+ | [3] |
| `gofmt` | Go 기본 | [10] |
| `go mod tidy` | Go 기본 | [11] |

## 하위 패키지 README

`gogin/auth/`, `gogin/boot/`, `gogin/ssac/`, `gogin/state/` 각각의 README 참조. 산출물 위치·생성 코드 상세는 거기에 있음.
