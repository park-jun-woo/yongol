# pkg/validate/tsx — TSX 자체 정합성 (T-*)

React `.tsx` SSOT 자체 검증 규칙 집합. 교차 검증은 `pkg/validate/tsx_openapi/` (XOT-*) 에 있다.

## 규칙

| Rule ID | Level | 설명 |
|---|---|---|
| `T-1` | WARNING | `@/components/` / `./components/` import 의 실제 파일 존재 |

## 방향성

TSX 단일 SSOT 내 문제만 본다 (OpenAPI 교차 검증 제외). npm 패키지 import 는 파서 단계에서 이미 걸러짐.
