# pkg/validate/query_rego

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

Rego `@ownership` 매핑이 대응하는 sqlc 쿼리 (`OwnerLookup<Resource>`) 를 가지는지 정적으로 강제. ssac/pkg/authz `interface.yaml` 의 canonical 명명 규약 준수.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XQP-30 | `xqp30OwnerLookupQuery` | `@ownership <res>: <table>.<col>` 매핑은 대응 sqlc 쿼리 `OwnerLookup<Resource>` 가 존재해야 함 (ERROR). 부재 시 advice 로 sqlc 쿼리 스텁 제공 (via 매핑은 JOIN 형식) | IF-ELSE | ✓ |

## Defeater

없음.
