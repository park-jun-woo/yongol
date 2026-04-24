# yongol Validation Rulebook

`yongol validate` 가 실행하는 전수 규칙 카탈로그. 매뉴얼(`manual-for-ai.md`)은 요약표만 유지하고 상세는 본 문서로 분리한다.

> 본 문서는 SSOT 간 계약(contract) 이 아닌 **파일 사실(fact)** 을 기준으로 구성한다. 대부분의 Rule 은 `pkg/validate/<domain>/` 에 Go 소스로 실재한다. 단, `<artifacts>` 경로가 필요한 generate-time 규칙은 `pkg/generate/<stack>/` 에 존재할 수 있다 (예: `Q-10`).

## Rule ID Prefix 범례

### 단일 SSOT

| Prefix | SSOT |
|---|---|
| `C-` | manifest.yaml |
| `D-`, `XDD-` | DDL (PostgreSQL CREATE TABLE + sqlc 쿼리) |
| `O-`, `XOO-` | OpenAPI |
| `S-`, `XSS-` | SSaC (Service-Sequence as Code) |
| `TM-` | STML (HTML-like page template) — Deprecated (STML 폐기 Phase008, XOT-* 로 대체) |
| `T-` | React TSX (frontend/**/*.tsx) |
| `ST-` | Mermaid stateDiagram |
| `P-`, `XPP-` | OPA Rego |
| `F-`, `XFF-` | Func 스펙 (Go AST) |
| `H-` | Hurl 시나리오 |
| `Q-` | sqlc 쿼리 |
| `CORS-` | manifest CORS 블록 |

### 교차 SSOT — `X<target><source>-<N>`

`<target>` = Lookup 키가 가리키는 SSOT (정답), `<source>` = 주장(claim)을 내는 SSOT.

| SSOT | Code | SSOT | Code |
|---|---|---|---|
| OpenAPI | `O` | DDL | `D` |
| SSaC | `S` | StateMachine | `M` |
| Rego | `P` | Manifest | `N` |
| Hurl | `H` | Func | `F` |
| Authz | `A` | sqlc | `Q` |
| TSX | `T` | | |

예: SSaC → OpenAPI (SSaC 가 주장, OpenAPI 가 정답) → `XOS-`.
예: TSX → OpenAPI (TSX 가 주장, OpenAPI 가 정답) → `XOT-`.

## Level

| Level | 동작 | 의미 |
|---|---|---|
| `ERROR` | `yongol validate` 실패 → exit code ≠ 0 | 정합성 위반. 반드시 수정 필요 |
| `WARNING` | 실패시키지 않음, 리포트에만 표시 | 설계 의도 재확인 권장 |

## Source

각 규칙 행의 `Source` 는 repo 루트 상대경로의 Go 파일이다. 예: `pkg/validate/ssac/s_27_var_declared.go`.

---

## A. SSaC Internal

