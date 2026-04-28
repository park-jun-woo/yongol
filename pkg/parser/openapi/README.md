# pkg/parser/openapi

## 변경이력

- 2026-04-28: 초기 작성

## 역할

이미 `kin-openapi` 로 로드된 `*openapi3.T` 문서에서 operationId 단위로 요청/응답 필드의 type / format / length / enum / required 제약을 추출한다. 별도로 yaml.v3 raw 파싱으로 `LineIndex` 를 만들어 각 필드 / operationId / path / schema property 의 줄 번호를 색인한다.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `ExtractRequestConstraints` | `(doc *openapi3.T, idx *LineIndex) map[op]map[field]FieldConstraint` | operationId 별 requestBody 제약 추출 |
| `ExtractResponseConstraints` | `(doc *openapi3.T, idx *LineIndex) map[op]map[field]FieldConstraint` | operationId 별 첫 2xx 응답 제약 추출 |
| `BuildLineIndex` | `(path string) (*LineIndex, error)` | openapi.yaml 을 yaml.v3 로 재파싱해 LineIndex 구축 |
| `DeriveSuccessStatus` | `(op *openapi3.Operation, method string) int` | HTTP method 관례에 따른 성공 2xx 응답 코드 선택 |
| `Declared2xx` | `(op *openapi3.Operation) []string` | operation 의 2xx 응답 코드 집합 |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `FieldConstraint` | 단일 OpenAPI schema property 제약 (`Type / Format / MaxLength / MinLength / Enum / Required / Line`) |
| `LineIndex` | OpenAPI 노드별 줄 번호 색인 (`OperationLine / RequestFieldLine / ResponseFieldLine / SchemaLine / SchemaPropertyLine / PathLine`) |

## 비고

- doc 이 이미 파싱된 상태이므로 본 패키지 자체는 에러를 발생시키지 않는다 (`(result, []Diagnostic)` 시그니처를 따르지 않는 예외).
- `LineIndex` 의 모든 lookup 메서드는 nil receiver 안전.
