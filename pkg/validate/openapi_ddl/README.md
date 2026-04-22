# pkg/validate/openapi_ddl

OpenAPI ↔ DDL 교차 검증.

> 규칙 전체 목록은 저장소 루트의 [`rulebook.md`](../../../rulebook.md) 참조.
> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 실 작동 또는 반례 확장 가능 / `IF-ELSE` = 단일 판정·Ground 조회

## RefExists (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XDO-9 | `DDL.column.<table>` | OpenAPI property → DDL column (ghost) | IF-ELSE |

## CoverageCheck

| 규칙 ID | LookupKey | 설명 | 구현 방식 | 예외 |
|---------|-----------|------|----------|------|
| XOD-10 | `OpenAPI.response.<op>` | DDL column → OpenAPI schema 포함 여부 | TOULMIN | `IsSensitiveCol` (`-- @sensitive`) defeater 실 작동 + `IsNoSensitive` 반례 |

## TypeMatch (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XDO-69 | `DDL.check.<table>` | DDL CHECK values ↔ OpenAPI enum | IF-ELSE |

## SchemaMatch (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XDO-67 | `DDL.varchar.<table>` | DDL VARCHAR(n) → OpenAPI maxLength | IF-ELSE |
| XDO-68 | `DDL.check.<table>` | DDL CHECK IN → OpenAPI enum | IF-ELSE |
| XDO-77 | `DDL.column.<table>` | DDL column 타입 ↔ OpenAPI field 타입 불일치 (ERROR) | IF-ELSE |

### XDO-77 타입 대조 테이블

| DDL 타입 | OpenAPI type | OpenAPI format |
|---|---|---|
| `BIGINT`, `BIGSERIAL` | `integer` | `int64` |
| `INTEGER`, `SERIAL`, `INT` | `integer` | (없거나 `int32`) |
| `SMALLINT` | `integer` | — |
| `TEXT`, `VARCHAR` | `string` | — |
| `BOOLEAN` | `boolean` | — |
| `TIMESTAMP`, `TIMESTAMPTZ` | `string` | `date-time` |
| `NUMERIC`, `DECIMAL`, `REAL`, `DOUBLE PRECISION` | `number` | — |

DDL `BIGINT`인데 OpenAPI `integer` (format 없음)이면 ERROR.
권고: OpenAPI에 `format: int64`를 추가하세요.

## 고유 함수 (IF-ELSE)

| 규칙 ID | 함수명 | 설명 | 구현 방식 |
|---------|--------|------|----------|
| XDO-70 | `MaxLengthExceedsVarchar` | OpenAPI maxLength > DDL VARCHAR(n) (WARNING) | IF-ELSE |

## NullabilityMatch (IF-ELSE)

OpenAPI optional 필드와 DDL NOT NULL 제약의 정합성 검증.

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---------|--------|------|----------|---------|
| XDO-75 | `OptionalNotNullNoDefault` | OpenAPI optional + DDL NOT NULL + DEFAULT 없음 (ERROR) | IF-ELSE | **누락** |
| XDO-76 | `RequiredNullable` | OpenAPI required + DDL nullable (WARNING) | IF-ELSE | **누락** |

XDO-76 면제: DDL 컬럼에 `-- @nullable` 어노테이션이 있으면 의도적 설계로 간주하여 WARNING 면제.

OpenAPI request field가 optional인데, 대응하는 DDL 컬럼이 NOT NULL이고,
DDL DEFAULT도 없으면 INSERT 시 실패한다. SSaC에서 기본값 삽입 로직이 있으면 면제.

| OpenAPI | DDL | DEFAULT / SSaC 기본값 | 판정 |
|---|---|---|---|
| optional | NOT NULL | DEFAULT 있음 | OK |
| optional | NOT NULL | SSaC 기본값 삽입 | OK |
| optional | NOT NULL | 둘 다 없음 | **ERROR** |
| optional | nullable | — | OK |
| required | NOT NULL | — | OK |
| required | nullable | — | **WARNING** — 의도적이지 않을 가능성 |

SSaC 기본값 삽입 판정: `@post` Input에서 해당 필드가 리터럴(`"draft"`, `0` 등)으로 지정되면 면제.

## Defeater

- `IsSensitiveCol` — XOD-10 에서 `-- @sensitive` 컬럼 스킵
- `IsNoSensitive` — sensitive 패턴 검사 면제용
- `IsNullableIntentional` — `-- @nullable` 어노테이션 (XDO-76 면제)
