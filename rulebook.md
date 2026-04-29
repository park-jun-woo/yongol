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
| `T-` | React TSX (frontend/**/*.tsx) |
| `ST-` | Mermaid stateDiagram |
| `P-`, `XPP-` | OPA Rego |
| `F-`, `XFF-` | Func spec (Go AST) |
| `H-` | Hurl scenarios |
| `Q-` | sqlc queries |
| `CORS-` | manifest CORS block |

### Cross SSOT — `X<target><source>-<N>`

`<target>` = the SSOT referenced by the Lookup key (ground truth), `<source>` = the SSOT making the claim.

| SSOT | Code | SSOT | Code |
|---|---|---|---|
| OpenAPI | `O` | DDL | `D` |
| SSaC | `S` | StateMachine | `M` |
| Rego | `P` | Manifest | `N` |
| Hurl | `H` | Func | `F` |
| Authz | `A` | sqlc | `Q` |
| TSX | `T` | | |

Example: SSaC → OpenAPI (SSaC is the claim, OpenAPI is the ground truth) → `XOS-`.
Example: TSX → OpenAPI (TSX is the claim, OpenAPI is the ground truth) → `XOT-`.

## Level

| Level | Behavior | Meaning |
|---|---|---|
| `ERROR` | `yongol validate` fails → exit code ≠ 0 | Consistency violation. Must be fixed. |
| `WARNING` | Does not fail; reported only | Design intent should be reconfirmed. |

## Source

The `Source` column of each rule row is a Go file path relative to the repo root. Example: `pkg/validate/ssac/s_27_var_declared.go`.

---

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
| S-37 | WARNING | FK-referencing `@get` should be followed by an `@empty` guard | `pkg/validate/ssac/s_37_fk_reference_guard.go` |
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
| S-61 | ERROR | Result variable names must not collide with codegen reserved names (`server`, `ctx`, `err`, etc.) | `pkg/validate/ssac/s_61_codegen_reserved_var.go` |
| S-62 | ERROR | Result variable is unreferenced in subsequent sequences | `pkg/validate/ssac/s_62_unused_result_var.go` |
| S-63 | WARNING | `@get []T` list endpoint has no pagination params and no `// @no-pagination` | `pkg/validate/ssac/s_63_list_no_pagination.go` |
| S-64 | ERROR | `@empty` / `@exists` Target must reference a Model (struct), not a scalar field | `pkg/validate/ssac/s_64_empty_exists_model_only.go` |
| S-67 | ERROR | `@eval` Func signature must be `func(req T) bool` | `pkg/validate/ssac/s_67_eval_func_signature.go` |
| S-68 | ERROR | `@eval` requires an explicit STATUS code (no default) | `pkg/validate/ssac/s_68_eval_status_required.go` |
| S-69 | ERROR | `@eval` Func must exist in Func Spec or built-in | `pkg/validate/ssac/s_69_eval_func_exists.go` |
| S-70 | ERROR | `@post` / `@put` Inputs value must not be a standalone reserved source (`currentUser`, `request`, `query`, `message`); use dotted form. `@call` exempt | `pkg/validate/ssac/s_70_post_put_blob_input_forbidden.go` |
| XSS-11 | WARNING | `@result` type is plural | `pkg/validate/ssac/xss_11_plural_result_type.go` |
| XSS-38 | ERROR | `@call` function name starts with a lowercase letter (uppercase recommended) | `pkg/validate/ssac/xss_38_call_func_lowercase.go` |
| XSS-47 | WARNING | `@call` argument source variable is undefined | `pkg/validate/ssac/xss_47_call_source_var_undefined.go` |
| XSS-57 | ERROR | `@publish` topic must have a matching `@subscribe` | `pkg/validate/ssac/xss_57_publish_to_subscribe.go` |
| XSS-58 | ERROR | `@subscribe` topic must have a matching `@publish` | `pkg/validate/ssac/xss_58_subscribe_to_publish.go` |
| XSS-59 | ERROR | `@subscribe` message fields must match `@publish` payload | `pkg/validate/ssac/xss_59_subscribe_fields.go` |

