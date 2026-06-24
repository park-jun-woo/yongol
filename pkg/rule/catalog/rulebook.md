# yongol Validation Rulebook

Complete catalog of rules executed by `yongol validate`. The manual (`manual-for-ai.md`) keeps only the summary table; detailed definitions live in this document.

> This document is organized around **file facts**, not contracts between SSOTs. Most rules are implemented as Go source under `pkg/validate/<domain>/`. Generate-time rules that require the `<artifacts>` path may reside under `pkg/generate/<stack>/` (for example, `Q-10`).

## Rule ID Prefix Legend

### Single SSOT

| Prefix | SSOT |
|---|---|
| `C-` | manifest.yaml |
| `D-`, `XDD-` | DDL (PostgreSQL CREATE TABLE + sqlc queries) |
| `O-`, `XOO-` | OpenAPI |
| `S-`, `XSS-` | SSaC (Service-Sequence as Code) |
| `ST-` | Mermaid stateDiagram |
| `P-`, `XPP-` | OPA Rego |
| `F-`, `XFF-` | Func spec (Go AST) |
| `H-` | Hurl scenarios |
| `Q-` | sqlc queries |
| `CORS-` | manifest CORS block |
| `V-` | DESIGN.md (design token spec) |
| `TM-` | STML ↔ OpenAPI cross-validation |
| `XMO-` | OpenAPI ↔ STML/frontend coverage (operationId 소비 강제 + 도메인 경계) |
| `FT-`, `XFO-`/`XOF-`, `XFD-`, `XFS-` | features.yaml (catalog; cross with OpenAPI/DDL/stateDiagram — uses SSOT code `F`) |
| `MIG-` | DDL migration diff |
| `OBS-` | manifest observability (metrics / tracing) |
| `SEC-` | manifest security (auth mode / CSRF / headers) |
| `PRV-` | Preserve contract (generated-file edits) |
| `XNC-` | manifest backend wiring ↔ DDL + sqlc |
| `XOE-` | OpenAPI ErrorResponse schema |
| `INI-` | Init check (project initialization) |
| `M-` | (retired) `model/` SSOT rules — superseded by sqlc row type. See Deprecated section. |

### Cross SSOT — `X<target><source>-<N>`

`<target>` = the SSOT referenced by the Lookup key (ground truth), `<source>` = the SSOT making the claim.

| SSOT | Code | SSOT | Code |
|---|---|---|---|
| OpenAPI | `O` | DDL | `D` |
| SSaC | `S` | StateMachine | `M` |
| Rego | `P` | Manifest | `N` |
| Hurl | `H` | Func | `F` |
| Authz | `A` | sqlc | `Q` |
| STML | `T` | Design | `V` |

Example: SSaC → OpenAPI (SSaC is the claim, OpenAPI is the ground truth) → `XOS-`.

> **Known prefix collision**: Func and features.yaml both use SSOT code `F`, so the `XFS-` prefix is shared by two categories — SSaC ↔ Func (`XFS-39` ~ `XFS-73`, section H) and Features ↔ stateDiagram (`XFS-01`, section Y). Rule IDs themselves are unique; resolve the category by the rule number.
## Level

| Level | Behavior | Meaning |
|---|---|---|
| `ERROR` | `yongol validate` fails → exit code ≠ 0 | Consistency violation. Must be fixed. |
| `WARNING` | Does not fail; reported only | Design intent should be reconfirmed. |

## Source

The `Source` column of each rule row is a Go file path relative to the repo root. Example: `pkg/validate/ssac/s_27_var_declared.go`.

## Rule count

**This catalog is the single source of truth for the rule set.** The official
total is the count of distinct rule IDs in the tables below — **398 rules across
60 prefixes**. This includes 27 retired rules (Deprecated section); the
**active** subset (rows in the non-deprecated tables) is **371**.

Counting note: a rule's ID is emitted either as a `[ID]` literal in the
diagnostic message **or** via a `RuleID:` field / builder. A naive
`grep '\[ID\]'` over `pkg/` therefore undercounts. Use this catalog — not source
grep — for the official total.

---

## INI. Init Check

Project initialization check — verifies that `specs/.yongol` exists before other validations run.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| INI-01 | WARNING | `.yongol` not found in specs directory — project not initialized | `pkg/validate/initcheck/ini_01_require_init.go` |

## A. SSaC Internal

SSaC self-consistency — required fields, variable flow, model references, @subscribe constraints, and pub/sub pair matching.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| S-1 | ERROR | `@get` Model field required | `pkg/validate/ssac/s_01_get_model.go` |
| S-2 | ERROR | `@get` Result field required | `pkg/validate/ssac/s_02_get_result.go` |
| S-3 | ERROR | `@post` Model field required | `pkg/validate/ssac/s_03_post_model.go` |
| S-4 | ERROR | `@post` Result field required | `pkg/validate/ssac/s_04_post_result.go` |
| S-5 | ERROR | `@post` Inputs field required | `pkg/validate/ssac/s_05_post_inputs.go` |
| S-6 | ERROR | `@put` Model field required | `pkg/validate/ssac/s_06_put_model.go` |
| S-7 | ERROR | `@put` Result field required | `pkg/validate/ssac/s_07_put_result.go` |
| S-8 | ERROR | `@put` Inputs field required | `pkg/validate/ssac/s_08_put_inputs.go` |
| S-9 | ERROR | `@delete` Model field required | `pkg/validate/ssac/s_09_delete_model.go` |
| S-10 | ERROR | `@delete` Result field required | `pkg/validate/ssac/s_10_delete_result.go` |
| S-11 | WARNING | `@delete` has no Inputs | `pkg/validate/ssac/s_11_delete_no_inputs.go` |
| S-12 | ERROR | `@empty` Target field required | `pkg/validate/ssac/s_12_empty_target.go` |
| S-13 | ERROR | `@empty` Message field required | `pkg/validate/ssac/s_13_empty_message.go` |
| S-14 | ERROR | `@state` DiagramID field required | `pkg/validate/ssac/s_14_state_diagram_id.go` |
| S-15 | ERROR | `@state` Inputs field required | `pkg/validate/ssac/s_15_state_inputs.go` |
| S-16 | ERROR | `@state` Transition field required | `pkg/validate/ssac/s_16_state_transition.go` |
| S-17 | ERROR | `@state` Message field required | `pkg/validate/ssac/s_17_state_message.go` |
| S-18 | ERROR | `@auth` Action field required | `pkg/validate/ssac/s_18_auth_action.go` |
| S-19 | ERROR | `@auth` Resource field required | `pkg/validate/ssac/s_19_auth_resource.go` |
| S-20 | ERROR | `@auth` Message field required | `pkg/validate/ssac/s_20_auth_message.go` |
| S-21 | ERROR | `@call` Model field required | `pkg/validate/ssac/s_21_call_model.go` |
| S-23 | ERROR | `@publish` Topic field required | `pkg/validate/ssac/s_23_publish_topic.go` |
| S-24 | ERROR | `@publish` Payload field required | `pkg/validate/ssac/s_24_publish_payload.go` |
| S-25 | ERROR | Unknown sequence type | `pkg/validate/ssac/s_25_unknown_seq_type.go` |
| S-26 | ERROR | `@call` value must be in `Model.Method` form | `pkg/validate/ssac/s_26_dot_method.go` |
| S-27 | ERROR | Variables must be declared before use (general identifiers) | `pkg/validate/ssac/s_27_var_declared.go` |
| S-28 | ERROR | Variables used in `@empty` Target must be declared | `pkg/validate/ssac/s_28_target_declared.go` |
| S-29 | ERROR | Variables used in Inputs values must be declared | `pkg/validate/ssac/s_29_input_declared.go` |
| S-30 | ERROR | Variables used in `@empty` Message must be declared | `pkg/validate/ssac/s_30_message_declared.go` |
| S-31 | ERROR | `ssac.config*` prefix is forbidden | `pkg/validate/ssac/s_31_config_prefix_forbidden.go` |
| S-32 | ERROR | Reserved words are forbidden in `@publish` topic | `pkg/validate/ssac/s_32_publish_forbidden.go` |
| S-33 | ERROR | Use of reserved source identifiers is forbidden | `pkg/validate/ssac/s_33_reserved_source.go` |
| S-34 | ERROR | Go reserved words must not be used as variable names | `pkg/validate/ssac/s_34_go_reserved_word.go` |
| S-35 | ERROR | Go reserved words must not be used as Model names | `pkg/validate/ssac/s_35_go_reserved_word_model.go` |
| S-36 | WARNING | `@response` used after `@put`/`@delete` without refresh (stale) | `pkg/validate/ssac/s_36_check_response_stale.go` |
| S-37 | WARNING | FK-referencing `@get` should be followed by an `@empty` guard (단일 Model 조회만 적용; scalar/배열 제외) | `pkg/validate/ssac/s_37_fk_reference_guard.go` |
| S-38 | ERROR | HTTP-only Inputs are forbidden in `@subscribe` | `pkg/validate/ssac/s_38_subscribe_no_http_inputs.go` |
| S-39 | ERROR | `@subscribe` Message is required | `pkg/validate/ssac/s_39_subscribe_message_required.go` |
| S-40 | ERROR | `@subscribe` handler must take a single parameter | `pkg/validate/ssac/s_40_subscribe_single_param.go` |
| S-41 | ERROR | `currentUser` is forbidden in `@subscribe` | `pkg/validate/ssac/s_41_subscribe_no_currentuser.go` |
| S-42 | ERROR | `request.*` is forbidden in `@subscribe` | `pkg/validate/ssac/s_42_subscribe_forbidden_request.go` |
| S-43 | ERROR | `query.*` is forbidden in `@subscribe` | `pkg/validate/ssac/s_43_subscribe_forbidden_query.go` |
| S-44 | ERROR | `message.*` is forbidden in HTTP functions | `pkg/validate/ssac/s_44_http_forbidden_message.go` |
| S-45 | ERROR | `@response` is forbidden in `@subscribe` | `pkg/validate/ssac/s_45_subscribe_no_response.go` |
| S-46 | ERROR | Result type must start with an uppercase letter | `pkg/validate/ssac/s_46_uppercase_start.go` |
| S-47 | ERROR | Package prefix is forbidden in `@model` | `pkg/validate/ssac/s_47_no_dot_prefix.go` |
| S-48 | ERROR | `@call` Model must exist in SymbolTable | `pkg/validate/ssac/s_48_symbol_table_model.go` |
| S-49 | ERROR | `Model.Method` — Method must exist in SymbolTable | `pkg/validate/ssac/s_49_symbol_table_method.go` |
| S-50 | ERROR | SSaC input must exist as an OpenAPI request field | `pkg/validate/ssac/s_50_openapi_request.go` |
| S-51 | WARNING | OpenAPI request field must be used in SSaC | `pkg/validate/ssac/s_51_request_usage.go` |
| S-57 | ERROR | `@call` input type must match FuncRequest field type | `pkg/validate/ssac/s_57_func_request_type.go` |
| S-58 | ERROR | Use of IANA-unregistered HTTP status codes is forbidden | `pkg/validate/ssac/s_58_invalid_err_status.go` |
| S-59 | ERROR | Dotted field reference format validation | `pkg/validate/ssac/s_59_dotted_field.go` |
| S-60 | ERROR | `request.<field>` references in SSaC must exist case-exactly in the OpenAPI operation's request schema | `pkg/validate/ssac/s_60_request_field_exact.go` |
| S-61 | ERROR | Result variable names must not collide with codegen reserved names (`server`, `ctx`, `err`, etc.) | `pkg/validate/ssac/s_61_codegen_reserved_var.go` |
| S-62 | ERROR | Result variable is unreferenced in subsequent sequences | `pkg/validate/ssac/s_62_unused_result_var.go` |
| S-63 | WARNING | `@get []T` list endpoint has no pagination params and no `// @no-pagination` | `pkg/validate/ssac/s_63_list_no_pagination.go` |
| S-64 | ERROR | `@empty` / `@exists` Target must reference a Model (struct), not a scalar field | `pkg/validate/ssac/s_64_empty_exists_model_only.go` |
| S-67 | ERROR | `@eval` Func signature must be `func(req T) bool` | `pkg/validate/ssac/s_67_eval_func_signature.go` |
| S-68 | ERROR | `@eval` requires an explicit STATUS code (no default) | `pkg/validate/ssac/s_68_eval_status_required.go` |
| S-69 | ERROR | `@eval` Func must exist in Func Spec or built-in | `pkg/validate/ssac/s_69_eval_func_exists.go` |
| S-70 | ERROR | `@post` / `@put` Inputs value must not be a standalone reserved source (`currentUser`, `request`, `query`, `message`); use dotted form. `@call` exempt | `pkg/validate/ssac/s_70_post_put_blob_input_forbidden.go` |
| S-71 | ERROR | SSaC Input 값의 변수 prefix 가 해당 시퀀스 지점에서 유효한 scope 에 없으면 ERROR | `pkg/validate/ssac/s_71_unknown_variable.go` |
| S-72 | ERROR | `@call`/`@eval` 참조 패키지에 대한 SSaC import 선언 필수 | `pkg/validate/ssac/s_72_call_eval_import_required.go` |
| S-73 | ERROR | SSaC import는 full Go import path 필수 (bare name 거부) | `pkg/validate/ssac/s_73_import_must_be_full_path.go` |
| S-74 | ERROR | SSaC 함수에 어노테이션이 하나도 없으면 ERROR (빈 함수 차단) | `pkg/parser/ssac/parse_func_decl.go` |
| XSS-11 | WARNING | `@result` type is plural | `pkg/validate/ssac/xss_11_plural_result_type.go` |
| XSS-38 | ERROR | `@call` function name starts with a lowercase letter (uppercase recommended) | `pkg/validate/ssac/xss_38_call_func_lowercase.go` |
| XSS-47 | WARNING | `@call` argument source variable is undefined | `pkg/validate/ssac/xss_47_call_source_var_undefined.go` |
| XSS-57 | ERROR | `@publish` topic must have a matching `@subscribe` | `pkg/validate/ssac/xss_57_publish_to_subscribe.go` |
| XSS-58 | ERROR | `@subscribe` topic must have a matching `@publish` | `pkg/validate/ssac/xss_58_subscribe_to_publish.go` |
| XSS-59 | ERROR | `@subscribe` message fields must match `@publish` payload | `pkg/validate/ssac/xss_59_subscribe_fields.go` |
| XSS-60 | WARNING | `@subscribe` message field type must be compatible with `@publish` payload inferred type | `pkg/validate/ssac/xss_60_subscribe_field_types.go` |

