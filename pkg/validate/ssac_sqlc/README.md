# pkg/validate/ssac_sqlc

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

SSaC ↔ sqlc 입출력 계약 정합성 검증 (XQS-*). Input key 이름·대소문자·타입·개수가 sqlc Params 와 정확히 일치하는지 + DB-using 빌트인 호출의 sqlc 쿼리 존재성 강제.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 작동 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XQS-14 | `InputKeyCase` | SSaC input key case ↔ sqlc param case (CRUD 전용, WARNING) | IF-ELSE | ✓ |
| XQS-15 | `InputKeyInitialism` | `@call` input key Go initialism 위반 (WARNING) | IF-ELSE | ✓ |
| XQS-16 | `InputKeyMissing` | SSaC Input key 가 sqlc Params 에 없음 (ERROR) | IF-ELSE | ✓ |
| XQS-17 | `ParamKeyMissing` | sqlc Params 필드가 SSaC Input 에 없음 (ERROR) | IF-ELSE | ✓ |
| XQS-18 | `ParamTypeMismatch` | OpenAPI param 타입 ↔ sqlc/DDL 타입 불일치 (ERROR) | IF-ELSE | ✓ |
| XQS-19 | `SsacBuiltinQueryRequired` | DB-using ssac 빌트인 호출 → 대응 sqlc 쿼리 존재 (ERROR) | IF-ELSE | ✓ |

## XQS-14 / XQS-15 역할 분담

| 시퀀스 타입 | 매핑 대상 | 네이밍 규칙 | 적용 규칙 |
|---|---|---|---|
| CRUD (`@get`/`@post`/`@put`/`@delete`) | sqlc Params struct | sqlc PascalCase (`url` → `Url`) | XQS-14 / XQS-16 / XQS-17 |
| `@call` | Go 패키지 Request struct | Go initialism (`url` → `URL`) | XQS-15 |

`seq.Type == "call"` 은 XQS-14/16/17/18 에서 스킵 — sqlc 가 아닌 Go 패키지 매핑.

## Defeater

없음.