SSaC 자체 정합성 — 필수 필드, 변수 흐름, 모델 참조, @subscribe 제약, pub/sub 쌍 매칭.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| S-1 | ERROR | `@get` Model 필드 필수 | `pkg/validate/ssac/s_01_get_model.go` |
| S-2 | ERROR | `@get` Result 필드 필수 | `pkg/validate/ssac/s_02_get_result.go` |
| S-3 | ERROR | `@post` Model 필드 필수 | `pkg/validate/ssac/s_03_post_model.go` |
| S-4 | ERROR | `@post` Result 필드 필수 | `pkg/validate/ssac/s_04_post_result.go` |
| S-5 | ERROR | `@post` Inputs 필드 필수 | `pkg/validate/ssac/s_05_post_inputs.go` |
| S-6 | ERROR | `@put` Model 필드 필수 | `pkg/validate/ssac/s_06_put_model.go` |
| S-7 | ERROR | `@put` Result 필드 필수 | `pkg/validate/ssac/s_07_put_result.go` |
| S-8 | ERROR | `@put` Inputs 필드 필수 | `pkg/validate/ssac/s_08_put_inputs.go` |
| S-9 | ERROR | `@delete` Model 필드 필수 | `pkg/validate/ssac/s_09_delete_model.go` |
| S-10 | ERROR | `@delete` Result 필드 필수 | `pkg/validate/ssac/s_10_delete_result.go` |
| S-11 | WARNING | `@delete` Inputs 없음 | `pkg/validate/ssac/s_11_delete_no_inputs.go` |
| S-12 | ERROR | `@empty` Target 필드 필수 | `pkg/validate/ssac/s_12_empty_target.go` |
| S-13 | ERROR | `@empty` Message 필드 필수 | `pkg/validate/ssac/s_13_empty_message.go` |
| S-14 | ERROR | `@state` DiagramID 필드 필수 | `pkg/validate/ssac/s_14_state_diagram_id.go` |
| S-15 | ERROR | `@state` Inputs 필드 필수 | `pkg/validate/ssac/s_15_state_inputs.go` |
| S-16 | ERROR | `@state` Transition 필드 필수 | `pkg/validate/ssac/s_16_state_transition.go` |
| S-17 | ERROR | `@state` Message 필드 필수 | `pkg/validate/ssac/s_17_state_message.go` |
| S-18 | ERROR | `@auth` Action 필드 필수 | `pkg/validate/ssac/s_18_auth_action.go` |
| S-19 | ERROR | `@auth` Resource 필드 필수 | `pkg/validate/ssac/s_19_auth_resource.go` |
| S-20 | ERROR | `@auth` Message 필드 필수 | `pkg/validate/ssac/s_20_auth_message.go` |
| S-21 | ERROR | `@call` Model 필드 필수 | `pkg/validate/ssac/s_21_call_model.go` |
| S-23 | ERROR | `@publish` Topic 필드 필수 | `pkg/validate/ssac/s_23_publish_topic.go` |
| S-24 | ERROR | `@publish` Payload 필드 필수 | `pkg/validate/ssac/s_24_publish_payload.go` |
| S-25 | ERROR | 알 수 없는 시퀀스 타입 | `pkg/validate/ssac/s_25_unknown_seq_type.go` |
| S-26 | ERROR | @call 값이 `Model.Method` 형식이어야 함 | `pkg/validate/ssac/s_26_dot_method.go` |
| S-27 | ERROR | 변수 선언 후 사용 (일반 식별자) | `pkg/validate/ssac/s_27_var_declared.go` |
| S-28 | ERROR | `@empty` Target 에 쓰인 변수 선언 필수 | `pkg/validate/ssac/s_28_target_declared.go` |
| S-29 | ERROR | Inputs 값에 쓰인 변수 선언 필수 | `pkg/validate/ssac/s_29_input_declared.go` |
| S-30 | ERROR | `@empty` Message 에 쓰인 변수 선언 필수 | `pkg/validate/ssac/s_30_message_declared.go` |
| S-31 | ERROR | `ssac.config*` prefix 금지 | `pkg/validate/ssac/s_31_config_prefix_forbidden.go` |
| S-32 | ERROR | `@publish` topic 에 예약어 금지 | `pkg/validate/ssac/s_32_publish_forbidden.go` |
| S-33 | ERROR | 예약된 source 식별자 사용 금지 | `pkg/validate/ssac/s_33_reserved_source.go` |
| S-34 | ERROR | Go 예약어를 변수명으로 사용 금지 | `pkg/validate/ssac/s_34_go_reserved_word.go` |
| S-35 | ERROR | Go 예약어를 Model 이름으로 사용 금지 | `pkg/validate/ssac/s_35_go_reserved_word_model.go` |
| S-36 | WARNING | `@put`/`@delete` 후 갱신 없이 `@response` 사용 (stale) | `pkg/validate/ssac/s_36_check_response_stale.go` |
| S-37 | WARNING | FK 참조 `@get` 후 `@empty` 가드 필요 | `pkg/validate/ssac/s_37_fk_reference_guard.go` |
| S-38 | ERROR | `@subscribe` 에 HTTP 전용 Inputs 금지 | `pkg/validate/ssac/s_38_subscribe_no_http_inputs.go` |
| S-39 | ERROR | `@subscribe` Message 필수 | `pkg/validate/ssac/s_39_subscribe_message_required.go` |
| S-40 | ERROR | `@subscribe` 함수 파라미터 단일 제약 | `pkg/validate/ssac/s_40_subscribe_single_param.go` |
| S-41 | ERROR | `@subscribe` 에 `currentUser` 사용 금지 | `pkg/validate/ssac/s_41_subscribe_no_currentuser.go` |
| S-42 | ERROR | `@subscribe` 에 `request.*` 사용 금지 | `pkg/validate/ssac/s_42_subscribe_forbidden_request.go` |
| S-43 | ERROR | `@subscribe` 에 `query.*` 사용 금지 | `pkg/validate/ssac/s_43_subscribe_forbidden_query.go` |
| S-44 | ERROR | HTTP 함수에 `message.*` 사용 금지 | `pkg/validate/ssac/s_44_http_forbidden_message.go` |
| S-45 | ERROR | `@subscribe` 에 `@response` 금지 | `pkg/validate/ssac/s_45_subscribe_no_response.go` |
| S-46 | ERROR | Result 타입은 대문자로 시작 | `pkg/validate/ssac/s_46_uppercase_start.go` |
| S-47 | ERROR | `@model` 에 패키지 prefix 금지 | `pkg/validate/ssac/s_47_no_dot_prefix.go` |
| S-48 | ERROR | `@call` Model 이 SymbolTable 에 존재 | `pkg/validate/ssac/s_48_symbol_table_model.go` |
| S-49 | ERROR | `Model.Method` 중 Method 가 SymbolTable 에 존재 | `pkg/validate/ssac/s_49_symbol_table_method.go` |
| S-50 | ERROR | SSaC input → OpenAPI request field 존재 | `pkg/validate/ssac/s_50_openapi_request.go` |
| S-51 | WARNING | OpenAPI request field 가 SSaC 에서 사용 | `pkg/validate/ssac/s_51_request_usage.go` |
| S-57 | ERROR | `@call` input type ↔ FuncRequest field type 일치 | `pkg/validate/ssac/s_57_func_request_type.go` |
| S-58 | ERROR | IANA 미등록 HTTP status 사용 금지 | `pkg/validate/ssac/s_58_invalid_err_status.go` |
| S-59 | ERROR | Dotted field 참조 형식 검증 | `pkg/validate/ssac/s_59_dotted_field.go` |
| S-61 | ERROR | Result 변수명이 코드젠 예약어(`server`, `ctx`, `err` 등)와 충돌 금지 | `pkg/validate/ssac/s_61_codegen_reserved_var.go` |
| S-62 | ERROR | Result 변수가 후속 시퀀스에서 미참조 | `pkg/validate/ssac/s_62_unused_result_var.go` |
| S-63 | WARNING | `@get []T` list 엔드포인트인데 pagination params 없고 `// @no-pagination` 없음 | `pkg/validate/ssac/s_63_list_no_pagination.go` |
| XSS-11 | WARNING | `@result` 타입이 복수형 | `pkg/validate/ssac/xss_11_plural_result_type.go` |
| XSS-38 | ERROR | `@call` 함수명 소문자 시작 (대문자 권고) | `pkg/validate/ssac/xss_38_call_func_lowercase.go` |
| XSS-47 | WARNING | `@call` 인자 source 변수 미정의 | `pkg/validate/ssac/xss_47_call_source_var_undefined.go` |
| XSS-57 | ERROR | `@publish` topic 에 대응하는 `@subscribe` 존재 | `pkg/validate/ssac/xss_57_publish_to_subscribe.go` |
| XSS-58 | ERROR | `@subscribe` topic 에 대응하는 `@publish` 존재 | `pkg/validate/ssac/xss_58_subscribe_to_publish.go` |
| XSS-59 | ERROR | `@subscribe` message 필드 ↔ `@publish` payload 일치 | `pkg/validate/ssac/xss_59_subscribe_fields.go` |

## B. Manifest

