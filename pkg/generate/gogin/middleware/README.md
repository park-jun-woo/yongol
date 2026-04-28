# pkg/generate/gogin/middleware

## 변경이력

- 2026-04-28: 초기 작성

## 역할

gogin 백엔드의 `internal/middleware/*.go` 를 emit 한다. request_id, body_limit, error_envelope, security_headers, csrf, prometheus, rate_limit, request_validator 등 manifest 설정에 따라 활성/비활성 분기.

## 진입점 (Generator)

| 함수 | 시그니처 | 산출 |
|---|---|---|
| `Generate` | `(fs *yongol.Fullstack, p prepared.State, artifactsDir string) error` | `request_validator.go` (+ `openapi.yaml` 복사) |
| `GenerateRequestID` | `(artifactsDir string) error` | `request_id.go` |
| `GenerateBodyLimit` | `(artifactsDir string) error` | `body_limit.go` |
| `GenerateErrorEnvelope` | `(artifactsDir string) error` | `error_envelope.go` |
| `GenerateSecurityHeaders` | `(fs *yongol.Fullstack, artifactsDir string) error` | `security_headers.go` |
| `GenerateCsrf` | `(a prepared.Auth, artifactsDir string) error` | `csrf.go` (cookie/hybrid 모드 시) |
| `GeneratePrometheus` | `(fs *yongol.Fullstack, artifactsDir string) error` | `prometheus.go` |
| `GenerateRateLimit` | `(fs *yongol.Fullstack, artifactsDir string) error` | `rate_limit.go` |

## 진입점 (Util)

| 함수 | 설명 |
|---|---|
| `ParseSize` | `"1MiB"` / `"32MiB"` 같은 human-readable size → bytes |

## 비고

각 `*_source.go` 는 산출 대상 Go 코드의 템플릿 상수. 일부 (request_id / error_envelope / security_headers / prometheus / csrf / body_limit) 는 정적 소스, rate_limit / prometheus 는 manifest 값 (`__MODULE__`, `__BUCKETS__`) 토큰 치환.
