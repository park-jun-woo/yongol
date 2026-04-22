# pkg/validate/openapi_hurl

Hurl 시나리오가 OpenAPI 에 정의된 path/method/status 를 사용하는지 확인.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 실 작동 또는 반례 확장 가능 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## RefExists (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XOH-35 | `OpenAPI.path` | Hurl path → OpenAPI | IF-ELSE |
| XOH-36 | `OpenAPI.method.<path>` | Hurl method → OpenAPI | IF-ELSE |

## 고유 함수 (IF-ELSE)

| 규칙 ID | 함수명 | 설명 | 구현 방식 |
|---------|--------|------|----------|
| XOH-37 | `HurlStatusNotDefined` | Hurl status code OpenAPI 미정의 (WARNING) | IF-ELSE |

## Defeater

없음.