## B. Manifest

`manifest.yaml` loading and base schema validation.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| C-2 | ERROR | `apiVersion` must be `yongol/v1` | `pkg/validate/manifest/c_02_api_version.go` |
| C-3 | ERROR | `kind` must be `Project` | `pkg/validate/manifest/c_03_kind.go` |
| C-4 | ERROR | `metadata.name` is required (empty value forbidden) | `pkg/validate/manifest/c_04_metadata_name.go` |
| C-5 | ERROR | `backend.module` is required (empty value forbidden) | `pkg/validate/manifest/c_05_backend_module.go` |
| C-6 | ERROR | `backend.auth` is required — yongol does not support auth-free backends (use a static site generator + CDN for public dynamic content) | `pkg/validate/manifest/c_06_backend_auth_required.go` |
| C-7 | ERROR | `backend.auth` requires `backend.rate_limit` — brute-force defense is mandatory | `pkg/validate/manifest/c_07_auth_requires_rate_limit.go` |
| C-8 | ERROR | `backend.rate_limit` requires a `Login` entry (primary brute-force target) | `pkg/validate/manifest/c_08_rate_limit_login_required.go` |
| C-9 | ERROR | Unsupported `backend.lang` + `backend.framework` combination | `pkg/validate/manifest/c_09_backend_lang_framework.go` |
| C-10 | ERROR | Every `backend.rate_limit` entry must have `rate` >= 1 and a `period` parseable by `time.ParseDuration` (blocks zero-value entries codegen would silently drop) | `pkg/validate/manifest/c_10_rate_limit_value_valid.go` |
| C-11 | WARNING | `backend.rate_limit` entry keyed by `ip` (or key unset → default `ip`) with `backend.http.trusted_proxies` unset — behind a reverse proxy `c.ClientIP()` always returns the proxy address, collapsing the IP-keyed limiter onto one key (BUG-117); directly exposed deployments may ignore | `pkg/validate/manifest/c_11_ipkey_requires_proxy.go` |
| C-12 | ERROR | Multi-domain: each entry under the top-level `domains` key must declare an `openapi` path — a domain without an API contract silently vanishes from the generated backend | `pkg/validate/manifest/c_12_domain_openapi_required.go` |
| C-13 | ERROR | Multi-domain: each `domains.<name>` must declare a `frontend` directory — each domain is an independent app whose STML source location is its own frontend dir | `pkg/validate/manifest/c_13_domain_frontend_required.go` |
| C-14 | ERROR | Multi-domain: two domains must not declare the same `route_prefix` — distinct domains must occupy distinct URL namespaces or their Gin route groups collide (empty prefixes ignored) | `pkg/validate/manifest/c_14_domain_route_prefix_unique.go` |
| C-15 | ERROR | Multi-domain: a domain's `auth_mode` override must be one of `cookie` / `bearer` / `hybrid` (reuses SEC-403's `validAuthModes`; empty = inherit `backend.auth.mode`) | `pkg/validate/manifest/c_15_domain_auth_mode_enum.go` |
| C-16 | WARNING | Multi-domain: a domain's `frontend` path resolving to the single-site STML root (`frontend`) collides with the legacy location — move the domain's pages into a dedicated subdir (e.g. `frontend/<name>`) | `pkg/validate/manifest/c_16_domain_frontend_conflict.go` |
| C-17 | ERROR | Multi-domain: a `domains:` block declaring exactly 1 domain is rejected — multi-site machinery only earns its complexity at ≥2 domains; a single domain should be a plain single-site project (top-level `openapi` + `frontend`) | `pkg/validate/manifest/c_17_domain_minimum_two.go` |
| CORS-01 | ERROR | `allow_origins=["*"]` combined with `allow_credentials=true` is forbidden | `pkg/validate/manifest/cors_01_wildcard_credentials.go` |
| OBS-001 | ERROR | `backend.observability.metrics.path` must be an absolute path starting with `/` | `pkg/validate/manifest/obs_01_metrics_path.go` |
| OBS-002 | ERROR | `backend.observability.metrics.path` must not collide with an OpenAPI path | `pkg/validate/manifest/obs_02_metrics_path_not_openapi.go` |
| OBS-003 | ERROR | `backend.observability.tracing.exporter` must be one of `otlp`/`stdout`/`noop` (when enabled=true) | `pkg/validate/manifest/obs_03_tracing_exporter.go` |
| OBS-004 | ERROR | `backend.observability.tracing.sample_rate` must be within `[0.0, 1.0]` (when enabled=true) | `pkg/validate/manifest/obs_04_tracing_sample_rate.go` |
| SEC-201 | ERROR | `backend.auth.mode=cookie\|hybrid` combined with `csrf.enabled=false` is forbidden (leaves CSRF attack surface exposed) | `pkg/validate/manifest/sec_201_cookie_without_csrf.go` |
| SEC-202 | WARNING | `backend.auth.mode=bearer` combined with `csrf.enabled=false` removes the runtime CSRF gate — the generated `BACKEND_AUTH_MODE` switch can flip the build to cookie/hybrid auth with no CSRF defense (BUG-116) | `pkg/validate/manifest/sec_202_runtime_mode_csrf.go` |
| SEC-301 | WARNING | `backend.security_headers.csp.directives.default-src` containing `*` or `'unsafe-eval'` weakens CSP protection | `pkg/validate/manifest/sec_301_csp_permissive.go` |
| SEC-302 | WARNING | `backend.security_headers.hsts.max_age < 15552000` (180 days) fails the HSTS preload minimum | `pkg/validate/manifest/sec_302_hsts_short.go` |
| SEC-401 | ERROR | Literal `backend.auth.secret` is forbidden (git leak / rotation impossible) — only `secret_env` is allowed | `pkg/validate/manifest/sec_401_jwt_secret_env_required.go` |
| SEC-402 | WARNING | `backend.auth.access_token_ttl > 30m` exceeds the OWASP recommended upper bound (expanded blast radius) | `pkg/validate/manifest/sec_402_access_ttl_upper_bound.go` |
| SEC-403 | ERROR | `backend.auth.mode` must be one of `cookie` / `bearer` / `hybrid` (defaults to `cookie` when unspecified) | `pkg/validate/manifest/sec_403_auth_mode_enum.go` |
| SEC-404 | ERROR | `frontend.auth.store` must be one of `localStorage` / `memory` (defaults to `localStorage` when unspecified); `cookie` is a mode, not a store — rejected with a pointer to `backend.auth.mode: cookie` | `pkg/validate/manifest/sec_404_frontend_auth_store_enum.go` |

## C. OpenAPI Internal

OpenAPI self-consistency (based on the document parsed by kin-openapi).

| Rule ID | Level | Description | Source |
|---|---|---|---|
| O-1 | ERROR | Path parameter name conflict | `pkg/validate/openapi/o_01_path_param_conflict.go` |
| O-2 | ERROR | Path parameter case conflict | `pkg/validate/openapi/o_02_path_param_case_conflict.go` |
| O-3 | ERROR | Path template parameter declaration missing | `pkg/validate/openapi/o_03_path_template_param.go` |
| O-4 | ERROR | Operation is missing `operationId` | `pkg/validate/openapi/o_04_op_id_required.go` |
| O-5 | ERROR | 4xx/5xx response is missing `content: application/json` + schema (204/304 exempt; 1xx-3xx out of scope) | `pkg/validate/openapi/o_05_response_body_required.go` |
| O-6 | ERROR | Schema `required` entry is not declared in that schema's `properties` (dangling required; checked for components + request/response inline schemas + nested) | `pkg/validate/openapi/o06_required_property_consistency.go` |
| XOO-71 | WARNING | Password-like fields have no `minLength` | `pkg/validate/openapi/xoo_71_password_no_min_length.go` |
| XOO-72 | WARNING | Email-like fields have no `format` | `pkg/validate/openapi/xoo_72_email_no_format.go` |
| XOE-01 | WARNING | ErrorResponse schema의 `error`/`code` 프로퍼티가 `required`에 없으면 oapi-codegen이 `*string`으로 생성하여 빌드 실패 | `pkg/validate/openapi/xoe_01_error_response_required.go` |
| XOE-02 | WARNING | codegen이 에러 표시에 쓰는 필드(`error` 우선, 차선 `message`)가 ErrorResponse schema에 string 속성으로 실재하지 않으면 생성 프론트엔드가 모든 액션 실패를 `String(err)` 폴백으로만 표시 | `pkg/validate/openapi/xoe_02_error_display_field.go` |

## D. Query / sqlc

sqlc query file self-consistency.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| Q-01 | ERROR | `-- name:` annotation is required | `pkg/validate/query/q_01_name_required.go` |
| Q-02 | ERROR | Cardinality (`:one` / `:many` / `:exec` / `:execrows`) is required | `pkg/validate/query/q_02_cardinality.go` |
| Q-03 | ERROR | Query name must be PascalCase | `pkg/validate/query/q_03_name_pascalcase.go` |
| Q-04 | WARNING | `:many` query is missing `LIMIT` | `pkg/validate/query/q_04_many_limit.go` |
| Q-05 | ERROR | `DELETE` statement requires `WHERE` | `pkg/validate/query/q_05_delete_where.go` |
| Q-06 | ERROR | `UPDATE` statement requires `WHERE` | `pkg/validate/query/q_06_update_where.go` |
| Q-07 | WARNING | `SELECT *` on a table that has `@sensitive` columns | `pkg/validate/query/q_07_select_star_sensitive.go` |
| Q-08 | ERROR | Declared parameter is unused in the query body | `pkg/validate/query/q_08_unused_param.go` |
| Q-09 | ERROR | `:exec` query returns `SELECT` | `pkg/validate/query/q_09_select_on_exec.go` |
| Q-10 | ERROR | `sql[].gen.go.out` in `sqlc.yaml` must resolve to `<artifacts>/backend/internal/db` (generate-time; requires `<artifacts>` CLI argument) | `pkg/generate/gogin/check_sqlc_out_path.go` |
| Q-11 | ERROR | `sql[].gen.go.sql_package` in `sqlc.yaml` must be `pgx/v5` (yongol backend codegen is unified on pgx/v5) | `pkg/validate/query/q_11_sql_package_pgx_v5.go` |
| Q-12 | ERROR | DDL has `UUID` column(s) but `sqlc.yaml` is missing the two `pgtype.UUID` overrides (NULL/NOT NULL). Implementation shares `checkPgtypeOverride` with the per-type Q-NN family below — the `NeedsOverride=true` flag on `types.GoTypeBinding` (`pkg/generate/gogin/types/`) is the single drift-free source. | `pkg/validate/query/q_12_pgtype_uuid_override.go` |
| Q-13 | ERROR | DDL has `NUMERIC` / `DECIMAL` column(s) but `sqlc.yaml` is missing the two `pgtype.Numeric` overrides | `pkg/validate/query/q_13_pgtype_numeric_override.go` |
| Q-14 | ERROR | DDL has `TIMESTAMPTZ` column(s) but `sqlc.yaml` is missing the two `pgtype.Timestamptz` overrides | `pkg/validate/query/q_14_pgtype_timestamptz_override.go` |
| Q-15 | ERROR | DDL has `TIMESTAMP` (no TZ) column(s) but `sqlc.yaml` is missing the two `pgtype.Timestamp` overrides | `pkg/validate/query/q_15_pgtype_timestamp_override.go` |
| Q-16 | ERROR | DDL has `DATE` column(s) but `sqlc.yaml` is missing the two `pgtype.Date` overrides | `pkg/validate/query/q_16_pgtype_date_override.go` |
| Q-17 | ERROR | DDL has `INET` / `CIDR` column(s) but `sqlc.yaml` is missing the two `pgtype.Inet` overrides | `pkg/validate/query/q_17_pgtype_inet_override.go` |
| Q-18 | ERROR | DDL has `INTERVAL` column(s) but `sqlc.yaml` is missing the two `pgtype.Interval` overrides | `pkg/validate/query/q_18_pgtype_interval_override.go` |

