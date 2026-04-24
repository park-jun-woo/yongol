# pkg/generate/gogin/boot

`artifacts/<project>/backend/cmd/main.go` 를 생성하는 Block Builder 엔진.

## 설계 원칙

1. **Block = 단위** — main.go 의 각 초기화 영역(DB, JWT, authz, queue, session, ...)은
   독립 블록. 블록은 **조건 + imports + lines** 로 구성.
2. **조건부 포함** — `Fullstack` (manifest, SSaC, policy) 를 보고 필요한 블록만 수집.
   불필요한 블록은 코드에 흔적 없이 제외.
3. **filefunc 준수** — 1 블록 = 1 파일 (block_*.go). orchestrator 1 파일 (generate.go).
   MainBlock type 1 파일.
4. **import dedup** — 블록별 import 를 합산 후 중복 제거. orchestrator 가 최종 import
   블록을 조립.
5. **순서 고정** — 블록 조립 순서는 orchestrator 가 하드코딩 (Go 의미에 영향:
   `defer` 위치, 변수 선언 순서).

## 런타임 라이브러리: `github.com/park-jun-woo/ssac/pkg/`

생성된 main.go 는 yongol 자체 코드가 아닌 **ssac 모듈의 런타임 패키지**를 import 한다.
인증, 인가, 세션, 캐시, 큐, 파일 등 런타임 기능은 모두 ssac/pkg/ 에 구현되어 있으며,
코드젠은 이를 **그대로 가져다 쓰는 코드**를 생성한다.

```
github.com/park-jun-woo/ssac/pkg/
├── auth/       hashPassword, verifyPassword, generateResetToken
├── authz/      Init(db, ownerships), Check(req)
├── cache/      Init(model), NewPostgresCache(ctx,db), NewMemoryCache()
├── file/       Init(model), NewLocalFile(root), NewS3File(client,bucket)
├── mail/       sendEmail, sendTemplateEmail
├── pagination/ Page[T], Cursor[T]
├── queue/      Init(ctx,backend,db), Publish(ctx,topic,payload), Subscribe(topic,handler), Start(ctx), Close()
├── session/    Init(model), NewPostgresSession(ctx,db), NewMemorySession()
├── crypto/     encrypt, decrypt, generateOTP, verifyOTP
├── storage/    uploadFile, deleteFile, presignURL
├── text/       generateSlug, sanitizeHTML, truncateText
└── image/      ogImage, thumbnail
```

**코드젠 규칙**: 블록이 생성하는 코드에서 위 패키지를 직접 호출한다. yongol 자체
(`github.com/park-jun-woo/yongol/`) 는 런타임에 import 되지 않음 — yongol 는
빌드 도구이지 런타임 의존이 아니다.

## MainBlock 구조

```go
type MainBlock struct {
    Name    string                           // "db-init", "jwt-secret", ...
    Active  func(fs *yongol.Fullstack) bool // nil = 항상 활성
    Imports []string                         // `"database/sql"`, `"github.com/park-jun-woo/ssac/pkg/authz"`
    Lines   []string                         // main() 본문 코드 라인
}
```

## 블록 목록 (조립 순서)