`manifest.yaml` 로드 및 기본 스키마 검증.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| C-2 | ERROR | `apiVersion` 값이 `yongol/v1` 이어야 함 | `pkg/validate/manifest/c_02_api_version.go` |
| C-3 | ERROR | `kind` 값이 `Project` 이어야 함 | `pkg/validate/manifest/c_03_kind.go` |
| C-4 | ERROR | `metadata.name` 필수 (빈 값 금지) | `pkg/validate/manifest/c_04_metadata_name.go` |
| C-5 | ERROR | `backend.module` 필수 (빈 값 금지) | `pkg/validate/manifest/c_05_backend_module.go` |
| CORS-01 | ERROR | `allow_origins=["*"]` + `allow_credentials=true` 동시 허용 금지 | `pkg/validate/manifest/cors_01_wildcard_credentials.go` |
| OBS-001 | ERROR | `backend.observability.metrics.path` 는 `/` 로 시작하는 절대 경로여야 함 | `pkg/validate/manifest/obs_01_metrics_path.go` |
| OBS-002 | ERROR | `backend.observability.metrics.path` 가 OpenAPI path 와 충돌 금지 | `pkg/validate/manifest/obs_02_metrics_path_not_openapi.go` |
| OBS-003 | ERROR | `backend.observability.tracing.exporter` 는 `otlp`/`stdout`/`noop` 중 하나여야 함 (enabled=true 시) | `pkg/validate/manifest/obs_03_tracing_exporter.go` |
| OBS-004 | ERROR | `backend.observability.tracing.sample_rate` 는 `[0.0, 1.0]` 범위여야 함 (enabled=true 시) | `pkg/validate/manifest/obs_04_tracing_sample_rate.go` |
| SEC-201 | ERROR | `backend.auth.mode=cookie\|hybrid` + `csrf.enabled=false` 조합 금지 (CSRF 공격 무방비) | `pkg/validate/manifest/sec_201_cookie_without_csrf.go` |
| SEC-301 | WARNING | `backend.security_headers.csp.directives.default-src` 에 `*` / `'unsafe-eval'` 포함 시 CSP 보호 약화 | `pkg/validate/manifest/sec_301_csp_permissive.go` |
| SEC-302 | WARNING | `backend.security_headers.hsts.max_age < 15552000` (180일) 시 HSTS preload 요구치 미달 | `pkg/validate/manifest/sec_302_hsts_short.go` |
| SEC-401 | ERROR | `backend.auth.secret` 리터럴 금지 (git 유출/rotation 불가) — `secret_env` 만 허용 | `pkg/validate/manifest/sec_401_jwt_secret_env_required.go` |
| SEC-402 | WARNING | `backend.auth.access_token_ttl > 30m` 은 OWASP 권고 상한 초과 (blast radius 확대) | `pkg/validate/manifest/sec_402_access_ttl_upper_bound.go` |
| SEC-403 | ERROR | `backend.auth.mode` 는 `cookie` / `bearer` / `hybrid` 중 하나여야 함 (미지정 시 `cookie`) | `pkg/validate/manifest/sec_403_auth_mode_enum.go` |

## C. OpenAPI Internal

OpenAPI 자체 정합성 (kin-openapi 로 파싱된 문서 기준).

| Rule ID | Level | Description | Source |
|---|---|---|---|
| O-1 | ERROR | path 파라미터명 충돌 | `pkg/validate/openapi/o_01_path_param_conflict.go` |
| O-2 | ERROR | path 파라미터 대소문자 충돌 | `pkg/validate/openapi/o_02_path_param_case_conflict.go` |
| O-3 | ERROR | path 템플릿 파라미터 선언 누락 | `pkg/validate/openapi/o_03_path_template_param.go` |
| O-4 | ERROR | operation 에 `operationId` 누락 | `pkg/validate/openapi/o_04_op_id_required.go` |
| XOO-71 | WARNING | password 계열 필드에 `minLength` 없음 | `pkg/validate/openapi/xoo_71_password_no_min_length.go` |
| XOO-72 | WARNING | email 계열 필드에 `format` 없음 | `pkg/validate/openapi/xoo_72_email_no_format.go` |

## D. Query / sqlc

sqlc 쿼리 파일 자체 정합성.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| Q-1 | ERROR | `-- name:` 어노테이션 필수 | `pkg/validate/query/q_01_name_required.go` |
| Q-2 | ERROR | cardinality (`:one` / `:many` / `:exec` / `:execrows`) 필수 | `pkg/validate/query/q_02_cardinality.go` |
| Q-3 | ERROR | 쿼리명 PascalCase 강제 | `pkg/validate/query/q_03_name_pascalcase.go` |
| Q-4 | WARNING | `:many` 쿼리에 `LIMIT` 누락 | `pkg/validate/query/q_04_many_limit.go` |
| Q-5 | ERROR | `DELETE` 문에 `WHERE` 필수 | `pkg/validate/query/q_05_delete_where.go` |
| Q-6 | ERROR | `UPDATE` 문에 `WHERE` 필수 | `pkg/validate/query/q_06_update_where.go` |
| Q-7 | WARNING | `SELECT *` + `@sensitive` 컬럼 보유 테이블 | `pkg/validate/query/q_07_select_star_sensitive.go` |
| Q-8 | ERROR | 선언된 파라미터가 본문에 미사용 | `pkg/validate/query/q_08_unused_param.go` |
| Q-9 | ERROR | `:exec` 쿼리에 `SELECT` 반환 | `pkg/validate/query/q_09_select_on_exec.go` |
| Q-10 | ERROR | `sqlc.yaml` 의 `sql[].gen.go.out` 이 `<artifacts>/backend/internal/db` 로 resolve 되어야 함 (generate-time, `<artifacts>` CLI 인자 필요) | `pkg/generate/gogin/check_sqlc_out_path.go` |

## E. DDL

DDL 자체 정합성 (PostgreSQL + sqlc 쿼리 정의).

| Rule ID | Level | Description | Source |
|---|---|---|---|
| D-1 | ERROR | sqlc query name 중복 | `pkg/validate/ddl/d_01_sqlc_query_duplicate.go` |
| D-2 | ERROR | `NOT NULL` 제약 누락 | `pkg/validate/ddl/d_02_nullable_column.go` |
| D-3 | ERROR | FK `DEFAULT 0` 센티널 레코드 누락 | `pkg/validate/ddl/d_03_sentinel_missing.go` |
| D-4 | ERROR | `db/sqlc.yaml` 파일 미존재 | `pkg/validate/ddl/d_04_sqlc_yaml_required.go` |
| D-5 | WARNING | `sqlc.yaml` 의 `schema` 경로가 DDL 디렉토리 미포함 | `pkg/validate/ddl/d_05_sqlc_yaml_schema_path.go` |
| D-6 | WARNING | `sqlc.yaml` 의 `queries` 경로가 `queries/` 미포함 | `pkg/validate/ddl/d_06_sqlc_yaml_queries_path.go` |
| D-7 | ERROR | sqlc 쿼리에 위치 파라미터(`$1`, `$2`) 사용 금지 | `pkg/validate/ddl/d_07_sqlc_positional_param.go` |
| XDD-61 | WARNING | 민감 패턴 (`password` / `secret` / `hash` / `token`) 컬럼에 `@sensitive` 어노테이션 누락 | `pkg/validate/ddl/xdd_61_sensitive_no_annotation.go` |

