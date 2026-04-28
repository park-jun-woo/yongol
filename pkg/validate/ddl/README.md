# pkg/validate/ddl

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

DDL (PostgreSQL `CREATE TABLE` / sqlc 쿼리 / `db/sqlc.yaml`) 자체 정합성 검증. 중복·NOT NULL 누락·SERIAL 금지·sentinel 규약·민감 패턴 컬럼 어노테이션.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| D-1 | `d01SqlcQueryDuplicate` | sqlc query name 중복 (ERROR) | IF-ELSE | ✓ |
| D-2 | `d02NullableColumn` | NOT NULL 누락 (ERROR, PK / `-- @nullable` 면제) | IF-ELSE | ✓ |
| D-3 | `d03SentinelMissing` | FK DEFAULT 0 sentinel 레코드 누락 (ERROR) | IF-ELSE | ✓ |
| D-4 | `d04SqlcYamlRequired` | `db/sqlc.yaml` 미존재 (ERROR) | IF-ELSE | ✓ |
| D-5 | `d05SqlcYamlSchemaPath` | `sqlc.yaml` schema 경로가 DDL 디렉토리 미포함 (WARNING) | IF-ELSE | ✓ |
| D-6 | `d06SqlcYamlQueriesPath` | `sqlc.yaml` queries 경로가 `queries/` 미포함 (WARNING) | IF-ELSE | ✓ |
| D-7 | `d07SqlcPositionalParam` | sqlc 쿼리에 위치 파라미터 (`$1`, `$2`) 전면 금지 (ERROR) | IF-ELSE | ✓ |
| D-8 | `d08SerialTypeBanned` | `SERIAL`/`BIGSERIAL`/`SMALLSERIAL` 금지, `GENERATED ALWAYS AS IDENTITY` 권고 (ERROR) | IF-ELSE | ✓ |
| D-9 | `d09TopLevelInsertWithoutSentinel` | top-level `INSERT` 에 `-- @sentinel` 어노테이션 부재 (ERROR) | IF-ELSE | ✓ |
| D-10 | `d10SentinelWithoutOnConflict` | `@sentinel` `INSERT` 에 `ON CONFLICT DO NOTHING` 부재 (ERROR) | IF-ELSE | ✓ |
| XDD-61 | `xdd61SensitiveNoAnnotation` | 민감 패턴 컬럼명 (password/secret/hash/token 등) `@sensitive` 없음 (WARNING) | TOULMIN | ✓ |

D-7 위치별 문법: `WHERE/SET/VALUES` 는 `@name`, `LIMIT/OFFSET` 은 `sqlc.arg(name)` 사용. XQS-16/17 이 named param 이름으로 SSaC Input key 대조.

## Defeater

| 이름 | 면제 warrant | 설명 |
|---|---|---|
| `IsNoSensitive` | XDD-61 | `-- @nosensitive` 컬럼은 패턴 매칭 면제 |

## internal 일치성 메모

- XDD-61: `-- @sensitive` / `-- @nosensitive` 어노테이션 보유 컬럼 스킵.