## E. DDL

DDL self-consistency (PostgreSQL + sqlc query definitions).

| Rule ID | Level | Description | Source |
|---|---|---|---|
| D-1 | ERROR | Duplicate sqlc query name | `pkg/validate/ddl/d_01_sqlc_query_duplicate.go` |
| D-2 | ERROR | `NOT NULL` constraint missing | `pkg/validate/ddl/d_02_nullable_column.go` |
| D-3 | ERROR | FK `DEFAULT 0` sentinel record missing | `pkg/validate/ddl/d_03_sentinel_missing.go` |
| D-4 | ERROR | `db/sqlc.yaml` file missing | `pkg/validate/ddl/d_04_sqlc_yaml_required.go` |
| D-5 | WARNING | `sqlc.yaml` `schema` path does not include the DDL directory | `pkg/validate/ddl/d_05_sqlc_yaml_schema_path.go` |
| D-6 | WARNING | `sqlc.yaml` `queries` path does not include `queries/` | `pkg/validate/ddl/d_06_sqlc_yaml_queries_path.go` |
| D-7 | ERROR | Positional parameters (`$1`, `$2`) are forbidden in sqlc queries | `pkg/validate/ddl/d_07_sqlc_positional_param.go` |
| D-8 | ERROR | `SERIAL` / `BIGSERIAL` / `SMALLSERIAL` column types are banned. Use `GENERATED ALWAYS AS IDENTITY`. | `pkg/validate/ddl/d_08_serial_type_banned.go` |
| D-9 | ERROR | Top-level `INSERT` in a DDL file must be preceded by `-- @sentinel` (otherwise migration would silently drop it) | `pkg/validate/ddl/d_09_top_level_insert_without_sentinel.go` |
| D-10 | ERROR | `@sentinel` `INSERT` must include `ON CONFLICT DO NOTHING` so repeated application is idempotent | `pkg/validate/ddl/d_10_sentinel_without_on_conflict.go` |
| D-11 | ERROR | Column uses an unsupported PG type — multi-word tokens (`DOUBLE PRECISION`, `TIMESTAMP WITH TIME ZONE`) or `CREATE TYPE` user-defined ENUMs. Use single-token aliases (`FLOAT8`, `TIMESTAMPTZ`) or inline `VARCHAR(N) + CHECK IN (...)`. | `pkg/validate/ddl/d_11_unsupported_pg_type.go` |
| D-15 | WARNING | FK 컬럼(`REFERENCES`)에 `NOT NULL`이 없고 `-- @nullable` 어노테이션도 없음 | `pkg/validate/ddl/d_15_fk_nullable.go` |
| XDD-61 | WARNING | Columns matching sensitive patterns (`password` / `secret` / `hash` / `token`) are missing the `@sensitive` annotation | `pkg/validate/ddl/xdd_61_sensitive_no_annotation.go` |

## F. DDL ↔ OpenAPI

Cross-consistency between DDL tables/columns and OpenAPI schemas / extension fields.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XDO-9 | ERROR | OpenAPI property does not correspond to any DDL table column (ghost) | `pkg/validate/openapi_ddl/xdo_09_ghost_property.go` |
| XDO-11 | ERROR | Two or more 2xx responses returning the same entity (same DDL table) expose different representations (field set / nesting / flat-vs-wrapper) — "one resource = one representation" violated (BUG-131). Entity identity: strategy A (SSaC `@response` var), B-1 (SSaC `@response` field base-var convergence), B-2 (DDL table/component guard) | `pkg/validate/openapi_ddl/canonical_response_repr.go` |
| XDO-12 | WARNING | An entity 2xx response is defined inline instead of sharing a `components.schemas.<Model>` `$ref` — consistent today but drift-prone | `pkg/validate/openapi_ddl/canonical_response_repr.go` |
| XDO-67 | ERROR | DDL `VARCHAR(n)` ↔ OpenAPI `maxLength` missing/inconsistent | `pkg/validate/openapi_ddl/xdo_67_max_length_varchar.go` |
| XDO-68 | ERROR | DDL `CHECK IN` ↔ OpenAPI `enum` missing/inconsistent | `pkg/validate/openapi_ddl/xdo_68_check_in_enum.go` |
| XDO-69 | ERROR | DDL `CHECK` allowed values ↔ OpenAPI `enum` values mismatch | `pkg/validate/openapi_ddl/xdo_69_check_values_enum.go` |
| XDO-70 | WARNING | OpenAPI `maxLength` exceeds DDL `VARCHAR(n)` | `pkg/validate/openapi_ddl/xdo_70_max_length_exceeds_varchar.go` |
| XDO-75 | ERROR | OpenAPI optional + DDL `NOT NULL` + no `DEFAULT` | `pkg/validate/openapi_ddl/xdo_75_optional_not_null_no_default.go` |
| XDO-76 | WARNING | OpenAPI required + DDL nullable | `pkg/validate/openapi_ddl/xdo_76_required_nullable.go` |
| XDO-77 | ERROR | DDL column type ↔ OpenAPI field type/format mismatch (incl. float columns require `format: double` — yongol maps every float column to `float64`, so formatless `number` = oapi-codegen `float32` breaks generate) | `pkg/validate/openapi_ddl/xdo_77_column_type_mismatch.go` |
| XDO-78 | ERROR | OpenAPI `enum` declared but DDL column has no matching `CHECK IN` constraint (reverse of XDO-68) | `pkg/validate/openapi_ddl/xdo_78_enum_no_check.go` |
| XOD-10 | WARNING | DDL column is missing from an OpenAPI response schema (coverage) | `pkg/validate/openapi_ddl/xod_10_ddl_to_response.go` |

## G. OpenAPI ↔ SSaC

Cross-consistency between OpenAPI operations/responses and SSaC functions / `@response`.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XOS-15 | ERROR | SSaC funcName must exist as an OpenAPI operationId | `pkg/validate/openapi_ssac/xos_15_func_name_op_id.go` |
| XOS-17 | ERROR | SSaC `@response` fields must match the OpenAPI response schema | `pkg/validate/openapi_ssac/xos_17_response_fields.go` |
| XOS-19 | ERROR | Shorthand `@response` must match the OpenAPI response schema | `pkg/validate/openapi_ssac/xos_19_shorthand_response.go` |
| XOS-21 | ERROR | ErrStatus of `@empty`/`@exists`/`@state`/`@auth`/`@call` must be defined in OpenAPI | `pkg/validate/openapi_ssac/xos_21_err_status_not_in_openapi.go` |
| XOS-22 | ERROR | SSaC has `@response` but OpenAPI has no 2xx response defined | `pkg/validate/openapi_ssac/xos_22_response_no_2xx.go` |
| XOS-66 | ERROR | Fields used in SSaC must be included in OpenAPI `required` | `pkg/validate/openapi_ssac/xos_66_used_fields_required.go` |
| XOS-67 | ERROR | Value type in `@response {key: value}` must be compatible with the expected OpenAPI response schema type | `pkg/validate/openapi_ssac/xos_67_response_field_type.go` |
| XOS-69 | WARNING | SSaC `@response` binds 0 fields but OpenAPI 200 response schema has properties | `pkg/validate/openapi_ssac/xos_69_response_empty_binding.go` |
| XSO-16 | ERROR | OpenAPI operationId must be used as a SSaC function (coverage) | `pkg/validate/openapi_ssac/xso_16_op_id_to_func.go` |
| XSO-18 | ERROR | OpenAPI response field must be used in a SSaC `@response` (coverage) | `pkg/validate/openapi_ssac/xso_18_response_field_used.go` |
| XSO-20 | ERROR | OpenAPI response field must be used in a shorthand `@response` (coverage) | `pkg/validate/openapi_ssac/xso_20_shorthand_field_used.go` |
| XOS-70 | ERROR | `@response` integer field (integer literal or variable binding, required or optional) requires `format: int64` in OpenAPI response schema — codegen binds integer fields to `int64` and formatless oapi-codegen `int` mismatches (covers non-DDL Func/COUNT integer responses; DDL-backed are already forced to int64 by XDO-77) | `pkg/validate/openapi_ssac/xos_70_response_literal_int_format.go` |
| XOS-80 | ERROR | HTTP-method-conventional success status (POST→201, PUT→200, DELETE→204, GET→200) is not declared in OpenAPI responses — codegen cannot derive the success status | `pkg/validate/openapi_ssac/xos_80_success_status_mismatch.go` |
| XOS-82 | WARNING | OpenAPI operation declares multiple 2xx responses but only the one selected by `DeriveSuccessStatus` is reachable from SSaC | `pkg/validate/openapi_ssac/xos_82_unreachable_success_status.go` |

## H. SSaC ↔ Func

Cross-consistency between SSaC `@call` and Func spec (Request/Response).

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XFS-39 | ERROR | `@call` must map to an existing Func spec | `pkg/validate/ssac_func/xfs_39_call_to_func_spec.go` |
| XFS-42 | ERROR | `@call` Inputs count must match FuncRequest field count | `pkg/validate/ssac_func/xfs_42_call_inputs_count.go` |
| XFS-43 | ERROR | `@call` Input field must exist in FuncRequest | `pkg/validate/ssac_func/xfs_43_call_input_fields.go` |
| XFS-44 | ERROR | `@call` Input type must be compatible with FuncRequest field type | `pkg/validate/ssac_func/xfs_44_call_input_type.go` |
| XFS-45 | ERROR | `@result` is declared but Func has no Response | `pkg/validate/ssac_func/xfs_45_call_result_missing.go` |
| XSF-46 | WARNING | Func has a Response but no `@result` is declared | `pkg/validate/ssac_func/xsf_46_call_result_ignored.go` |
| XFS-63 | ERROR | `@call` Func signature must return `(Response, error)` | `pkg/validate/ssac_func/xfs_63_call_func_signature.go` |
| XSF-62 | WARNING | Func spec must be used in SSaC (coverage) | `pkg/validate/ssac_func/xsf_62_func_spec_used.go` |
| XFS-70 | ERROR | `@auth` input value type must be string-compatible | `pkg/validate/ssac_func/xfs_70_auth_input_type.go` |
| XFS-73 | ERROR | `@call` input `request.*` OpenAPI param type must match Func Request field type | `pkg/validate/ssac_func/xfs_73_call_request_param_type.go` |

## I. SSaC ↔ StateMachine

Cross-consistency between SSaC `@state` and Mermaid stateDiagram.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XMS-24 | ERROR | `@state` DiagramID must exist in stateDiagram | `pkg/validate/ssac_statemachine/xms_24_state_diagram_exists.go` |
| XMS-25 | ERROR | `@state` transition event must be defined in stateDiagram | `pkg/validate/ssac_statemachine/xms_25_state_event.go` |
| XSM-23 | ERROR | stateDiagram transition event must exist as a SSaC function | `pkg/validate/ssac_statemachine/xsm_23_transition_to_func.go` |
| XSM-26 | WARNING | Function participating in a state transition has no `@state` declaration | `pkg/validate/ssac_statemachine/xsm_26_missing_state_guard.go` |
| XSM-27 | WARNING | POST/PUT/DELETE on a stateful resource must declare either `@state` or `// @state-neutral` | `pkg/validate/ssac_statemachine/xsm_27_state_intent_declaration.go` |
| XSM-71 | ERROR | `@state` input value type must be string-compatible | `pkg/validate/ssac_statemachine/xsm_71_state_input_type.go` |

## J. SSaC ↔ Rego

Bidirectional matching of `action:resource` pairs between SSaC `@auth` and Rego allow rules.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XPS-28 | ERROR | SSaC `@auth (action:resource)` must have a matching Rego allow rule | `pkg/validate/ssac_rego/xps_28_ssac_auth_to_rego.go` |
| XSP-29 | ERROR | Rego allow `(action:resource)` must be used by a SSaC `@auth` | `pkg/validate/ssac_rego/xsp_29_rego_allow_to_ssac.go` |

## K. DDL ↔ Rego