| # | 파일 | 블록 이름 | 활성 조건 | ssac/pkg import | 생성 코드 |
|---|---|---|---|---|---|
| 0 | `block_logger_init.go` | `logger-init` | 항상 (최상단) | — | `slog.SetDefault(slog.New(handler))` — `LOG_LEVEL`/`LOG_FORMAT` 환경변수 기반 |
| 0.5 | `block_env_helpers.go` | `env-helpers` | 항상 | — | top-level `envInt`, `envDuration`, `envStringList`, `envBool` (main 외부) |
| 1 | `block_db_init.go` | `db-init` | 항상 | — | `pgxpool.NewWithConfig(ctx, poolCfg)` + `stdlib.OpenDBFromPool(pool)` 브릿지, `db.New(pool)` (Phase005 pgx/v5 refit) |
| 2 | `block_jwt_secret.go` | `jwt-secret` | `manifest.backend.auth` 존재 | — | `os.Getenv(manifest.auth.secret_env)` |
| 3 | `block_authz_init.go` | `authz-init` | SSaC 에 `@auth` 사용 | `ssac/pkg/authz` | `authz.Init(conn, []authz.OwnershipMapping{...})` |
| 4 | `block_queue_init.go` | `queue-init` | `manifest.queue.backend` 존재 | `ssac/pkg/queue` | `queue.Init(ctx, "postgres", conn)` + `queue.Subscribe(...)` + `queue.Start(ctx)` + `defer queue.Close()` |
| 5 | `block_session_init.go` | `session-init` | `manifest.session.backend` 존재 | `ssac/pkg/session` | postgres: `session.Init(session.NewPostgresSession(ctx, conn))` / memory: `session.Init(session.NewMemorySession())` |
| 6 | `block_cache_init.go` | `cache-init` | `manifest.cache.backend` 존재 | `ssac/pkg/cache` | postgres: `cache.Init(cache.NewPostgresCache(ctx, conn))` / memory: `cache.Init(cache.NewMemoryCache())` |
| 7 | `block_file_init.go` | `file-init` | `manifest.file.backend` 존재 | `ssac/pkg/file` | local: `file.Init(file.NewLocalFile("./uploads"))` / s3: `file.Init(file.NewS3File(s3Client, bucket))` |
| 8 | `block_server_struct.go` | `server` | 항상 | — | `&Server{Queries: queries, ...}` |
| 9 | `block_middleware.go` | `middleware` | `manifest.backend.middleware` 에 `bearerAuth` | — (자체 생성된 middleware 패키지) | `r.Use(middleware.BearerAuth(jwtSecret))` |
| 10 | `block_router.go` | `router` | 항상 | — | `gin.Default()`, `api.RegisterHandlers(r, srv)` |
| 10.3 | `block_cors.go` | `cors` | `manifest.backend.cors.enabled=true` | — | `r.Use(cors.New(corsCfg))` — gin-contrib/cors 기반, env 오버라이드 포함 |
| 10.5 | `block_health.go` | `health` | 항상 | — | `r.GET("/health", ...)` + `r.GET("/ready", ...)` (DDL 있을 때 `conn.PingContext` 포함) |
| 11 | `block_gin_run.go` | `gin-run` | 항상 | — | `http.Server` + SIGINT/SIGTERM graceful shutdown (`httpSrv.Shutdown(ctx)` 10s) |

## ssac/pkg 호출 패턴 상세

### authz (block #3)

```go
import "github.com/park-jun-woo/ssac/pkg/authz"

// Rego @ownership 에서 추출한 리터럴
if err := authz.Init(conn, []authz.OwnershipMapping{
    {Resource: "workflow", Table: "workflows", Column: "org_id"},
    {Resource: "webhook", Table: "webhooks", Column: "org_id"},
    // ... ParsedPolicies 에서 동적 생성
}); err != nil {
    slog.Error("authz init", "err", err)
    os.Exit(1)
}
```

`authz.OwnershipMapping` 의 `Resource/Table/Column/JoinTable/JoinFK` 필드는
`fs.ParsedPolicies[].Ownerships` 에서 그대로 가져온다.

### queue (block #4)

```go
import "github.com/park-jun-woo/ssac/pkg/queue"

ctx := context.Background()
if err := queue.Init(ctx, "postgres", conn); err != nil {
    slog.Error("queue init", "err", err)
    os.Exit(1)
}
defer queue.Close()

// @subscribe 함수마다 1개
queue.Subscribe("workflow.executed", srv.OnWorkflowExecuted)
// ... SSaC @subscribe 에서 동적 생성

go func() {
    if err := queue.Start(ctx); err != nil {
        slog.Error("queue start", "err", err)
    }
}()
```

