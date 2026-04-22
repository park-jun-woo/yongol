# pkg/validate/rego_manifest

Rego 정책의 `input.claims` / role 참조가 `manifest.yaml` 의 claims/roles 정의와 일치하는지, 역으로 Manifest 가 선언한 claims/roles 가 Rego 에서 사용되는지 확인.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 실 작동 또는 반례 확장 가능 / `IF-ELSE` = 단일 판정·Ground 조회

## RefExists (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XNP-53 | `Manifest.claims.values` | Rego input.claims → claims 값 | IF-ELSE |
| XNP-63 | `Manifest.roles` | Rego role → Manifest roles | IF-ELSE |

## CoverageCheck

| 규칙 ID | LookupKey | 설명 | 구현 방식 | 예외 |
|---------|-----------|------|----------|------|
| XPN-54 | `Rego.claims` | Manifest claims → Rego 참조 여부 | TOULMIN | "middleware/response 참조도 인정 가능" 검토 중 — 반례 추가 가능성 |
| XPN-64 | `Rego.roles` | Manifest roles → Rego 사용 여부 | TOULMIN | XPN-54와 대칭 — 동일 확장 후보군 |

## Defeater

없음.

## internal 필수 예외

- XPN-54: Rego 외 middleware/response 참조도 인정 가능 (검토 중) — `check_claims.go` forward+reverse
