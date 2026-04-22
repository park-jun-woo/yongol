# Preserve Contract

`yongol generate` does not overwrite user-modified generated code on re-runs. Preservation is **file-level**: if a user edits a generated file, it enters the **preserved** state and subsequent generates skip it. Drift is caught by `yongol validate <specs> <arts>` as ERROR, not by the filesystem.

## `//ff:checked` Annotation

Injected by yongol at generation. Users do not write it.

```go
//ff:func feature=service type=handler control=sequence
//ff:what ActivateWorkflow — HTTP handler
//ff:checked llm=yongol-gen hash=35ed498e
package service
```

- `llm=yongol-gen` — writer identity. Distinguishes yongol-gen output from filefunc llmc.
- `hash=<8hex>` — first 4 bytes of SHA-256 over the first `func`'s signature + body (excluding `init`). Compatible with filefunc `CalcBodyHash`.
- Normalization: CRLF -> LF, BOM removed, trailing newline ensured. Annotation block + generator banner excluded.

Files with no hashable function (type-only / const-only) receive no annotation and are excluded from preserve judgement.

### `//ff:preserve reason=` (Optional)

User may annotate preservation intent above the package declaration.

```go
//ff:preserve reason="custom observability log for premium accounts, 2026-04-21"
//ff:checked llm=yongol-gen hash=540810dd
```

Not required. `yongol status` flags preserved files without a reason as `reason: <none>`.

### No Block Markers

No `// BEGIN PRESERVE` / `// END PRESERVE`. Filefunc `F1` (1 file, 1 function) -> file = one function -> block markers unnecessary. Edit partially by leaving the entire file preserved and managing it manually.

## Flow

```
1. Does target path exist?
   - No  -> write (os.WriteFile)
   - Yes -> DetectPreserved
            - No annotation -> write (external artifact)
            - Hash matches  -> write (untouched)
            - Hash mismatch -> SKIP + slog.Info "skipping preserved file"
```

After: non-preserved files have their `//ff:checked` line refreshed with the new hash; preserved files are untouched so they remain preserved at the next validate.

sqlc-generated `internal/db/querier.go` / `db.go` and `specs/func/` mirror files are never given a preserve marker.

## PRV Rules

Run during `yongol validate <specs> <arts>` or `yongol generate` when `<arts>` is present.

### Contract Drift

| Rule ID | Level | Description |
|---|---|---|
| PRV-01 | ERROR | Preserved-file function signature drifts from SSOT. Expected: `(ctx context.Context, request api.<OpID>RequestObject) (api.<OpID>ResponseObject, error)` |
| PRV-02 | ERROR | Preserved file references a sqlc query / `@call` target / DDL field that no longer exists |

### Runtime Safety Guards

Applied only to preserved files.

| Rule ID | Level | Description |
|---|---|---|
| PRV-10 | ERROR | Use of `panic(` (exempt: inside `init()`, or `// nolint:panic`) |
| PRV-11 | ERROR | `ctx.Value("currentUser").(T)` not in comma-ok form |
| PRV-12 | ERROR | Ignored error from `json.Unmarshal` / `yaml.Unmarshal` |
| PRV-13 | ERROR | Ignored error from `sql.Rows.Scan` / `sql.Row.Scan` |
| PRV-14 | ERROR | `x[0]` without `len` guard |
| PRV-15 | ERROR | Inline `m[k].Field` without comma-ok |
| PRV-16 | ERROR | Field access on `Get*()` / `Find*()` return without nil check |
| PRV-17 | ERROR | Missing `defer Close` after `os.Open` / `db.Query` / `http.Get` |

### Exemptions

- Inside `init()` — PRV-10
- `// nolint:panic` on same or previous line — PRV-10
- `// nolint:prv-NN` on same or previous line — specific rule

Full rule catalog: [rulebook.md](../rulebook.md) Section S.

## Limitations

- Renaming an `operationId` changes the generated file name — the previous preserved file becomes an orphan. Move modifications manually before renaming.
- `internal/db/querier.go` / `db.go` are sqlc-owned; edits are overwritten on regeneration.
- The unit is the whole file. Import-only edits still preserve the entire file.
- PRV-02 can yield false positives on oapi-codegen request objects; suppress with `// nolint:prv-02` and file an issue.

## Releasing Preservation

Delete the file; next generate rewrites from the SSOT.

```bash
rm arts/backend/internal/service/activate_workflow.go
yongol generate specs arts
```

No CLI flag, no metadata file.

## CLI Impact

### `yongol status <specs> [<arts>]`

SSOT summary + preserved/drift dashboard when `<arts>` is given. Read-only.

### `yongol validate <specs> [<arts>]`

With `<arts>`, PRV-01/02 contract drift and PRV-10~17 runtime guards run as the final step. Any ERROR exits 1.

### `yongol generate <specs> <arts>` (default)

Preserved files are skipped. Skip logs via `slog.Info` to stderr:

```
INFO skipping preserved file path=arts/backend/internal/service/add_action.go
```