Validates that Rego `@ownership` / roles reference DDL tables, columns, and CHECK constraints correctly.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XDP-31 | ERROR | `@ownership` table must exist in DDL | `pkg/validate/ddl_rego/xdp_31_ownership_table.go` |
| XDP-32 | ERROR | `@ownership` column must exist in DDL | `pkg/validate/ddl_rego/xdp_32_ownership_column.go` |
| XDP-33 | ERROR | `@ownership via` join table must exist in DDL | `pkg/validate/ddl_rego/xdp_33_ownership_join_table.go` |
| XDP-34 | ERROR | `@ownership via` join column must exist in DDL | `pkg/validate/ddl_rego/xdp_34_ownership_join_column.go` |
| XDP-65 | ERROR | Rego role must be included in DDL `CHECK` constraint allowed values | `pkg/validate/ddl_rego/xdp_65_role_ddl_check.go` |

## L. Rego ↔ Manifest

Bidirectional matching between Rego `input.claims` / role references and the claims/roles defined in `manifest.yaml`.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XNP-53 | ERROR | Rego `input.claims` reference value must exist in manifest claims | `pkg/validate/rego_manifest/xnp_53_input_claims_values.go` |
| XNP-63 | ERROR | Rego role must exist in manifest roles | `pkg/validate/rego_manifest/xnp_63_role_manifest.go` |
| XPN-54 | WARNING | manifest claims must be referenced in Rego (coverage) | `pkg/validate/rego_manifest/xpn_54_claims_to_rego.go` |
| XPN-64 | WARNING | manifest roles must be used in Rego (coverage) | `pkg/validate/rego_manifest/xpn_64_roles_to_rego.go` |

## M. SSaC ↔ Manifest

Cross-consistency between SSaC `currentUser` / `@publish` / `@subscribe` / JWT `@call` and the manifest claims / queue / auth configuration.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XNS-48 | ERROR | Use of `currentUser` requires `backend.auth.claims` in manifest | `pkg/validate/ssac_manifest/xns_48_current_user_claims.go` |
| XNS-49 | ERROR | `currentUser.<field>` must exist in manifest claims | `pkg/validate/ssac_manifest/xns_49_current_user_field.go` |
| XNS-56 | ERROR | Use of `@publish`/`@subscribe` requires `queue.backend` configuration in manifest | `pkg/validate/ssac_manifest/xns_56_queue_required.go` |
| XNS-57 | WARNING | `queue.backend: memory` combined with `@publish` attached to `@post/@put/@delete` (tx-bound publish) — the memory backend does not support `PublishTx`, causing runtime failure | `pkg/validate/ssac_manifest/xns_57_memory_tx_publish.go` |
| XNS-73 | ERROR | JWT `@call` input field must exist in manifest claims fields | `pkg/validate/ssac_manifest/xns_73_jwt_call_claims.go` |
| XNS-80 | ERROR | `@response` field `manifest.*` reference must resolve to an existing manifest.yaml value | `pkg/validate/ssac_manifest/xns_80_manifest_ref.go` |
| XSA-70 | ERROR | SSaC uses `session.*` built-ins but `manifest.session.backend` is not declared | `pkg/validate/ssac_manifest/xsa_70_session_backend_required.go` |
| XSA-71 | ERROR | SSaC uses `cache.*` built-ins but `manifest.cache.backend` is not declared | `pkg/validate/ssac_manifest/xsa_71_cache_backend_required.go` |
| XSA-72 | ERROR | SSaC uses `file.*` / `storage.*` built-ins but `manifest.file.backend` is not declared | `pkg/validate/ssac_manifest/xsa_72_file_backend_required.go` |
| XSA-74 | WARNING | `manifest.session.backend` is declared but no SSaC function uses `session.*` | `pkg/validate/ssac_manifest/xsa_74_session_backend_unused.go` |
| XSA-75 | WARNING | `manifest.cache.backend` is declared but no SSaC function uses `cache.*` | `pkg/validate/ssac_manifest/xsa_75_cache_backend_unused.go` |
| XSA-76 | WARNING | `manifest.file.backend` is declared but no SSaC function uses `file.*` / `storage.*` | `pkg/validate/ssac_manifest/xsa_76_file_backend_unused.go` |
| XSA-77 | WARNING | `manifest.queue.backend` is declared but no SSaC function uses `@publish` / `@subscribe` | `pkg/validate/ssac_manifest/xsa_77_queue_backend_unused.go` |
| XAS-60 | ERROR | `@auth` input field must exist on the Authz `CheckRequest` struct | `pkg/validate/ssac_authz/xas_60_auth_input_field.go` |

## N. OpenAPI ↔ Manifest

Cross-consistency between OpenAPI security schemes and manifest middleware configuration.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XNO-50 | ERROR | OpenAPI securityScheme must map to existing manifest middleware | `pkg/validate/openapi_manifest/xno_50_security_scheme_middleware.go` |
| XNO-52 | ERROR | Endpoint security must reference an existing manifest middleware name, and the middleware block itself must exist | `pkg/validate/openapi_manifest/xno_52_security_middleware.go` |
| XON-51 | ERROR | manifest middleware must map to an existing OpenAPI securityScheme (coverage) | `pkg/validate/openapi_manifest/xon_51_middleware_security_scheme.go` |
| XON-60 | ERROR | `frontend.auth.token_field` (and `refresh_field` when declared) must exist as a property of at least one OpenAPI 2xx response schema; `refresh_op` (when declared) must name an existing operationId whose 2xx response carries `token_field`. Exactly one exemption (plans/stml/sitemap Phase005): a **role_field-only** block (no token_field / refresh_field / refresh_op / store — the `manifest.FrontendAuth.RoleFieldOnly` predicate, shared verbatim with TM-24) skips the token_field requirement, because cookie-mode logins put no token in the response body and the block only wires the `data-roles` menu filter; any token-related key restores the full check — XON-60 is the single enforcer of token_field presence, so an unconditional relaxation would weaken the bearer protection | `pkg/validate/openapi_manifest/xon_60_frontend_auth_token_field.go` |
| SEC-04 | ERROR | The `<key>` of `backend.http.overrides.<key>` must exist as an OpenAPI operationId | `pkg/validate/openapi_manifest/sec_04_http_overrides_operation_id.go` |
| SEC-05 | ERROR | The `<key>` of `backend.rate_limit.<key>` must map to an OpenAPI route (its operationId must exist), otherwise codegen silently omits the rate limiter | `pkg/validate/openapi_manifest/sec_05_rate_limit_op_routable.go` |
| SEC-101 | ERROR | generate-time: the generated main.go must register the request_id and error_envelope middleware immediately after the router, in that order | `pkg/generate/gogin/boot/collect_active_blocks.go` |

## N2. Manifest auth claims ↔ DDL columns (XDN-*)

Cross-consistency between `manifest.backend.auth` (user_table + claims mapping) and the DDL user table parsed from `db/*.sql`. JWT claims are a fixed, finite key set; they must map to typed columns on a named table — not to a generic JSONB blob.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XDN-01 | ERROR | `backend.auth.user_table` is required when auth is active (`auth.type != "none"`) | `pkg/validate/manifest_ddl/xdn_01_user_table_required.go` |
| XDN-02 | ERROR | `backend.auth.user_table` must reference a table parsed from `db/*.sql` | `pkg/validate/manifest_ddl/xdn_02_user_table_exists.go` |
| XDN-03 | ERROR | Each `backend.auth.claims.<Field>: <col>[:<type>]` mapping's column must exist on the user_table | `pkg/validate/manifest_ddl/xdn_03_claim_column_exists.go` |
| XDN-05 | ERROR | Each `backend.auth.claims.<Field>` value must use `<col>:<type>` format (type declaration required). Allowed types: `string`, `int64`, `int32`, `bool`, `uuid` | `pkg/validate/manifest_ddl/xdn_05_claim_type_required.go` |
| XDN-06 | ERROR | Each claim's declared type must match the user_table column's DDL type per the compatibility matrix (uuid↔UUID, string↔TEXT/VARCHAR, int64↔BIGINT/INT8, int32↔INTEGER/INT/INT4, bool↔BOOLEAN/BOOL) | `pkg/validate/manifest_ddl/xdn_06_claim_ddl_type.go` |

> XDN-04 (claim Go type ↔ column DDL-derived Go type) is **deprecated — superseded by XDN-06** and is not registered in `manifest_ddl.Run`. See the Deprecated section.

## O. SSaC ↔ sqlc

Validates that the names, case, and count of SSaC Input keys match sqlc Params exactly.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XQS-14 | WARNING | SSaC input key case must match sqlc param case | `pkg/validate/ssac_sqlc/xqs_14_input_key_case.go` |
| XQS-15 | WARNING | `@call` SSaC input key violates Go initialism rules | `pkg/validate/ssac_sqlc/xqs_15_input_key_initialism.go` |
| XQS-16 | ERROR | SSaC Input key must exist in sqlc Params | `pkg/validate/ssac_sqlc/xqs_16_input_key_missing.go` |
| XQS-17 | ERROR | sqlc Params field must be provided by SSaC Input | `pkg/validate/ssac_sqlc/xqs_17_param_key_missing.go` |
| XQS-18 | ERROR | OpenAPI param type must be compatible with the sqlc param Go type | `pkg/validate/ssac_sqlc/xqs_18_param_type_mismatch.go` |
| XQS-19 | ERROR | SSaC call to a DB-using ssac built-in requires the corresponding sqlc query (per ssac `interface.yaml`) | `pkg/validate/ssac_sqlc/xqs_19_ssac_builtin_query_required.go` |
| XQS-20 | ERROR | SSaC declared return type must match sqlc query RETURNING shape (Model ↔ full RETURNING, `<QueryName>Row` ↔ partial RETURNING) | `pkg/validate/ssac_sqlc/xqs_20_return_type_match.go` |
| XQS-21 | ERROR | `@verify-password` requires the sqlc query `<Model>FindBy<Col>` to exist | `pkg/validate/ssac_sqlc/xqs_21_verify_password_query.go` |
| XQS-72 | ERROR | OpenAPI query param int width must match sqlc param int width | `pkg/validate/ssac_sqlc/xqs_72_query_param_int_width.go` |
| XQS-73 | ERROR | SSaC field reference on partial SELECT query result must exist in SELECT column list | `pkg/validate/ssac_sqlc/xqs_73_partial_select_field.go` |
| XQS-74 | ERROR | `@empty` / `@exists` guard target model's DDL primary key must be an integer type (codegen emits `var.ID == 0`) | `pkg/validate/ssac_sqlc/xqs_74_empty_non_integer_pk.go` |
| XQS-75 | ERROR | `@put` / `@delete` expects `:exec` sqlc query but references a `:one` / `:many` query (assignment mismatch) | `pkg/validate/ssac_sqlc/xqs_75_put_delete_exec_cardinality.go` |
| XQS-76 | ERROR | `@get` / `@post` expects `:one` / `:many` sqlc query but references a `:exec` / `:execrows` / `:execresult` query (assignment mismatch) | `pkg/validate/ssac_sqlc/xqs_76_get_post_exec_cardinality.go` |

## P. SSaC ↔ DDL

Bidirectional validation that SSaC `@result` / `@input` matches DDL tables/columns.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XDS-12 | WARNING | `@result` type must match a sqlc row type or DDL table. @result↔DDL 매칭은 단수 정규화(canonical = 단수 lower-snake, `canonicalTableKey`)로 하므로 단/복수 명명 모두 허용(`AppConfig` ↔ `app_config` / `app_configs`) — XSD-55와 일관 | `pkg/validate/ssac_ddl/xds_12_result_no_ddl_table.go` |
| XSD-55 | ERROR | DDL table must be referenced in a SSaC `@model` (coverage). 모델↔테이블은 단수 정규화(canonical = 단수 lower-snake, `caseconv.TableSingular`)로 매칭하므로 단/복수 명명 모두 허용(`AppConfig` ↔ `app_config` / `app_configs`). Exempt: `-- @func-managed` (RPC/함수가 관리하는 활성 테이블 — `@call`로 위임되어 SSaC `@model`/`@result`에 직접 안 나타남), `-- @archived` (미사용/폐기 테이블) | `pkg/validate/ssac_ddl/xsd_55_ddl_to_model_ref.go` |

## P2. sqlc ↔ Rego

Phase003 (ssac/purify) — every Rego `@ownership` annotation must be backed
by a user-authored `OwnerLookup<Resource>` sqlc query. ssac/pkg/authz is
DB-free: the handler performs the owner lookup and injects the result
into `authz.CheckRequest.Owners` before calling authz.Check.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XQP-30 | ERROR | `@ownership <res>: <table>.<col>` requires sqlc query `OwnerLookup<Pascal(res)>` to exist | `pkg/validate/query_rego/xqp_30_owner_lookup_query.go` |

## P3. Manifest-driven DB requirements (XN*-90)

