# pkg/validate/tsx_openapi

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

React `.tsx` SSOT 가 OpenAPI 계약과 일치하는지 단방향 (TSX → OpenAPI) 검증 (XOT-*). OpenAPI → TSX 미소비 경고는 다소비처 false positive 노이즈로 검사하지 않는다.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 작동 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XOT-1 | `OperationId` | `apiClient.<op>()` 의 `<op>` → OpenAPI operationId 집합 (ERROR) | IF-ELSE | ✓ |
| XOT-2 | `ParameterMatch` | apiClient 호출 path/query 인자 키 → OpenAPI parameters (ERROR) | IF-ELSE | ✓ |
| XOT-3 | `FormField` | `useForm().register('x')` → 페이지 mutation 의 OpenAPI request body schema (WARNING) | IF-ELSE | ✓ |

## Defeater

| defeater | 면제 warrant | 조건 |
|---|---|---|
| `IsTransportKey` | XOT-2 | `body`/`data`/`payload`/`json` 등 transport wrapper key |
| operationId 부재 | XOT-2 | XOT-1 이 이미 커버 — XOT-2 는 skip |
| body 없음 | XOT-3 | mutation 호출에 body 없음 — skip |

## Ground Lookup 키

- `OpenAPI.operationId` — 전체 operationId 집합
- `OpenAPI.param.<opID>` — operation 별 path/query parameter 이름
- `OpenAPI.request.<opID>` — operation 별 request body 스키마 필드
