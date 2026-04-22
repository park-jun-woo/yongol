# pkg/validate/ssac_sqlc

SSaC ↔ sqlc 입출력 계약 정합성 검증. Input key 이름·대소문자·개수가 sqlc Params와 정확히 일치하는지 확인.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 실 작동 또는 반례 확장 가능 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## TypeMatch (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 | pkg 구현 |
|---------|-----------|------|----------|---------|
| XQS-14 | `SQLc.param.<model>` | SSaC input key case ↔ sqlc param | IF-ELSE | ✓ |

## Initialism (IF-ELSE)

| 규칙 ID | 함수명 | 설명 | 구현 방식 | 적용 대상 | pkg 구현 |
|---------|--------|------|----------|----------|---------|
| XQS-15 | `InputKeyInitialism` | SSaC input key Go initialism 위반 (WARNING) | IF-ELSE | **`@call` 전용** | ✓ |

### XQS-14 / XQS-15 / XQS-16 역할 분담

SSaC Input key의 네이밍 규칙은 매핑 대상에 따라 다르다:

| 시퀀스 타입 | 매핑 대상 | 네이밍 규칙 | 검증 규칙 |
|---|---|---|---|
| CRUD (`@get`/`@post`/`@put`/`@delete`) | sqlc Params struct | sqlc PascalCase (`url` → `Url`) | **XQS-14** (case) / **XQS-16** (missing) |
| `@call` | Go 패키지 Request struct | Go initialism (`url` → `URL`) | **XQS-15** |

sqlc는 `id` → `ID`만 약어 처리하고, `url` → `Url`, `api` → `Api`로 생성한다.
Go 관례는 `url` → `URL`, `api` → `API`다. 두 규칙이 다르므로 적용 대상을 분리한다.

## ParamMatch (IF-ELSE)

SSaC CRUD Input keys ↔ sqlc Params 필드 양방향 완전 일치 검증.

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---------|--------|------|----------|---------|
| XQS-16 | `InputKeyMissing` | SSaC Input key가 sqlc Params에 없음 (ERROR) | IF-ELSE | **누락** |
| XQS-17 | `ParamKeyMissing` | sqlc Params 필드가 SSaC Input에 없음 (ERROR) | IF-ELSE | **누락** |

XQS-16: SSaC에서 `{OrgID: ..., Page: ...}`를 넘기는데 sqlc Params에 `Page` 필드가 없으면 ERROR.
XQS-17: sqlc Params에 `OrgID` 필드가 있는데 SSaC Input에서 안 넘기면 ERROR.
양방향으로 잡으면 이름 + 개수 + 정확한 매칭을 보장.

`seq.Type == "call"` 스킵 — @call은 sqlc가 아닌 Go 패키지 매핑.

sqlc Params 필드명은 `@name` → PascalCase, `sqlc.arg(name)` → PascalCase로 결정된다.
LIMIT/OFFSET의 `sqlc.arg(per_page)` → `PerPage` 필드가 되므로 SSaC Input key와 일치.

## ParamTypeMatch — OpenAPI ↔ sqlc 타입 교차검증 (IF-ELSE)

SSaC Input이 `request.*`로 OpenAPI param을 참조하고, 동일 key가 sqlc Params에도 있을 때,
OpenAPI param 타입과 sqlc param 타입이 일치하는지 검증.

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---------|--------|------|----------|---------|
| XQS-18 | `ParamTypeMismatch` | OpenAPI param 타입 ↔ sqlc param 타입 불일치 (ERROR) | IF-ELSE | **누락** |

타입 호환 테이블:

| OpenAPI type | 호환 sqlc Go type |
|---|---|
| `integer` (int32/int64) | `int32`, `int64`, `int` |
| `string` | `string` |
| `boolean` | `bool` |
| `number` | `float32`, `float64` |

OpenAPI `string`인데 sqlc `bigint`(→`int64`)이면 ERROR.
권고: OpenAPI param 타입을 sqlc에 맞추거나, sqlc 쿼리의 cast를 변경.

`seq.Type == "call"` 스킵 — @call은 sqlc 매핑이 아님.

## Defeater

없음.
