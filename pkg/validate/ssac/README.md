# pkg/validate/ssac

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

SSaC (Service-Sequence as Code) 자체 정합성 검증 (S-*, XSS-*). 필수 필드·변수 흐름·모델 참조·`@subscribe` 제약·pub/sub 쌍 매칭 등.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 작동 / `IF-ELSE` = 단일 판정·Ground 조회

## FieldRequired (TOULMIN, `IsSubscribe` defeater)

| 규칙 ID | SeqType | Field | Required | pkg 구현 |
|---|---|---|---|---|
| S-1 / S-2 | `@get` | Model / Result | true / true | ✓ |
| S-3 / S-4 / S-5 | `@post` | Model / Result / Inputs | true | ✓ |
| S-6 / S-7 / S-8 | `@put` | Model / Result / Inputs | true / **false** / true | ✓ |
| S-9 / S-10 | `@delete` | Model / Result | true / **false** | ✓ |
| S-12 / S-13 | `@empty` | Target / Message | true | ✓ |
| S-14 / S-15 / S-16 / S-17 | `@state` | DiagramID / Inputs / Transition / Message | true | ✓ |
| S-18 / S-19 / S-20 | `@auth` | Action / Resource / Message | true | ✓ |
| S-21 | `@call` | Model | true | ✓ |
| S-23 / S-24 | `@publish` | Topic / Payload | true | ✓ |
| S-39 | `@subscribe` | Message | true | ✓ |

## VarDeclared (TOULMIN, `IsImplicitVar` defeater)

| 규칙 ID | 함수명 | 설명 | pkg 구현 |
|---|---|---|---|
| S-27 | `VarDeclared` | 일반 식별자 변수 선언 (ERROR) | ✓ |
| S-28 | `TargetDeclared` | `@empty` Target 변수 선언 (ERROR) | ✓ |
| S-29 | `InputDeclared` | Inputs 값 변수 선언 (ERROR) | ✓ |
| S-30 | `MessageDeclared` | `@response` Fields 변수 선언 (ERROR) | ✓ |

## NameFormat (IF-ELSE)

| 규칙 ID | 함수명 | 설명 | pkg 구현 |
|---|---|---|---|
| S-26 | `DotMethod` | `Model.Method` 형식 강제 (ERROR) | ✓ |
| S-46 | `UppercaseStart` | Result 타입 첫 글자 대문자 (ERROR) | ✓ |
| S-47 | `NoDotPrefix` | `@model` package prefix 금지 (ERROR) | ✓ |
| XSS-38 | `CallFuncLowercase` | `@call` 함수명 소문자 시작 권고 (ERROR) | ✓ |

## ForbiddenRef (IF-ELSE)

| 규칙 ID | 함수명 | 설명 | pkg 구현 |
|---|---|---|---|
| S-31 | `ConfigPrefixForbidden` | `ssac.config*` prefix 금지 | ✓ |
| S-32 | `PublishForbidden` | `@publish` query 참조 금지 | ✓ |
| S-33 | `ReservedSource` | reserved source 를 result var 로 금지 | ✓ |
| S-34 / S-35 | `GoReservedWord` (var/Model) | Go 예약어 금지 | ✓ |
| S-42 | `SubscribeForbiddenRequest` | `@subscribe` 내 `request.*` 금지 | ✓ |
| S-43 | `SubscribeForbiddenQuery` | `@subscribe` 내 `query.*` 금지 | ✓ |
| S-44 | `HttpForbiddenMessage` | HTTP 함수 내 `message.*` 금지 | ✓ |

## RefExists (IF-ELSE)

| 규칙 ID | 함수명 | 설명 | pkg 구현 |
|---|---|---|---|
| S-48 | `SymbolTableModel` | `@call` Model 심볼 테이블 존재 (ERROR) | ✓ |
| S-49 | `SymbolTableMethod` | `Model.Method` 의 Method 존재 (ERROR) | ✓ |
| S-50 | `OpenAPIRequest` | SSaC input → OpenAPI request field (ERROR) | ✓ |
| S-60 | `RequestFieldExact` | `request.<field>` case-exact 일치 (ERROR) | ✓ |
| XSS-47 | `CallSourceVarUndefined` | `@call` arg source 미정의 (WARNING) | ✓ |