## F. DDL ↔ OpenAPI

DDL 테이블/컬럼과 OpenAPI 스키마·확장 필드의 교차 정합성.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XDO-9 | ERROR | OpenAPI property 가 어느 DDL 테이블 컬럼에도 대응되지 않음 (ghost) | `pkg/validate/openapi_ddl/xdo_09_ghost_property.go` |
| XDO-67 | ERROR | DDL `VARCHAR(n)` ↔ OpenAPI `maxLength` 누락/불일치 | `pkg/validate/openapi_ddl/xdo_67_max_length_varchar.go` |
| XDO-68 | ERROR | DDL `CHECK IN` ↔ OpenAPI `enum` 누락/불일치 | `pkg/validate/openapi_ddl/xdo_68_check_in_enum.go` |
| XDO-69 | ERROR | DDL `CHECK` 허용 값 ↔ OpenAPI `enum` 값 불일치 | `pkg/validate/openapi_ddl/xdo_69_check_values_enum.go` |
| XDO-70 | WARNING | OpenAPI `maxLength` 가 DDL `VARCHAR(n)` 초과 | `pkg/validate/openapi_ddl/xdo_70_max_length_exceeds_varchar.go` |
| XDO-75 | ERROR | OpenAPI optional + DDL `NOT NULL` + `DEFAULT` 없음 | `pkg/validate/openapi_ddl/xdo_75_optional_not_null_no_default.go` |
| XDO-76 | WARNING | OpenAPI required + DDL nullable | `pkg/validate/openapi_ddl/xdo_76_required_nullable.go` |
| XDO-77 | ERROR | DDL 컬럼 타입 ↔ OpenAPI field 타입 불일치 | `pkg/validate/openapi_ddl/xdo_77_column_type_mismatch.go` |
| XOD-10 | WARNING | DDL 컬럼이 OpenAPI response schema 에 누락 (커버리지) | `pkg/validate/openapi_ddl/xod_10_ddl_to_response.go` |

## G. OpenAPI ↔ SSaC

OpenAPI operation/response 와 SSaC 함수/`@response` 의 교차 정합성.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XOS-15 | ERROR | SSaC funcName 이 OpenAPI operationId 에 존재 | `pkg/validate/openapi_ssac/xos_15_func_name_op_id.go` |
| XOS-17 | ERROR | SSaC `@response` fields ↔ OpenAPI response schema 일치 | `pkg/validate/openapi_ssac/xos_17_response_fields.go` |
| XOS-19 | ERROR | shorthand `@response` ↔ OpenAPI response schema 일치 | `pkg/validate/openapi_ssac/xos_19_shorthand_response.go` |
| XOS-21 | ERROR | `@empty`/`@exists`/`@state`/`@auth`/`@call` ErrStatus 가 OpenAPI 에 정의됨 | `pkg/validate/openapi_ssac/xos_21_err_status_not_in_openapi.go` |
| XOS-22 | ERROR | SSaC `@response` 가 있는데 OpenAPI 2xx 응답 정의 없음 | `pkg/validate/openapi_ssac/xos_22_response_no_2xx.go` |
| XOS-66 | ERROR | SSaC 에서 사용되는 field 가 OpenAPI `required` 에 포함 | `pkg/validate/openapi_ssac/xos_66_used_fields_required.go` |
| XOS-67 | ERROR | `@response {key: value}` 의 값 type 이 OpenAPI response schema 기대 type 과 호환 | `pkg/validate/openapi_ssac/xos_67_response_field_type.go` |
| XSO-16 | ERROR | OpenAPI operationId 가 SSaC 함수로 사용됨 (커버리지) | `pkg/validate/openapi_ssac/xso_16_op_id_to_func.go` |
| XSO-18 | ERROR | OpenAPI response field 가 SSaC `@response` 에서 사용됨 (커버리지) | `pkg/validate/openapi_ssac/xso_18_response_field_used.go` |
| XSO-20 | ERROR | OpenAPI response field 가 shorthand `@response` 에서 사용됨 (커버리지) | `pkg/validate/openapi_ssac/xso_20_shorthand_field_used.go` |

## H. SSaC ↔ Func

SSaC `@call` 과 Func 스펙(Request/Response) 의 교차 정합성.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XFS-39 | ERROR | `@call` → Func 스펙 존재 | `pkg/validate/ssac_func/xfs_39_call_to_func_spec.go` |
| XFS-42 | ERROR | `@call` Inputs 개수 ↔ FuncRequest 필드 개수 | `pkg/validate/ssac_func/xfs_42_call_inputs_count.go` |
| XFS-43 | ERROR | `@call` Input field 가 FuncRequest 에 존재 | `pkg/validate/ssac_func/xfs_43_call_input_fields.go` |
| XFS-44 | ERROR | `@call` Input type ↔ FuncRequest field type 호환 | `pkg/validate/ssac_func/xfs_44_call_input_type.go` |
| XFS-45 | ERROR | `@result` 있는데 Func Response 없음 | `pkg/validate/ssac_func/xfs_45_call_result_missing.go` |
| XSF-46 | WARNING | `@result` 없는데 Func Response 있음 | `pkg/validate/ssac_func/xsf_46_call_result_ignored.go` |
| XSF-62 | WARNING | Func 스펙이 SSaC 에서 사용됨 (커버리지) | `pkg/validate/ssac_func/xsf_62_func_spec_used.go` |

## I. SSaC ↔ StateMachine

