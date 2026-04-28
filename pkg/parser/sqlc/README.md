# pkg/parser/sqlc

## 변경이력

- 2026-04-28: 초기 작성

## 역할

sqlc 쿼리 파일 (`queries/*.sql`) 에서 `-- name: <Method> :<cardinality>` 매크로를 스캔해 `QuerySpec` 목록을 추출한다. 파일명에서 모델명 (단수형 PascalCase) 을 도출하고, SQL 본문의 `@name` 및 `sqlc.arg(name)` named parameter 를 수집한다. `$N` 위치 파라미터는 후속 검증 단계에서 D-7 ERROR 로 차단.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `ParseDir` | `(dir string) ([]QuerySpec, []diagnostic.Diagnostic)` | 디렉토리 내 모든 `*.sql` 을 sqlc 쿼리로 파싱 (디렉토리 부재는 silent OK) |
| `ParseFile` | `(path string) ([]QuerySpec, []diagnostic.Diagnostic)` | 단일 `.sql` 파일에서 QuerySpec 목록 추출 (named params 포함) |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `QuerySpec` | sqlc 쿼리 1건 (`File / Model / Method / Cardinality / RowType / Params / Line`). 파일명 → 모델, name 주석 → 메서드 |

## 비고

- 본 패키지는 sqlc 자체를 호출하지 않고 `-- name:` 헤더 주석만 정규식으로 스캔한다.
- 파일명 → 모델명 변환은 `modelFromFilename` (예: `users.sql` → `User`).
