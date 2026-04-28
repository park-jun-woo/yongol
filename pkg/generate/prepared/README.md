# pkg/generate/prepared

## 변경이력

- 2026-04-28: 초기 작성

## 역할

generate 진입점에서 한 번 계산되는 불변 파생 상태. manifest + SSaC 사용 여부를 합쳐 auth/session/cache/file/queue/middleware/route 의 활성 여부 + 기본값을 미리 해석해 하위 generator 들이 동일한 결정을 공유하게 한다.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `New` | `(fs *yongol.Fullstack) State` | Fullstack → prepared.State 변환 (infallible) |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `State` | 모든 파생 상태 모음 (Auth / ActiveBackends / Routes / Middlewares …) |
| `Auth` | 인증 설정 (Mode 기본값 해석 + CsrfRequired) |
| `ActiveBackends` | session / cache / file / queue 활성 여부 (nil = 비활성) |
| `Session` / `Cache` / `File` / `Queue` | 백엔드별 파생 상태 (non-zero 보장) |
| `Route` / `Middleware` | placeholder (Stage 5 확장) |

## 활성 판정 규칙

- session/cache/file: `manifest.<X>.backend` 선언 또는 SSaC 가 `@call <x>.*` 사용 → 활성
- queue: manifest 선언 또는 SSaC `@subscribe` / `publish` 사용
- auth: `manifest.backend.auth` 존재 시 Mode 기본값 cookie, type=jwt → bearer 매핑. CsrfRequired = mode ∈ {cookie, hybrid}
