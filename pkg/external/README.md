# pkg/external

## 변경이력

- 2026-04-28: 초기 작성

## 역할

`yongol import <openapi-source> <output-dir>` CLI 명령의 백엔드. 외부 OpenAPI 문서 (Stripe, GitHub 등) 를 파일 경로 또는 HTTP(S) URL 로 받아 Go 클라이언트 패키지 (인터페이스 / 메서드 / 응답 구조체 / HTTP 헬퍼) 코드를 산출한다.

## 공개 함수

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Generate` | `Generate(source, outputDir string) error` | OpenAPI 소스를 읽어 Go 클라이언트 코드를 생성하고 `<outputDir>/<service>.go` 로 저장 |
| `GenerateBytes` | `GenerateBytes(source string, data []byte) ([]byte, error)` | 바이트로 받은 OpenAPI 데이터에서 코드를 생성해 바이트로 반환 (테스트용) |

## 비고

내부 단계는 모두 unexported. 흐름:

1. `readSource` — 파일 경로 / HTTP(S) URL 자동 분기. HTTP 4xx/5xx 는 body snippet 1KiB truncate 포함 에러
2. `kin-openapi` 로 OpenAPI 파싱 + `doc.Validate` 실행
3. `inferServiceName` — OpenAPI title 또는 파일명에서 서비스명 추론
4. `generateCode` → `extractMethods` → `buildMethodInfo` → `schemaToGoType` 로 Go 타입 / 메서드 시그니처 / 구현 / 응답 구조체 합성

생성된 클라이언트는 SSaC 에서 `@call <pkg>.<Func>({...})` 로 호출한다.
