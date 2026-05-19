# features.yaml — Feature Catalog

Optional SSOT that lists every project feature as a flat YAML array keyed by `operationId`. Serves as a human-readable checklist and cross-validates bidirectionally against OpenAPI.

## Location

`<project-root>/features.yaml`

## Format

```yaml
features:
  - op: CreateWorkflow
    path: POST /workflows
    desc: Create a new workflow in draft state

  - op: ActivateWorkflow
    path: POST /workflows/{id}/activate
    desc: Transition workflow to active (requires credits > 0)

  - op: ListWorkflows
    path: GET /workflows
    desc: List workflows for the current org
```

## Fields

| Field | Required | Description |
|---|---|---|
| `op` | Yes | `operationId` (PascalCase). Must match an OpenAPI `operationId`. |
| `path` | Yes | HTTP method + URI pattern (e.g. `POST /workflows/{id}/activate`). |
| `desc` | Yes | One-line human-readable description of the feature. |

## Validation Rules

### Internal

| Rule | Level | Description |
|---|---|---|
| FT-01 | ERROR | Duplicate `op` in features.yaml. |
| FT-02 | ERROR | Duplicate `path` in features.yaml. |
| FT-03 | ERROR | `specs/.yongol` missing or SHA-256 hash mismatch with features.yaml. |

### Cross-validation (features ↔ OpenAPI)

| Rule | Level | Description |
|---|---|---|
| XFO-01 | ERROR | `op` in features.yaml has no matching `operationId` in OpenAPI. The feature is declared but not implemented. |
| XOF-01 | ERROR | `operationId` in OpenAPI is not listed in features.yaml. The endpoint exists but is not declared as a feature. |

Both directions are ERROR level. When `features.yaml` is present, the feature list and OpenAPI must be in exact agreement.

## Hash Lock (`specs/.yongol`)

`yongol init` generates a `specs/.yongol` file that stores the SHA-256 hash of the features.yaml used to create the project:

```yaml
hashes:
  features.yaml: sha256:<64hex>
```

`yongol validate` checks this hash via FT-03:

- **features.yaml exists, .yongol missing** → FT-03 ERROR: `specs/.yongol not found`
- **features.yaml exists, .yongol exists, hashes differ** → FT-03 ERROR: `features.yaml was modified after baseline`
- **features.yaml exists, hashes match** → pass
- **features.yaml absent** → skip (no error)

Only `yongol init` creates `.yongol`; `yongol validate` reads it but never writes. To reset after intentional features.yaml changes: `rm specs/.yongol` then re-run `yongol init`.

## Scaffolding via `yongol init`

```bash
yongol init MyApp features.yaml "My workflow SaaS"
```

Reads features.yaml and generates SSOT stubs for each feature:

- **OpenAPI**: `specs/api/openapi.yaml` with path + operationId stubs
- **SSaC**: `specs/service/{domain}/{Op}.ssac` stub files
- **Rego**: `specs/policy/authz.rego` with allow rule stubs
- **Hurl**: `specs/tests/smoke.hurl` with request stubs
- **Hash lock**: `specs/.yongol` with SHA-256

Does NOT generate: DDL, sqlc queries, states, func specs, STML.

## When features.yaml is absent

If `features.yaml` does not exist in the project root, yongol skips features validation entirely (`? features  SSOT not detected`). No errors, no warnings. All other SSOT validations proceed normally.

## Purpose

### For humans

`features.yaml` is the minimum artifact a human must review. Three fields per feature — operation name, HTTP path, one-line description. A product owner can read it without knowing OpenAPI, DDL, or SSaC.

### For AI agents

`features.yaml` serves as an implementation checklist. An agent can work through the list one feature at a time, and `yongol validate` confirms completeness:

- XFO-01 fires when a feature is declared but its OpenAPI endpoint is missing → "implement this next"
- XOF-01 fires when an endpoint exists but isn't in the feature list → "register this feature"

### For progress tracking

Feature count is a natural progress metric:

```
features.yaml: 32 features
OpenAPI: 28 operationIds
XFO-01 fires on 4 features → 87.5% complete
```

## Example

See [`examples/zenflow/zenflow.yaml`](../examples/zenflow/zenflow.yaml) for a complete 32-feature catalog covering auth, CRUD, state machines, webhooks, templates, scheduling, audit, dashboard, batch operations, external API integration, and conditional updates.
