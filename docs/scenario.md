# Hurl — Scenario and Invariant Tests

End-to-end HTTP tests in Hurl syntax. yongol adds no DSL on top — see the [Hurl manual](https://hurl.dev/docs/manual.html) for syntax.

## Location

```
<project-root>/tests/
├── scenario-*.hurl    # business flows (user-authored)
└── invariant-*.hurl   # cross-endpoint invariants (user-authored)
```

`yongol generate` additionally emits a `smoke.hurl` endpoint smoke test.

## Naming Rules

- `scenario-*.hurl` — business flow across multiple endpoints (signup -> login -> order -> pay).
- `invariant-*.hurl` — verify invariants still hold after state changes (balance recomputation, session invalidation).
- `.feature` files (Gherkin) are not supported — H-1 ERROR.

## Execution

```bash
cd artifacts/<project>
hurl --test --variable host=http://localhost:8080 tests/*.hurl
```

Each `.hurl` runs independently; captures are not shared across files.

## Cross-SSOT Links (Hurl -> OpenAPI, one-way)

| Rule | Level |
|---|---|
| Path in `.hurl` exists in OpenAPI `paths` | ERROR |
| HTTP method defined on that path in OpenAPI | ERROR |
| Expected status code declared in OpenAPI responses | WARNING |

## Further Reading

- [Hurl manual](https://hurl.dev/docs/manual.html)
- [docs/openapi.md](./openapi.md)
- [docs/ssac.md](./ssac.md)
- [rulebook.md](../rulebook.md)
