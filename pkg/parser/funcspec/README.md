# pkg/parser/funcspec

## 변경이력

- 2026-04-28: 초기 작성

## 역할

Go AST 기반으로 사용자 정의 Func 스펙 (`@func` 어노테이션이 달린 `*.go` 파일) 을 파싱해 함수명 / Request·Response 구조체 / Imports / 본문 보유 여부 (HasBody) 를 수집한다. 한 파일에 `@func` 가 둘 이상이면 ERROR. Request/Response 가 별도 파일에 있어도 패키지 단위로 병합한다.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `ParseDir` | `(dir string) ([]FuncSpec, []diagnostic.Diagnostic)` | 디렉토리 내 모든 `.go` 파일을 재귀 파싱해 FuncSpec 슬라이스 반환 |
| `ParseFile` | `(path string) (*FuncSpec, []diagnostic.Diagnostic)` | 단일 `.go` 파일에서 `@func` 어노테이션 1건 추출 |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `FuncSpec` | 파싱된 func spec 1건 (`Package / Name / Description / Error / Imports / RequestFields / ResponseFields / HasBody / ReturnTypes / Line`) |
| `Field` | struct 필드 (`Name / Type / JSONName`) |

## 비고

- Go 표준 라이브러리 `go/parser`, `go/ast` 를 직접 사용.
- HasBody 판정: 빈 본문 / `panic("TODO")` / 모든 return 이 zero value 면 false.
- `@func` / `@error` / `@description` 만 어노테이션으로 인정.
