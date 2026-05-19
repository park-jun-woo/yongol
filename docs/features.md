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

### Cross-validation (features ↔ OpenAPI)

| Rule | Level | Description |
|---|---|---|
| XFO-01 | ERROR | `op` in features.yaml has no matching `operationId` in OpenAPI. The feature is declared but not implemented. |
| XOF-01 | ERROR | `operationId` in OpenAPI is not listed in features.yaml. The endpoint exists but is not declared as a feature. |

Both directions are ERROR level. When `features.yaml` is present, the feature list and OpenAPI must be in exact agreement.

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
