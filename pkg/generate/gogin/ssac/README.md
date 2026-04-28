# pkg/generate/gogin/ssac

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

SSaC 함수 선언을 oapi-codegen `StrictServerInterface` 메서드(Service 계층) 로 변환. Controller 는 oapi-codegen strict-server 가 자동 생성하므로 본 패키지는 비즈니스 로직만 emit. 활성 조건: `len(fs.ServiceFuncs) > 0`.

> 상위: [`pkg/generate/gogin/README.md`](../README.md) [10].

## 공개 함수

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Generate` | `(fs *yongol.Fullstack, artifactsDir string) error` | 진입점. server.go + 함수별 method 파일 emit |
| `generateServerGo` | `(fs, artifactsDir) error` | `service/server.go` (Server struct + helper) |
| `generateServerHelpers` | `(fs, artifactsDir) error` | strPtr, neutralCode, log helper 등 |
| `generateHTTPMethod` | `(fn, fs, artifactsDir) error` | StrictServerInterface 메서드 1 개 |
| `generateSubscribeMethod` | `(fn, fs, artifactsDir) error` | `@subscribe` handler 1 개 (HTTP 외) |
| `needsTransaction` | `(fn) bool` | mutating seq (@post/@put/@delete) 1 개 이상 → true |
| `buildSequence` / `buildGet/Post/Put/Delete/Empty/Exists/State/Auth/Call/Publish/Eval/Response/Field` | `(...)` | 시퀀스별 코드 빌더 |
| `resolveSqlcMethod` | `(model, method) string` | `Model.Method` → `Queries.<Prefix><Method>` |
| `resolveErrStatus` / `resolveCallErrStatus` | `(...)` | @error annotation → response status |
| `pgtypeRowUnwrap` / `pgtypeHelpersEmit` / `emitConvertFuncFile` / `emitConvertListFile` | `(...)` | pgtype → API 타입 변환 helper emit |
| `collectOwnerships` | `(fs) []OwnershipMapping` | Rego @ownership 추출 (XAS-60) |

## 산출물

```
arts/backend/internal/service/
├── server.go                  Server struct (StrictServerInterface 구현)
├── helpers.go (외 helper)     strPtr / neutral code / log tag
├── <func_snake>.go            함수별 1 파일 (전부 package service)
└── pgtype/*.go                pg_numeric/pg_uuid → string 변환 helper
```

## 트랜잭션 정책

| 조건 | 처리 |
|---|---|
| mutating seq 0 | tx 없음 — `server.Queries` 직접 사용 |
| mutating seq ≥ 1 | `pool.Begin(ctx)` + `defer Rollback` + `Commit` (pgx/v5), `qtx := server.Queries.WithTx(tx)` |

`@publish` 는 commit 후 best-effort, return 은 typed `ResponseObject`.

## SSaC guard → typed response

| guard | HTTP | 반환 |
|---|---|---|
| `@empty` | 404 | `api.Xxx404JSONResponse{Error: ...}` |
| `@exists` / `@state` | 409 | `api.Xxx409JSONResponse{...}` |
| `@auth` | 403 | `api.Xxx403JSONResponse{...}` |
| `@call` `@error NNN` | NNN | `api.XxxNNNJSONResponse{...}` |
| 내부 에러 | 500 | `return nil, err` (adapter 처리) |

`gin.H` 미사용. ServiceError 타입 미사용.

## Server struct (Phase005 pgx/v5)

```go
type Server struct {
    DB      *pgxpool.Pool
    Queries *db.Queries
}
```

Server 는 `api.StrictServerInterface` 구현. JWT 처리는 middleware 책임. main.go 배선은 `boot/` 가 담당.

## currentUser 접근

strict-server 는 gin.Context 미사용. `StrictMiddlewareFunc` 가 `context.WithValue(ctx, "currentUser", ...)` 주입 → service method 에서 `ctx.Value("currentUser").(*model.CurrentUser)`.

## 외부 의존 / import

| 종류 | 경로 |
|---|---|
| oapi-codegen types + response | `<module>/internal/api` |
| sqlc queries | `<module>/internal/db` |
| ssac runtime | `github.com/park-jun-woo/ssac/pkg/<authz|queue|session|cache|file|auth>` |
| project custom func | `<module>/internal/<pkg>` |
| statemachine | `<module>/internal/statemachine` |
| model | `<module>/internal/model` |
| pgx 가드 | `github.com/jackc/pgx/v5` |

## 비고

handler/ 와의 관계: 동일 산출물 (`internal/service/**/*.go`). 본 ssac/ 가 SSaC 시퀀스 → Go body 변환의 최종 패키지 (handler/ 는 설계 배경 문서로만 유지).
