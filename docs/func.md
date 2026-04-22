# Func Spec — Custom Business Functions

Optional SSOT for custom Go functions called from SSaC via `@call`. Use for domain-specific logic not covered by the built-ins (`auth`, `session`, `cache`, `file`, `crypto`, `mail`, `text`, `image`, `storage`).

## Location

```
<project-root>/
└── func/<pkg>/*.go     # Func Spec (optional)
```

## Fixed Signature

```
func FuncName(req FuncNameRequest) (FuncNameResponse, error)
```

Request / Response types live in the **same file** as Go structs.

## Annotations

### `@func`

`// @func camelCaseName` above the function. The camelCase name must match the SSaC `@call` reference. **One `@func` per file** — split multi-function packages into one file per function (`set.go`, `get.go`, ...).

```go
package billing

// @func holdEscrow
// @description Locks funds in escrow

type HoldEscrowRequest struct {
    GigID    int64
    Amount   int64
    ClientID int64
}

type HoldEscrowResponse struct {
    TransactionID int64
}

func HoldEscrow(req HoldEscrowRequest) (HoldEscrowResponse, error) {
    return HoldEscrowResponse{TransactionID: req.GigID}, nil
}
```

SSaC reference:

```go
// @call billing.HoldEscrowResponse r = billing.HoldEscrow({GigID: gig.ID, Amount: gig.Budget, ClientID: gig.ClientID})
```

### `@error`

`// @error NNN` sets the default HTTP error status when `@call` fails.

```go
// @func verifyPassword
// @error 401
```

Priority (high -> low):
1. Explicit in `.ssac`: `@call auth.VerifyPassword({...}) 500`
2. `@error` annotation
3. Default 500

## Purity Rule

Applies to every `@call` function:

- Allowed: file I/O (`io`, `bufio`, `os`), session/cache read/write.
- Forbidden: DB (`database/sql`, `lib/pq`, `jackc/pgx`).
- Forbidden: network (`net/http`, `net/rpc`, `grpc`).

Uniform — no per-package exceptions.

## Import-Path Convention

SSaC imports `internal/<pkg>`. Specs live under `specs/<project>/func/<pkg>/`; `yongol generate` copies them to `artifacts/<project>/backend/internal/<pkg>/`.

## Fallback Chain

1. `specs/<project>/func/<pkg>/` — project-specific
2. `ssac/pkg/<pkg>/` — built-in
3. Neither -> ERROR + skeleton suggestion

## Cross-SSOT Links

| Link | Validation |
|---|---|
| `@func camelCaseName` -> SSaC `@call pkg.Func` | Identical |
| `Request` fields -> SSaC args | Name + type |
| `Response` fields -> SSaC result-variable fields | Match |
| Purity rule | Forbidden DB/network imports |
| Struct fields -> OpenAPI response schema | Field-name match |

## Prefer Built-ins

For common concerns, check the built-ins first (password hashing, JWT, TOTP, file upload, email, slug). Full list in [docs/ssac.md](./ssac.md) under "Built-in Function Packages".

## Further Reading

- [docs/ssac.md](./ssac.md)
- [docs/ddl.md](./ddl.md)
- [docs/manifest.md](./manifest.md)
- [docs/openapi.md](./openapi.md)
- [rulebook.md](../rulebook.md)
