# pkg/util/caseconv

## 변경이력

- 2026-04-28: 초기 작성

## 역할

문자열 case 변환 유틸. snake_case ↔ PascalCase / camelCase / kebab-case 사이 변환과 sqlc 네이밍 규칙(`id`/`ids` 전체 대문자) 변형을 제공한다. 코드젠 / 검증 양쪽에서 식별자 정규화에 사용.

## 공개 함수

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `SnakeToPascal` | `SnakeToPascal(s string) string` | snake_case → PascalCase. plain (capitalize-first per part). `"user_id"` → `"UserId"` |
| `SnakeToPascalSqlc` | `SnakeToPascalSqlc(s string) string` | snake_case → PascalCase. sqlc 규칙: `"id"` / `"ids"` 부분은 전체 대문자. `"org_id"` → `"OrgID"` |
| `PascalToSnake` | `PascalToSnake(s string) string` | PascalCase / camelCase → snake_case. `ettle/strcase` 라이브러리 경유 |
| `KebabToCamel` | `KebabToCamel(s string) string` | kebab-case → camelCase. dash 가 없으면 입력 그대로 |

## 의존성

- `github.com/ettle/strcase` (`PascalToSnake` 한정)
