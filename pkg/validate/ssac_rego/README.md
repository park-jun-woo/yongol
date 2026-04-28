# pkg/validate/ssac_rego

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

SSaC `@auth` 와 Rego allow 규칙의 `(action, resource)` 쌍 양방향 매칭 검증 (XPS-*, XSP-*).

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 작동 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XPS-28 | `SSaCAuthToRego` | SSaC `@auth (action:resource)` → Rego allow (ERROR) | IF-ELSE | ✓ |
| XSP-29 | `RegoAllowToSSaC` | Rego allow `(action:resource)` → SSaC `@auth` (ERROR) | IF-ELSE | ✓ |

## Defeater

없음.
