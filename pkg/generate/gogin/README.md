# pkg/generate/gogin

**Go + Gin 백엔드** 코드 생성 orchestrator. yongol generate 의 `--backend go-gin`
경로. 외부 도구(oapi-codegen, sqlc) 호출 + 하위 패키지 생성기 조율 + 자체 glue
(func copy, go.mod) 를 책임진다.

## 진입점

```go
// generate.go
func Generate(fs *yongol.Fullstack, artifactsDir string) error
```

`pkg/generate/run_backend.go` 에서 `BackendType == GoGin` 일 때 호출.

## 파이프라인 (실행 순서)

```
Generate(fs, artifactsDir)
  [1]  generateOpenAPIGoGin     — oapi-codegen 외부 실행
  [2]  postgresql.Generate      — 동적 pagination/sort/filter SQL 생성
  [3]  generateSQLCGoGin        — sqlc 외부 실행
  [4]  auth.Generate            — JWT issue/verify/refresh + reexport
  [5]  middleware.Generate      — bearerAuth gin middleware
  [6]  model.Generate           — CurrentUser struct
  [7]  statemachine.Generate    — CanTransition 함수
  [8]  copyFuncSpecs            — specs/func/ → internal/<pkg>/ 복사
  [9]  boot.Generate            — cmd/main.go (Block Builder)
  [10] handler.Generate         — ServerInterface method body (SSaC 기반)
  [11] generateGoMod            — go.mod + go mod tidy
```

### 순서 근거

| 단계 | 의존 | 이유 |
|---|---|---|
| [1] oapi-codegen | 없음 | OpenAPI → types + ServerInterface 생성 |
| [2] postgresql | [1] 의 OpenAPI x-extensions 파싱 결과 | pagination SQL 은 sqlc 이전에 생성 |
| [3] sqlc | [2] 의 *_generated.sql | 정적 + 동적 쿼리 모두 Go 변환 |
| [4] auth | manifest claims | JWT 함수 — middleware 가 import |
| [5] middleware | [4] auth | VerifyToken 호출 코드 |
| [6] model | manifest claims | CurrentUser struct — handler/middleware 가 import |
| [7] statemachine | states/ | CanTransition — handler 가 import |
| [8] func copy | specs/func/ | custom func — handler 가 import |
| [9] root | [4]~[8] 존재 전제 | main.go 가 모든 internal 패키지 import |
| [10] handler | [1]~[9] 전부 | ServerInterface 구현 — 모든 internal 패키지 사용 |
| [11] go.mod | [1]~[10] 전부 | require 수집 후 `go mod tidy` |

## 산출물 맵

```
arts/backend/
├── cmd/
│   └── main.go                         ← [9] root/
├── internal/
│   ├── api/
│   │   └── server.gen.go               ← [1] oapi-codegen
│   ├── db/
│   │   ├── db.go                       ← [3] sqlc
│   │   ├── models.go                   ← [3] sqlc
│   │   └── *.sql.go                    ← [3] sqlc (정적 + [2] 동적)
│   ├── auth/
│   │   ├── issue_token.go              ← [4] auth/
│   │   ├── verify_token.go             ← [4]
│   │   ├── refresh_token.go            ← [4]
│   │   └── reexport.go                 ← [4] ssac/pkg/auth 재export
│   ├── middleware/
│   │   └── bearerauth.go              ← [5] middleware/
│   ├── model/
│   │   └── current_user.go            ← [6] model/
│   ├── statemachine/
│   │   └── workflow.go                ← [7] statemachine/
│   ├── <custom-func-pkg>/
│   │   └── *.go                       ← [8] func copy
│   └── service/
│       ├── server.go                  ← [10] handler/
│       ├── workflow/
│       │   ├── create_workflow.go     ← [10]
│       │   ├── activate_workflow.go   ← [10]
│       │   └── ...
│       └── ...
└── go.mod                              ← [11] go.mod
```

## 자체 담당 (이 패키지 직접 구현)

하위 패키지에 위임하지 않고 `gogin/` 에서 직접 처리하는 항목:

