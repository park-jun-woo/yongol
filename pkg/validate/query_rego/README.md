# validate/query_rego — sqlc ↔ Rego cross-validation

Rule prefix: `XQP-` (target = `Q` sqlc, source = `P` Rego policy).

## Rules

| Rule ID | Level | Description |
|---|---|---|
| XQP-30 | ERROR | `@ownership <res>: <table>.<col>` 매핑은 대응 sqlc 쿼리 `OwnerLookup<Resource>` 가 존재해야 함 |

## Scope

Every Rego `@ownership` annotation declares an ownership lookup that the
handler must perform before calling `authz.Check`. Phase003 (ssac/purify)
removed ssac's internal `fmt.Sprintf("SELECT %s FROM %s...")` dynamic SQL;
the handler now invokes a user-authored sqlc query whose canonical name
follows the `OwnerLookup<Resource>` convention declared in
`ssac/pkg/authz/interface.yaml`.

XQP-30 enforces that convention statically so missing queries surface at
`yongol validate` rather than at runtime.
