# pkg/validate/openapi_ssac

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

OpenAPI ↔ SSaC 교차 검증. operationId/funcName 매칭, `@response` 필드 ↔ schema, ErrStatus ↔ responses, 성공 상태 관례 검사.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XOS-15 | `xos15FuncNameOpId` | SSaC funcName → OpenAPI operationId 존재 (ERROR) | IF-ELSE | ✓ |
| XSO-16 | `xso16OpIdToFunc` | OpenAPI operationId → SSaC funcName 구현 존재 (ERROR) | IF-ELSE | ✓ |
| XOS-17 | `xos17ResponseFields` | SSaC `@response` 필드 → OpenAPI response schema 포함 (ERROR) | IF-ELSE | ✓ |
| XSO-18 | `xso18ResponseFieldUsed` | OpenAPI response 필드 → SSaC 명시 `@response` 사용 (WARNING) | IF-ELSE | ✓ |
| XOS-19 | `xos19ShorthandResponse` | shorthand `@response` 필드 → OpenAPI response 포함 (ERROR) | IF-ELSE | ✓ |
| XSO-20 | `xso20ShorthandFieldUsed` | OpenAPI response 필드 → shorthand `@response` 변수 타입에 사용 (WARNING) | IF-ELSE | ✓ |
| XOS-21 | `xos21ErrStatusNotInOpenAPI` | SSaC `@empty/@exists/@state/@auth/@call` ErrStatus 가 OpenAPI responses 미정의 (ERROR) | IF-ELSE | ✓ |
| XOS-22 | `xos22ResponseNo2xx` | `@response` 있는데 OpenAPI 에 2xx 없음 (ERROR) | IF-ELSE | ✓ |
| XOS-66 | `xos66UsedFieldsRequired` | SSaC 에서 참조한 request 필드가 OpenAPI requestBody `required` 에 등재 (ERROR) | IF-ELSE | ✓ |
| XOS-67 | `xos67ResponseFieldType` | `@response` 필드 값 타입 ↔ OpenAPI response schema 기대 타입 호환 (ERROR) | IF-ELSE | ✓ |
| XOS-80 | `xos80SuccessStatusMismatch` | HTTP method 관례 성공 상태가 OpenAPI responses 부재 (ERROR) | IF-ELSE | ✓ |
| XOS-82 | `xos82UnreachableSuccessStatus` | OpenAPI 선언 2xx 중 yongol 이 emit 하지 않는 상태 존재 (WARNING) | IF-ELSE | ✓ |

## Defeater

없음.

## internal 일치성 메모

- ~~XSO-20 `Result.Wrapper != ""` 스킵~~ — 폐기. wrapper(Page[T]/Cursor[T]/[]T) 도 OpenAPI fields 매칭 필수.
- 폐기 규칙 이력은 [`rulebook.md`](../../../rulebook.md) Deprecated 섹션 참조.
