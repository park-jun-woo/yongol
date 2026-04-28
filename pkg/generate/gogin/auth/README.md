# pkg/generate/gogin/auth

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

JWT 인증 + bearerAuth gin 미들웨어 + CurrentUser 모델을 한 번에 생성. 셋 모두 `manifest.backend.auth.claims` 에서 파생되며 `manifest.backend.auth == nil || len(claims) == 0` 일 때 전체 skip.

> 상위: [`pkg/generate/gogin/README.md`](../README.md) 의 [4]+[5]+[6] 통합.

## 공개 함수

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Generate` | `(fs *yongol.Fullstack, artifactsDir string) error` | 진입점. parseClaims → user_claim/issue/verify/refresh/reexport/bearer_auth 순으로 emit |
| `parseClaims` | `(manifest) []ClaimField` | manifest claims → `ClaimField{Name, Key, GoType}` 슬라이스 |
| `generateUserClaim` | `(fields, artifactsDir) error` | `internal/model/current_user.go` |
| `generateBearerAuth` | `(fields, artifactsDir) error` | `internal/middleware/bearerauth.go` (template 기반) |
| `cleanStaleAuthFiles` | `(artifactsDir) error` | 더이상 필요 없는 auth 산출물 제거 |

## 공개 타입

| 타입 | 설명 |
|---|---|
| `ClaimField` | `{Name string; Key string; GoType string}` — manifest claims 한 줄 (`ID: user_id:int64` → `{ID, user_id, int64}`) |

## 산출물 (5 파일)

```
arts/backend/internal/
├── auth/
│   ├── issue_token.go       JWT 발급 (access, 24h)
│   ├── verify_token.go      JWT 검증 → claims 추출 (float64 → int64 등 type assert)
│   ├── refresh_token.go     JWT 갱신 (refresh, 7d)
│   └── reexport.go          ssac/pkg/auth 재export (HashPassword/VerifyPassword/GenerateResetToken …)
├── middleware/bearerauth.go bearer/cookie/hybrid 모드 지원, c.Set("currentUser", ...)
└── model/current_user.go    CurrentUser struct (claims 필드, manifest 선언 순)
```

생성 코드 형식은 `template_bearer_auth.go` 및 `generate_*.go` 소스 참조.

## 입력

| 소스 | 데이터 |
|---|---|
| `manifest.backend.auth.claims` | 필드명 + claim key + Go 타입 (예: `ID: user_id:int64`) |
| `manifest.backend.auth.secret_env` | JWT secret 환경변수명 (기본 `JWT_SECRET`) |
| `manifest.backend.auth.cookie` | cookie/hybrid 모드 설정 |
| `manifest.backend.middleware` | `bearerAuth` 포함 여부 |
| `manifest.backend.module` | import 경로 (`<module>/internal/{auth,model}`) |

## 외부 의존

| 라이브러리 | 용도 | import 위치 |
|---|---|---|
| `golang-jwt/jwt/v5` | JWT sign/parse | `internal/auth/` |
| `ssac/pkg/auth` | bcrypt 등 | `internal/auth/reexport.go` |
| `gin-gonic/gin` | middleware | `internal/middleware/` |

## 비고

SSaC `@call auth.HashPassword(...)` 가 `import "<module>/internal/auth"` 하나로 JWT + bcrypt 모두 접근하도록 `reexport.go` 가 ssac runtime 패키지를 alias 한다.
