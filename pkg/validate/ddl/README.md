# pkg/validate/ddl

DDL (PostgreSQL CREATE TABLE / sqlc 쿼리) 자체 정합성 검증.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule 공통 함수 + defeater 그래프 / `IF-ELSE` = 단일 구조·흐름·휴리스틱 검사

## 고유 함수 (구조 검증)

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---------|--------|------|----------|---------|
| D-1 | `SqlcQueryDuplicate` | sqlc query name 중복 | IF-ELSE | ✓ |
| D-2 | `NullableColumn` | NOT NULL 누락 | IF-ELSE | ✓ |
| D-3 | `SentinelMissing` | FK DEFAULT 0 센티널 레코드 누락 | IF-ELSE | ✓ |
| D-4 | `SqlcYamlRequired` | `db/sqlc.yaml` 미존재 시 ERROR | IF-ELSE | ⏳ |
| D-5 | `SqlcYamlSchemaPath` | `sqlc.yaml` 의 schema 경로가 DDL 디렉토리를 포함하지 않으면 WARNING | IF-ELSE | ⏳ |
| D-6 | `SqlcYamlQueriesPath` | `sqlc.yaml` 의 queries 경로가 `queries/` 를 포함하지 않으면 WARNING | IF-ELSE | ⏳ |
| D-8 | `FKFileOrder` | FK 참조 대상 테이블의 DDL 파일이 알파벳순으로 현재 파일보다 뒤에 위치하면 ERROR. `fs.DDLTables`의 `table.ForeignKeys`를 사용해 FK 의존관계를 분석하고, 올바른 파일 순서를 권고 메시지에 포함 (예: `001_organizations.sql`, `002_users.sql`, `003_workflows.sql`). | IF-ELSE | ⏳ |

## 고유 함수 (sqlc 쿼리 제약)

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---------|--------|------|----------|---------|
| D-7 | `SqlcPositionalParam` | sqlc 쿼리에 위치 파라미터($1, $2 등) 전면 금지 (ERROR) | IF-ELSE | ✓ |

`$N` 위치 파라미터 전면 금지. 위치별 문법:
- WHERE/SET/VALUES: `@name` 사용
- LIMIT/OFFSET: `sqlc.arg(name)` 사용 (`@name` shorthand는 sqlc 제약으로 불가)

XQS-16/17이 named param 이름으로 SSaC Input key 대조.

## 고유 함수 (민감 패턴 — 구 pkg/crosscheck/ddl)

| 규칙 ID | 함수명 | 설명 | 구현 방식 | 조건 / 예외 |
|---------|--------|------|----------|------|
| XDD-61 | `SensitiveNoAnnotation` | 민감 패턴 컬럼 @sensitive 없음 (WARNING) | TOULMIN | `IsNoSensitive` defeater로 `-- @nosensitive` 컬럼 면제. 패턴 매칭은 password/secret/hash/token 등 |

## pkg/rule 사용

없음.

## Defeater

- `IsNoSensitive` — `-- @nosensitive` 컬럼 (패턴 매칭 면제)

## internal 일치성 메모

- XDD-61: `-- @sensitive` / `-- @nosensitive` 어노테이션 보유 컬럼 스킵 — `check_table_sensitive_columns.go:10-15`
