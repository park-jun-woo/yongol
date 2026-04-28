# pkg/validate/rego

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

OPA Rego 정책 문서 자체 정합성 검증 (P-*, XPP-*). 파싱 오류와 `resource_owner` 참조 시 `@ownership` 어노테이션 누락을 잡는다.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| P-1 | `RegoParse` | Rego 파일 재파싱으로 구조 오류 감지 (ERROR) | IF-ELSE | ✓ |
| XPP-30 | `OwnershipNoAnnotation` | `resource_owner` 참조 시 `@ownership` 누락 (ERROR) | IF-ELSE | ✓ |

## Defeater

없음.
