# pkg/validate/openapi

OpenAPI 문서 자체 정합성 검증 (kin-openapi 로 파싱된 `*openapi3.T` 기반).

> 규칙 전체 목록은 저장소 루트의 [`rulebook.md`](../../../rulebook.md) 참조.
> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule 공통 함수 + defeater 그래프 / `IF-ELSE` = 단일 구조·흐름·휴리스틱 검사

## 고유 함수 (구조 검증)

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---------|--------|------|----------|---------|
| O-1 | `PathParamConflict` | path 파라미터명 충돌 | IF-ELSE | ✓ |
| O-4 | `OpIdRequired` | operation 에 `operationId` 누락 (ERROR) | IF-ELSE | ✓ |

## 고유 함수 (의미 검증)

| 규칙 ID | 함수명 | 설명 | 구현 방식 | 조건 |
|---------|--------|------|----------|------|
| XOO-71 | `PasswordNoMinLength` | password 필드 minLength 없음 (WARNING) | IF-ELSE | OpenAPI 필드명 `password` / `new_password` / `current_password` 등 |
| XOO-72 | `EmailNoFormat` | email 필드 format 없음 (WARNING) | IF-ELSE | OpenAPI 필드명 `email` 계열 |

## pkg/rule 사용

없음.

## Defeater

없음.