`queue.Subscribe` 의 topic 과 handler 함수명은 SSaC `@subscribe "topic"` + func name
에서 추출.

### session (block #5)

```go
import "github.com/park-jun-woo/ssac/pkg/session"

// manifest.session.backend == "postgres"
sm, err := session.NewPostgresSession(ctx, conn)
if err != nil {
    slog.Error("session init", "err", err)
    os.Exit(1)
}
session.Init(sm)

// manifest.session.backend == "memory"
session.Init(session.NewMemorySession())
```

### cache (block #6)

```go
import "github.com/park-jun-woo/ssac/pkg/cache"

// manifest.cache.backend == "postgres"
cm, err := cache.NewPostgresCache(ctx, conn)
if err != nil {
    slog.Error("cache init", "err", err)
    os.Exit(1)
}
cache.Init(cm)

// manifest.cache.backend == "memory"
cache.Init(cache.NewMemoryCache())
```

### file (block #7)

```go
import "github.com/park-jun-woo/ssac/pkg/file"

// manifest.file.backend == "local"
file.Init(file.NewLocalFile(os.Getenv("FILE_ROOT")))

// manifest.file.backend == "s3"
s3Client := ... // AWS SDK 초기화
file.Init(file.NewS3File(s3Client, os.Getenv("S3_BUCKET")))
```

## SSaC handler 에서의 ssac/pkg 사용

main.go 외에도 SSaC handler codegen 이 생성하는 코드에서 ssac/pkg 를 직접 호출한다:

| SSaC 시퀀스 | 생성 코드 | ssac/pkg import |
|---|---|---|
| `@auth "Action" "Resource" {ResourceID: x}` | `authz.Check(authz.CheckRequest{...})` | `ssac/pkg/authz` |
| `@call auth.HashPassword({...})` | `auth.HashPassword(auth.HashPasswordRequest{...})` | `ssac/pkg/auth` |
| `@call session.Set({...})` | `session.Set(session.SetRequest{...})` | `ssac/pkg/session` |
| `@call cache.Get({...})` | `cache.Get(cache.GetRequest{...})` | `ssac/pkg/cache` |
| `@call file.Upload({...})` | `file.Upload(file.UploadRequest{...})` | `ssac/pkg/file` |
| `@publish "topic" {payload}` | `queue.Publish(ctx, "topic", payload)` | `ssac/pkg/queue` |

## 산출물 예시 (zenflow try-02 기준)

