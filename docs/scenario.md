# Hurl — User-authored smoke / scenario / invariant tests

All hurl files live under `specs/tests/` and are **user-authored**. yongol
does not auto-generate any hurl file. Instead, `yongol validate`
cross-checks your hurl against OpenAPI / state machine / manifest.auth
(rules XOH-01 ~ XOH-09). yongol adds no DSL on top of Hurl — see the
[Hurl manual](https://hurl.dev/docs/manual.html) for syntax.

## Location

```
<project-root>/specs/tests/
├── smoke.hurl         # endpoint smoke (user-authored)
├── scenario-*.hurl    # business flows (user-authored)
└── invariant-*.hurl   # cross-endpoint invariants (user-authored)
```

`yongol generate` mirrors `specs/tests/` → `arts/tests/` without
modification. Run hurl against the mirrored copy (or directly against
specs — content is identical).

## Naming Rules

- `smoke.hurl` (singular, fixed filename) — project-root smoke. One only.
- `scenario-<name>.hurl` — business flow across multiple endpoints
  (signup → login → order → pay).
- `invariant-<name>.hurl` — verify invariants still hold after state
  changes (balance recomputation, session invalidation).
- `.feature` files (Gherkin) are not supported — H-1 ERROR.

## Execution

```bash
cd artifacts/<project>
hurl --test --variable host=http://localhost:8080 tests/*.hurl
```

Each `.hurl` runs independently; captures are not shared across files.

## Authoring templates

Copy-paste starting points. Adjust operationIds, field names, and
credentials to match your project.

### Cookie + CSRF (`backend.auth.mode: cookie`, the 2026 default)

The generated middleware sets the JS-readable CSRF cookie (`XSRF-TOKEN`)
on safe requests (GET/HEAD/OPTIONS) — it never sends the token as a
response header. Capture the cookie after a safe request, then duplicate
it into the `X-XSRF-TOKEN` header (double-submit) on every state-changing
request — including the auth POSTs, since the generated CSRF middleware
is registered globally with no default exemptions. Cookie and header
names follow `backend.auth.csrf.cookie_name` / `header_name` overrides
when set.

```hurl
GET {{host}}/api/workflows
HTTP 200
[Captures]
csrf: cookie "XSRF-TOKEN"

POST {{host}}/auth/register
Content-Type: application/json
X-XSRF-TOKEN: {{csrf}}
{ "email": "smoke+{{newUuid}}@example.com", "password": "p@ssw0rd!" }
HTTP 201
[Asserts]
jsonpath "$.user.id" isInteger

POST {{host}}/auth/login
Content-Type: application/json
X-XSRF-TOKEN: {{csrf}}
{ "email": "smoke+{{newUuid}}@example.com", "password": "p@ssw0rd!" }
HTTP 200

POST {{host}}/api/workflows
Content-Type: application/json
X-XSRF-TOKEN: {{csrf}}
{ "name": "Demo" }
HTTP 201
```

### Bearer (`backend.auth.mode: bearer`)

```hurl
POST {{host}}/auth/login
Content-Type: application/json
{ "email": "smoke@example.com", "password": "p@ssw0rd!" }
HTTP 200
[Captures]
token: jsonpath "$.access_token"

GET {{host}}/api/workflows
Authorization: Bearer {{token}}
HTTP 200
```

## Cross-validation (Hurl ↔ other SSOTs)

| Rule | Level | Check |
|---|---|---|
| XOH-01 | ERROR | Request URL + method declared in OpenAPI |
| XOH-02 | ERROR | Response status declared in OpenAPI responses |
| XOH-03 | ERROR | Request body field in OpenAPI schema |
| XOH-04 | ERROR | Assert jsonpath reachable in response schema |
| XOH-05 | WARNING | Call order satisfies state transitions |
| XOH-06 | WARNING | Protected endpoint preceded by an auth step |
| XOH-07 | WARNING | Cookie-mode mutation carries the manifest-resolved CSRF header (default `X-XSRF-TOKEN`) |
| XOH-08 | ERROR | Capture jsonpath reachable in response schema |
| XOH-09 | WARNING | Captured variable is used later |

## Further reading

- [Hurl manual](https://hurl.dev/docs/manual.html)
- [rulebook.md — Section R / R2 / R3 / R4](../rulebook.md)
- [docs/openapi.md](./openapi.md)
- [docs/ssac.md](./ssac.md)
