# pkg/generate/gogin/ssac

SSaC 함수 선언을 oapi-codegen **`StrictServerInterface` 메서드** 로 변환하는 코드젠.
MVC+Service 에서 **Service 계층** — Controller (HTTP 파싱/응답) 는 oapi-codegen 의
strict-server 가 자동 처리하므로 이 패키지는 **비즈니스 로직만** 생성한다.

## 아키텍처 (MVC+Service)

```
oapi-codegen strict-server (Controller — 자동)
  ├─ request body 파싱 → typed RequestObject
  ├─ path param 추출 → RequestObject.Id
  ├─ response 포매팅 → c.JSON(status, body)
  └─ StrictServerInterface 메서드 호출 ──→ ssac/ (Service — 이 패키지)
                                              ├─ BeginTx + defer Rollback
                                              ├─ @get  (qtx)
                                              ├─ @empty / @exists / @auth / @state
                                              ├─ @post / @put / @delete (qtx)
                                              ├─ tx.Commit()
                                              ├─ @publish (commit 후)
                                              └─ return ResponseObject, nil

sqlc Queries (Model — 자동)
  └─ DB 접근 전부
```

**handler/ 불필요** — oapi-codegen strict-server 가 Controller 전체를 자동 생성.

## oapi-codegen 모드

```bash
# 이전 (gin-server) — ServerInterface 에 gin.Context 포함
oapi-codegen -generate types,gin-server

# 현재 (strict-server) — StrictServerInterface 에 gin.Context 없음
oapi-codegen -generate types,strict-server,gin
```

### oapi-codegen 이 생성하는 것 (strict-server,gin)

```go
// 1. StrictServerInterface — 순수 비즈니스 인터페이스 (gin.Context 없음)
type StrictServerInterface interface {
    CreateWorkflow(ctx context.Context, request CreateWorkflowRequestObject) (CreateWorkflowResponseObject, error)
    ActivateWorkflow(ctx context.Context, request ActivateWorkflowRequestObject) (ActivateWorkflowResponseObject, error)
    // ...
}

// 2. typed Request Object — body + path param 자동 포함
type CreateWorkflowRequestObject struct {
    Body *CreateWorkflowJSONRequestBody
}
type ActivateWorkflowRequestObject struct {
    Id int64  // path param {id} → typed field
}

// 3. typed Response Object — status code 별
type CreateWorkflow200JSONResponse struct { ... }
type CreateWorkflow403JSONResponse struct { ... }

// 4. gin adapter — HTTP 파싱/응답 전부 자동
func NewStrictHandlerWithOptions(si StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface
```

## 활성 조건

`len(fs.ServiceFuncs) > 0`

## 진입점

```go
// generate.go
func Generate(fs *yongol.Fullstack, artifactsDir string) error
```

## 산출물

```
arts/backend/internal/service/
├── server.go                          ← Server struct (StrictServerInterface 구현)
├── create_workflow.go                 ← func (server *Server) CreateWorkflow(ctx, req) (resp, error)
├── activate_workflow.go
├── execute_workflow.go
├── on_workflow_executed.go            ← subscribe handler
└── ...
```

**전부 `package service`** (flat). Server 메서드이므로 같은 패키지 필수.
파일명은 `snake_case(funcName).go`.

## Server struct (Phase005 pgx/v5 refit)

```go
package service

import (
    "github.com/jackc/pgx/v5/pgxpool"
    "<module>/internal/db"
)

type Server struct {
    DB      *pgxpool.Pool
    Queries *db.Queries
}
```

Server 는 `api.StrictServerInterface` 를 구현. `JWTSecret` 은 Server 필드에 없음 —
middleware 가 gin context 에서 처리하고, StrictServerInterface 메서드에는 gin.Context
가 없으므로 JWT secret 이 필요한 경우 `@call auth.IssueToken` 이 내부에서 `os.Getenv` 로 읽음.

## main.go 배선 (boot/)

```go
srv := &service.Server{DB: pool, Queries: queries}
strictHandler := api.NewStrictHandlerWithOptions(srv, nil, api.StrictHTTPServerOptions{})
r := gin.Default()
r.Use(middleware.BearerAuth(jwtSecret))
api.RegisterHandlers(r, strictHandler)
```

## 트랜잭션 정책