```go
package main

import (
    "context"
    "database/sql"
    "log/slog"
    "net/http"
    "os"
    "os/signal"
    "strings"
    "syscall"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/jackc/pgx/v5/stdlib"

    "github.com/park-jun-woo/ssac/pkg/authz"
    "github.com/park-jun-woo/ssac/pkg/cache"
    "github.com/park-jun-woo/ssac/pkg/file"
    "github.com/park-jun-woo/ssac/pkg/queue"
    "github.com/park-jun-woo/ssac/pkg/session"

    "github.com/park-jun-woo/zenflow/internal/api"
    "github.com/park-jun-woo/zenflow/internal/db"
    "github.com/park-jun-woo/zenflow/internal/middleware"
)

func main() {
    // [0] logger-init — slog 기본 핸들러. LOG_LEVEL (DEBUG/INFO/WARN/ERROR),
    //     LOG_FORMAT (JSON 기본/TEXT) 환경변수로 제어.
    logLevel := slog.LevelInfo
    switch strings.ToUpper(os.Getenv("LOG_LEVEL")) {
    case "DEBUG":
        logLevel = slog.LevelDebug
    case "WARN":
        logLevel = slog.LevelWarn
    case "ERROR":
        logLevel = slog.LevelError
    }
    var handler slog.Handler
    if strings.ToUpper(os.Getenv("LOG_FORMAT")) == "TEXT" {
        handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
    } else {
        handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
    }
    slog.SetDefault(slog.New(handler))

    // [1] db-init (Phase005 pgx/v5 refit)
    ctx := context.Background()
    slog.Info("connecting to database")
    poolCfg, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
    if err != nil {
        slog.Error("db init: parse DATABASE_URL", "err", err)
        os.Exit(1)
    }
    pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
    if err != nil {
        slog.Error("db init", "err", err)
        os.Exit(1)
    }
    defer pool.Close()
    // Bridge pool → *sql.DB for ssac packages (auth / queue / authz / cache / session).
    var conn *sql.DB = stdlib.OpenDBFromPool(pool)
    defer func() { _ = conn.Close() }()
    queries := db.New(pool)
    slog.Info("database connected")

    // [2] jwt-secret
    jwtSecret := os.Getenv("JWT_SECRET")

    // [3] authz-init
    slog.Info("initializing authz")
    if err := authz.Init(conn, []authz.OwnershipMapping{
        {Resource: "workflow", Table: "workflows", Column: "org_id"},
        {Resource: "webhook", Table: "webhooks", Column: "org_id"},
        {Resource: "template", Table: "templates", Column: "org_id"},
        {Resource: "execution_log", Table: "execution_logs", Column: "org_id"},
    }); err != nil {
        slog.Error("authz init", "err", err)
        os.Exit(1)
    }

    // [4] queue-init
    slog.Info("initializing queue")
    if err := queue.Init(ctx, "postgres", conn); err != nil {
        slog.Error("queue init", "err", err)
        os.Exit(1)
    }
    defer queue.Close()
    queue.Subscribe("workflow.executed", srv.OnWorkflowExecuted)
    go func() {
        if err := queue.Start(ctx); err != nil {
            slog.Error("queue start", "err", err)
        }
    }()

    // [5] session-init
    slog.Info("initializing session (postgres)")
    sm, err := session.NewPostgresSession(ctx, conn)
    if err != nil {
        slog.Error("session init", "err", err)
        os.Exit(1)
    }
    session.Init(sm)

    // [6] cache-init
    slog.Info("initializing cache (postgres)")
    cm, err := cache.NewPostgresCache(ctx, conn)
    if err != nil {
        slog.Error("cache init", "err", err)
        os.Exit(1)
    }
    cache.Init(cm)

    // [7] file-init
    file.Init(file.NewLocalFile(os.Getenv("FILE_ROOT")))

    // [8] server
    srv := &Server{
        Queries:   queries,
        JWTSecret: jwtSecret,
    }

    // [9] middleware + [10] router
    r := gin.Default()
    r.Use(middleware.BearerAuth(jwtSecret))
    api.RegisterHandlers(r, srv)

    // [11] gin-run — http.Server + SIGINT/SIGTERM graceful shutdown.
    //      SIGINT/SIGTERM 수신 시 최대 10초 동안 진행 중 요청을 마친 뒤
    //      main 이 정상 return → defer 로 queue.Close()/conn.Close() 수행.
    httpSrv := &http.Server{Addr: ":8080", Handler: r}
    go func() {
        slog.Info("server starting", "addr", ":8080")
        if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            slog.Error("server", "err", err)
            os.Exit(1)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    slog.Info("shutting down")

    shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := httpSrv.Shutdown(shutdownCtx); err != nil {
        slog.Error("server shutdown", "err", err)
    }
    slog.Info("shutdown complete")
}
```

### 환경변수

| 변수 | 기본 | 효과 |
|------|------|------|
| `LOG_LEVEL` | `INFO` | `DEBUG` / `INFO` / `WARN` / `ERROR` |
| `LOG_FORMAT` | `JSON` | `TEXT` 로 설정 시 사람 가독형 |
| `DB_MAX_OPEN_CONNS` | `25` | `conn.SetMaxOpenConns` 상한 |
| `DB_MAX_IDLE_CONNS` | `5` | `conn.SetMaxIdleConns` 풀 유지 수 |
| `DB_CONN_MAX_LIFETIME` | `5m` | `conn.SetConnMaxLifetime` (`time.ParseDuration`) |
| `CORS_ALLOW_ORIGINS` | (manifest 값) | `,` 구분 문자열로 manifest allow_origins 덮어쓰기 |
| `CORS_ALLOW_METHODS` | (manifest 값) | `,` 구분 문자열로 allow_methods 덮어쓰기 |
| `CORS_ALLOW_CREDENTIALS` | (manifest 값) | `true`/`false`/`1` 로 allow_credentials 덮어쓰기 |

