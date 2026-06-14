# pkg/validate/openapi_ddl

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

OpenAPI ↔ DDL 교차 검증. property/column 매칭, 길이·enum·CHECK 제약, 타입·Nullability 정합.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XDO-9 | `xdo09GhostProperty` | OpenAPI property → 대응 DDL 컬럼 부재 (ERROR, ghost) | IF-ELSE | ✓ |
| XOD-10 | `xod10DDLToResponse` | DDL 컬럼이 OpenAPI components.schemas 에 노출 안 됨 (WARNING) | TOULMIN | ✓ |
| XDO-67 | `xdo67MaxLengthVarchar` | DDL VARCHAR(n) 컬럼의 OpenAPI 요청 필드에 `maxLength` 없음 (ERROR) | IF-ELSE | ✓ |
| XDO-68 | `xdo68CheckInEnum` | DDL `CHECK IN(...)` 컬럼에 OpenAPI enum 없음 (ERROR) | IF-ELSE | ✓ |
| XDO-69 | `xdo69CheckValuesEnum` | DDL CHECK IN 값 ↔ OpenAPI enum 값 셋 불일치 (ERROR) | IF-ELSE | ✓ |
| XDO-70 | `xdo70MaxLengthExceedsVarchar` | OpenAPI `maxLength` > DDL `VARCHAR(n)` (WARNING) | IF-ELSE | ✓ |
| XDO-75 | `xdo75OptionalNotNullNoDefault` | OpenAPI optional + DDL NOT NULL + DEFAULT 없음 (ERROR) | IF-ELSE | ✓ |
| XDO-76 | `xdo76RequiredNullable` | OpenAPI required + DDL nullable (WARNING, `-- @nullable` 면제) | IF-ELSE | ✓ |
| XDO-77 | `xdo77ColumnTypeMismatch` | DDL column 타입 ↔ OpenAPI field 타입 불일치 (ERROR) | IF-ELSE | ✓ |
| XDO-78 | `xdo78EnumNoCheck` | OpenAPI enum 요청 필드에 대응 DDL CHECK IN 제약 없음 (ERROR) | IF-ELSE | ✓ |
| XDO-11 | `canonicalResponseRepr` | 같은 엔티티 2xx 응답이 서로 다른 표현 노출 (ERROR, 리소스 1개 = 표현 1개) | IF-ELSE | ✓ |
| XDO-12 | `canonicalResponseRepr` | 엔티티 응답을 `$ref` 공유 없이 inline 정의 (WARNING, drift 위험) | IF-ELSE | ✓ |

## XDO-77 타입 대조

| DDL 타입 | OpenAPI type | format |
|---|---|---|
| `BIGINT`, `BIGSERIAL` | `integer` | `int64` |
| `INTEGER`, `SERIAL`, `INT` | `integer` | (없거나 `int32`) |
| `SMALLINT` | `integer` | — |
| `TEXT`, `VARCHAR` | `string` | — |
| `BOOLEAN` | `boolean` | — |
| `TIMESTAMP`, `TIMESTAMPTZ` | `string` | `date-time` |
| `NUMERIC`, `DECIMAL`, `REAL`, `DOUBLE PRECISION` | `number` | — |

## XDO-75 판정표

| OpenAPI | DDL | DEFAULT / SSaC 기본값 | 판정 |
|---|---|---|---|
| optional | NOT NULL | DEFAULT 있음 | OK |
| optional | NOT NULL | SSaC 리터럴 기본값 | OK |
| optional | NOT NULL | 둘 다 없음 | ERROR |
| optional | nullable | — | OK |
| required | NOT NULL | — | OK |
| required | nullable | — | WARNING (XDO-76) |

## Defeater

| 이름 | 면제 warrant | 설명 |
|---|---|---|
| `IsSensitiveCol` | XOD-10 | DDL `-- @sensitive` 컬럼 스킵 |
| `IsNoSensitive` | XOD-10 (반례) | `-- @nosensitive` 으로 sensitive 패턴 면제 |
| `IsNullableIntentional` | XDO-76 | DDL `-- @nullable` 어노테이션 |
