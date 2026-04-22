# pkg/validate/ssac_rego

SSaC @auth 와 Rego allow 규칙의 `action:resource` 쌍 양방향 매칭.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 실 작동 또는 반례 확장 가능 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## PairMatch (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XPS-28 | `Policy.auth` | SSaC @auth (action:resource) → Rego allow | IF-ELSE |
| XSP-29 | `SSaC.auth` | Rego allow (action:resource) → SSaC @auth | IF-ELSE |

## Defeater

없음.