## CoverageCheck / TypeMatch / 흐름 분석 (IF-ELSE)

| 규칙 ID | 함수명 | 설명 | pkg 구현 |
|---|---|---|---|
| S-11 | `DeleteNoInputs` | `@delete` Inputs 없음 (WARNING) | ✓ |
| S-25 | `UnknownSeqType` | 알 수 없는 시퀀스 타입 (ERROR) | ✓ |
| S-36 | `StaleResponse` | `@put`/`@delete` 후 갱신 없이 `@response` (WARNING) | ✓ |
| S-37 | `FKReferenceGuard` | FK 참조 `@get` 후 `@empty` 가드 권고 (WARNING) | ✓ |
| S-40 | `SubscribeSingleParam` | `@subscribe` 단일 `message` 파라미터 (ERROR) | ✓ |
| S-41 | `SubscribeNoCurrentUser` | `@subscribe` 내 `currentUser` 금지 (ERROR) | ✓ |
| S-45 | `SubscribeNoResponse` | `@subscribe` 내 `@response` 금지 (ERROR) | ✓ |
| S-51 | `RequestUsage` | OpenAPI request field → SSaC 사용 (WARNING, coverage) | ✓ |
| S-57 | `FuncRequestType` | `@call` input type ↔ FuncRequest field type (ERROR) | ✓ |
| S-58 | `InvalidErrStatus` | IANA 미등록 HTTP status (ERROR) | ✓ |
| S-59 | `DottedField` | `var.field` 의 field 가 type 의 실제 field 인지 (ERROR) | ✓ |
| S-61 | `CodegenReservedVar` | result 변수명이 codegen 예약어 (`server`/`ctx`/`err`) 충돌 (ERROR) | ✓ |
| S-62 | `UnusedResultVar` | result 변수가 후속 시퀀스 미참조 (ERROR) | ✓ |
| S-63 | `ListNoPagination` | `@get []T` 에 pagination 없고 `// @no-pagination` 없음 (WARNING) | ✓ |
| S-64 | `EmptyExistsModelOnly` | `@empty`/`@exists` Target 은 Model 변수 (ERROR) | ✓ |
| S-67 | `EvalFuncSignature` | `@eval` Func 은 `func(req T) bool` (ERROR) | ✓ |
| S-68 | `EvalStatusRequired` | `@eval` STATUS 명시 필수 (ERROR) | ✓ |
| S-69 | `EvalFuncExists` | `@eval` Func 이 Func Spec/빌트인 존재 (ERROR) | ✓ |
| S-70 | `PostPutBlobInputForbidden` | `@post`/`@put` Inputs reserved source 단독 참조 금지 (ERROR) | ✓ |
| XSS-11 | `PluralResultType` | `@result` 타입 복수형 (WARNING) | ✓ |
| XSS-57 | `PublishToSubscribe` | `@publish` topic ↔ `@subscribe` 매칭 (ERROR) | ✓ |
| XSS-58 | `SubscribeToPublish` | `@subscribe` topic ↔ `@publish` 매칭 (ERROR) | ✓ |
| XSS-59 | `SubscribeFields` | `@subscribe` message fields ↔ `@publish` payload (ERROR) | ✓ |

## Defeater

| defeater | 면제 warrant | 조건 |
|---|---|---|
| `IsImplicitVar` | S-27~S-30, XSS-47 | `currentUser`/`request`/`query`/`message` 예약어 |
| `IsSubscribe` | FieldRequired (HTTP 전용) | `@subscribe` 함수는 HTTP 필드 요구 제외 |
| `seq.Type=="call"` / `seq.Package!=""` | XSS-11 | 외부 패키지 @call 결과는 plural 검사 제외 |
| primitive Go 타입 | XSS-11 | primitive 타입은 plural 검사 제외 |

## internal 일치성 메모

- `IsImplicitVar` 예약어: `currentUser`, `request`, `query`, `message` — `is_implicit_var.go`.
- `IsSubscribe`: `@subscribe` 함수는 모든 HTTP 전용 FieldRequired 에서 스킵.
- XSS-11: `seq.Type=="call"` + `seq.Package!=""` 스킵 — `xss_11_plural_result_type.go`.
- 폐기: S-52~S-56, XDS-13/14, XNS-77 (rulebook.md Deprecated 섹션 참조).
