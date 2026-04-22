# pkg/generate/gogin/auth

JWT 인증 + bearerAuth 미들웨어 + CurrentUser 모델을 **한 번에** 생성한다.
셋 모두 `manifest.backend.auth.claims` 에서 파생되며, claims 가 없으면 전체 skip.

## 활성 조건

`manifest.backend.auth != nil && len(claims) > 0`

## 진입점

```go
// generate.go
func Generate(fs *yongol.Fullstack, artifactsDir string) error
```

gogin 파이프라인 [4]+[5]+[6] 을 통합.

## 입력

| 소스 | 데이터 |
|---|---|
| `manifest.backend.auth.claims` | 필드명 + JWT claim key + Go 타입. 예: `ID: user_id:int64` |
| `manifest.backend.auth.secret_env` | JWT secret 환경변수 이름 (기본 `JWT_SECRET`) |
| `manifest.backend.middleware` | `bearerAuth` 포함 여부 → middleware 생성 판단 |
| `manifest.backend.module` | import 경로 (`<module>/internal/auth`, `<module>/internal/model`) |

## 산출물 (5 파일)

```
arts/backend/internal/
├── auth/
│   ├── issue_token.go       ← JWT 발급 (access, 24h)
│   ├── verify_token.go      ← JWT 검증 → claims 추출
│   ├── refresh_token.go     ← JWT 갱신 (refresh, 7d)
│   └── reexport.go          ← ssac/pkg/auth 재export (HashPassword, VerifyPassword, ...)
├── middleware/
│   └── bearerauth.go        ← gin middleware: Authorization → VerifyToken → currentUser
└── model/
    └── current_user.go      ← CurrentUser struct (claims 필드)
```

## claims → 코드 매핑

manifest:
```yaml
claims:
  ID: user_id:int64
  Email: email           # 타입 생략 = string
  Role: role
  OrgID: org_id:int64
```

### CurrentUser (`internal/model/current_user.go`)

```go
package model

type CurrentUser struct {
    ID    int64
    Email string
    Role  string
    OrgID int64
}
```

필드 순서: manifest claims 선언 순. 필드명 = claim 키의 왼쪽 (PascalCase).
Go 타입 = `claim_key:go_type` 에서 `:` 뒤 (없으면 `string`).

### IssueToken (`internal/auth/issue_token.go`)

```go
package auth

import (
    "os"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

type IssueTokenRequest struct {
    ID    int64
    Email string
    Role  string
    OrgID int64
}

type IssueTokenResponse struct {
    AccessToken string
}

func IssueToken(req IssueTokenRequest) (IssueTokenResponse, error) {
    secret := os.Getenv("JWT_SECRET")   // manifest.auth.secret_env
    claims := jwt.MapClaims{
        "user_id": req.ID,              // claim key from manifest
        "email":   req.Email,
        "role":    req.Role,
        "org_id":  req.OrgID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signed, err := token.SignedString([]byte(secret))
    return IssueTokenResponse{AccessToken: signed}, err
}
```

### VerifyToken (`internal/auth/verify_token.go`)

```go
package auth

import (
    "fmt"
    "github.com/golang-jwt/jwt/v5"
)

type VerifyTokenRequest struct {
    Token  string
    Secret string
}

type VerifyTokenResponse struct {
    ID    int64
    Email string
    Role  string
    OrgID int64
}

func VerifyToken(req VerifyTokenRequest) (VerifyTokenResponse, error) {
    token, err := jwt.Parse(req.Token, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
        }
        return []byte(req.Secret), nil
    })
    if err != nil { return VerifyTokenResponse{}, err }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid { return VerifyTokenResponse{}, fmt.Errorf("invalid token") }

    // claims 추출 — 타입별 assertion
    id, _ := claims["user_id"].(float64)      // JWT 는 숫자를 float64 로
    email, _ := claims["email"].(string)
    role, _ := claims["role"].(string)
    orgID, _ := claims["org_id"].(float64)

    return VerifyTokenResponse{
        ID:    int64(id),
        Email: email,
        Role:  role,
        OrgID: int64(orgID),
    }, nil
}
```

타입 assertion 규칙:
- `int64` → `claims[key].(float64)` → `int64(v)` (JWT 숫자는 float64)
- `string` → `claims[key].(string)`
- `bool` → `claims[key].(bool)`

### RefreshToken (`internal/auth/refresh_token.go`)

IssueToken 과 동일 구조, 만료 7일, Response 필드 `RefreshToken`.

### Reexport (`internal/auth/reexport.go`)