SSaC `@state` 와 Mermaid stateDiagram 의 교차 정합성.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XMS-24 | ERROR | `@state` DiagramID 가 stateDiagram 에 존재 | `pkg/validate/ssac_statemachine/xms_24_state_diagram_exists.go` |
| XMS-25 | ERROR | `@state` transition event 가 stateDiagram 에 정의 | `pkg/validate/ssac_statemachine/xms_25_state_event.go` |
| XSM-23 | ERROR | stateDiagram transition event 가 SSaC 함수로 존재 | `pkg/validate/ssac_statemachine/xsm_23_transition_to_func.go` |
| XSM-26 | WARNING | 상태 전이 참여 함수에 `@state` 선언 없음 | `pkg/validate/ssac_statemachine/xsm_26_missing_state_guard.go` |

## J. SSaC ↔ Rego

SSaC `@auth` 와 Rego allow 규칙의 `action:resource` 쌍 양방향 매칭.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XPS-28 | ERROR | SSaC `@auth (action:resource)` → Rego allow 정의 존재 | `pkg/validate/ssac_rego/xps_28_ssac_auth_to_rego.go` |
| XSP-29 | ERROR | Rego allow `(action:resource)` → SSaC `@auth` 에서 사용 | `pkg/validate/ssac_rego/xsp_29_rego_allow_to_ssac.go` |

## K. DDL ↔ Rego

Rego `@ownership` / role 이 DDL 테이블·컬럼·CHECK 제약을 정확히 참조하는지 검증.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XDP-31 | ERROR | `@ownership` table 이 DDL 에 존재 | `pkg/validate/ddl_rego/xdp_31_ownership_table.go` |
| XDP-32 | ERROR | `@ownership` column 이 DDL 에 존재 | `pkg/validate/ddl_rego/xdp_32_ownership_column.go` |
| XDP-33 | ERROR | `@ownership via` join table 이 DDL 에 존재 | `pkg/validate/ddl_rego/xdp_33_ownership_join_table.go` |
| XDP-34 | ERROR | `@ownership via` join column 이 DDL 에 존재 | `pkg/validate/ddl_rego/xdp_34_ownership_join_column.go` |
| XDP-65 | ERROR | Rego role 이 DDL `CHECK` 제약 허용 값에 포함 | `pkg/validate/ddl_rego/xdp_65_role_ddl_check.go` |

## L. Rego ↔ Manifest

Rego `input.claims` / role 참조와 `manifest.yaml` 의 claims/roles 정의의 양방향 매칭.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XNP-53 | ERROR | Rego `input.claims` 참조 값이 manifest claims 에 존재 | `pkg/validate/rego_manifest/xnp_53_input_claims_values.go` |
| XNP-63 | ERROR | Rego role 이 manifest roles 에 존재 | `pkg/validate/rego_manifest/xnp_63_role_manifest.go` |
| XPN-54 | WARNING | manifest claims 가 Rego 에서 참조됨 (커버리지) | `pkg/validate/rego_manifest/xpn_54_claims_to_rego.go` |
| XPN-64 | WARNING | manifest roles 가 Rego 에서 사용됨 (커버리지) | `pkg/validate/rego_manifest/xpn_64_roles_to_rego.go` |

## M. SSaC ↔ Manifest

SSaC 의 `currentUser` / `@publish` / `@subscribe` / JWT `@call` 과 manifest claims / queue / auth 설정의 교차 정합성.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XNS-48 | ERROR | `currentUser` 사용 → manifest `backend.auth.claims` 필수 | `pkg/validate/ssac_manifest/xns_48_current_user_claims.go` |
| XNS-49 | ERROR | `currentUser.<field>` 이 manifest claims 에 존재 | `pkg/validate/ssac_manifest/xns_49_current_user_field.go` |
| XNS-56 | ERROR | `@publish`/`@subscribe` 사용 → manifest `queue.backend` 설정 필수 | `pkg/validate/ssac_manifest/xns_56_queue_required.go` |
| XNS-57 | WARNING | `queue.backend: memory` + `@post/@put/@delete` 에 동반된 `@publish` 조합 (tx-bound 발행) — memory 백엔드는 `PublishTx` 미지원이므로 런타임 실패 | `pkg/validate/ssac_manifest/xns_57_memory_tx_publish.go` |
| XNS-73 | ERROR | JWT `@call` input field 가 manifest claims fields 에 존재 | `pkg/validate/ssac_manifest/xns_73_jwt_call_claims.go` |
| XAS-60 | ERROR | `@auth` input field 가 Authz `CheckRequest` 구조체에 존재 | `pkg/validate/ssac_authz/xas_60_auth_input_field.go` |

## N. OpenAPI ↔ Manifest

OpenAPI security scheme 과 manifest middleware 설정의 교차 정합성.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XNO-50 | ERROR | OpenAPI securityScheme → manifest middleware 존재 | `pkg/validate/openapi_manifest/xno_50_security_scheme_middleware.go` |
| XNO-52 | ERROR | endpoint security → manifest middleware 이름 존재 + middleware 블록 자체 존재 | `pkg/validate/openapi_manifest/xno_52_security_middleware.go` |
| XON-51 | ERROR | manifest middleware → OpenAPI securityScheme 존재 (커버리지) | `pkg/validate/openapi_manifest/xon_51_middleware_security_scheme.go` |
| SEC-04 | ERROR | `backend.http.overrides.<key>` 의 `<key>` 가 OpenAPI operationId 에 존재 | `pkg/validate/openapi_manifest/sec_04_http_overrides_operation_id.go` |
| SEC-101 | ERROR | generate-time: 생성된 main.go 가 request_id + error_envelope 미들웨어를 router 직후 순서대로 등록 (Phase004) | `pkg/generate/gogin/boot/collect_active_blocks.go` |

## O. SSaC ↔ sqlc

SSaC Input key 이름·대소문자·개수가 sqlc Params 와 정확히 일치하는지 검증.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XQS-14 | WARNING | SSaC input key case ↔ sqlc param case 일치 | `pkg/validate/ssac_sqlc/xqs_14_input_key_case.go` |
| XQS-15 | WARNING | `@call` SSaC input key 가 Go initialism 위반 | `pkg/validate/ssac_sqlc/xqs_15_input_key_initialism.go` |
| XQS-16 | ERROR | SSaC Input key 가 sqlc Params 에 존재 | `pkg/validate/ssac_sqlc/xqs_16_input_key_missing.go` |
| XQS-17 | ERROR | sqlc Params 필드가 SSaC Input 에서 전달됨 | `pkg/validate/ssac_sqlc/xqs_17_param_key_missing.go` |
| XQS-18 | ERROR | OpenAPI param 타입 ↔ sqlc param Go 타입 호환 | `pkg/validate/ssac_sqlc/xqs_18_param_type_mismatch.go` |