Phase004 (ssac/purify) — when the manifest opts into a DB-backed built-in
(cache / session / queue / auth refresh), the user must provide the
canonical DDL + sqlc queries declared in the matching
`ssac/pkg/<x>/interface.yaml`. Missing entries surface as ERROR with a
copy-pasteable advice block sourced from the interface.yaml's
`canonical_ddl` + `canonical_queries`.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XNC-90 | ERROR | `manifest.cache.backend=postgres` requires `fullend_cache` DDL + `CacheSet/CacheGet/CacheDelete` sqlc queries | `pkg/validate/manifest/xnc_90_cache_backend_requires_sqlc.go` |
| XNS-90 | ERROR | `manifest.session.backend=postgres` requires `fullend_sessions` DDL + `SessionSet/SessionGet/SessionDelete` sqlc queries | `pkg/validate/manifest/xns_90_session_backend_requires_sqlc.go` |
| XNQ-90 | ERROR | `manifest.queue.backend=postgres` requires `fullend_queue` DDL + `QueuePublish/QueuePoll/QueueAck` sqlc queries | `pkg/validate/manifest/xnq_90_queue_backend_requires_sqlc.go` |
| XNA-90 | ERROR | `manifest.backend.auth` configured requires `refresh_tokens` DDL + `RefreshTokenInsert/FindByHash/Revoke/RevokeAll` sqlc queries | `pkg/validate/manifest/xna_90_refresh_requires_sqlc.go` |

## R. Hurl Internal

| Rule ID | Level | Description | Source |
|---|---|---|---|
| H-1 | ERROR | `.feature` files exist (deprecated; use Hurl `.hurl`) | `pkg/validate/hurl/h_01_deprecated_feature.go` |
| H-2 | WARNING | `tests/` directory is empty | `pkg/validate/hurl/h_02_empty_tests_dir.go` |

## R2. Hurl ↔ OpenAPI

All Hurl files under `specs/tests/` are user-authored. yongol does not emit any Hurl; `generate` mirrors `specs/tests/` → `arts/tests/` verbatim. These rules catch drift between user-authored Hurl and the OpenAPI SSOT at validate time.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XOH-01 | ERROR | Hurl URL + method declared in OpenAPI | `pkg/validate/hurl_openapi/xoh_01_url_method.go` |
| XOH-02 | ERROR | Hurl `HTTP <status>` declared in OpenAPI responses | `pkg/validate/hurl_openapi/xoh_02_status_declared.go` |
| XOH-03 | ERROR | Hurl JSON body field in OpenAPI request schema | `pkg/validate/hurl_openapi/xoh_03_request_field_in_schema.go` |
| XOH-04 | ERROR | Hurl assert jsonpath reachable in OpenAPI response schema | `pkg/validate/hurl_openapi/xoh_04_assert_path_in_schema.go` |
| XOH-08 | ERROR | Hurl capture jsonpath reachable in OpenAPI response schema | `pkg/validate/hurl_openapi/xoh_08_capture_path_in_schema.go` |
| XOH-09 | WARNING | Hurl captured variable is referenced later in the file | `pkg/validate/hurl_openapi/xoh_09_unused_capture.go` |
| XOH-10 | ERROR | smoke.hurl is required in specs/tests/ | `pkg/validate/hurl_openapi/xoh_10_smoke_required.go` |
| XOH-11 | ERROR | smoke.hurl must cover all OpenAPI operationIds | `pkg/validate/hurl_openapi/xoh_11_smoke_coverage.go` |
| XOH-12 | WARNING | OpenAPI declared status codes covered by hurl tests (5xx excluded) | `pkg/validate/hurl_openapi/xoh_12_status_coverage.go` |
| XOH-13 | WARNING | SSaC guard ErrStatus + @response happy path covered by hurl tests | `pkg/validate/hurl_openapi/xoh_13_guard_coverage.go` |

## R3. Hurl ↔ State Machine

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XOH-05 | WARNING | Hurl call order satisfies state machine transitions | `pkg/validate/hurl_statemachine/xoh_05_state_transition_order.go` |

## R4. Hurl ↔ Manifest

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XOH-06 | WARNING | Protected Hurl call preceded by an auth step | `pkg/validate/hurl_manifest/xoh_06_auth_precondition.go` |
| XOH-07 | WARNING | Cookie-mode mutation includes the manifest-resolved CSRF header (default `X-XSRF-TOKEN`) | `pkg/validate/hurl_manifest/xoh_07_csrf_on_mutation.go` |

## R5. State Machine / Rego / Func

| Rule ID | Level | Description | Source |
|---|---|---|---|
| ST-1 | ERROR | Mermaid stateDiagram parsing validation | `pkg/validate/statemachine/st_01_parse.go` |
| P-1 | ERROR | Rego policy parsing validation | `pkg/validate/rego/p_01_parse.go` |
| XPP-30 | ERROR | Rego references `resource_owner` but no `@ownership` annotation is present | `pkg/validate/rego/xpp_30_ownership_no_annotation.go` |
| F-1 | WARNING | Func name collides with a built-in package name (`auth`/`session`/`cache`/`file`) | `pkg/validate/funcspec/f_01_builtin_override.go` |
| XFF-40 | ERROR | Func body is unimplemented (`panic("TODO")` / `// TODO` / empty body) | `pkg/validate/funcspec/xff_40_func_body_todo.go` |
| XFF-41 | ERROR | Func body must not import I/O packages (`database/sql`, `net/http`, `grpc`, etc.) | `pkg/validate/funcspec/xff_41_func_forbidden_import.go` |
| XDM-27 | ERROR | `@state` field must exist as a DDL column | `pkg/validate/ddl_statemachine/xdm_27_state_field_column.go` |
| XDM-28 | ERROR | stateDiagram `[*] → X` initial transition must match DDL `DEFAULT 'X'` | `pkg/validate/ddl_statemachine/xdm_28_default_initial_state.go` |

## Z1. Design Internal (`V-*`)

DESIGN.md self-consistency — name, colors, typography, dimensions, token references, headings, and component properties.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| V-01 | ERROR | `name` field is required in DESIGN.md | `pkg/validate/design/v_01_name_required.go` |
| V-02 | ERROR | Color value must be a valid hex code (`#` followed by 3, 4, 6, or 8 hex digits) | `pkg/validate/design/v_02_hex_valid.go` |
| V-03 | ERROR | Typography token must have `fontFamily`, `fontSize`, and `fontWeight` fields | `pkg/validate/design/v_03_typography_required.go` |
| V-04 | ERROR | `rounded` / `spacing` values must be valid dimensions (number optionally followed by `px`, `em`, or `rem`) | `pkg/validate/design/v_04_dimension_valid.go` |
| V-05 | ERROR | `{group.token}` references in component props must resolve to actual tokens in the same DESIGN.md | `pkg/validate/design/v_05_token_ref_resolve.go` |
| V-06 | ERROR | Duplicate `##` section heading in DESIGN.md body | `pkg/validate/design/v_06_duplicate_heading.go` |
| V-07 | WARNING | Component property name is not in the known set (possible typo or unknown prop) | `pkg/validate/design/v_07_unknown_prop.go` |

## Z2. Design ↔ Manifest (`XNV-*`)

Cross-consistency between `manifest.yaml` `frontend.design` and DESIGN.md file existence.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XNV-01 | ERROR | `manifest.frontend.design` path does not resolve to an existing file | `pkg/validate/design_manifest/xnv_01_path_exists.go` |
| XNV-02 | WARNING | DESIGN.md (or `*.design.md`) file exists under `specs/frontend/` but is not declared in `manifest.frontend.design` | `pkg/validate/design_manifest/xnv_02_undeclared.go` |

## Z3. STML ↔ Design (`XVM-*`, `XMV-*`)

Cross-validation between STML template classes/attributes and DESIGN.md design tokens. `XVM-*` rules check STML → Design (claim direction), `XMV-*` rules check Design → STML (coverage direction, dead tokens).

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XVM-01 | WARNING | Color token name used in STML class is not defined in DESIGN.md `colors` | `pkg/validate/stml_design/xvm_01_color.go` |
| XVM-02 | WARNING | Rounded token name used in STML class is not defined in DESIGN.md `rounded` | `pkg/validate/stml_design/xvm_02_rounded.go` |
| XVM-03 | WARNING | Spacing token name used in STML class is not defined in DESIGN.md `spacing` | `pkg/validate/stml_design/xvm_03_spacing.go` |
| XVM-04 | WARNING | Font name used in STML class does not match any `fontFamily` in DESIGN.md `typography` (case-insensitive) | `pkg/validate/stml_design/xvm_04_font.go` |
| XVM-05 | WARNING | Inline `style` attribute contains a hardcoded hex color that matches a DESIGN.md color token value — use the token instead | `pkg/validate/stml_design/xvm_05_inline.go` |
| XVM-06 | ERROR | `data-component` used in STML but not defined in DESIGN.md `components` | `pkg/validate/stml_design/xvm_06_component_design_required.go` |
| XMV-10 | WARNING | DESIGN.md color token is defined but not referenced in any STML page (dead token) | `pkg/validate/stml_design/xmv_10_dead_color.go` |
| XMV-11 | WARNING | DESIGN.md typography token's `fontFamily` is defined but not referenced in any STML page (dead token) | `pkg/validate/stml_design/xmv_11_dead_typography.go` |
| XMV-12 | WARNING | DESIGN.md component token is defined but not referenced by any STML `data-component` (dead token) | `pkg/validate/stml_design/xmv_12_dead_component.go` |

## Z4. Domain Security (`XDO-90`, `XDS-80/81/82`, `XMO-20/21/22`)

Multi-domain OpenAPI security rules. Applies only when `manifest.yaml` declares multiple domain configurations (e.g. `public`, `admin`, `internal` — these key names are reserved semantic markers, see C-12~C-17 in section B). Validates operationId uniqueness across domains, domain-specific access control, and STML consumption coverage per domain.

> **Domain-mode coverage (post BUG-141).** A multi-domain project has no top-level `api/openapi.yaml`; before BUG-141 the validate step gate silently skipped every OpenAPI-dependent step — including `domain_security` — yielding a false `0 errors`. After the fix, in domain mode the OpenAPI-dependent steps run over the per-domain specs and `domain_security` executes (it is internally guarded by `len(fs.Manifest.Domains) > 0`).
>
> **XMO-10 vs XMO-11/12 / XMO-20/21/22 in domain mode.** The single-site coverage rule **XMO-10** (`stml_openapi`) is single-site-only — it does not fire in domain mode, where per-domain consumption is instead enforced by **XMO-20/21** (this section). **XMO-11/12** (`stml_openapi`) still run, iterating over *all* domains' STML pages / OpenAPI docs (`AllSTMLPages()` / `AllOpenAPIDocs()`). XMO-20/21/22 here are the domain-scoped operationId-consumption and cross-domain-call boundary checks.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XDO-90 | ERROR | operationId is declared in more than one domain's OpenAPI spec (cross-domain duplicate) | `pkg/validate/domain_security/xdo_90.go` |
| XDS-80 | ERROR | Admin domain endpoint allows public access via `security: []` | `pkg/validate/domain_security/xds_80.go` |
| XDS-81 | WARNING | Internal domain endpoint has explicit security declarations (typically unnecessary for service-to-service calls) | `pkg/validate/domain_security/xds_81.go` |
| XDS-82 | ERROR | Public domain DELETE operation has no corresponding admin Rego allow rule with `delete` action | `pkg/validate/domain_security/xds_82.go` |
| XMO-20 | ERROR | Public domain OpenAPI operationId is not consumed by any STML page in the public frontend directory | `pkg/validate/domain_security/xmo_20.go` |
| XMO-21 | ERROR | Admin domain OpenAPI operationId is not consumed by any STML page in the admin frontend directory (only when admin frontend is configured) | `pkg/validate/domain_security/xmo_21.go` |
| XMO-22 | WARNING | STML page calls an operationId belonging to a different domain's OpenAPI spec (domain boundary violation) | `pkg/validate/domain_security/xmo_22.go` |

## S. Preserve (`PRV-*`) — Preserved file contract / runtime safety guards

Applies only to `.go` files whose `//ff:checked hash` no longer matches (i.e., the user has edited them). Runs only when the arts directory is provided, as in `yongol validate <specs> <arts>`.

### Contract drift (PRV-01 ~ PRV-09)

| Rule ID | Level | Description | Source |
|---|---|---|---|
| PRV-01 | ERROR | Preserved file's function signature has drifted from the SSOT expectation | `pkg/validate/contract/prv_01_signature_drift.go` |
| PRV-02 | ERROR | Preserved file references a sqlc query / @call / DDL field that does not exist in the SSOT | `pkg/validate/contract/prv_02_external_symbol_drift.go` |

### Runtime safety guards (PRV-10 ~ PRV-19)

