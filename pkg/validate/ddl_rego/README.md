# pkg/validate/ddl_rego

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

Rego `@ownership` / role 이 DDL 실제 테이블·컬럼·CHECK 제약을 정확히 참조하는지 확인.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙

| 규칙 ID | 함수명 | LookupKey | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|---|
| XDP-31 | `xdp31OwnershipTable` | `DDL.table` | `@ownership` table → DDL 테이블 존재 | IF-ELSE | ✓ |
| XDP-32 | `xdp32OwnershipColumn` | `DDL.column.<table>` | `@ownership` column → DDL 컬럼 존재 | IF-ELSE | ✓ |
| XDP-33 | `xdp33OwnershipJoinTable` | `DDL.table` | `@ownership via` join table → DDL | IF-ELSE | ✓ |
| XDP-34 | `xdp34OwnershipJoinColumn` | `DDL.column.<table>` | `@ownership via` join column → DDL | IF-ELSE | ✓ |
| XDP-65 | `xdp65RoleDDLCheck` | `DDL.check.<table>` | Rego role → DDL CHECK 제약 | IF-ELSE | ✓ |

## Defeater

없음.
