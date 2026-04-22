# pkg/validate/ddl_rego

Rego @ownership / role 이 DDL 실제 테이블·컬럼·CHECK 제약을 정확히 참조하는지 확인.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 실 작동 또는 반례 확장 가능 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## RefExists (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XDP-31 | `DDL.table` | @ownership table → DDL | IF-ELSE |
| XDP-32 | `DDL.column.<table>` | @ownership column → DDL | IF-ELSE |
| XDP-33 | `DDL.table` | @ownership via join table → DDL | IF-ELSE |
| XDP-34 | `DDL.column.<table>` | @ownership via join column → DDL | IF-ELSE |
| XDP-65 | `DDL.check.<table>` | Rego role → DDL CHECK 제약 | IF-ELSE |

## Defeater

없음.