## orchestrator 흐름

```go
// generate.go
func Generate(fs *yongol.Fullstack, artifactsDir, modulePath string) error {
    blocks := collectActiveBlocks(fs)   // Active 조건 필터링
    imports := deduplicateImports(blocks, modulePath)
    body := assembleBody(blocks)
    return writeMainGo(artifactsDir, imports, body)
}
```

## Server struct 결정 로직

`block_server_struct.go` 는 `Fullstack` 에서 다음을 수집:
- `Queries *db.Queries` — 항상
- `JWTSecret string` — manifest.auth 있을 때
- custom func 필드 — SSaC `@call` 이 참조하는 **프로젝트 func** 패키지별
  (ssac/pkg/ 함수는 Server 필드가 아님 — 패키지 레벨 직접 호출)

이 struct 는 **oapi-codegen 의 `ServerInterface` 를 구현**해야 하므로,
별도 `server.go` 파일에 type + 빈 메서드 stub 도 함께 생성 (SSaC handler codegen 이
채우기 전까지 컴파일 되도록).

## import 경로 규칙

| 카테고리 | import 경로 | 예시 |
|---|---|---|
| ssac 런타임 | `github.com/park-jun-woo/ssac/pkg/<pkg>` | `ssac/pkg/authz`, `ssac/pkg/queue` |
| 프로젝트 내부 | `<manifest.backend.module>/internal/<pkg>` | `zenflow/internal/api`, `zenflow/internal/db` |
| 프로젝트 custom func | `<manifest.backend.module>/internal/<func-pkg>` | `zenflow/internal/billing` |
| 표준 라이브러리 | 그대로 | `"database/sql"` (ssac 브릿지 전용), `"context"`, `"os"` |
| 외부 의존 | 그대로 | `"github.com/gin-gonic/gin"`, `"github.com/jackc/pgx/v5/pgxpool"`, `"github.com/jackc/pgx/v5/stdlib"` |

생성된 프로젝트의 `go.mod` 에는 `require github.com/park-jun-woo/ssac` 가 포함되어야 함.
`go.mod` 생성도 후속 블록 또는 별도 Phase 로 처리.

## 의존성 (입력)

| 입력 | 소스 | 블록에서 사용 |
|---|---|---|
| `fs.Manifest` | manifest.yaml | auth, session, cache, queue, file, middleware 판단 |
| `fs.ServiceFuncs` | SSaC | @auth/@call/@publish/@subscribe 사용 여부 |
| `fs.ParsedPolicies` | Rego | @ownership mapping 리터럴 (authz-init 블록) |
| `fs.StateDiagrams` | states/ | state machine init (후속) |
| `modulePath` | manifest.backend.module | import 경로 조립 |
| `artifactsDir` | CLI 인자 | 출력 경로 |

## 후속 Phase 연결

- **SSaC handler codegen** (`pkg/generate/gogin/handler/`) — `ServerInterface` 메서드 본문
  생성. ssac/pkg/ 를 직접 호출하는 코드 포함 (`authz.Check`, `auth.HashPassword` 등).
- **middleware codegen** (`pkg/generate/gogin/middleware/`) — bearerAuth Go 파일 생성.
  ssac/pkg/auth 의 VerifyToken 을 호출하는 middleware.
- **authz codegen** (`pkg/generate/gogin/authz/`) — Rego 복사 + Init 리터럴 생성.
- **go.mod codegen** — `require github.com/park-jun-woo/ssac` 포함.