## P. SSaC ↔ DDL

SSaC `@result` / `@input` 이 DDL 테이블/컬럼과 일치하는지 양방향 검증.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XDS-12 | WARNING | `@result` 타입에 대응하는 sqlc row type 또는 DDL 테이블 매칭 | `pkg/validate/ssac_ddl/xds_12_result_no_ddl_table.go` |
| XSD-55 | ERROR | DDL 테이블이 SSaC `@model` 에서 참조됨 (커버리지) | `pkg/validate/ssac_ddl/xsd_55_ddl_to_model_ref.go` |

## R. Hurl 내부

| Rule ID | Level | Description | Source |
|---|---|---|---|
| H-1 | ERROR | `.feature` 파일 존재 (deprecated, Hurl `.hurl` 사용) | `pkg/validate/hurl/h_01_deprecated_feature.go` |
| H-2 | WARNING | `tests/` 디렉토리 비어있음 | `pkg/validate/hurl/h_02_empty_tests_dir.go` |

## R2. Hurl ↔ OpenAPI

`specs/tests/` 의 Hurl 파일은 전부 사용자 소유. yongol 은 Hurl 파일을 생성하지 않으며, `generate` 는 `specs/tests/` → `arts/tests/` 를 그대로 미러링한다. 본 규칙들은 작성된 Hurl 과 OpenAPI SSOT 사이의 드리프트를 validate 단계에서 잡는다.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XOH-01 | ERROR | Hurl URL + method 가 OpenAPI operation 으로 선언됨 | `pkg/validate/hurl_openapi/xoh_01_url_method.go` |
| XOH-02 | ERROR | Hurl `HTTP <status>` 가 OpenAPI responses 에 선언됨 | `pkg/validate/hurl_openapi/xoh_02_status_declared.go` |
| XOH-03 | ERROR | Hurl 요청 JSON 필드가 OpenAPI request schema 에 존재 | `pkg/validate/hurl_openapi/xoh_03_request_field_in_schema.go` |
| XOH-04 | ERROR | Hurl assert jsonpath 이 OpenAPI response schema 에 도달 가능 | `pkg/validate/hurl_openapi/xoh_04_assert_path_in_schema.go` |
| XOH-08 | ERROR | Hurl capture jsonpath 이 OpenAPI response schema 에 존재 | `pkg/validate/hurl_openapi/xoh_08_capture_path_in_schema.go` |
| XOH-09 | WARNING | Hurl [Captures] 로 저장한 변수가 같은 파일에서 사용됨 | `pkg/validate/hurl_openapi/xoh_09_unused_capture.go` |

## R3. Hurl ↔ State Machine

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XOH-05 | WARNING | Hurl 호출 순서가 state machine 전이 규칙을 준수 | `pkg/validate/hurl_statemachine/xoh_05_state_transition_order.go` |

## R4. Hurl ↔ Manifest

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XOH-06 | WARNING | 보호 구간 Hurl 호출 전에 인증 스텝 선행 | `pkg/validate/hurl_manifest/xoh_06_auth_precondition.go` |
| XOH-07 | WARNING | cookie 모드 mutation 요청에 `X-CSRF-Token` 헤더 포함 | `pkg/validate/hurl_manifest/xoh_07_csrf_on_mutation.go` |

## R5. State Machine / Rego / Func

| Rule ID | Level | Description | Source |
|---|---|---|---|
| ST-1 | ERROR | Mermaid stateDiagram 파싱 검증 | `pkg/validate/statemachine/st_01_parse.go` |
| P-1 | ERROR | Rego 정책 파싱 검증 | `pkg/validate/rego/p_01_parse.go` |
| XPP-30 | ERROR | Rego 가 `resource_owner` 참조하는데 `@ownership` 어노테이션 없음 | `pkg/validate/rego/xpp_30_ownership_no_annotation.go` |
| F-1 | WARNING | Func 이름이 built-in 패키지명(`auth`/`session`/`cache`/`file`)과 충돌 | `pkg/validate/funcspec/f_01_builtin_override.go` |
| XFF-40 | ERROR | Func 본체 미구현 (`panic("TODO")` / `// TODO` / 빈 본체) | `pkg/validate/funcspec/xff_40_func_body_todo.go` |
| XFF-41 | ERROR | Func 본체에 I/O 패키지 (`database/sql`, `net/http`, `grpc` 등) import 금지 | `pkg/validate/funcspec/xff_41_func_forbidden_import.go` |
| XDM-27 | ERROR | `@state` field 가 DDL 컬럼에 존재 | `pkg/validate/ddl_statemachine/xdm_27_state_field_column.go` |
| XDM-28 | ERROR | stateDiagram `[*] → X` 초기 전이 ↔ DDL `DEFAULT 'X'` 일치 | `pkg/validate/ddl_statemachine/xdm_28_default_initial_state.go` |

## Q. TSX (React 프론트엔드)

React `.tsx` (frontend/**/*.tsx) 자체 정합성.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| T-1 | WARNING | `@/components/` 또는 상대 경로 import 대상 파일이 존재 | `pkg/validate/tsx/t_01_component_file.go` |

## Q2. TSX ↔ OpenAPI

TSX 의 `apiClient.<op>()` 호출이 OpenAPI 계약과 일치하는지 검증. 단방향(TSX → OpenAPI). OpenAPI 미소비 operationId 는 의도적으로 보고하지 않는다(모바일 / CLI / 파트너 등 다양한 소비처 고려).

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XOT-1 | ERROR | `apiClient.<op>()` 의 `<op>` 가 OpenAPI operationId 집합에 존재 | `pkg/validate/tsx_openapi/xot_01_operation_id.go` |
| XOT-2 | ERROR | apiClient 호출의 path/query 인자 객체 키가 OpenAPI parameters 에 존재 | `pkg/validate/tsx_openapi/xot_02_parameter_match.go` |
| XOT-3 | WARNING | `useForm().register('x')` 필드가 해당 페이지 mutation 의 OpenAPI request body schema 에 존재 | `pkg/validate/tsx_openapi/xot_03_form_field.go` |