## B. Manifest

`manifest.yaml` loading and base schema validation.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| C-2 | ERROR | `apiVersion` must be `yongol/v1` | `pkg/validate/manifest/c_02_api_version.go` |
| C-3 | ERROR | `kind` must be `Project` | `pkg/validate/manifest/c_03_kind.go` |
| C-4 | ERROR | `metadata.name` is required (empty value forbidden) | `pkg/validate/manifest/c_04_metadata_name.go` |
| C-5 | ERROR | `backend.module` is required (empty value forbidden) | `pkg/validate/manifest/c_05_backend_module.go` |
| C-6 | ERROR | `backend.auth` is required — yongol does not support auth-free backends (use a static site generator + CDN for public dynamic content) | `pkg/validate/manifest/c_06_backend_auth_required.go` |
| CORS-01 | ERROR | `allow_origins=["*"]` combined with `allow_credentials=true` is forbidden | `pkg/validate/manifest/cors_01_wildcard_credentials.go` |
| OBS-001 | ERROR | `backend.observability.metrics.path` must be an absolute path starting with `/` | `pkg/validate/manifest/obs_01_metrics_path.go` |
| OBS-002 | ERROR | `backend.observability.metrics.path` must not collide with an OpenAPI path | `pkg/validate/manifest/obs_02_metrics_path_not_openapi.go` |
| OBS-003 | ERROR | `backend.observability.tracing.exporter` must be one of `otlp`/`stdout`/`noop` (when enabled=true) | `pkg/validate/manifest/obs_03_tracing_exporter.go` |
| OBS-004 | ERROR | `backend.observability.tracing.sample_rate` must be within `[0.0, 1.0]` (when enabled=true) | `pkg/validate/manifest/obs_04_tracing_sample_rate.go` |
| SEC-201 | ERROR | `backend.auth.mode=cookie\|hybrid` combined with `csrf.enabled=false` is forbidden (leaves CSRF attack surface exposed) | `pkg/validate/manifest/sec_201_cookie_without_csrf.go` |
| SEC-301 | WARNING | `backend.security_headers.csp.directives.default-src` containing `*` or `'unsafe-eval'` weakens CSP protection | `pkg/validate/manifest/sec_301_csp_permissive.go` |
| SEC-302 | WARNING | `backend.security_headers.hsts.max_age < 15552000` (180 days) fails the HSTS preload minimum | `pkg/validate/manifest/sec_302_hsts_short.go` |
| SEC-401 | ERROR | Literal `backend.auth.secret` is forbidden (git leak / rotation impossible) — only `secret_env` is allowed | `pkg/validate/manifest/sec_401_jwt_secret_env_required.go` |
| SEC-402 | WARNING | `backend.auth.access_token_ttl > 30m` exceeds the OWASP recommended upper bound (expanded blast radius) | `pkg/validate/manifest/sec_402_access_ttl_upper_bound.go` |
| SEC-403 | ERROR | `backend.auth.mode` must be one of `cookie` / `bearer` / `hybrid` (defaults to `cookie` when unspecified) | `pkg/validate/manifest/sec_403_auth_mode_enum.go` |

## C. OpenAPI Internal

OpenAPI self-consistency (based on the document parsed by kin-openapi).

| Rule ID | Level | Description | Source |
|---|---|---|---|
| O-1 | ERROR | Path parameter name conflict | `pkg/validate/openapi/o_01_path_param_conflict.go` |
| O-2 | ERROR | Path parameter case conflict | `pkg/validate/openapi/o_02_path_param_case_conflict.go` |
| O-3 | ERROR | Path template parameter declaration missing | `pkg/validate/openapi/o_03_path_template_param.go` |
| O-4 | ERROR | Operation is missing `operationId` | `pkg/validate/openapi/o_04_op_id_required.go` |
| XOO-71 | WARNING | Password-like fields have no `minLength` | `pkg/validate/openapi/xoo_71_password_no_min_length.go` |
| XOO-72 | WARNING | Email-like fields have no `format` | `pkg/validate/openapi/xoo_72_email_no_format.go` |

