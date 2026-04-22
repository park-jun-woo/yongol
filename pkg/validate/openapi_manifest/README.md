# pkg/validate/openapi_manifest

OpenAPI security scheme 과 `manifest.yaml` middleware 설정의 정합성 검증.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 실 작동 또는 반례 확장 가능 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## RefExists (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XNO-52 | `Manifest.middleware` | endpoint security → middleware | IF-ELSE |

## PairMatch (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XNO-50 | `Manifest.middleware` | OpenAPI securityScheme → Manifest middleware | IF-ELSE |
| XON-51 | `OpenAPI.security` | Manifest middleware → OpenAPI securityScheme | IF-ELSE |

## ConfigRequired (IF-ELSE)

| 규칙 ID | ConfigKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XNO-52 | `backend.middleware` | endpoint security → middleware 필수 | IF-ELSE |

> XNO-52 는 RefExists(middleware 이름 존재) + ConfigRequired(middleware 블록 자체 존재) 두 단계로 검사한다.

## Defeater

없음.
