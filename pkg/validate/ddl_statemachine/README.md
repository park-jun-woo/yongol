# pkg/validate/ddl_statemachine

StateDiagram 전이에 등장하는 필드가 DDL 컬럼에 존재하는지 확인.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 실 작동 또는 반례 확장 가능 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## RefExists (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XDM-27 | `DDL.column.<table>` | @state field → DDL column | IF-ELSE |

## Defeater

없음.
