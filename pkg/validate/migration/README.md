# pkg/validate/migration

## 변경이력

- 2026-04-28: 초기 작성

## 역할

DDL 마이그레이션 안전성 검증. 이전 스냅샷 (`arts/db/.latest_schema.sql`) 과 현재 DDL 을 비교해 rename / NOT NULL 승격 / 파괴적 변경 / 위험한 타입 변경 / 데이터 마이그레이션 누락 / 스냅샷 hash drift 를 잡아낸다. ERROR 가 하나라도 있으면 generate 파이프라인에서 마이그레이션 emit 이 차단된다.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `IF-ELSE` = 두 스키마 AST diff + 힌트 어노테이션 매칭

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| MIG-001 | `mig_001_rename_mismatch` | `@rename from=...` 의 from 이 이전 스냅샷에 없거나 to 가 현재 DDL 에 없음 (ERROR) | IF-ELSE | ✓ |
| MIG-002 | `mig_002_not_null_without_backfill` | NOT NULL 추가인데 `@backfill` 도 DEFAULT 도 없음 → emit 차단 (ERROR) | IF-ELSE | ✓ |
| MIG-003 | `mig_003_data_migration_missing` | `@data_migration file=...` 가 가리키는 sidecar SQL 이 없음 (ERROR) | IF-ELSE | ✓ |
| MIG-004 | `mig_004_destructive_without_allow` | DROP TABLE / DROP COLUMN 인데 `@allow_destructive` 없음 (WARNING) | IF-ELSE | ✓ |
| MIG-005 | `mig_005_cast_missing` | `INTEGER↔TEXT` 등 위험한 타입 변경에 `@cast using=...` 힌트 없음 (WARNING) | IF-ELSE | ✓ |
| MIG-006 | `mig_006_snapshot_drift` | `.generated_schema.sql` 의 `YONGOL_SCHEMA_HASH` 헤더가 본문 sha256 과 불일치 (ERROR) | IF-ELSE | ✓ |

## 주요 함수

| 함수 | 설명 |
|---|---|
| `Run(fs, specsDir)` | MIG-001~006 전체 실행 (generate 파이프라인이 호출) |

## 보조 헬퍼

`hasColumn`, `splitHashHeader`, `emitByRule`, `mig001CheckRenameTable`, `mig001CheckRenameColumn`.
