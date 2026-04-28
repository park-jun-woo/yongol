# pkg/parser/tsx

## 변경이력

- 2026-04-28: 초기 작성

## 역할

React `.tsx` 파일을 외부 swc (`@swc/core`) 바이너리로 파싱해 AST JSON 을 받아오고, `apiClient.<op>(...)` 호출 / `register('field', { required })` 폼 필드 / 로컬 component import 를 추출한다. 결과는 페이지별 `PageSpec` 으로 묶여 XOT-* (TSX ↔ OpenAPI) 교차 검증에 공급된다.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Parse` | `(file string) (PageSpec, error)` | 단일 `.tsx` 파일 → PageSpec |
| `ParseDir` | `(root string) ([]PageSpec, error)` | 디렉토리 재귀 순회. `node_modules / dist / build / .*` 스킵 |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `PageSpec` | 단일 `.tsx` 파일 추출 결과 (`File / Calls / FormFields / ComponentImports`) |
| `APICall` | `apiClient.<op>(...)` 호출 1건 (operationId + 인자 바인딩 + 위치) |
| `ArgBinding` | apiClient 호출 인자 객체의 key/value 한 쌍 |
| `FormField` | `register('name', opts)` 선언 1건 (Required 포함) |
| `ComponentImport` | 로컬 component import 1건 (`Name / Path / Line`) |

## 비고

- swc 실행 디렉토리는 `$YONGOL_NODE` env override → `node_modules` 조상 → `package.json` 조상 순으로 결정.
- 본 패키지는 `(result, error)` 반환 (다른 파서들의 `[]Diagnostic` 시그니처와 다름) — 외부 spawn 의 fail-fast 특성 때문.
- npm 패키지 import 는 무시하고 `@/components/`, `./components/`, sibling 경로만 수집.
