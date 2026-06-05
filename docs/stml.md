# STML — Statechart-aware Decision Binding (enhancement reference)

This is the reference for the STML *state enhancement*: the attributes and
guard syntax that bind UI decisions to yongol's existing state-machine SSOT.
STML stays decision-only — it references the state machine in
`states/*.md` (Mermaid `stateDiagram-v2`) and the OpenAPI contract; it never
defines a second state machine. The generated React is a disposable
projection.

For the base STML model (`data-fetch` / `data-action` / `data-bind` / `data-each`
/ `data-param-*` / `data-field` / `data-component`), see the STML section of
`manual-for-ai.md`.

## Location

`<project-root>/frontend/*.html` — flat, no subdirectories.

## New / extended attributes

| Attribute | Purpose | Example |
|---|---|---|
| `data-enabled-when` | Action enablement decision — the action is disabled unless the guard holds | `<button data-action="ActivateWorkflow" data-enabled-when="workflow.status=draft">` |
| `data-invalidates` | Effect declaration — GET queries to refetch after the action succeeds (space-separated operationIds) | `<div data-action="CreateWorkflow" data-invalidates="ListWorkflows">` |
| `data-state` (extended) | Conditional display, now accepting the full guard syntax below | `data-state="workflow.status=active && currentUser.Role=owner"` |

- `data-enabled-when` declares *when an action is available*. Codegen renders
  `disabled={!(...)}`; the source of truth is the guard, not the rendered code.
- `data-invalidates` declares *what goes stale on success*. Each listed
  operationId must be a GET and is refetched (TanStack Query invalidation).
  The target is an `operationId` (already a contract) — not a hardcoded HTTP
  call.

## Guard syntax

`data-state` and `data-enabled-when` accept a deliberately restricted,
Turing-incomplete expression language: comparisons, logical combinators,
negation, and parentheses only. No function calls, arithmetic, or ternaries —
so guards stay finite and statically verifiable.

```
guard     := term (("&&" | "||") term)*
term      := "!"? atom
atom      := ref op value | ref "." lifecycle | "(" guard ")"
ref       := <model> "." <Field>            // workflow.status, currentUser.Role
op        := "=" | "!=" | ">" | "<" | ">=" | "<="
value     := <state-id> | <number> | <quoted-string> | <enum-literal>
lifecycle := "loading" | "error" | "empty"
```

Examples:

```html
<!-- resource-state branch — references a state in states/workflow.md -->
<div data-state="workflow.status = active">
  <span data-bind="workflow.activatedAt"></span>
</div>

<!-- fetch lifecycle (TanStack Query vocabulary) -->
<div data-state=".loading">Loading…</div>
<div data-state=".error">Failed to load</div>
<div data-state="workflows.empty">No workflows</div>

<!-- composite guard -->
<button data-action="ApproveOrder"
        data-enabled-when="order.status=pending && currentUser.Role=manager">
  Approve
</button>
```

### Backward compatibility

A single comparison (`field=value`), a lifecycle suffix
(`.loading` / `.error` / `.empty`, including `items.empty`), and a bare field
keep their existing behavior unchanged. Only conditions containing a combinator
(`&&`, `||`), a leading `!`, or parentheses are routed through the guard parser
and validated by TM-17.

## Cross-validation rules (summary)

| Rule | Level | Cross target | Contract |
|---|---|---|---|
| TM-14 | ERROR | OpenAPI | `data-enabled-when` guard ref model is a top-level property of some page `data-fetch` response schema |
| TM-15 | ERROR | stateDiagram | guard comparison state value exists in the matching stateDiagram |
| TM-16 | ERROR | OpenAPI | `data-invalidates` operationId is defined in OpenAPI and is a GET |
| TM-17 | ERROR | STML internal | `data-state` guard with a combinator parses under the guard EBNF |
| TM-18 | WARNING | stateDiagram | the `data-action` transition is legal from the state its `data-enabled-when` requires |
| XMO-10 | ERROR | OpenAPI | Frontend ON & operationId is consumed by some STML page/component **or** tagged `no-front` |
| XMO-11 | ERROR | manifest | Frontend ON requires at least one STML page (else `frontend.enabled: false`) |
| XMO-12 | WARNING | OpenAPI | operationId tagged `no-front` must not actually be consumed (stale tag) |

An operationId counts as **consumed** when an STML `data-fetch`/`data-action`
references it, **or** when a referenced `data-component` calls
`api.<operationId>(` inside its `.tsx` file under `frontend/components/`. So a
form whose inner widget is a custom component still consumes the op it calls.

TM-14, TM-16, TM-17 are implemented in `pkg/validate/stml_openapi/`; the
stateDiagram cross-checks TM-15 and TM-18 in `pkg/validate/stml_statemachine/`.
See `rulebook.md` (section U) for the authoritative rows and source paths.

## Attribution (homage)

STML invents no new formalism. It borrows the *semantics* of proven standards
and the *ideas* of prior art, and names every source — hidden borrowing is
plagiarism, acknowledged borrowing is homage.

| STML element | Source | Borrowed | Not borrowed (boundary) |
|---|---|---|---|
| `data-state` state-reference semantics | **Harel Statecharts** (David Harel, 1987) / **W3C SCXML** (Recommendation, 2015) / the **Mermaid `stateDiagram-v2`** yongol already uses | the *structural semantics* of state · transition · guard · hierarchy · parallelism; "the truth of state lives in the state machine" | SCXML's XML serialization and `<datamodel>` / `<script>` / `<assign>` / `<send>` executable content (implementation leakage) |
| Turing-incomplete guard syntax | **MEL** (Manifesto, Jungwoo Jung / eggplantiny) — "deliberate Turing-incompleteness, guaranteed termination" + SCXML `cond` | restricting guards to *finite, termination-guaranteed* expressions rather than arbitrary code | MEL's arbitrary expressions and runtime paths (`$runtime.*`) |
| `data-enabled-when` (action enablement decision) | **MEL** `available when` / `dispatchable when` guards | declaring "when is this action available" as a *decision* | the MEL DSL itself and `@meta` UI-framework leakage |
| `data-invalidates` (effect declaration) | **MEL** "effect is declaration, not execution" boundary + **TanStack Query** (React Query, Tanner Linsley) query invalidation | declaring the post-success effect (invalidation) as a *decision* | MEL's `effect api.fetch({url,method})` HTTP hardcoding |
| fetch lifecycle `.loading` / `.error` / `.empty` | **TanStack Query** state vocabulary | the loading / error / empty *UI-branch decision* | — (already present in STML; source named only) |
| form validation (existing) | **React Hook Form** + **Zod** + OpenAPI `requestBody` | (existing Phase009 / 021) | — |
| design philosophy "Agent proposes, World verifies" | **Manifesto / MEL** (Jungwoo Jung) | deterministic gate = STML decides, validate judges | — |

> Source: `files/stml/09-STML-보강설계.md` §1. The same table appears in the
> STML section of `manual-for-ai.md` so the attribution is visible in both code
> and docs.
