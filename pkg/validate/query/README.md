# pkg/validate/query

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

sqlc 쿼리 + `db/sqlc.yaml` 자체 검증 (Q-*). DDL/OpenAPI/Rego 와의 교차 검사는 `query_rego/` 등 페어 폴더에 위치.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| Q-1 | `q01NameRequired` | `-- name:` 어노테이션 필수 (ERROR) | IF-ELSE | ✓ |
| Q-2 | `q02Cardinality` | cardinality (`:one`/`:many`/`:exec`/`:execrows`) 필수 (ERROR) | IF-ELSE | ✓ |
| Q-3 | `q03NamePascalCase` | query name PascalCase 강제 (ERROR) | IF-ELSE | ✓ |
| Q-4 | `q04ManyLimit` | `:many` 쿼리에 `LIMIT` 누락 (WARNING) | IF-ELSE | ✓ |
| Q-5 | `q05DeleteWhere` | `DELETE` 문에 `WHERE` 필수 (ERROR) | IF-ELSE | ✓ |
| Q-6 | `q06UpdateWhere` | `UPDATE` 문에 `WHERE` 필수 (ERROR) | IF-ELSE | ✓ |
| Q-7 | `q07SelectStarSensitive` | `@sensitive` 컬럼 보유 테이블에 `SELECT *` (WARNING) | IF-ELSE | ✓ |
| Q-8 | `q08UnusedParam` | 선언된 파라미터가 본문에서 미참조 (ERROR) | IF-ELSE | ✓ |
| Q-9 | `q09SelectOnExec` | `:exec` 쿼리에 `SELECT`/`RETURNING` 반환 (ERROR) | IF-ELSE | ✓ |
| Q-11 | `q11SqlPackagePgxV5` | `sqlc.yaml` `sql[].gen.go.sql_package` 가 `pgx/v5` (ERROR) | IF-ELSE | ✓ |
| Q-12 | `q12PgtypeUuidOverride` | DDL UUID 컬럼 존재 시 `sqlc.yaml` 에 `pgtype.UUID` overrides (NULL/NOT NULL) 강제 (ERROR) | IF-ELSE | ✓ |

## Defeater

없음.

## 비고

- Q-10 은 generate-time 규칙 (artifacts 인자 필요) — `pkg/generate/gogin/check_sqlc_out_path.go` 에 위치.