---

## S. Preserve (`PRV-*`) — preserved 파일 계약 / 런타임 안전 가드

`//ff:checked hash` 가 어긋난 (= 사용자가 수정한) `.go` 파일만 대상. `yongol validate <specs> <arts>` 처럼 arts 디렉토리를 같이 넘길 때만 실행된다.

### Contract drift (PRV-01 ~ PRV-09)

| Rule ID | Level | Description | Source |
|---|---|---|---|
| PRV-01 | ERROR | preserved 파일의 함수 시그니처가 SSOT 기대값과 drift | `pkg/validate/contract/prv_01_signature_drift.go` |
| PRV-02 | ERROR | preserved 파일이 SSOT 에 존재하지 않는 sqlc query / @call / DDL field 참조 | `pkg/validate/contract/prv_02_external_symbol_drift.go` |

### Runtime safety guards (PRV-10 ~ PRV-19)

| Rule ID | Level | Description | Source |
|---|---|---|---|
| PRV-10 | ERROR | preserved 파일에 허용되지 않은 `panic(` 존재 (init() / `// nolint:panic` 제외) | `pkg/validate/contract/prv_10_preserved_panic.go` |
| PRV-11 | ERROR | preserved 파일의 `ctx.Value("currentUser").(T)` 가 comma-ok 형태 아님 | `pkg/validate/contract/prv_11_preserved_currentuser_assertion.go` |
| PRV-12 | ERROR | preserved 파일에서 `json.Unmarshal` / `yaml.Unmarshal` 에러 무시 | `pkg/validate/contract/prv_12_preserved_unmarshal_err.go` |
| PRV-13 | ERROR | preserved 파일에서 `sql.Rows.Scan` / `sql.Row.Scan` 에러 무시 | `pkg/validate/contract/prv_13_preserved_scan_err.go` |
| PRV-14 | ERROR | preserved 파일에서 slice 첫 요소 접근(`x[0]`) 전 `len` 가드 없음 | `pkg/validate/contract/prv_14_preserved_slice_bounds.go` |
| PRV-15 | ERROR | preserved 파일에서 `m[k].Field` 같은 inline selector 접근 (comma-ok 가드 없음) | `pkg/validate/contract/prv_15_preserved_map_access.go` |
| PRV-16 | ERROR | preserved 파일에서 `Get*()`/`Find*()` 반환값을 즉시 필드 접근 (nil 가드 없음) | `pkg/validate/contract/prv_16_preserved_nil_deref.go` |
| PRV-17 | ERROR | preserved 파일에서 `os.Open` / `db.Query` / `http.Get` 등 리소스 획득 후 `defer Close` 누락 | `pkg/validate/contract/prv_17_preserved_missing_close.go` |

Allowlist:
- `init()` 함수 내부 — PRV-10 면제
- 위 라인 (또는 본 라인) 에 `// nolint:panic` — PRV-10 면제
- 위 라인 (또는 본 라인) 에 `// nolint:prv-NN` — 해당 규칙 면제

---

## T. Migration (`MIG-*`) — DDL 자동 마이그레이션

`yongol generate` 의 DDL 마이그레이션 단계(`pkg/generate/migration/`)에서 수집되는 힌트 · 스냅샷 · 파괴성 관련 규칙. DDL 주석 힌트(`-- @rename`, `-- @cast`, `-- @backfill`, `-- @data_migration`, `-- @allow_destructive`) 문법은 `manual-for-ai.md` 의 "DDL 힌트" 섹션과 `docs/MIGRATION.md` 참조.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| MIG-001 | ERROR | `@rename from=...` 의 from 이 이전 스냅샷에 없거나 to 가 현재 DDL 에 없음 (mismatch) | `pkg/validate/migration/mig_001_rename_mismatch.go` |
| MIG-002 | ERROR | NOT NULL 컬럼 추가 + `@backfill default=...` 힌트 없음 + DEFAULT 절 없음 → 기존 row 가 NULL 인 위험으로 emit 차단 | `pkg/validate/migration/mig_002_not_null_without_backfill.go` |
| MIG-003 | ERROR | `@data_migration file=<path>` 의 sidecar SQL 파일이 실제 존재하지 않음 | `pkg/validate/migration/mig_003_data_migration_missing.go` |
| MIG-004 | WARNING | DROP TABLE / DROP COLUMN 이 발생하는데 대상 테이블에 `@allow_destructive` 없음 (의도 재확인 권장) | `pkg/validate/migration/mig_004_destructive_without_allow.go` |
| MIG-005 | WARNING | 위험한 타입 변경(`INTEGER↔TEXT`, `VARCHAR(N)` 축소 등)에 `@cast using=<expr>` 힌트 없음 | `pkg/validate/migration/mig_005_cast_missing.go` |
| MIG-006 | ERROR | `specs/db/.generated_schema.sql` 의 `-- YONGOL_SCHEMA_HASH:` 헤더가 본문 sha256 과 일치하지 않음 (사용자 수정 = drift) | `pkg/validate/migration/mig_006_snapshot_drift.go` |

---

## Deprecated

코드에서 이미 제거되었거나 Phase002 에서 제거 예정인 규칙. 현재 `yongol validate` 의 일반 카탈로그에서 제외된다.

### Phase002 에서 제거 예정 — OpenAPI `x-pagination` / `x-sort` / `x-filter` / `x-include` 폐기 후속 정리