## D. Query / sqlc

sqlc query file self-consistency.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| Q-1 | ERROR | `-- name:` annotation is required | `pkg/validate/query/q_01_name_required.go` |
| Q-2 | ERROR | Cardinality (`:one` / `:many` / `:exec` / `:execrows`) is required | `pkg/validate/query/q_02_cardinality.go` |
| Q-3 | ERROR | Query name must be PascalCase | `pkg/validate/query/q_03_name_pascalcase.go` |
| Q-4 | WARNING | `:many` query is missing `LIMIT` | `pkg/validate/query/q_04_many_limit.go` |
| Q-5 | ERROR | `DELETE` statement requires `WHERE` | `pkg/validate/query/q_05_delete_where.go` |
| Q-6 | ERROR | `UPDATE` statement requires `WHERE` | `pkg/validate/query/q_06_update_where.go` |
| Q-7 | WARNING | `SELECT *` on a table that has `@sensitive` columns | `pkg/validate/query/q_07_select_star_sensitive.go` |
| Q-8 | ERROR | Declared parameter is unused in the query body | `pkg/validate/query/q_08_unused_param.go` |
| Q-9 | ERROR | `:exec` query returns `SELECT` | `pkg/validate/query/q_09_select_on_exec.go` |
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
| XDD-61 | WARNING | Columns matching sensitive patterns (`password` / `secret` / `hash` / `token`) are missing the `@sensitive` annotation | `pkg/validate/ddl/xdd_61_sensitive_no_annotation.go` |

## F. DDL ↔ OpenAPI

Cross-consistency between DDL tables/columns and OpenAPI schemas / extension fields.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XDO-9 | ERROR | OpenAPI property does not correspond to any DDL table column (ghost) | `pkg/validate/openapi_ddl/xdo_09_ghost_property.go` |
| XDO-67 | ERROR | DDL `VARCHAR(n)` ↔ OpenAPI `maxLength` missing/inconsistent | `pkg/validate/openapi_ddl/xdo_67_max_length_varchar.go` |
| XDO-68 | ERROR | DDL `CHECK IN` ↔ OpenAPI `enum` missing/inconsistent | `pkg/validate/openapi_ddl/xdo_68_check_in_enum.go` |
| XDO-69 | ERROR | DDL `CHECK` allowed values ↔ OpenAPI `enum` values mismatch | `pkg/validate/openapi_ddl/xdo_69_check_values_enum.go` |
| XDO-70 | WARNING | OpenAPI `maxLength` exceeds DDL `VARCHAR(n)` | `pkg/validate/openapi_ddl/xdo_70_max_length_exceeds_varchar.go` |
| XDO-75 | ERROR | OpenAPI optional + DDL `NOT NULL` + no `DEFAULT` | `pkg/validate/openapi_ddl/xdo_75_optional_not_null_no_default.go` |
| XDO-76 | WARNING | OpenAPI required + DDL nullable | `pkg/validate/openapi_ddl/xdo_76_required_nullable.go` |
| XDO-77 | ERROR | DDL column type ↔ OpenAPI field type mismatch | `pkg/validate/openapi_ddl/xdo_77_column_type_mismatch.go` |
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
| XSO-16 | ERROR | OpenAPI operationId must be used as a SSaC function (coverage) | `pkg/validate/openapi_ssac/xso_16_op_id_to_func.go` |
| XSO-18 | ERROR | OpenAPI response field must be used in a SSaC `@response` (coverage) | `pkg/validate/openapi_ssac/xso_18_response_field_used.go` |
| XSO-20 | ERROR | OpenAPI response field must be used in a shorthand `@response` (coverage) | `pkg/validate/openapi_ssac/xso_20_shorthand_field_used.go` |

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
| XSF-62 | WARNING | Func spec must be used in SSaC (coverage) | `pkg/validate/ssac_func/xsf_62_func_spec_used.go` |

## I. SSaC ↔ StateMachine

