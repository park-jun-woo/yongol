# pkg/validate/openapi_manifest

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

OpenAPI security scheme ↔ `manifest.yaml` middleware 정합성 + `backend.http.overrides` 키와 OpenAPI operationId 매칭 검증.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙

| 규칙 ID | 함수명 | LookupKey | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|---|
| XNO-50 | `xno50SecuritySchemeMiddleware` | `Manifest.middleware` | OpenAPI securityScheme → Manifest middleware 매칭 | IF-ELSE | ✓ |
| XON-51 | `xon51MiddlewareSecurityScheme` | `OpenAPI.security` | Manifest middleware → OpenAPI securityScheme 매칭 | IF-ELSE | ✓ |
| XNO-52 | `xno52SecurityMiddleware` | `Manifest.middleware` | endpoint security 참조 → middleware 존재 + `backend.middleware` 블록 필수 | IF-ELSE | ✓ |
| SEC-04 | `sec04HttpOverridesOperationId` | `OpenAPI.operationId` | `backend.http.overrides.<key>` → OpenAPI operationId 존재 (ERROR) | IF-ELSE | ✓ |

## Defeater

없음.
