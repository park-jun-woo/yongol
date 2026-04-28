# pkg/validate/ssac_authz

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

SSaC `@auth` input field 가 Authz `CheckRequest` 구조체 필드 집합에 존재하는지 확인 (XAS-*).

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 작동 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XAS-60 | `AuthInputField` | `@auth` input field → CheckRequest 필드 (ERROR) | IF-ELSE | ✓ |

## Defeater

CheckRequest 필드가 비어 있으면 (커스텀 authz) 규칙 침묵 — `test_xas60_skips_when_custom_authz_test.go`.
