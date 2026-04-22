# pkg/validate/ssac

SSaC (Service-Sequence as Code) 자체 정합성 검증. 필수 필드·변수 흐름·모델 참조·@subscribe 제약·pub/sub 쌍 매칭 등.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **상태**: 50/58 규칙 구현, 테스트 0개 (internal 114 + 77). 재구현 완성도 검증 시 internal 테스트 포팅이 안전 경로.
> **구현 방식 범례**: `TOULMIN` = defeater가 실제로 판정을 뒤집거나 반례 확장 가능성이 있는 규칙 / `IF-ELSE` = 단일 판정·Ground 조회만 (pkg/ground 에서 적재된 map/set 조회 후 단순 if 체크)

## FieldRequired (TOULMIN)

SeqType별 Field 존재/부재 검증. `IsSubscribe` defeater가 HTTP 전용 필드 요구를 면제 — 22개 규칙이 동일 shape + defeater 공유 → Toulmin 정당화.

| 규칙 ID | SeqType | Field | Required | 구현 방식 | pkg 구현 |
|---|---|---|---|---|---|
| S-1 | @get | Model | true | TOULMIN | ✓ |
| S-2 | @get | Result | true | TOULMIN | ✓ |
| S-3 | @post | Model | true | TOULMIN | ✓ |
| S-4 | @post | Result | true | TOULMIN | ✓ |
| S-5 | @post | Inputs | true | TOULMIN | ✓ |
| S-6 | @put | Model | true | TOULMIN | ✓ |
| S-7 | @put | Result | **false** | TOULMIN | ✓ |
| S-8 | @put | Inputs | true | TOULMIN | ✓ |
| S-9 | @delete | Model | true | TOULMIN | ✓ |
| S-10 | @delete | Result | **false** | TOULMIN | ✓ |
| S-12 | @empty | Target | true | TOULMIN | ✓ |
| S-13 | @empty | Message | true | TOULMIN | ✓ |
| S-14 | @state | DiagramID | true | TOULMIN | ✓ |
| S-15 | @state | Inputs | true | TOULMIN | ✓ |
| S-16 | @state | Transition | true | TOULMIN | ✓ |
| S-17 | @state | Message | true | TOULMIN | ✓ |
| S-18 | @auth | Action | true | TOULMIN | ✓ |
| S-19 | @auth | Resource | true | TOULMIN | ✓ |
| S-20 | @auth | Message | true | TOULMIN | ✓ |
| S-21 | @call | Model | true | TOULMIN | ✓ |
| S-23 | @publish | Topic | true | TOULMIN | ✓ |
| S-24 | @publish | Payload | true | TOULMIN | ✓ |

## VarDeclared (TOULMIN)

`IsImplicitVar` defeater가 예약어(`currentUser`, `request`, `message`)를 면제. 새 예약어 추가 가능성 있음.

| 규칙 ID | backing | 구현 방식 | pkg 구현 |
|---|---|---|---|
| S-27 | Ground.Vars 조회 | TOULMIN | ✓ |
| S-28~S-30 | Ground.Vars 조회 (target/input/message) | TOULMIN | **누락** |

## NameFormat (IF-ELSE)

Pattern 파라미터로 문자열 규칙만 체크. defeater 없음, 반례 가능성 낮음.

| 규칙 ID | Pattern | 구현 방식 | pkg 구현 |
|---|---|---|---|
| S-26 | `dot-method` (Model.Method 형식) | IF-ELSE | ✓ |
| S-46 | `uppercase-start` (Result 타입) | IF-ELSE | ✓ |
| S-47 | `no-dot-prefix` (@model package-prefix 금지) | IF-ELSE | ✓ |

## ForbiddenRef (IF-ELSE)

Ground.Lookup 에서 금지 set 조회만 하면 끝. defeater 없음.

| 규칙 ID | LookupKey | 구현 방식 | pkg 구현 |
|---|---|---|---|
| S-31 | `ssac.configPrefix` | IF-ELSE | ✓ |
| S-32 | `publish.forbidden` | IF-ELSE | ✓ |
| S-33 | `ssac.reservedSource` | IF-ELSE | ✓ |
| S-34~S-35 | `go.reserved` | IF-ELSE | ✓ |
| S-42 | `subscribe.forbidden` | IF-ELSE | ✓ |
| S-43 | `subscribe.forbidden` (http) | IF-ELSE | **누락** |
| S-44 | `http.forbidden` | IF-ELSE | ✓ |

## RefExists (IF-ELSE)

Ground.Lookup set 존재 여부 확인. defeater 없음.

