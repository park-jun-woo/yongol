# pkg/validate/ddl_statemachine

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

stateDiagram 전이에 등장하는 필드가 DDL 컬럼에 존재하는지 + 초기 전이가 DDL DEFAULT 값과 일치하는지 검증.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙

| 규칙 ID | 함수명 | LookupKey | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|---|
| XDM-27 | `xdm27StateFieldColumn` | `DDL.column.<table>` | `@state` field → DDL column 존재 | IF-ELSE | ✓ |
| XDM-28 | `xdm28DefaultInitialState` | — | stateDiagram `[*] → X` 초기 전이와 DDL DEFAULT 'X' 일치 | IF-ELSE | ✓ |

## Defeater

없음.