### [1] generateOpenAPIGoGin (`generate_openapi_gogin.go`)

```bash
oapi-codegen -package api -generate types,gin-server -o <out> <openapi.yaml>
```

### [3] generateSQLCGoGin (`generate_sqlc_gogin.go`)

```bash
sqlc generate -f sqlc.yaml    # cmd.Dir = specs/db/
```

### [8] copyFuncSpecs (`copy_func_specs.go` — 예정)

`specs/func/<pkg>/*.go` → `arts/backend/internal/<pkg>/*.go` 단순 복사.

- 파일 내용 변경 없음 (사용자 코드 그대로)
- `package <pkg>` 선언 유지
- import 경로: SSaC `@call` 이 `import "<module>/internal/<pkg>"` 로 참조
- 디렉토리 구조 보존: `func/billing/spend.go` → `internal/billing/spend.go`

### [11] generateGoMod (`generate_go_mod.go` — 예정)

`arts/backend/go.mod` 생성 + `go mod tidy` 실행.

```go
module <manifest.backend.module>

go 1.22

require (
    github.com/gin-gonic/gin v1.x
    github.com/jackc/pgx/v5 v5.x           // pgxpool + stdlib bridge (ssac 호환)
    github.com/park-jun-woo/ssac v0.x     // authz, queue, session, cache, file, auth
    github.com/golang-jwt/jwt/v5 v5.x     // auth codegen
    github.com/oapi-codegen/runtime v1.x   // oapi-codegen generated code
)
```

require 목록은 [1]~[10] 에서 생성된 import 를 스캔하거나, 필요 패키지를 하드코딩.
이후 `go mod tidy` 로 정리.

## 하위 패키지 위임

| 단계 | 패키지 | README |
|---|---|---|
| [2] | `pkg/generate/postgresql/` | ✅ |
| [4] | `pkg/generate/gogin/auth/` | ⏳ (향후) |
| [5] | `pkg/generate/gogin/middleware/` | ⏳ |
| [6] | `pkg/generate/gogin/model/` | ⏳ |
| [7] | `pkg/generate/gogin/statemachine/` | ⏳ |
| [9] | `pkg/generate/gogin/boot/` | ✅ 구현 |
| [10] | `pkg/generate/gogin/handler/` | ✅ README |
| — | `pkg/generate/gogin/ssac/` | ✅ README (handler 와 통합 검토) |

### ssac/ vs handler/ 관계

`gogin/ssac/` 와 `gogin/handler/` 는 동일 산출물 (`internal/service/**/*.go`) 을 다룸.
**handler/ 가 SSaC 시퀀스 → Go handler body 변환의 최종 패키지**. ssac/ README 는
설계 배경 문서로 유지하되, 구현은 handler/ 에 집중.

## 조건부 단계

| 단계 | 조건 | 미충족 시 |
|---|---|---|
| [4] auth | `manifest.backend.auth` 존재 | skip |
| [5] middleware | `manifest.backend.middleware` 에 `bearerAuth` | skip |
| [6] model | `manifest.backend.auth` 존재 (CurrentUser 생성) | skip |
| [7] statemachine | `fs.StateDiagrams` 1개 이상 | skip |
| [8] func copy | `specs/func/` 디렉토리 존재 | skip |

나머지 ([1], [2], [3], [9], [10], [11]) 는 항상 실행.

## 에러 처리

각 단계가 error 반환 시 즉시 중단 + `fmt.Errorf("<step>: %w", err)` 로 래핑.
부분 산출물은 남아있을 수 있음 — `--clean` 플래그 (향후) 로 arts/ 전체 삭제 후 재생성.

## 외부 도구 의존

| 도구 | 버전 | 설치 |
|---|---|---|
| `oapi-codegen` | v2.6+ | `go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest` |
| `sqlc` | v1.30+ | `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest` |
| `gofmt` | Go 기본 | Go 설치 시 포함 |
| `go mod tidy` | Go 기본 | Go 설치 시 포함 |
