# pkg/validate/tsx_openapi — TSX ↔ OpenAPI 교차 검증 (XOT-*)

React `.tsx` SSOT 가 OpenAPI 계약과 일치하는지 단방향(TSX → OpenAPI) 검증.

## 규칙

| Rule ID | Level | 설명 |
|---|---|---|
| `XOT-1` | ERROR | `apiClient.<op>()` 의 `<op>` 가 OpenAPI operationId 집합에 존재 |
| `XOT-2` | ERROR | apiClient 호출의 path/query 인자 객체 키가 OpenAPI parameters 에 존재 |
| `XOT-3` | WARNING | `useForm().register('x')` 필드가 해당 페이지 mutation 의 OpenAPI request body schema 에 존재 |

## 방향성 — 단방향 검증

```
TSX (주장)  →  OpenAPI (정답)     : 검증 (XOT-*)
OpenAPI (정답)  →  TSX (소비 여부)  : 검증 안 함
```

OpenAPI operationId 는 모바일 앱 / CLI / 파트너 / 배치 등 여러 소비처를 가질 수 있어 TSX 미소비 경고는 false positive 노이즈다.

## Ground Lookup 키

- `OpenAPI.operationId` — 전체 operationId 집합
- `OpenAPI.param.<opID>` — operation 별 path/query parameter 이름
- `OpenAPI.request.<opID>` — operation 별 request body 스키마 필드