Cross-consistency between SSaC `@state` and Mermaid stateDiagram.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XMS-24 | ERROR | `@state` DiagramID must exist in stateDiagram | `pkg/validate/ssac_statemachine/xms_24_state_diagram_exists.go` |
| XMS-25 | ERROR | `@state` transition event must be defined in stateDiagram | `pkg/validate/ssac_statemachine/xms_25_state_event.go` |
| XSM-23 | ERROR | stateDiagram transition event must exist as a SSaC function | `pkg/validate/ssac_statemachine/xsm_23_transition_to_func.go` |
| XSM-26 | WARNING | Function participating in a state transition has no `@state` declaration | `pkg/validate/ssac_statemachine/xsm_26_missing_state_guard.go` |
| XSM-27 | WARNING | POST/PUT/DELETE on a stateful resource must declare either `@state` or `// @state-neutral` | `pkg/validate/ssac_statemachine/xsm_27_state_intent_declaration.go` |

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
| XAS-60 | ERROR | `@auth` input field must exist on the Authz `CheckRequest` struct | `pkg/validate/ssac_authz/xas_60_auth_input_field.go` |

## N. OpenAPI ↔ Manifest

Cross-consistency between OpenAPI security schemes and manifest middleware configuration.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XNO-50 | ERROR | OpenAPI securityScheme must map to existing manifest middleware | `pkg/validate/openapi_manifest/xno_50_security_scheme_middleware.go` |
| XNO-52 | ERROR | Endpoint security must reference an existing manifest middleware name, and the middleware block itself must exist | `pkg/validate/openapi_manifest/xno_52_security_middleware.go` |
| XON-51 | ERROR | manifest middleware must map to an existing OpenAPI securityScheme (coverage) | `pkg/validate/openapi_manifest/xon_51_middleware_security_scheme.go` |
| SEC-04 | ERROR | The `<key>` of `backend.http.overrides.<key>` must exist as an OpenAPI operationId | `pkg/validate/openapi_manifest/sec_04_http_overrides_operation_id.go` |
| SEC-101 | ERROR | generate-time: the generated main.go must register the request_id and error_envelope middleware immediately after the router, in that order | `pkg/generate/gogin/boot/collect_active_blocks.go` |

## N2. Manifest auth claims ↔ DDL columns (XDN-*)

Cross-consistency between `manifest.backend.auth` (user_table + claims mapping) and the DDL user table parsed from `db/*.sql`. JWT claims are a fixed, finite key set; they must map to typed columns on a named table — not to a generic JSONB blob.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XDN-01 | ERROR | `backend.auth.user_table` is required when auth is active (`auth.type != "none"`) | `pkg/validate/manifest_ddl/xdn_01_user_table_required.go` |
| XDN-02 | ERROR | `backend.auth.user_table` must reference a table parsed from `db/*.sql` | `pkg/validate/manifest_ddl/xdn_02_user_table_exists.go` |
| XDN-03 | ERROR | Each `backend.auth.claims.<Field>: <col>[:<type>]` mapping's column must exist on the user_table | `pkg/validate/manifest_ddl/xdn_03_claim_column_exists.go` |
| XDN-04 | ERROR | Each claim's Go type (`int64` / `string` / `bool`, default `string`) must match the user_table column's DDL-derived Go type | `pkg/validate/manifest_ddl/xdn_04_claim_column_type.go` |

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

## P. SSaC ↔ DDL

Bidirectional validation that SSaC `@result` / `@input` matches DDL tables/columns.

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XDS-12 | WARNING | `@result` type must match a sqlc row type or DDL table | `pkg/validate/ssac_ddl/xds_12_result_no_ddl_table.go` |
| XSD-55 | ERROR | DDL table must be referenced in a SSaC `@model` (coverage) | `pkg/validate/ssac_ddl/xsd_55_ddl_to_model_ref.go` |

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

## R3. Hurl ↔ State Machine

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XOH-05 | WARNING | Hurl call order satisfies state machine transitions | `pkg/validate/hurl_statemachine/xoh_05_state_transition_order.go` |

