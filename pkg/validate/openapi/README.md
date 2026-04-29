# pkg/validate/openapi

## 변경이력

- 2026-04-29: O-5 (4xx/5xx response body 필수) 신설 — BUG-040 사전 차단
- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

OpenAPI 문서 (kin-openapi `*openapi3.T`) 자체 정합성 검증. path 파라미터 충돌 / operationId 누락 / password·email 보안 휴리스틱.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| O-1 | `o01PathParamConflict` | path 내 `{param}` 이름 중복 (ERROR) | IF-ELSE | ✓ |
| O-2 | `o02PathParamCaseConflict` | 여러 path 에 걸쳐 case-only 차이의 path parameter (ERROR) | IF-ELSE | ✓ |
| O-3 | `o03PathTemplateParam` | path 템플릿 변수와 `parameters[].name` 불일치 (ERROR) | IF-ELSE | ✓ |
| O-4 | `o04OpIdRequired` | operation 에 `operationId` 누락 (ERROR) | IF-ELSE | ✓ |
| O-5 | `o05ResponseBodyRequired` | 4xx/5xx response 가 `content: application/json` + schema 미선언 (ERROR; 204/304 예외) | IF-ELSE | ✓ |
| XOO-71 | `xoo71PasswordNoMinLength` | password 계열 필드 `minLength` 없음 (WARNING) | IF-ELSE | ✓ |
| XOO-72 | `xoo72EmailNoFormat` | email 계열 필드 `format: email` 없음 (WARNING) | IF-ELSE | ✓ |

## Defeater

없음.
