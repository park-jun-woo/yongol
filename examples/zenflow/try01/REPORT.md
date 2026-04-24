# ZenFlow try01 — End-to-End Report

## Summary

| Stage | Result |
|---|---|
| `yongol validate specs` | ✅ 0 errors, 0 warnings |
| `yongol generate specs arts` | ✅ initial migration `0001_initial.sql` (16 ops), backend + frontend emitted |
| `go build ./...` (backend) | ✅ clean |
| Postgres + migrations | ✅ applied via `psql` (golang-migrate expects paired up/down files; single-file migration applied directly) |
| `hurl --test smoke.hurl` | ✅ 8/8 requests passed, 99 ms |

## Timing

- Start (first `manifest.yaml` line): **2026-04-23T15:04:17Z**
- End (all 4 steps green): **2026-04-23T15:14:59Z**
- Wall-clock: **~10 min 42 s** (single green run, no retries needed after the one fix described below)

## Scope

Scoped to base `zenflow.md` entities and flow (no add01–add07 addons):

- 5 DDL tables — organizations, users, workflows, actions, execution_logs (all `BIGINT GENERATED ALWAYS AS IDENTITY` per D-8)
- 8 endpoints — Signup, Login, GetCurrentUser, CreateWorkflow, ListWorkflows, AddAction, ActivateWorkflow, ExecuteWorkflow
- `states/workflow.md` — draft → active, paused → active, active → active (ExecuteWorkflow)
- `policy/authz.rego` — 5 allow rules (CreateWorkflow, ListWorkflows, AddAction, ActivateWorkflow, ExecuteWorkflow)
- `func/billing/` — `checkCredits` (402 when balance ≤ 0), `spend` (402 when insufficient)
- 1 TSX page — `frontend/pages/WorkflowsPage.tsx`
- 3 scenario hurl files + auto-generated `smoke.hurl`

## Issues Encountered

### 1. Codegen: `api.SignupJSONBodyPlanType` named-enum type

- `plan_type: { type: string, enum: [free, pro, enterprise] }` in OpenAPI generated a named string type `api.SignupJSONBodyPlanType`, which could not be assigned to the DDL-derived plain `string` field `OrganizationCreateParams.PlanType`.
- Workaround: dropped the `enum:` facet from the request body, keeping `maxLength: 32`. SSaC passes `request.plan_type` through unmodified.
- Root cause is yongol-side: the SSaC codegen does not inject a string cast when the OpenAPI field has been narrowed into an enum type. Worth filing a BUG-NNN if the enum round-trip should be supported.

### 2. `golang-migrate` expects paired `.up.sql` / `.down.sql`

- Yongol emits a single `NNNN_<desc>.sql` wrapped in `BEGIN; ... COMMIT;` — the golang-migrate CLI rejected it with `error: first .: file does not exist` (its filename convention is `NNNN_<desc>.up.sql`).
- Applied directly via `psql -f` — clean BEGIN/COMMIT so it is transactional.
- Manual `manual-for-ai.md` says "DB application is the user's responsibility — use golang-migrate, flyway, or similar." Consider documenting the filename mismatch, or adopting `.up.sql` naming.

### 3. Seed `id=0` sentinel rows are DDL-level `INSERT`s, not captured by migration diff

- `db/organizations.sql` and `db/users.sql` contain `INSERT INTO ... (id, ...) OVERRIDING SYSTEM VALUE VALUES (0, ...)` seed statements. The migration diff only emits schema changes, so sentinel rows are not applied by `0001_initial.sql`.
- Worked around by running the INSERTs manually after migration. For ordinary user signup this does not matter (FK points at a real org), but documentation should make the gap explicit, or the migration generator should pick DML up.

### 4. smoke.hurl Login precedes Signup

- Auto-generated `smoke.hurl` starts with `POST /auth/login {"email":"{{smoke_email}}",...}` before its Signup block. Requires a pre-seeded user whose email matches the `smoke_email` hurl variable, plus sufficient credits to pass `ActivateWorkflow`/`ExecuteWorkflow` later in the flow.
- Handled by pre-seeding via `POST /auth/signup` with email `smoke@zenflow.test`, password `Password1234!`, `credits_balance: 100`, then running `hurl --variable smoke_email=smoke@zenflow.test`.

## Commands Used (reference)

```bash
# Validate + generate
cd ~/.clari/repos/fullend/yongol/examples/zenflow/try01
yongol validate specs
yongol generate specs arts

# Postgres
docker run -d --rm --name zenflow_pg \
  -e POSTGRES_PASSWORD=zenflow -e POSTGRES_USER=zenflow -e POSTGRES_DB=zenflow \
  -p 5441:5432 postgres:16
PGPASSWORD=zenflow psql -h localhost -p 5441 -U zenflow -d zenflow \
  -v ON_ERROR_STOP=1 -f arts/db/migrations/0001_initial.sql
PGPASSWORD=zenflow psql -h localhost -p 5441 -U zenflow -d zenflow \
  -c "INSERT INTO organizations (id,name,plan_type,credits_balance) OVERRIDING SYSTEM VALUE VALUES (0,'system','free',0) ON CONFLICT DO NOTHING;
      INSERT INTO users (id,org_id,email,password_hash,role,name) OVERRIDING SYSTEM VALUE VALUES (0,0,'nobody@system','','system','Nobody') ON CONFLICT DO NOTHING;"

# Server
cd arts/backend && go build -o /tmp/zenflow_server ./cmd
DATABASE_URL='postgres://zenflow:zenflow@localhost:5441/zenflow?sslmode=disable' \
  JWT_SECRET='test-jwt-secret-0123456789abcdef0123456789' \
  OPA_POLICY_PATH=./policy/authz.rego \
  /tmp/zenflow_server &

# Seed smoke user
curl -X POST http://localhost:8080/auth/signup \
  -H 'Content-Type: application/json' \
  -d '{"email":"smoke@zenflow.test","password":"Password1234!","name":"Smoke","org_name":"SmokeCo","plan_type":"pro","credits_balance":100}'

# Smoke test
hurl --test --variable host=http://localhost:8080 \
  --variable smoke_email=smoke@zenflow.test \
  arts/tests/smoke.hurl
```

## Result

```
file: arts/tests/smoke.hurl
Success (8 request(s) in 97 ms)
Executed files:    1
Executed requests: 8 (80.8/s)
Succeeded files:   1 (100.0%)
Failed files:      0 (0.0%)
```
