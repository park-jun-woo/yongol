# pkg/chain

## 변경이력

- 2026-04-28: 초기 작성

## 역할

`yongol chain <operationId>` CLI 명령의 백엔드. 단일 operationId 를 기준으로 OpenAPI / SSaC / DDL / Rego / StateDiagram / FuncSpec / Hurl 등 SSOT 노드를 추적해 한 줄씩 출력 가능한 `Link` 슬라이스를 만든다.

## 공개 함수

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Chain` | `Chain(fs *yongol.Fullstack, operationID string) ([]Link, error)` | operationId 에 연결된 모든 SSOT 노드를 모아 Link 슬라이스로 반환. OpenAPIDoc 미파싱 / operationId 미존재는 에러 |
| `Print` | `Print(w io.Writer, operationID string, links []Link)` | Link 슬라이스를 헤더 + SSOT 섹션 + Artifacts 섹션으로 포맷 출력 |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `Link` | 체인의 단일 노드. `Kind` (OpenAPI/SSaC/DDL/Rego/StateDiag/FuncSpec/Hurl/Handler/Model/Authz/Types), `File` (specs 상대 경로), `Line`, `Summary`, `Ownership` ("/gen/preserve") |

## 비고

내부 trace 함수 (`traceOpenAPI`, `traceSSaC`, `traceDDL`, `tracePolicy`, `traceStates`, `traceFuncSpecs`, `traceHurl`) 은 unexported. 호출 진입점은 `Chain` 단일.
