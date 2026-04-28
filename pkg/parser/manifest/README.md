# pkg/parser/manifest

## 변경이력

- 2026-04-28: 초기 작성

## 역할

`manifest.yaml` 한 파일을 yaml.v3 로 읽어 `ProjectConfig` 구조체로 변환한다. 1차 패스는 일반 unmarshal, 2차 패스는 yaml.Node 트리를 직접 순회해 backend.auth.claims / roles / user_table 등 주요 키의 줄 번호를 후속 진단용으로 채운다.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Load` | `(specsDir string) (*ProjectConfig, []diagnostic.Diagnostic)` | `<specsDir>/manifest.yaml` 로드 + 파싱 + claim/role 줄 번호 채움 |
| `FindAuthNode` | `(data []byte) *yaml.Node` | manifest.yaml 원본에서 `backend.auth` yaml.Node 반환 (외부 검증기 보조) |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `ProjectConfig` | manifest.yaml 최상위 (Metadata / Backend / Frontend / Deploy 등) |
| `Metadata` | 프로젝트 메타데이터 |
| `Backend` | 백엔드 섹션 |
| `Frontend`, `FrontendTheme` | 프런트엔드 섹션 + shadcn 테마 |
| `Deploy` | 배포 설정 |
| `Auth`, `ClaimDef`, `CookieConfig`, `CsrfConfig` | JWT 인증 / 클레임 / 쿠키 / CSRF 설정 |
| `AuthzConfig` | 인가 패키지 설정 |
| `CORSConfig`, `SecurityHeadersConfig`, `CSPConfig`, `HSTSConfig` | 웹 보안 헤더 설정 |
| `HTTPConfig`, `HTTPOverride` | backend.http (limit + 오버라이드 맵) |
| `ErrorConfig`, `RequestIDConfig` | error envelope + request_id 정책 |
| `Observability`, `ObservabilityMetrics`, `ObservabilityTracing` | 메트릭 / 트레이싱 설정 |
| `BuiltinBackend`, `FileBackend`, `LocalConfig`, `S3Config`, `QueueBackend` | 세션/캐시/파일/큐 백엔드 설정 |

## 비고

- 모든 YAML 파싱은 표준 라이브러리 `gopkg.in/yaml.v3` 에 위임.
- 줄 번호 보강은 검증 단계에서 `Diagnostic.Line` 정확도를 위한 것.