| Rule ID | 이전 설명 | 폐기 사유 |
|---|---|---|
| XDO-1 | `x-sort` column → DDL 참조 | `x-sort` 확장 폐기. 표준 OpenAPI query parameters + `sort_by` enum 로 대체 |
| XDO-2 | `x-sort` column 에 인덱스 없음 (WARNING) | 동일 |
| XDO-3 | `x-filter` column → DDL 참조 | `x-filter` 확장 폐기. 표준 OpenAPI 컬럼별 query parameter 로 대체 |
| XDO-5 | `x-include` target table → DDL 참조 | `x-include` 확장 폐기 |
| XDO-6 | `x-include` FK 제약 누락 (WARNING) | 동일 |
| XDO-8 | cursor sort 기본 컬럼이 UNIQUE 아님 | `x-pagination` cursor 모드 폐기 |
| XOO-4 | `x-include` 형식 오류 | `x-include` 확장 폐기 |
| XOO-7 | cursor 모드 `x-sort.allowed` 2개 이상 | `x-pagination` cursor 모드 폐기 |
| TM-9 | STML `data-paginate` + OpenAPI `x-pagination` 필수 | `x-pagination` 폐기. STML 은 표준 query parameter 참조 |
| TM-10 | STML `data-sort` ref → OpenAPI `x-sort` 허용 컬럼 | `x-sort` 폐기 |
| TM-11 | STML `data-filter` ref → OpenAPI `x-filter` 허용 컬럼 | `x-filter` 폐기 |

### Phase008 — STML (HTML+data-*) 전면 폐기 (TSX SSOT 로 대체)

| Rule ID | 이전 설명 | 폐기 사유 |
|---|---|---|
| TM-1 | top-level `data-fetch` operationId → OpenAPI (GET) 존재 | STML 폐기. `XOT-1` 로 대체 (apiClient.<op>() 기반) |
| TM-2 | `data-action` operationId → OpenAPI 존재 | STML 폐기. `XOT-1` 로 통합 |
| TM-3 | 중첩 `data-fetch` operationId → OpenAPI 존재 | STML 폐기. `XOT-1` 로 통합 |
| TM-4 | `data-param` 참조가 OpenAPI parameter 에 존재 | STML 폐기. `XOT-2` 로 대체 |
| TM-5 | `data-action` payload 필드가 OpenAPI request 에 존재 | STML 폐기. `XOT-3` 로 대체 (WARNING, TypeScript 가 이중 방어) |
| TM-6 | `data-bind` 필드가 OpenAPI response 에 존재 | STML 폐기. TypeScript + `openapi-typescript` 타입 체커가 담당 |
| TM-7 | `data-each` 필드 타입이 배열이어야 함 | STML 폐기. TypeScript `Array<T>` 가 담당 |
| TM-8 | `data-bind` 필드 미발견 (custom.ts export 시 면제) | STML 폐기. TypeScript 담당 |
| TM-12 | `data-component` 에 참조된 컴포넌트 파일이 존재 | STML 폐기. `T-1` 로 대체 (import 기반) |

### 이미 코드에서 폐기 처리된 규칙

| Rule ID | 이전 설명 | 폐기 사유 |
|---|---|---|
| XOO-73 | `x-pagination` 필수 query params | 주석으로 폐기 표시, `x-pagination` 전면 폐기 |
| XOO-74 | `x-pagination` 필수 response fields | 동일 |
| XOS-68 | OpenAPI `x-pagination` 커버리지 | 표준 parameters + `XSO-18` / `XQS-16` / `XQS-17` 로 대체 |
| S-52 | `QueryUsageMismatch` — OpenAPI `x-pagination` ↔ SSaC query 불일치 | `XOS-68` 로 일시 대체 후, `x-pagination` 폐기와 함께 제거 |
| S-53 | SSaC query usage 커버리지 | 동일 |
| S-54 | `Page[T]`/`Cursor[T]` wrapper → `x-pagination` 필수 | Page/Cursor wrapper 타입 폐기 |
| S-55 | `x-pagination` 옵션 매칭 | wrapper + `x-pagination` 폐기 |
| S-56 | `x-pagination` 옵션 매칭 (보조) | 동일 |
| XDS-13 | SSaC input 이 DDL 컬럼에 없음 (WARNING) | Replaced by XQS-14/16 — sqlc Params 기준이 더 엄밀 |
| XDS-14 | SSaC CRUD Input key 가 sqlc Go 필드명 (PascalCase) 과 불일치 (ERROR) | Replaced by XQS-14/15/16 — 실제 sqlc Params 집합 기준이 더 엄밀 |
| M-1 | `model/` 디렉토리 및 `*.go` 파일 존재 | `model/` SSOT 및 `@dto` 전면 폐기 — sqlc 합성 row type 이 모델 역할 수행 (Phase001 ModelSSOTDeprecation) |
| M-2 | `model/*.go` struct 타입이 `@dto` 또는 DDL 테이블 중 하나와 매칭됨 | 동일 |
| XNS-77 | manifest `auth.claims` 있는데 SSaC 에 `auth.IssueToken` 호출 없음 (WARNING) | Login 누락은 true positive 극히 희박 + 검증자-전용 마이크로서비스에서 false positive. 런타임 첫 로그인 시도로 즉시 드러나므로 정적 검증 가치 낮음 (`plans/deprecated/Phase005-RemoveXNS77.md`) |
| SEC-03 | `backend.rate_limit.endpoints.<key>` 의 `<key>` 가 OpenAPI operationId 에 존재 (ERROR) | 애플리케이션 계층 rate_limit 자체 폐기 — CDN/WAF/Gateway 계층 책임으로 이관. `FixedRateLimit` 하드코딩 가드(/auth/refresh)만 유지 (`plans/deprecated/Phase006-DeprecateAppLayerRateLimit.md`) |
| XOH-35 | Hurl path → OpenAPI path 존재 | XOH-01 로 흡수 (2026-04-24, hurl_openapi 재편). path + method 를 단일 판정. |
| XOH-36 | Hurl method → OpenAPI method 존재 | XOH-01 로 흡수 (2026-04-24). |
| XOH-37 | Hurl status code → OpenAPI responses | XOH-02 로 이관 (2026-04-24). 심각도 WARNING → ERROR. |
| `pkg/generate/hurl/` | 자동 smoke/scenario 생성 엔진 | 2026-04-24 디렉토리 통째 삭제 (`plans/gen/hurl/Phase001`). Hurl 은 사용자 소유, yongol 은 `specs/tests/` → `arts/tests/` 미러링만 수행. |

---

## 참고

- 규칙 설계 철학, Toulmin defeats 그래프, Ground 매핑: `pkg/validate/README.md`
- 각 카테고리 세부 설명: `pkg/validate/<domain>/README.md`
- SSOT 문법과 교차 검증 규칙 요약: `manual-for-ai.md` → "Cross-Validation Rules Catalog" 섹션
