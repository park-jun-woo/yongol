# pkg/generate

yongol 코드 생성 최상위 orchestrator. `yongol generate <specs> <arts>` 가 호출하는
`Generate(fs, artifactsDir, backend, frontend)` 의 진입점.

## 진입점

```go
// generate.go
func Generate(fs *yongol.Fullstack, artifactsDir string, backend BackendType, frontend FrontendType) error
```

`cmd/yongol/generate_cmd.go` → `pkg/generate/generate.go` → backend/frontend/hurl/opa 순으로 위임.

## 파이프라인 (전체)

```
Generate(fs, artifactsDir, GoGin, React)
  ├─ [A] runBackend    → gogin.Generate
  │     [1]  oapi-codegen          → backend/internal/api/server.gen.go
  │     [2]  postgresql.Generate   → specs/db/queries/*_generated.sql
  │     [3]  sqlc generate         → backend/internal/db/*.go
  │     [4]  auth.Generate         → backend/internal/auth/ + middleware/ + model/
  │     [7]  state.Generate        → backend/internal/statemachine/*.go
  │     [8]  copyFuncSpecs         → backend/internal/<pkg>/*.go
  │     [9]  boot.Generate         → backend/cmd/main.go
  │     [10] handler.Generate      → backend/internal/service/**/*.go
  │     [11] generateGoMod         → backend/go.mod + go mod tidy
  │
  ├─ [B] runFrontend   → react.Generate (향후)
  │     → frontend/
  │
  ├─ [C] hurl.Generate → tests/smoke.hurl
  │
  └─ [D] copyOPARego   → backend/policy/*.rego
```

## 산출물 전체 맵

```
arts/
├── backend/
│   ├── cmd/
│   │   └── main.go                         ← [9] boot/
│   ├── internal/
│   │   ├── api/
│   │   │   └── server.gen.go               ← [1] oapi-codegen
│   │   ├── db/
│   │   │   ├── db.go                       ← [3] sqlc
│   │   │   ├── models.go                   ← [3] sqlc
│   │   │   └── *.sql.go                    ← [3] sqlc (정적 + [2] 동적)
│   │   ├── auth/
│   │   │   ├── issue_token.go              ← [4] auth/
│   │   │   ├── verify_token.go             ← [4]
│   │   │   ├── refresh_token.go            ← [4]
│   │   │   └── reexport.go                ← [4] ssac/pkg/auth 재export
│   │   ├── middleware/
│   │   │   └── bearerauth.go              ← [4] auth/
│   │   ├── model/
│   │   │   └── current_user.go            ← [4] auth/
│   │   ├── statemachine/
│   │   │   └── workflow.go                ← [7] state/
│   │   ├── <custom-func-pkg>/
│   │   │   └── *.go                       ← [8] func copy
│   │   └── service/
│   │       ├── server.go                  ← [10] handler/
│   │       └── <feature>/<func>.go        ← [10] handler/
│   ├── policy/
│   │   └── authz.rego                     ← [D] OPA rego copy
│   └── go.mod                             ← [11] go.mod
├── frontend/                               ← [B] react (향후)
└── tests/
    └── smoke.hurl                          ← [C] hurl/
```

## 하위 패키지

| 패키지 | 역할 | 상태 |
|---|---|---|
| `pkg/generate/gogin/` | Go+Gin 백엔드 orchestrator | ✅ 구현 + 📝 README |
| `pkg/generate/gogin/boot/` | cmd/main.go Block Builder | ✅ 구현 |
| `pkg/generate/gogin/handler/` | ServerInterface method body (SSaC 기반) | 📝 README |
| `pkg/generate/gogin/auth/` | JWT + middleware + CurrentUser (통합) | 📝 README |
| `pkg/generate/gogin/state/` | stateDiagram → CanTransition | 📝 README |
| `pkg/generate/gogin/ssac/` | SSaC codegen 설계 배경 문서 | 📝 README |
| `pkg/generate/postgresql/` | 동적 pagination/sort/filter SQL 생성 | 📝 README |
| `pkg/generate/react/` | React 프론트엔드 (향후) | ⏳ 빈 껍질 |
| `pkg/generate/hurl/` | Hurl 스모크 테스트 | 📝 README |

## 자체 담당 (이 패키지 직접 구현)

### generate.go — orchestrator

```go
func Generate(fs, artifactsDir, backend, frontend) error {
    runBackend(fs, artifactsDir, backend)   // [A]
    runFrontend(fs, frontend)               // [B]
    hurl.Generate(fs, artifactsDir)         // [C]
    copyOPARego(fs, artifactsDir)           // [D]
}
```

### run_backend.go — BackendType 분기

```go
func runBackend(fs, artifactsDir, backend) error {
    switch backend {
    case GoGin: return gogin.Generate(fs, artifactsDir)
    }
}
```

### run_frontend.go — FrontendType 분기

```go
func runFrontend(fs, frontend) error {
    switch frontend {
    case React: return react.Generate(fs)
    }
}
```

### copy_opa_rego.go — [D] OPA Rego 복사 (예정)

`specs/policy/*.rego` → `arts/backend/policy/*.rego` 단순 복사.

- 파일 내용 변경 없음 (사용자 작성 Rego 그대로)
- main.go 의 `OPA_POLICY_PATH` 환경변수가 이 복사된 경로를 가리킴
- `authz.Init(conn, ownerships)` 는 런타임에 이 .rego 를 로드
- 복사 이유: arts/ 를 배포 단위로 패키징할 때 specs/ 의존 없이 자족 가능

```go
func copyOPARego(fs *yongol.Fullstack, artifactsDir string) error {
    // specs/policy/*.rego → arts/backend/policy/*.rego
}
```

## BackendType / FrontendType

```go
type BackendType string
const GoGin BackendType = "go-gin"

type FrontendType string
const React FrontendType = "react"
```

CLI 플래그 `--backend go-gin` `--frontend react` 로 선택. 현재 각 1종만 지원.

## generate 게이트

`cmd/yongol/generate_cmd.go` 가 Generate 호출 전에:
1. `ParseAll` → parser diagnostic 0 확인
2. `validate.Validate` → ERROR 0 확인
3. `printReport` 2번째 반환값(warnings) → WARNING 0 확인

**ERROR 또는 WARNING 이 하나라도 있으면 Generate 호출 안 됨** (validate 보다 엄격).

## 에러 처리

각 하위 단계가 error 반환 시 즉시 중단. `generate_cmd.go` 가 에러를 출력하고 exit 1.
부분 산출물은 남아있을 수 있음.

## 외부 도구

| 도구 | 호출 위치 | 용도 |
|---|---|---|
| `oapi-codegen` | gogin [1] | OpenAPI → Go types + gin-server |
| `sqlc` | gogin [3] | DDL + queries → Go model + query methods |
| `gofmt` | handler [10] | 생성된 Go 코드 포매팅 |
| `go mod tidy` | gogin [11] | go.mod 의존성 정리 |