## R4. Hurl ↔ Manifest

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XOH-06 | WARNING | Protected Hurl call preceded by an auth step | `pkg/validate/hurl_manifest/xoh_06_auth_precondition.go` |
| XOH-07 | WARNING | Cookie-mode mutation includes `X-CSRF-Token` header | `pkg/validate/hurl_manifest/xoh_07_csrf_on_mutation.go` |

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

## Q. TSX (React frontend)

Self-consistency for React `.tsx` files (frontend/**/*.tsx).

| Rule ID | Level | Description | Source |
|---|---|---|---|
| T-1 | WARNING | Import target file for `@/components/` or relative paths must exist | `pkg/validate/tsx/t_01_component_file.go` |

## Q2. TSX ↔ OpenAPI

Validates that `apiClient.<op>()` calls in TSX match the OpenAPI contract. One-directional (TSX → OpenAPI). operationIds not consumed by OpenAPI are intentionally not reported, since there are many other possible consumers (mobile, CLI, partners, etc.).

| Rule ID | Level | Description | Source |
|---|---|---|---|
| XOT-1 | ERROR | `<op>` in `apiClient.<op>()` must exist in the OpenAPI operationId set | `pkg/validate/tsx_openapi/xot_01_operation_id.go` |
| XOT-2 | ERROR | Path/query argument object keys of apiClient calls must exist in OpenAPI parameters | `pkg/validate/tsx_openapi/xot_02_parameter_match.go` |
| XOT-3 | WARNING | `useForm().register('x')` field must exist in the OpenAPI request body schema of that page's mutation | `pkg/validate/tsx_openapi/xot_03_form_field.go` |

---

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
| MIG-006 | ERROR | The `-- YONGOL_SCHEMA_HASH:` header in `specs/db/.generated_schema.sql` does not match the sha256 of the body (user-edited = drift) | `pkg/validate/migration/mig_006_snapshot_drift.go` |

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
| XDS-13 | SSaC input not present in DDL column (WARNING) | Replaced by XQS-14/16 — the sqlc Params basis is stricter. |
| XDS-14 | SSaC CRUD Input key does not match sqlc Go field name (PascalCase) (ERROR) | Replaced by XQS-14/15/16 — the actual sqlc Params set is a stricter basis. |
| M-1 | `model/` directory and `*.go` files exist | `model/` SSOT and `@dto` fully retired — the sqlc-synthesized row type takes over the model role. |
| M-2 | `model/*.go` struct type matches either a `@dto` or a DDL table | Same as above. |
| XNS-77 | manifest `auth.claims` present but no `auth.IssueToken` call in SSaC (WARNING) | A missing login is rarely a true positive and produces false positives in verifier-only microservices. It surfaces immediately on the first runtime login attempt, so the static check has little value. |
| SEC-03 | The `<key>` of `backend.rate_limit.endpoints.<key>` must exist as an OpenAPI operationId (ERROR) | Application-layer rate_limit retired altogether — responsibility shifted to the CDN/WAF/Gateway layer. Only the hardcoded `FixedRateLimit` guard (/auth/refresh) is kept. |
| XOH-35 | Hurl path → OpenAPI path exists | Merged into XOH-01 on 2026-04-24 (hurl_openapi re-org); path + method are judged together. |
| XOH-36 | Hurl method → OpenAPI method exists | Merged into XOH-01 on 2026-04-24 — a single diagnostic covers both path and method. |
| XOH-37 | Hurl status code → OpenAPI responses | Moved to XOH-02 on 2026-04-24 and upgraded from WARNING to ERROR. |
| `pkg/generate/hurl/` | Auto-generated smoke/scenario Hurl | Entire package removed 2026-04-24 (plans/gen/hurl/Phase001). Hurl files are now user-owned; yongol only mirrors `specs/tests/` → `arts/tests/`. |

---

## References

- Rule design philosophy, Toulmin defeats graph, Ground mapping: `pkg/validate/README.md`
- Per-category detail: `pkg/validate/<domain>/README.md`
- SSOT syntax and cross-validation rule summary: `manual-for-ai.md` → "Cross-Validation Rules Catalog" section