| Rule ID | Level | Description | Source |
|---|---|---|---|
| PRV-10 | ERROR | Preserved file contains a disallowed `panic(` (excluding init() and `// nolint:panic`) | `pkg/validate/contract/prv_10_preserved_panic.go` |
| PRV-11 | ERROR | Preserved file's `ctx.Value("currentUser").(T)` is not in comma-ok form | `pkg/validate/contract/prv_11_preserved_currentuser_assertion.go` |
| PRV-12 | ERROR | Preserved file ignores errors from `json.Unmarshal` / `yaml.Unmarshal` | `pkg/validate/contract/prv_12_preserved_unmarshal_err.go` |
| PRV-13 | ERROR | Preserved file ignores errors from `sql.Rows.Scan` / `sql.Row.Scan` | `pkg/validate/contract/prv_13_preserved_scan_err.go` |
| PRV-14 | ERROR | Preserved file accesses the first slice element (`x[0]`) without a `len` guard | `pkg/validate/contract/prv_14_preserved_slice_bounds.go` |
| PRV-15 | ERROR | Preserved file uses an inline selector access such as `m[k].Field` (no comma-ok guard) | `pkg/validate/contract/prv_15_preserved_map_access.go` |
| PRV-16 | ERROR | Preserved file accesses fields of a `Get*()`/`Find*()` return value directly (no nil guard) | `pkg/validate/contract/prv_16_preserved_nil_deref.go` |
| PRV-17 | ERROR | Preserved file acquires a resource (`os.Open` / `db.Query` / `http.Get`, etc.) but is missing `defer Close` | `pkg/validate/contract/prv_17_preserved_missing_close.go` |

Allowlist:
- Inside `init()` functions — exempt from PRV-10
- A `// nolint:panic` on the preceding line (or the same line) — exempt from PRV-10
- A `// nolint:prv-NN` on the preceding line (or the same line) — exempt from the corresponding rule

---

## T. Migration (`MIG-*`) — DDL auto migration

Rules collected during the DDL migration phase of `yongol generate` (`pkg/generate/migration/`), covering hints, snapshots, and destructive operations. For the syntax of DDL comment hints (`-- @rename`, `-- @cast`, `-- @backfill`, `-- @data_migration`, `-- @allow_destructive`), see the "DDL hints" section of `manual-for-ai.md` and `docs/MIGRATION.md`.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| MIG-001 | ERROR | The `from` of `@rename from=...` is missing in the previous snapshot, or `to` is missing in the current DDL (mismatch) | `pkg/validate/migration/mig_001_rename_mismatch.go` |
| MIG-002 | ERROR | NOT NULL column addition with no `@backfill default=...` hint and no DEFAULT clause → emit is blocked because existing rows would be NULL | `pkg/validate/migration/mig_002_not_null_without_backfill.go` |
| MIG-003 | ERROR | The sidecar SQL file referenced by `@data_migration file=<path>` does not exist | `pkg/validate/migration/mig_003_data_migration_missing.go` |
| MIG-004 | WARNING | DROP TABLE / DROP COLUMN occurs but the target table has no `@allow_destructive` (intent reconfirmation recommended) | `pkg/validate/migration/mig_004_destructive_without_allow.go` |
| MIG-005 | WARNING | Risky type change (`INTEGER↔TEXT`, narrowing `VARCHAR(N)`, etc.) has no `@cast using=<expr>` hint | `pkg/validate/migration/mig_005_cast_missing.go` |
| MIG-006 | ERROR | The `-- YONGOL_SCHEMA_HASH:` header in `arts/db/.latest_schema.sql` does not match the sha256 of the body (user-edited = drift) | `pkg/validate/migration/mig_006_snapshot_drift.go` |

## U. STML ↔ OpenAPI / stateDiagram (`TM-*`)