| 규칙 ID | LookupKey | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| S-48 | `SymbolTable.model` | @call Model 이 심볼 테이블에 존재 | IF-ELSE | ✓ |
| S-49 | `SymbolTable.method.<Model>` | Model.Method 중 Method 존재 | IF-ELSE | ✓ |
| S-50 | `OpenAPI.request.<operationId>` | SSaC input → OpenAPI request field | IF-ELSE | ✓ |
| XSS-38 | (함수명 검사) | @call 함수명 소문자 시작 | IF-ELSE | — |

## CoverageCheck (IF-ELSE)

Ground set diff. defeater 없음.

| 규칙 ID | LookupKey | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| S-51 | `SSaC.requestUsage.<operationId>` | OpenAPI request field 가 SSaC 에서 사용 | IF-ELSE | **누락** |

## TypeMatch (IF-ELSE)

Ground.Types 문자열 비교. defeater 없음.

| 규칙 ID | LookupKey | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| S-57 | `Func.request.<funcName>` | @call input type ↔ FuncRequest field type | IF-ELSE | ✓ |

## PairMatch pub/sub (IF-ELSE)

topic 쌍 매칭. 반례 가능성 낮음.

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---|---|---|---|
| XSS-57 | `SSaC.subscribe` | @publish topic → @subscribe | IF-ELSE |
| XSS-58 | `SSaC.publish` | @subscribe topic → @publish | IF-ELSE |

## SchemaMatch pub/sub payload (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---|---|---|---|
| XSS-59 | `SSaC.publish.<topic>` | @subscribe message fields → @publish payload | IF-ELSE |

## 고유 함수

| 규칙 ID | 함수명 | 설명 | 구현 방식 | 비고 |
|---------|--------|------|----------|------|
| S-11 | `DeleteNoInputs` | @delete Inputs 없음 WARNING | IF-ELSE | 단일 조건 |
| S-25 | `UnknownSeqType` | 알 수 없는 시퀀스 타입 | IF-ELSE | 단일 조건 |
| S-36 | `StaleResponse` | @put/@delete 후 갱신 없이 @response 사용 | IF-ELSE | 시퀀스 흐름 분석 |
| S-37 | `FKReferenceGuard` | FK 참조 @get 후 @empty 가드 필요 | IF-ELSE | 시퀀스 흐름 분석 |
| S-38~S-41, S-45 | `SubscribeConstraints` | @subscribe 제약 (파라미터, message struct, @response 금지) | IF-ELSE | 다조건 검사지만 defeater 없음 |
| S-58 | `InvalidErrStatus` | IANA 미등록 HTTP status | IF-ELSE | 테이블 조회 |
| XSS-11 | `PluralResultType` | @result 타입 복수형 (WARNING) | TOULMIN | primitive 스킵, `seq.Type=="call"`, `seq.Package!=""` 등 스킵 조건 |
| XSS-47 | `CallSourceVarUndefined` | @call arg source 미정의 (WARNING) | TOULMIN | `IsImplicitVar` defeater — 새 예약어 추가 가능성 |
| S-61 | `CodegenReservedVar` | result 변수명이 코드젠 예약어(`server`, `ctx`, `err` 등)와 충돌 | IF-ELSE | 단일 조건 |
| S-62 | `UnusedResultVar` | result 변수가 후속 시퀀스에서 미참조 (ERROR) | IF-ELSE | 시퀀스 흐름 분석 |
| S-63 | `ListNoPagination` | `@get []T` list 엔드포인트인데 pagination params 없고 `// @no-pagination` 없음 (WARNING) | IF-ELSE | 단일 조건 |

## ListPagination (IF-ELSE)

`@get []T` (배열 리턴)인데 Inputs에 pagination key(`Page`/`PerPage`/`Cursor`)가 없으면 WARNING.
의도적으로 pagination 없는 list면 SSaC 함수 주석에 `// @no-pagination`을 붙여 면제.

## CodegenReserved (IF-ELSE)

코드젠이 사용하는 예약 변수명과 충돌 방지. defeater 없음.

| 규칙 ID | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|
| S-61 | result 변수명이 코드젠 예약어와 충돌 | IF-ELSE | ✓ |

## Defeater

| defeater | 면제 warrant | 조건 |
|---|---|---|
| `IsImplicitVar` | VarDeclared (S-27~S-30), XSS-47 | `currentUser`, `request`, `message` 예약어 |
| `IsSubscribe` | FieldRequired (HTTP 전용) | @subscribe 함수는 HTTP 필드 요구에서 제외 |

## internal 일치성 메모

- **IsImplicitVar 예약어 제외**: `currentUser`, `request`, `message` — VarDeclared 계열에서 반드시 스킵
- **IsSubscribe 분기**: @subscribe 는 HTTP 전용 필드 요구 제외
- XSS-11: `seq.Type == "call"` 스킵, `seq.Package != ""` 스킵 — `check_ssac_ddl_func.go:13-18`
- internal `test_validate_*.go` (94개) 케이스는 실전 엣지 케이스이므로 재구현 완성도 검증 시 포팅 권장