### 규칙

**SSaC 함수에 `@post` / `@put` / `@delete` 가 1개라도 있으면** 전체 시퀀스를
DB 트랜잭션으로 감싼다.

| 조건 | 트랜잭션 |
|---|---|
| mutating seq (@post/@put/@delete) **0개** | ❌ tx 없음 — `server.Queries` 직접 사용 |
| mutating seq **1개 이상** | ✅ `Begin(ctx)` + `defer Rollback(ctx)` + `Commit(ctx)` (pgx/v5) |

### 생성 코드 (mutating — ActivateWorkflow 예시, Phase005 pgx/v5 refit)

```go
func (server *Server) ActivateWorkflow(ctx context.Context, request api.ActivateWorkflowRequestObject) (api.ActivateWorkflowResponseObject, error) {
    // ── Begin Transaction (pgx.Tx) ──
    tx, err := server.DB.Begin(ctx)
    if err != nil {
        return nil, err
    }
    defer func() {
        if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
            slog.Warn("rollback failed", "op", "ActivateWorkflow", "err", err)
        }
    }()
    qtx := server.Queries.WithTx(tx)

    // @get Workflow wf = Workflow.FindByID({ID: request.id})
    wf, err := qtx.WorkflowFindByID(ctx, request.Id)
    if err != nil {
        return nil, err
    }

    // @empty wf "Workflow not found" 404
    if wf.ID == 0 {
        return api.ActivateWorkflow404JSONResponse{Error: strPtr("Workflow not found")}, nil
    }

    // @auth "ActivateWorkflow" "workflow" {ResourceID: wf.ID} "Forbidden" 403
    if _, err := authz.Check(authz.CheckRequest{
        Action: "ActivateWorkflow", Resource: "workflow",
        UserID: currentUser.ID, Role: currentUser.Role,
        ResourceID: wf.ID,
    }); err != nil {
        return api.ActivateWorkflow403JSONResponse{Error: strPtr("Forbidden")}, nil
    }

    // @state Workflow {status: wf.Status} "ActivateWorkflow" "Cannot activate" 409
    if !statemachine.WorkflowCanTransition(wf.Status, "ActivateWorkflow") {
        return api.ActivateWorkflow409JSONResponse{Error: strPtr("Cannot activate")}, nil
    }

    // @put Workflow.UpdateStatus({ID: wf.ID, Status: "active"})
    if err := qtx.WorkflowUpdateStatus(ctx, db.WorkflowUpdateStatusParams{
        ID: wf.ID, Status: "active",
    }); err != nil {
        return nil, err
    }

    // @get Workflow updated = Workflow.FindByID({ID: wf.ID})
    updated, err := qtx.WorkflowFindByID(ctx, wf.ID)
    if err != nil {
        return nil, err
    }

    // ── Commit ──
    if err := tx.Commit(ctx); err != nil {
        return nil, err
    }

    // ── Post-commit: @publish ──
    // (이 함수에는 @publish 없음)

    return api.ActivateWorkflow200JSONResponse{Workflow: &updated}, nil
}
```

### 생성 코드 (read-only — ListWorkflows)

```go
func (server *Server) ListWorkflows(ctx context.Context, request api.ListWorkflowsRequestObject) (api.ListWorkflowsResponseObject, error) {
    workflows, err := server.Queries.WorkflowListByOrgID(ctx, currentUser.OrgID)
    if err != nil {
        return nil, err
    }
    return api.ListWorkflows200JSONResponse{Workflows: workflows}, nil
}
```

### 트랜잭션 경계

```
┌─ tx 영역 ────────────────────────────┐
│  @get (read within tx)               │
│  @empty / @exists (guard)            │
│  @auth (guard)                       │
│  @state (guard)                      │
│  @call (pure func — DB 접근 없음)     │
│  @post / @put / @delete (mutating)   │
│  tx.Commit()                         │
└──────────────────────────────────────┘
   @publish (commit 후 side effect)
   return ResponseObject
```

## 에러 → Response Object 매핑

**ServiceError type 불필요.** oapi-codegen 의 typed response 로 직접 반환:

