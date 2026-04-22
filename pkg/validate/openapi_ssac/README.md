# pkg/validate/openapi_ssac

OpenAPI ↔ SSaC 교차 검증.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 실 작동 또는 반례 확장 가능 / `IF-ELSE` = 단일 판정·Ground 조회

## RefExists (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XOS-15 | `OpenAPI.operationId` | SSaC funcName → OpenAPI | IF-ELSE |

## CoverageCheck

| 규칙 ID | LookupKey | 설명 | 구현 방식 | 예외 |
|---------|-----------|------|----------|------|
| XSO-16 | `SSaC.funcName` | OpenAPI operationId → SSaC 함수 사용 여부 | IF-ELSE | — |
| XSO-18 | `SSaC.response.<funcName>` | OpenAPI response field → @response 사용 여부 | IF-ELSE | — |
| XSO-20 | `SSaC.response.<funcName>` | OpenAPI field → shorthand @response 사용 여부 | IF-ELSE | wrapper 면제 제거 — shorthand도 OpenAPI fields 매칭 필수 |

## SchemaMatch (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XOS-17 | `OpenAPI.response.<op>` | SSaC @response fields → OpenAPI response | IF-ELSE |
| XOS-19 | `OpenAPI.response.<op>` | shorthand @response → OpenAPI response | IF-ELSE |
| XOS-66 | `OpenAPI.required.<op>` | SSaC used fields → OpenAPI required | IF-ELSE |

## 고유 함수 (IF-ELSE)

| 규칙 ID | 함수명 | 설명 | 구현 방식 |
|---------|--------|------|----------|
| XOS-21 | `ErrStatusNotInOpenAPI` | @empty/@exists/@state/@auth/@call ErrStatus OpenAPI 미정의 (ERROR) | IF-ELSE |
| XOS-22 | `ResponseNo2xx` | @response 있는데 OpenAPI 2xx 없음 (ERROR) | IF-ELSE |

## 폐기

폐기된 규칙의 이력은 저장소 루트 [`rulebook.md`](../../../rulebook.md) Deprecated 섹션 참조.

## Defeater

없음.

## internal 필수 예외

- ~~XSO-20: `Result.Wrapper != ""` 스킵~~ — **폐기**: wrapper 면제 제거
