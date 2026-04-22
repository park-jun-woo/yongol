# OPA Rego — Authorization Policy

Authorization policies in OPA Rego. Every SSaC `@auth` emits an auto-generated `authz.Check(...)` call and must have a matching `allow` rule.

## Location

`<project-root>/policy/*.rego` — **OPA v1 syntax** (every rule uses `if`).

## yongol Conventions

- Every `allow` rule must declare both `input.action` and `input.resource`. Missing either -> XPS-28 ERROR (crosscheck can't pair it with an `@auth`).
- `@ownership` annotation declares DB-based owner lookup.
- Fixed input schema: `input.action`, `input.resource`, `input.resource_id`, `input.claims.<field>` (mirrors `manifest.backend.auth.claims`).
- `data.owners.<resource>` is loaded from DB per request following `@ownership`.

## @ownership

```rego
# @ownership course: courses.instructor_id
# @ownership lesson: courses.instructor_id via lessons.course_id
# @ownership review: reviews.user_id
```

| Format | Meaning |
|---|---|
| `resource: table.column` | Direct lookup |
| `resource: table.column via join_table.fk` | JOIN lookup |

Default authz package: `github.com/park-jun-woo/ssac/pkg/authz`. Override via `manifest.authz.package`.

## 5 Allow Patterns

| Pattern | Conditions |
|---|---|
| unconditional | `input.action` + `input.resource` only |
| role-based | + `input.claims.role` |
| owner-based | + `data.owners.<resource>[input.resource_id] == input.claims.user_id` |
| role + owner | both |
| multiple actions | use a set: `input.action in {"A", "B"}` |

## Example

```rego
package authz

# @ownership gig: gigs.owner_id
# @ownership proposal: gigs.owner_id via proposals.gig_id

default allow := false

allow if {
    input.claims.role == "admin"
}

allow if {
    input.action in {"UpdateGig", "DeleteGig", "PublishGig"}
    input.resource == "gig"
    input.claims.role == "client"
    data.owners.gig[input.resource_id] == input.claims.user_id
}

allow if {
    input.action in {"AcceptProposal", "RejectProposal"}
    input.resource == "proposal"
    data.owners.proposal[input.resource_id] == input.claims.user_id
}
```

## Runtime

- `OPA_POLICY_PATH` — path to the `.rego` file. Server startup fails if unset.

## Cross-SSOT Links

| Link | Validation |
|---|---|
| `input.action` + `input.resource` pairs -> SSaC `@auth` | Every `@auth` must have a matching allow rule (XPS-28) |
| `@ownership table.column` -> DDL | Table/column existence |
| `@ownership ... via join_table.fk` -> DDL | Join table / FK existence |
| `input.claims.<field>` -> manifest claims | Declaration |
| Role literals -> `manifest.backend.auth.roles` | Declaration |

## Further Reading

- [OPA Policy Language](https://www.openpolicyagent.org/docs/policy-language)
- [docs/ssac.md](./ssac.md)
- [docs/manifest.md](./manifest.md)
- [docs/ddl.md](./ddl.md)
- [rulebook.md](../rulebook.md)