| SSaC guard | HTTP status | 반환 |
|---|---|---|
| `@empty` 404 | 404 | `api.Xxx404JSONResponse{Error: "msg"}` |
| `@exists` 409 | 409 | `api.Xxx409JSONResponse{Error: "msg"}` |
| `@state` 409 | 409 | `api.Xxx409JSONResponse{Error: "msg"}` |
| `@auth` 403 | 403 | `api.Xxx403JSONResponse{Error: "msg"}` |
| `@call` @error NNN | NNN | `api.XxxNNNJSONResponse{Error: "msg"}` |
| 내부 에러 | 500 | `return nil, err` (adapter 가 500 처리) |

**gin.H 없음, ServiceError 없음.** 모든 에러 응답이 typed.

## currentUser 접근

strict-server 에서 gin.Context 가 없으므로 `c.MustGet("currentUser")` 불가.
대신 **context.Value** 또는 **StrictMiddlewareFunc** 로 주입:

```go
// middleware 에서 context 에 주입
func BearerAuthStrictMiddleware(f StrictHandlerFunc, operationID string) StrictHandlerFunc {
    return func(ctx context.Context, w http.ResponseWriter, r *http.Request, args interface{}) (interface{}, error) {
        // token 검증 → currentUser 생성 → context 에 주입
        ctx = context.WithValue(ctx, "currentUser", currentUser)
        return f(ctx, w, r, args)
    }
}

// service method 에서 꺼내기
currentUser := ctx.Value("currentUser").(*model.CurrentUser)
```

## Subscribe handler

subscribe 함수는 `StrictServerInterface` 에 포함되지 않음 (HTTP 가 아니므로).
별도 Server 메서드로 생성:

```go
func (server *Server) OnWorkflowExecuted(ctx context.Context, msg []byte) error {
    var message struct { ... }
    json.Unmarshal(msg, &message)
    // 시퀀스 ...
    return nil
}
```

main.go: `queue.Subscribe("workflow.executed", srv.OnWorkflowExecuted)`

## import 경로

| 종류 | import |
|---|---|
| oapi-codegen types + response | `<module>/internal/api` |
| sqlc queries | `<module>/internal/db` |
| ssac authz | `github.com/park-jun-woo/ssac/pkg/authz` |
| ssac queue | `github.com/park-jun-woo/ssac/pkg/queue` |
| ssac session/cache/file | `github.com/park-jun-woo/ssac/pkg/<pkg>` |
| project custom func | `<module>/internal/<pkg>` |
| statemachine | `<module>/internal/statemachine` |
| model (CurrentUser) | `<module>/internal/model` |
| 표준 | `"context"`, `"encoding/json"` (+ pgx 가드에 `"github.com/jackc/pgx/v5"`) |

## 파일 구조 (예정)

```
pkg/generate/gogin/ssac/
├── README.md
├── generate.go                     ← orchestrator
├── generate_server_go.go           ← Server struct
├── generate_http_method.go         ← HTTP StrictServerInterface method 1개
├── generate_subscribe_method.go    ← Subscribe method 1개
├── needs_transaction.go            ← mutating seq 판정
├── build_tx_open.go                ← BeginTx + defer Rollback + WithTx
├── build_tx_commit.go              ← tx.Commit()
├── build_sequence_get.go           ← @get → qtx.XxxFind/List
├── build_sequence_post.go          ← @post → qtx.XxxCreate
├── build_sequence_put.go           ← @put → qtx.XxxUpdate
├── build_sequence_delete.go        ← @delete → qtx.XxxDelete
├── build_sequence_empty.go         ← @empty → typed 404 response
├── build_sequence_exists.go        ← @exists → typed 409 response
├── build_sequence_state.go         ← @state → CanTransition
├── build_sequence_auth.go          ← @auth → authz.Check → typed 403 response
├── build_sequence_call.go          ← @call → func invocation
├── build_sequence_publish.go       ← @publish → queue.Publish (post-commit)
├── build_response.go              ← @response → typed ResponseObject
├── map_args.go                     ← SSaC Args → Go code
├── resolve_sqlc_method.go          ← Model.Method → Queries.<Prefix><Method>
├── resolve_import.go               ← @call pkg → import path
├── resolve_err_status.go           ← @error annotation → response type
├── zero_value_check.go             ← @empty/@exists zero-value 비교
├── assemble_method.go              ← imports + body → Go source + gofmt
└── write_method_file.go            ← os.WriteFile
```