```go
package auth

import pkgauth "github.com/park-jun-woo/ssac/pkg/auth"

var HashPassword = pkgauth.HashPassword
var VerifyPassword = pkgauth.VerifyPassword
var GenerateResetToken = pkgauth.GenerateResetToken

type HashPasswordRequest = pkgauth.HashPasswordRequest
type HashPasswordResponse = pkgauth.HashPasswordResponse
type VerifyPasswordRequest = pkgauth.VerifyPasswordRequest
type VerifyPasswordResponse = pkgauth.VerifyPasswordResponse
type GenerateResetTokenRequest = pkgauth.GenerateResetTokenRequest
type GenerateResetTokenResponse = pkgauth.GenerateResetTokenResponse
```

SSaC `@call auth.HashPassword(...)` → `import "<module>/internal/auth"` 하나로
JWT + bcrypt 전부 접근. import 분산 방지.

### BearerAuth Middleware (`internal/middleware/bearerauth.go`)

```go
package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"

    "<module>/internal/auth"
    "<module>/internal/model"
)

func BearerAuth(secret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        header := c.GetHeader("Authorization")
        if !strings.HasPrefix(header, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            return
        }
        token := strings.TrimPrefix(header, "Bearer ")
        out, err := auth.VerifyToken(auth.VerifyTokenRequest{Token: token, Secret: secret})
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
            return
        }
        c.Set("currentUser", &model.CurrentUser{
            ID:    out.ID,
            Email: out.Email,
            Role:  out.Role,
            OrgID: out.OrgID,
        })
        c.Next()
    }
}
```

`c.Set("currentUser", ...)` → handler 에서 `c.MustGet("currentUser").(*model.CurrentUser)`.

## 생성 흐름

```
auth.Generate(fs, artifactsDir)
  ├─ parseClaims(manifest)        → []ClaimField{Name, Key, GoType}
  ├─ generateCurrentUser(fields)  → internal/model/current_user.go
  ├─ generateIssueToken(fields)   → internal/auth/issue_token.go
  ├─ generateVerifyToken(fields)  → internal/auth/verify_token.go
  ├─ generateRefreshToken(fields) → internal/auth/refresh_token.go
  ├─ generateReexport()           → internal/auth/reexport.go
  └─ generateBearerAuth(fields)   → internal/middleware/bearerauth.go
```

모든 하위 함수가 **같은 `[]ClaimField`** 를 입력으로 받음 — claims 가 유일한 변수.

## ClaimField 구조

```go
type ClaimField struct {
    Name   string // "ID", "Email", "Role", "OrgID" (Go struct 필드명)
    Key    string // "user_id", "email", "role", "org_id" (JWT claim key)
    GoType string // "int64", "string", "bool"
}
```

`manifest.backend.auth.claims` 를 파싱해서 생성. 파싱 로직은 이미
`pkg/parser/manifest/` 에 `ClaimDef{Key, GoType}` 로 존재 — 재사용.

## 외부 의존

| 라이브러리 | 용도 | import |
|---|---|---|
| `golang-jwt/jwt/v5` | JWT sign/parse | `internal/auth/` |
| `ssac/pkg/auth` | bcrypt (HashPassword, VerifyPassword) | `internal/auth/reexport.go` |
| `gin-gonic/gin` | middleware | `internal/middleware/` |

## 파일 구조 (예정)

```
pkg/generate/gogin/auth/
├── README.md
├── generate.go                  ← orchestrator
├── claim_field.go               ← ClaimField type
├── parse_claims.go              ← manifest claims → []ClaimField
├── generate_current_user.go     ← model/current_user.go 생성
├── generate_issue_token.go      ← auth/issue_token.go 생성
├── generate_verify_token.go     ← auth/verify_token.go 생성
├── generate_refresh_token.go    ← auth/refresh_token.go 생성
├── generate_reexport.go         ← auth/reexport.go 생성
├── generate_bearer_auth.go      ← middleware/bearerauth.go 생성
└── claim_extract_line.go        ← JWT float64 → Go type assertion 코드
```

## gogin README 반영

gogin 파이프라인 [4][5][6] → **[4] auth.Generate** 로 통합:

```
[1]  oapi-codegen
[2]  postgresql.Generate
[3]  sqlc generate
[4]  auth.Generate          ← JWT + middleware + CurrentUser 통합
[7]  statemachine.Generate
[8]  copyFuncSpecs
[9]  boot.Generate
[10] handler.Generate
[11] generateGoMod
```