Cross-validation between STML template attributes (`data-fetch`, `data-action`, `data-param`, `data-field`, `data-bind`, `data-each`, `data-component`, `data-layout`, `data-state`, `data-enabled-when`, `data-invalidates`, `data-capture`, `data-redirect`, `data-redirect-params`, `data-on-error`, `data-link`, `data-link-params`, plus the layout vocabulary `data-nav`/`data-outlet`/`data-logout` and the sitemap vocabulary `data-sitemap`/`data-page`/`data-index`/`data-entry`/`data-menu`/`data-icon`/`data-roles`/`data-crumb-field` and the dynamic menu group vocabulary `data-fetch`/`data-each`/`data-link`/`data-link-params`/`data-label-field` on a group's nested `<ul>`) and the OpenAPI spec, layouts, the sitemap, and Mermaid stateDiagrams. Ensures that STML references resolve to valid OpenAPI operations, parameters, request/response fields, component files, layouts, and statechart states/transitions. Most rules live in `pkg/validate/stml_openapi/`; the stateDiagram cross-checks (TM-15, TM-18, TM-23) live in `pkg/validate/stml_statemachine/`. TM-20~26 are the runtime twins of the Hurl flow rules (XOH-05/06/07/08/09) for the auth session flow (plans/stml/auth-flow Phase002). TM-39~42/49 validate `frontend/sitemap.html` (plans/stml/sitemap Phase001); TM-43 runs the edge-based reachability (orphan page) check over it (Phase002 — listing in the sitemap is a node, not an edge); TM-44 enforces the menu's single source of truth once the sitemap exists (Phase003 — the layout menu is emitted from the sitemap tree, so a surviving layout `data-nav` is an ERROR); TM-46/47 validate the `data-roles` role-based menu filter and its claim wiring (Phase005 — menu hiding is UX, not security; access blocking stays Rego's concern); TM-50 validates the `data-crumb-field` dynamic breadcrumb label against the page's first fetch response (Phase006); TM-51 is the inverse of TM-49 — a sitemap derives a menu but no layout hosts it (`layouts/` empty + no `defaultLayout` + no nav `data-layout`), so the menu/breadcrumb never render (Phase008, BUG-129); TM-48 validates the dynamic menu groups of Phase007 (`data-fetch`/`data-each` sitemap groups — never in a `data-entry` block, required vocabulary complete), whose field-level judgments reuse the sitemap extensions of TM-01/07/08/30/31/32. TM-45 (reserved-attribute warning) retired in Phase007 — the whole reserved vocabulary graduated.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| TM-01 | ERROR | `data-fetch` operationId is not defined in OpenAPI — page fetches (`tm_01_fetch_op_not_found.go`) and sitemap dynamic menu groups (Phase007, `tm_01_sitemap_group_fetch.go`) alike | `pkg/validate/stml_openapi/tm_01_fetch_op_not_found.go` |
| TM-02 | ERROR | `data-action` operationId is not defined in OpenAPI | `pkg/validate/stml_openapi/tm_02_action_op_not_found.go` |
| TM-03 | ERROR | `data-action` references a GET method endpoint (actions require POST/PUT/DELETE) | `pkg/validate/stml_openapi/tm_03_action_get_method.go` |
| TM-04 | ERROR | `data-param` name is not declared in the OpenAPI operation's parameters | `pkg/validate/stml_openapi/tm_04_param_not_found.go` |
| TM-05 | ERROR | `data-field` name is not in the OpenAPI operation's request body schema | `pkg/validate/stml_openapi/tm_05_field_not_found.go` |
| TM-06 | ERROR | `data-bind` field is not in the OpenAPI operation's success 2xx response schema (`responseFields`, scanned 200 → 201 — the codegen `Res<K>` priority; a 204/no-body op fails every bind) | `pkg/validate/stml_openapi/tm_06_bind_not_found.go` |
| TM-07 | ERROR | `data-each` field is not in the OpenAPI operation's response schema — page fetches and sitemap dynamic menu groups (Phase007, `tm_07_sitemap_group_each.go`) alike | `pkg/validate/stml_openapi/tm_0708_each.go` |
| TM-08 | ERROR | `data-each` field exists in the response schema but is not an array type (sitemap dynamic menu groups included) | `pkg/validate/stml_openapi/tm_0708_each.go` |
| TM-09 | ERROR | `data-component` references a `.tsx` file that does not exist | `pkg/validate/stml_openapi/tm_09_component_not_found.go` |
| TM-10 | ERROR | STML element uses a `class` attribute directly — use `<!-- @override class="..." -->` comment instead | `pkg/validate/stml_openapi/tm_10_class_prohibited.go` |
| TM-11 | ERROR | `data-layout` value on a page does not match any layout defined in `layouts/` | `pkg/validate/stml_openapi/tm_11_layout_not_found.go` |
| TM-12 | ERROR | `manifest.frontend.defaultLayout` value does not match any layout defined in `layouts/` | `pkg/validate/stml_openapi/tm_12_default_layout_not_found.go` |
| TM-13 | WARNING | Layout defined in `layouts/` is not referenced by any page's `data-layout`, `manifest.frontend.defaultLayout`, or a sitemap `<nav data-sitemap data-layout>` block (the Phase003 menu emitter and route builder use the sitemap assignment, so it counts as used — validation and emission must not drift) | `pkg/validate/stml_openapi/tm_13_unused_layout.go` |
| TM-14 | ERROR | `data-enabled-when` guard references a model that is not a top-level property of any page `data-fetch` response schema | `pkg/validate/stml_openapi/tm_14_enabled_when_ref_not_found.go` |
| TM-15 | ERROR | `data-enabled-when` guard comparison references a state value that does not exist in the matching Mermaid stateDiagram | `pkg/validate/stml_statemachine/tm_15_state_value_in_diagram.go` |
| TM-16 | ERROR | `data-invalidates` operationId is not defined in OpenAPI, or is not a GET endpoint (only GET queries can be invalidated) | `pkg/validate/stml_openapi/tm_16_invalidates_op_not_found.go` |
| TM-17 | ERROR | `data-state` guard using a combinator (`&&`, `\|\|`, leading `!`, or parentheses) is not valid guard syntax (§3.4 EBNF; no function calls, arithmetic, or ternaries) | `pkg/validate/stml_openapi/tm_17_guard_syntax.go` |
| TM-18 | WARNING | `data-action` transition is not legal from the state its `data-enabled-when` guard requires, per the Mermaid stateDiagram | `pkg/validate/stml_statemachine/tm_18_transition_validity.go` |
| TM-19 | WARNING | `data-field` binds an `object`(map) type request body field to a plain text input — the generated key-value data cannot be entered through a single text input | `pkg/validate/stml_openapi/tm_19_map_field_text_input.go` |
| TM-20 | ERROR | `data-capture` syntax violation (must be `<respField> -> <sink>` with sink `auth.token`/`auth.refresh`/`auth.claims.<name>` — the claims sink ships with plans/stml/sitemap Phase005), or a captured respField (claims captures included) is not in the operation's OpenAPI 2xx response schema (↔ XOH-08) | `pkg/validate/stml_openapi/tm_20_capture_field_in_response.go` |
| TM-21 | WARNING | bearer mode but no STML page captures `auth.token`, or captures exist but no page calls a security-protected operation — the captured token is never consumed (↔ XOH-09) | `pkg/validate/stml_openapi/tm_21_capture_sink_unused.go` |
| TM-22 | ERROR | bearer mode + a page calls a `security`-protected operation + no STML page captures `auth.token` — every protected screen is guaranteed a 401 (↔ XOH-06) | `pkg/validate/stml_openapi/tm_22_protected_op_no_token_supply.go` |
| TM-23 | WARNING | `data-redirect` target page's `data-state` guard (`=` comparison on the same stateDiagram) requires a state that is not an arrival state of the action's transition (↔ XOH-05); not-comparable guards stay silent | `pkg/validate/stml_statemachine/tm_23_redirect_state_conflict.go` |
| TM-24 | WARNING | cookie mode but an `auth.token`/`auth.refresh` `data-capture` or a manifest `frontend.auth` block with token keys is declared — httpOnly cookies cannot be captured (↔ XOH-07 mode consistency). Two Phase005 exemptions, both grounded in "the blocking reason does not apply": `auth.claims.*` captures read the login response *body* (not the cookie) and stay first-class in cookie mode, and a role_field-only `frontend.auth` block (the `manifest.FrontendAuth.RoleFieldOnly` predicate, shared verbatim with XON-60 so the two judgments cannot drift) declares no token contract; a mixed block (token keys + role_field) still warns, with the advice asking to remove only the token-related keys | `pkg/validate/stml_openapi/tm_24_cookie_mode_capture_conflict.go` |
| TM-25 | ERROR | `data-on-error` is outside any `data-action` block, or `data-capture`/`data-redirect` sits on an element without `data-action` | `pkg/validate/stml_openapi/tm_25_flow_attr_placement.go` |
| TM-26 | ERROR | `data-redirect` does not resolve to any STML page: a `/`-prefixed value is a static path matched against the resolved route patterns (`/` is allowed as the index route), any other value is a page-name reference (STML filename without `.html`) that must name an existing page | `pkg/validate/stml_openapi/tm_26_redirect_route_exists.go` |
| TM-27 | ERROR | a `route.<Name>` param the page consumes has no same-named `:Name`/`:Name?` segment in the page's resolved route (`stml.RoutePaths`; case-exact) — the param is always `undefined` at runtime | `pkg/validate/stml_openapi/tm_27_route_param_missing.go` |
| TM-28 | WARNING | a `:Name`/`:Name?` segment of the page's resolved route is consumed by no `data-param-*` binding — dead segment (URL design vs page implementation drift) | `pkg/validate/stml_openapi/tm_28_route_segment_unused.go` |
| TM-29 | WARNING | the `data-action` operation declares a 4xx/5xx response but the action block has no `data-on-error` element — the server error falls back to the default inline slot (`role="alert"`); declare `data-on-error` to decide where it appears | `pkg/validate/stml_openapi/tm_29_action_on_error_missing.go` |
| TM-30 | ERROR | an `item.<Field>` `data-param-*` source is used outside any `data-each` block (no row is in scope), or the field is not in the item schema of the enclosing `data-each` array (OpenAPI response; innermost each); a sitemap dynamic menu group (Phase007, `tm_30_sitemap_group_label_field.go`) **requires** `data-label-field`, which must exist in the group's `data-each` item schema as a string/integer/number scalar — the menu items' label source | `pkg/validate/stml_openapi/tm_30_item_source.go` |
| TM-31 | ERROR | `data-link` target page name does not match any STML page (filename without `.html`) — the emitted `<Link>` would navigate into the void; sitemap dynamic menu groups included (Phase007, `tm_31_sitemap_group_link.go` — every fetched item links there) | `pkg/validate/stml_openapi/tm_31_link_target_not_found.go` |
| TM-32 | ERROR | `data-link-params` violation: syntax (`<source> -> <SegmentName>`, comma-separated; sources `item.*`/`route.*` only), a **required** segment of the target page's resolved route left unmapped, a SegmentName absent from the target route, an `item.*` source outside `data-each` or not in the enclosing each's item schema (TM-30 infrastructure), a `route.*` source absent from this page's resolved route (TM-27 infrastructure), or the elided form (`item.id`) used when the target route does not have exactly one required segment; unmapped **optional** segments are legal. Sitemap dynamic menu groups (Phase007, `tm_32_sitemap_group_link_params.go`) get the same judgment with the group's each item schema in scope and `route.*` sources rejected outright — the layout menu renders on every route, so no own-route value exists | `pkg/validate/stml_openapi/tm_32_link_params_unsatisfied.go` |
| TM-33 | ERROR | `data-redirect-params` violation: declared on a static (`/`-prefixed) `data-redirect` (contradiction — substitution applies only to page-name references), syntax (`<respField> -> <SegmentName>`, comma-separated; sources are unprefixed 2xx respFields or `route.*`), a respField absent from the action operation's OpenAPI 2xx response schema (TM-20 infrastructure; `route.*` sources are exempt), a SegmentName absent from the redirect target page's resolved route, a **required** target segment left unmapped (also when no params are declared at all), or the elided form used when the target route does not have exactly one required segment; unmapped **optional** segments are legal | `pkg/validate/stml_openapi/tm_33_redirect_params.go` |
| TM-34 | ERROR | `manifest.frontend.index` violation: the page name (STML filename without `.html`) matches no STML page, the target page's resolved route carries a **required** parameter segment (a redirect has no value to fill it; optional-only `:Name?` pages are legal — segments are stripped), or a page already mounts `/` via `data-route` (mount vs redirect declared at once — contradiction) | `pkg/validate/stml_openapi/tm_34_index_target.go` |
| TM-35 | WARNING | frontend ON with at least one STML page, but no index is declared (no page mounts `/`, `manifest.frontend.index` is absent, and no sitemap entry carries `data-index`) — the `/` route falls back to the first public page in file-name sort order, so an accident, not a declaration, decides the first screen; the advice names the picked fallback page and all three declaration vehicles | `pkg/validate/stml_openapi/tm_35_index_fallback.go` |
| TM-36 | ERROR | (sitemap-absent path — with a sitemap, `data-nav` itself is TM-44's ERROR) layout `data-nav` target does not resolve: a `/`-prefixed value is a static path matched against every page's resolved route patterns (`/` is allowed as the index route), any other value is a page-name reference (STML filename without `.html`) that must name an existing page **and** resolve to a route without a **required** parameter segment — a static menu link has no value to fill it (optional `:Name?` segments are legal, the emitter strips them; parameterized navigation belongs to `data-link`) | `pkg/validate/stml_openapi/tm_36_nav_target.go` |
| TM-37 | ERROR | layout `data-logout` operationId is not defined in OpenAPI, or references a GET endpoint (session-ending operations require POST/PUT/DELETE — the TM-02/03 contract applied to the layout) | `pkg/validate/stml_openapi/tm_37_logout_op.go` |
| TM-38 | WARNING | dead or under-powered `data-logout`: the project declares no backend.auth (there is no session to end — no logout UI is emitted), or the effective auth mode (prepared.AuthFor — the emitter's own derivation) is non-bearer (cookie/hybrid) with a *valueless* `data-logout` — the session lives in an httpOnly cookie only a server operation can end | `pkg/validate/stml_openapi/tm_38_logout_mode.go` |
| TM-39 | ERROR | sitemap `data-page` does not name any STML page (filename without `.html`), an entry declares both `data-page` and an `<a href>` external link — the two vehicles are mutually exclusive (internal page vs external URL) —, or `data-crumb-field` sits on an entry without `data-page` (group label / external link) — the dynamic crumb label is read from the page's fetch response, so the attribute is page-item-only (plans/stml/sitemap Phase006) | `pkg/validate/stml_openapi/tm_39_sitemap_page_not_found.go` |
| TM-40 | ERROR | a page appears more than once across the whole sitemap (nav blocks included) — a page has exactly one canonical position (breadcrumbs and active highlighting require a unique parent); the message names both positions; cross-references belong to `data-link` on the referring page | `pkg/validate/stml_openapi/tm_40_sitemap_duplicate_page.go` |
| TM-41 | ERROR | sitemap `<nav data-layout>` value does not match any layout defined in `layouts/` (the TM-11 existence judgment applied to the sitemap; navs without `data-layout` delegate to `defaultLayout` and are skipped) | `pkg/validate/stml_openapi/tm_41_sitemap_layout_not_found.go` |
| TM-42 | ERROR | `data-index` violation: declared more than once across the sitemap (`/` redirects to exactly one page), sits on an entry without `data-page`, the marked page's resolved route carries a **required** parameter segment (the TM-34 judgment — a redirect has no value to fill it; optional `:Name?` segments are legal), or `manifest.frontend.index` is also declared and names a different page (the same decision stated twice, in contradiction; a nonexistent `data-page` is TM-39's finding) | `pkg/validate/stml_openapi/tm_42_sitemap_index_conflict.go` |
| TM-43 | WARNING | a page is unreachable from the roots (sitemap `data-index` pages ∪ every page of a `data-entry` block ∪ `manifest.frontend.index` ∪ `data-route="/"` mounts) by BFS over the actual movement edges: menu-rendered sitemap entries (`MenuRenderable` — depth ≤ 2, no required route param, no `data-menu="false"`; the same judgment the Phase003 menu emitter consumes), `data-link` targets, resolvable `data-redirect` targets and the breadcrumb up-edges of plans/stml/sitemap Phase004 (DESIGN §4.10 edge (d) — every sitemap-listed page of depth ≥ 2 links its `MenuRenderable` ancestors through the generated `<Breadcrumb>`, the exact crumbs the emitter gives an href), and the `data-link` targets of menu-rendered dynamic menu groups (Phase007 — every fetched item is a menu NavLink there, folded into the roots like edge (a)). **Listing in the sitemap is a node, not an edge** (DESIGN §4.10) — a listed entry that does not render in the menu still needs an incoming link, so listing alone never silences the warning. The message classifies the cause (listed-but-not-menu-rendered with the reason vs not listed; no incoming edge vs only unreachable sources). Active only when `frontend/sitemap.html` exists (it is where roots/`data-entry` opt-outs are declared; its absence is TM-49's finding) — together they close BUG-122 | `pkg/validate/stml_openapi/tm_43_unreachable_page.go` |
| TM-44 | ERROR | a layout HTML still declares `data-nav` while `frontend/sitemap.html` exists (plans/stml/sitemap Phase003, DESIGN §4.9). From Phase003 on the layout menu is emitted from the sitemap tree, so a surviving `data-nav` is a second, drifting menu truth — the message carries the migration guidance (메뉴는 sitemap.html 로 이동; the layout keeps only the menu position). Coexistence is an ERROR, not a WARNING, because tolerated drift contradicts the validate philosophy. Without a sitemap the `data-nav` path stays fully supported (TM-36 resolves its targets) | `pkg/validate/stml_openapi/tm_44_data_nav_with_sitemap.go` |
| TM-46 | ERROR | a sitemap `data-roles` value is not in manifest `backend.auth.roles` — the typo'd role would never match any user's claim and the menu entry would be hidden from everyone (plans/stml/sitemap Phase005). The message states the industry separation explicitly: **menu hiding is not security** — access blocking is Rego's (backend) concern, so a wrong role here only mis-renders the menu. Silent when `backend.auth.roles` is empty (that is TM-47's finding) | `pkg/validate/stml_openapi/tm_46_sitemap_role_unknown.go` |
| TM-47 | ERROR | the sitemap uses `data-roles` but the role-claim wiring is incomplete (plans/stml/sitemap Phase005): `frontend.auth.role_field` is not declared (the menu filter reads `claims[role_field]`), or no action captures `auth.claims.<role_field>` (the claim would never be filled — every role-gated entry stays hidden), or `backend.auth.roles` is empty (no role vocabulary). One ERROR per missing link; silent without any `data-roles` use | `pkg/validate/stml_openapi/tm_47_roles_wiring_missing.go` |
| TM-48 | ERROR | sitemap dynamic menu group structure (plans/stml/sitemap Phase007, DESIGN §4.11 (a)): a group declaring any of the dynamic vocabulary must not sit in a `data-entry` block (the public entry layout renders for signed-out visitors, where the list fetch can never be satisfied) and must declare `data-fetch`, `data-each` and `data-link` together (`data-label-field` is TM-30's required-attribute finding). Field-level judgments live in the TM-01/07/08/30/31/32 sitemap extensions; a complete group's `data-link` target also counts as a TM-43 reachability edge, and its `data-fetch` as an XMO-10/12 consumer | `pkg/validate/stml_openapi/tm_48_sitemap_dynamic_group.go` |
| TM-49 | WARNING | frontend ON with at least one STML page but no `frontend/sitemap.html` — the site structure is undeclared, so menu/breadcrumb derivation and reachability (orphan page) validation are inactive; a central file's absence is itself detectable (one axis of closing BUG-122) | `pkg/validate/stml_openapi/tm_49_sitemap_absent.go` |
| TM-50 | ERROR | a sitemap `data-crumb-field` declaration is unsatisfiable (plans/stml/sitemap Phase006): (a) the page declares no `data-fetch` — the dynamic crumb label is read from fetch data, so it could never fill; (b) the field is not a top-level property of the **first** fetch operation's 2xx response schema (the `responseFields` judgment TM-20 shares — the emitter reads exactly that fetch's data variable); or (c) the field is not a string/integer/number scalar — an object or array cannot render as a crumb label. An unknown page is TM-39's finding, an unknown operationId TM-01's; `data-crumb-field` on a group `<li>` is TM-39's placement rejection | `pkg/validate/stml_openapi/tm_50_crumb_field.go` |
| TM-51 | WARNING | a `frontend/sitemap.html` exists and derives a menu (≥1 menu-rendered entry, the same hidden-subtree-aware `collectMenuRendered` judgment the Phase003 emitter consumes) but no layout exists to host the menu/breadcrumb (`layouts/` empty ∧ `manifest.frontend.defaultLayout` unset ∧ no nav declares `data-layout`), so the derived menu and breadcrumb never render and the emitted `<Breadcrumb>` would be dead code — the inverse of TM-49 (absence vs hostless), closing BUG-129 (plans/stml/sitemap Phase008). Stays silent where TM-12 (defaultLayout set + empty `layouts/`) or TM-41 (nav `data-layout` + empty `layouts/`) already raise an ERROR — no double diagnosis | `pkg/validate/stml_openapi/tm_51_sitemap_no_layout_host.go` |
| TM-52 | WARNING | two or more `data-action` forms on one page declare a same-named `data-field` (collected via `CollectChildActions(page.Children)`, so nested update/create forms inside a `data-fetch` are included — the BUG-127 case `page.Actions` misses). Before BUG-127's codegen fix the generated `id`/`htmlFor` were the bare field name, producing a duplicate DOM id and broken label-for; the id is now form-scoped to `{operationId}-{field}`, so this is a non-blocking advisory that the bare field name is not a stable DOM id for external selectors (E2E/CSS). `data-component:` fields emit no id and are excluded (plans/gen/frontend Phase038, BUG-127) | `pkg/validate/stml_openapi/tm_52_duplicate_field_across_forms.go` |
| TM-53 | WARNING | a `data-bind` cannot render as readable content (plans/gen/frontend Phase037, BUG-126): (a) the field type is **object/array** bound as text — React shows `[object Object]` or a comma-joined string (use a dotted path `User.Name` for an object, `data-each` for an array); (b) the `data-bind` is on a void/media tag codegen cannot bind — `<input>`/`<br>`/`<video>` etc., everything except `<img>` (move it to a text tag or `<img>`); (c) `<img data-bind>` whose field is not a string URL. `boolean` is excluded — the codegen now renders it as `Yes`/`No`. Direct fetch binds consult the response schema (`responseFields`); `data-each` binds consult the array item schema; unknown fields stay silent (TM-06/07 own those). All WARNING — a display-quality issue, not a contract mismatch, so it does not block codegen | `pkg/validate/stml_openapi/tm_53_unrenderable_bind.go` |
| TM-54 | ERROR / WARNING | a `data-action` form's `data-prefill` is unsatisfiable (plans/gen/frontend Phase035, BUG-124). ERROR: the value is not a `data-fetch` operationId on the same page (nested fetches included) — the codegen reads that fetch's data variable (`toLowerFirst(op)+"Data"`), so a typo or a foreign-page op leaves the variable out of scope and the form cannot be wired. WARNING: a form requestBody `data-field` is absent from the prefill 2xx response top-level (the `responseFields` judgment TM-20/TM-50 share) — the codegen fills it with a type-appropriate empty literal so the build still passes, but that input opens blank. An untyped/void prefill response makes no coverage claim; an unknown fetch operationId is TM-01's finding | `pkg/validate/stml_openapi/tm_54_prefill_source.go` |
| TM-55 | WARNING | the canonical edit page forgets prefill (plans/gen/frontend Phase035, BUG-124): the page has a GET-by-id `data-fetch` (a GET that consumes a `route.` path param) and a PUT/PATCH `data-action` carrying `data-field` inputs, yet declares no `data-prefill` — so the form is generated empty and a single-field edit forces re-entering every field. It makes that blank-form generation visible and points to `data-prefill`. Forms already declaring `data-prefill`, field-less actions and non-PUT/PATCH actions stay silent; an unknown operationId is TM-02's finding | `pkg/validate/stml_openapi/tm_55_edit_form_no_prefill.go` |
| TM-56 | WARNING | a PATCH `data-action` form consumes an operation whose requestBody fields are all required (plans/gen/frontend Phase035, BUG-124) — PATCH means partial update, but all-required forces resending every field. The codegen never relaxes zod on its own (required is the OpenAPI decision), so it points back to OpenAPI: mark the optional fields not required and the existing `zod_chain` `.optional()` path applies. Scoped to operations actually consumed by an STML form (de-duplicated by operationId) to avoid noise on non-frontend PATCH APIs | `pkg/validate/stml_openapi/tm_56_patch_all_required.go` |
| TM-57 | ERROR | a state-changing mutation `data-action` (OpenAPI POST/PUT/PATCH/DELETE) does not declare `data-redirect` — where to navigate on success (plans/gen/frontend Phase040, BUG-132). "Where to go after create/update/delete" is an author decision, not a heuristic, so the codegen requires it: with a declared `data-redirect` the generated onSuccess always navigates (and combines invalidate/removeQueries), otherwise the CRUD screen stays on the same form and delete refetches the deleted resource. A bearer login capture action (`data-capture`) is exempt (it drives its own navigation); a GET `data-action` and an unknown operationId (TM-02 reports it) stay silent | `pkg/validate/stml_openapi/tm_57_mutation_redirect_required.go` |
| TM-58 | WARNING | bearer 모드에서 layout `data-logout`에 operationId가 미지정(valueless)이고 OpenAPI에 logout-like operation(operationId에 "logout" 포함, case-insensitive, auth 필요)이 존재하면 서버 logout이 호출되지 않아 refresh token이 revoke되지 않을 수 있음 — TM-38의 bearer-mode 대칭 (plans/gen/frontend Phase045, BUG-145) | `pkg/validate/stml_openapi/tm_58_bearer_logout_op_hint.go` |
| XMO-10 | ERROR | Frontend ON & OpenAPI operationId is never consumed by any STML `data-fetch`, `data-action`, or component `api.<Op>(` call, and is not tagged `no-front` (auth endpoints are no longer auto-excluded) | `pkg/validate/stml_openapi/xmo_10_unconsumed.go` |
| XMO-11 | ERROR | Frontend ON but no STML pages were found (set `frontend.enabled: false` for a backend-only project) | `pkg/validate/stml_openapi/xmo_11_no_stml.go` |
| XMO-12 | WARNING | OpenAPI operationId is tagged `no-front` but is actually consumed by an STML page or component (stale or wrong tag) | `pkg/validate/stml_openapi/xmo_12_no_front_consumed.go` |

## V. Features Internal (`FT-*`)

Internal validation for `features.yaml`. Ensures no duplicate entries.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| FT-01 | ERROR | Duplicate `op` in features.yaml | `pkg/validate/features/ft_01_duplicate_op.go` |
| FT-02 | ERROR | Duplicate `path` in features.yaml | `pkg/validate/features/ft_02_duplicate_path.go` |
| FT-03 | ERROR | features.yaml hash mismatch with specs/.yongol (or .yongol missing) | `pkg/validate/features/ft_03_hash_mismatch.go` |
| FT-10 | ERROR | `has_many` references a table not defined in `tables` | `pkg/validate/features/ft_10_has_many_ref.go` |
| FT-11 | ERROR | `belongs_to` references a table not defined in `tables` | `pkg/validate/features/ft_11_belongs_to_ref.go` |
| FT-12 | WARNING | `has_many` without matching `belongs_to` on the child table | `pkg/validate/features/ft_12_bidirectional.go` |
| FT-13 | ERROR | Feature `table` references a table not defined in `tables` | `pkg/validate/features/ft_13_feature_table_ref.go` |
| FT-16 | ERROR | `features.yaml` missing required `tables` section | `pkg/validate/features/ft_16_tables_required.go` |
| FT-17 | ERROR | Feature missing required `table` field | `pkg/validate/features/ft_17_feature_table_required.go` |

## W. Features ↔ OpenAPI (`XFO-*` / `XOF-*`)

Cross-validation between `features.yaml` and the OpenAPI spec. Ensures that every feature op maps to a real OpenAPI operationId and vice versa.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XFO-01 | ERROR | Features op has no matching OpenAPI operationId | `pkg/validate/features_openapi/xfo_01_op_not_in_openapi.go` |
| XOF-01 | ERROR | OpenAPI operationId is not listed in features.yaml | `pkg/validate/features_openapi/xof_01_op_id_not_in_features.go` |

## X. Features ↔ DDL (`XFD-*`)

Cross-validation between `features.yaml` tables section and DDL. Ensures that declared tables and foreign-key relationships are backed by actual DDL definitions.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XFD-01 | ERROR | Features table has no corresponding DDL file | `pkg/validate/features_ddl/xfd_01_table_exists.go` |
| XFD-02 | ERROR | `belongs_to` relationship has no FK column in child DDL table | `pkg/validate/features_ddl/xfd_02_fk_exists.go` |

## Y. Features ↔ StateMachine (`XFS-*`)

Cross-validation between `features.yaml` tables section and Mermaid stateDiagram. Ensures that declared state values exist in the corresponding stateDiagram. (Note: the `XFS-` prefix here collides with section H's SSaC ↔ Func rules — see the prefix-collision note in the legend.)

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XFS-01 | ERROR | Features table declares a state not present in stateDiagram | `pkg/validate/features_statemachine/xfs_01_states_in_diagram.go` |

---

## Deprecated

Rules that have already been removed from the code or are scheduled for removal. They are currently excluded from the general `yongol validate` catalog.

### Scheduled for removal — follow-up cleanup after retiring OpenAPI `x-pagination` / `x-sort` / `x-filter` / `x-include`

| Rule ID | Previous Description | Reason for Removal |
|---|---|---|
| XDO-1 | `x-sort` column → DDL reference | `x-sort` extension retired. Replaced by standard OpenAPI query parameters + a `sort_by` enum. |
| XDO-2 | `x-sort` column has no index (WARNING) | Same as above. |
| XDO-3 | `x-filter` column → DDL reference | `x-filter` extension retired. Replaced by standard OpenAPI per-column query parameters. |
| XDO-5 | `x-include` target table → DDL reference | `x-include` extension retired. |
| XDO-6 | `x-include` FK constraint missing (WARNING) | Same as above. |
| XDO-8 | Cursor sort default column is not UNIQUE | `x-pagination` cursor mode retired. |
| XOO-4 | `x-include` format error | `x-include` extension retired. |
| XOO-7 | Two or more `x-sort.allowed` entries in cursor mode | `x-pagination` cursor mode retired. |

### Rules already deprecated in code

| Rule ID | Previous Description | Reason for Removal |
|---|---|---|
| XOO-73 | `x-pagination` required query params | Marked deprecated in comments; `x-pagination` fully retired. |
| XOO-74 | `x-pagination` required response fields | Same as above. |
| XOS-68 | OpenAPI `x-pagination` coverage | Replaced by standard parameters + `XSO-18` / `XQS-16` / `XQS-17`. |
| S-52 | `QueryUsageMismatch` — OpenAPI `x-pagination` ↔ SSaC query mismatch | Temporarily replaced by `XOS-68`, then removed alongside the retirement of `x-pagination`. |
| S-53 | SSaC query usage coverage | Same as above. |
| S-54 | `Page[T]`/`Cursor[T]` wrapper → `x-pagination` required | Page/Cursor wrapper types retired. |
| S-55 | `x-pagination` option matching | Wrapper and `x-pagination` retired. |
| S-56 | `x-pagination` option matching (auxiliary) | Same as above. |
| XDN-04 | Each claim's Go type must match the user_table column's DDL-derived Go type (ERROR) | Superseded by XDN-06 (compatibility-matrix based). Not registered in `manifest_ddl.Run`; source files (`xdn_04_*.go`) remain but are unreachable. |
| XDS-13 | SSaC input not present in DDL column (WARNING) | Replaced by XQS-14/16 — the sqlc Params basis is stricter. |
| XDS-14 | SSaC CRUD Input key does not match sqlc Go field name (PascalCase) (ERROR) | Replaced by XQS-14/15/16 — the actual sqlc Params set is a stricter basis. |
| M-1 | `model/` directory and `*.go` files exist | `model/` SSOT and `@dto` fully retired — the sqlc-synthesized row type takes over the model role. |
| M-2 | `model/*.go` struct type matches either a `@dto` or a DDL table | Same as above. |
| TM-45 | Reserved sitemap attribute used (declared but not yet supported) — honest "no effect yet" WARNING | Retired in plans/stml/sitemap Phase007: every reserved attribute graduated to a first-class judgment — `data-roles` (Phase005, TM-46/47), `data-crumb-field` (Phase006, TM-50/TM-39), the dynamic-group vocabulary (Phase007, TM-48 + TM-01/07/08/30/31/32 sitemap extensions). Nothing is left to reserve; the parser keeps no reserved-attribute record. |
| XNS-77 | manifest `auth.claims` present but no `auth.IssueToken` call in SSaC (WARNING) | A missing login is rarely a true positive and produces false positives in verifier-only microservices. It surfaces immediately on the first runtime login attempt, so the static check has little value. |
| SEC-03 | The `<key>` of `backend.rate_limit.endpoints.<key>` must exist as an OpenAPI operationId (ERROR) | Original application-layer rate_limit retired. Replaced by C-7/C-8 (auth-scoped rate_limit mandatory) + codegen `RouteRateLimit` wiring (Phase003). |
| XOH-35 | Hurl path → OpenAPI path exists | Merged into XOH-01 on 2026-04-24 (hurl_openapi re-org); path + method are judged together. |
| XOH-36 | Hurl method → OpenAPI method exists | Merged into XOH-01 on 2026-04-24 — a single diagnostic covers both path and method. |
| XOH-37 | Hurl status code → OpenAPI responses | Moved to XOH-02 on 2026-04-24 and upgraded from WARNING to ERROR. |
| `pkg/generate/hurl/` | Auto-generated smoke/scenario Hurl | Entire package removed 2026-04-24 (plans/gen/hurl/Phase001). Hurl files are now user-owned; yongol only mirrors `specs/tests/` → `arts/tests/`. |

---

## References

- Rule design philosophy, Toulmin defeats graph, Ground mapping: `pkg/validate/README.md`
- Per-category detail: `pkg/validate/<domain>/README.md`
- SSOT syntax and cross-validation rule summary: `manual-for-ai.md` → "Validation" section
