# pkg/validate/rego_manifest

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

Rego 정책의 `input.claims` / role 참조와 `manifest.yaml` claims/roles 간 양방향 정합성 검증 (XNP-*, XPN-*).

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 작동 또는 반례 확장 가능 / `IF-ELSE` = 단일 판정·Ground 조회

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XNP-53 | `InputClaimsValues` | Rego `input.claims` 값 → manifest claims | IF-ELSE | ✓ |
| XNP-63 | `RoleManifest` | Rego role → manifest roles | IF-ELSE | ✓ |
| XPN-54 | `ClaimsToRego` | manifest claim → Rego/middleware/response 참조 (WARNING, coverage) | TOULMIN | ✓ |
| XPN-64 | `RolesToRego` | manifest role → Rego allow 참조 (WARNING, coverage) | TOULMIN | ✓ |

## Defeater

없음 (XPN-54/64 는 middleware/response 참조도 인정 — 향후 반례 추가 가능).

## internal 일치성 메모

- XPN-54: Rego 외 middleware/response 참조도 인정 — `xpn_54_claims_to_rego.go` forward+reverse 검사.
