# zenflow try01 — 완료 보고서

## 요약

- **작성 일자**: 2026-04-24
- **담당 에이전트 모델**: Claude Opus 4.7 (1M context)
- **최종 상태**: **PASS**

clean-room 방식으로 `examples/zenflow/try01/` 를 처음부터 구축하여 `yongol
validate` → `yongol generate` → `go build` → `migrate up` → `hurl --test
smoke.hurl` 의 전체 체인을 한 번에 통과시켰다.

## 단계별 결과

| 단계 | 결과 | 비고 |
|---|---|---|
| validate | PASS | 0 errors, 0 warnings |
| generate | PASS | 202 artifact 파일 생성 (backend + frontend + tests mirror + db migration) |
| go build | PASS | `go build ./...` clean, 경고 없음 |
| DB migrate | PASS | `0001_initial.up.sql` 1 개 |
| app 기동 | PASS | `:8080` 에서 정상 기동, gin 라우팅 7 경로 |
| hurl smoke | PASS | smoke.hurl 7/7 requests succeeded (183 ms) |
| hurl scenario-happy-path | PASS (bonus) | 5/5 requests succeeded (96 ms) |

## 작성한 SSOT 목록

### endpoint 5 개 (모두 `operationId`)
1. `Register` — `POST /auth/register`
2. `Login` — `POST /auth/login`
3. `CreateWorkflow` — `POST /workflows` (protected, admin)
4. `ListWorkflows` — `GET /workflows` (protected, admin, offset pagination)
5. `ActivateWorkflow` — `POST /workflows/{id}/activate` (protected, admin, state-guarded)

### table 3 개
1. `users` — (id, email, password_hash, role, created_at) + @sensitive on password_hash + id=0 sentinel row
2. `workflows` — (id, owner_id, title, trigger_event, status, created_at) + status CHECK + FK to users(id)
3. `refresh_tokens` — canonical from `ssac/pkg/auth/interface.yaml`, `@archived` (XNA-90 canonical)

### state diagram 1 개
- `states/workflow.md` — `[*] --> draft --> active: ActivateWorkflow`

### authz rule 1 개 (admin 통합)
- `CreateWorkflow / ListWorkflows / ActivateWorkflow` on resource `workflow` — role=admin

### ssac 시퀀스 5 개
- `service/auth/register.ssac`
- `service/auth/login.ssac`
- `service/workflow/create_workflow.ssac`
- `service/workflow/list_workflows.ssac`
- `service/workflow/activate_workflow.ssac`

### hurl 파일 2 개
- `specs/tests/smoke.hurl` — 7 steps: register(201) → bad login(401) → register(201) → login(200) → CreateWorkflow(201) → ListWorkflows(200) → ActivateWorkflow(200)
- `specs/tests/scenario-happy-path.hurl` — 5 steps (register → login → create → list → activate)
  - `scenario-*.hurl` 파일명이 `KindScenario` 감지에 필요 (H-2 경고 해소용)

## manifest 설정

- `backend.auth.mode: bearer` — 선택 근거: cookie+CSRF 보다 smoke hurl 작성 단순 (`Authorization: Bearer {{token}}` 한 헤더로 커버). CSRF 관련 XOH-07 경고도 회피
- `backend.auth.claims: {ID: user_id:int64, Email: email, Role: role}` — id 가 int64 (GENERATED ALWAYS AS IDENTITY) 와 정합
- `backend.auth.roles: [admin]` — XPN-64 회피를 위해 `member` 는 제거 (rego 에서 사용 안 함)
- `backend.middleware: [bearerAuth]`

infra 블록(session/cache/queue/file) 은 전부 선언하지 않음 → 메모리 기본. 스코프에 불필요.

## yongol 버그 발견

### 후보 1: `CREATE INDEX ... USING GIN` 의 USING 절이 migration 에서 누락

- 위치: `pkg/generate/migration/` 또는 `pkg/parser/ddl/`
- 증상: `refresh_tokens.sql` 의 `CREATE INDEX ... USING GIN (claims)` 가 migration 에서 `CREATE INDEX ... (claims)` 로 emit (USING 절 증발)
- 영향: 본 smoke 범위에선 `@>` 연산을 쓰지 않아 기능 실패 없음. 운영 환경에서는 `auth.RefreshTokenRevokeAll` 성능 저하 가능
- 회피: 없음 (기록만)

### 후보 2: `@auth` + `@ownership` 조합 시 생성형 endpoint 에서 PgxErrNoRows → 403

- 위치: `pkg/generate/gogin/` 핸들러 빌더
- 증상: `@auth "CreateWorkflow" "workflow" {}` + rego `@ownership workflow: workflows.owner_id` 이 있으면 `create_workflow.go` 가 `qtx.OwnerLookupWorkflow(ctx, 0)` 를 먼저 호출. id=0 workflow 가 없어서 `no rows` → 403. 신규 리소스 생성에서 ownership 조회는 논리적 모순
- 회피 (본 프로젝트에 적용): rego 에서 `@ownership workflow: ...` 를 제거. 역할만으로 허가하므로 `data.owners.workflow` 가 필요 없었음
- 제안: `@auth` 에 `ResourceID` 가 없거나 0 이면 ownerLookup 을 생략하는 게 합리적

### 후보 3: snapshot 이 남아있으면 arts 삭제 후 재생성 시 migration 이 regenerate 되지 않음

- 위치: `pkg/generate/migration/`
- 증상: 첫 generate 후 `specs/db/.generated_schema.sql` 이 남은 상태에서 `arts/` 만 삭제 후 재 generate 하면 `mode=noop` 으로 판정, migration 파일이 재생성 안 됨 → 새 DB 를 세팅할 방법이 사라짐
- 회피: snapshot 도 함께 삭제 후 generate
- 제안: `arts/db/migrations/0001_initial.up.sql` 의 물리적 존재 여부도 같이 고려

## 계획 대비 차이

- zenflow.md 의 `credits_balance`/`plan_type`/`organizations` 는 범위 밖으로 둠 (402 Payment Required, 복수 조직 tenant 분리 생략)
- `ExecuteWorkflow`, `actions`, `execution_logs` 생략 — smoke 통과엔 create+activate 면 충분
- zenflow.md rego 의 `ListWorkflows: is_same_org` 는 admin 전체 허가로 단순화 (org 개념이 스코프에 없음)
- frontend TSX 는 작성하지 않음 — smoke 에 무관. `? tsx SSOT not detected` 로 통과

## 재현 방법

### 1. Postgres 기동

```bash
docker run --rm -d \
  --name zenflow-try01-pg \
  -e POSTGRES_PASSWORD=testpass \
  -e POSTGRES_DB=zenflow \
  -p 55432:5432 \
  postgres:16-alpine

# 준비 대기
until docker exec zenflow-try01-pg pg_isready -U postgres -d zenflow; do sleep 1; done
```

### 2. validate + generate

```bash
cd ~/.clari/repos/fullend/yongol
go run ./cmd/yongol validate ~/.clari/repos/fullend/yongol/examples/zenflow/try01/specs
go run ./cmd/yongol generate  ~/.clari/repos/fullend/yongol/examples/zenflow/try01/specs \
                               ~/.clari/repos/fullend/yongol/examples/zenflow/try01/arts
```

### 3. Migrate + Build + Run

```bash
export DATABASE_URL='postgres://postgres:testpass@localhost:55432/zenflow?sslmode=disable'
export JWT_SECRET='smoke-test-secret-at-least-32-characters-long!!'
export OPA_POLICY_PATH=~/.clari/repos/fullend/yongol/examples/zenflow/try01/arts/backend/policy/authz.rego
export BACKEND_OBSERVABILITY_METRICS_ENABLED=false

migrate -path ~/.clari/repos/fullend/yongol/examples/zenflow/try01/arts/db/migrations \
        -database "$DATABASE_URL" up

cd ~/.clari/repos/fullend/yongol/examples/zenflow/try01/arts/backend
go mod tidy
go build -o /tmp/zenflow-try01-server ./cmd/
/tmp/zenflow-try01-server &
```

### 4. Hurl smoke

```bash
hurl --test --variable host=http://localhost:8080 \
     ~/.clari/repos/fullend/yongol/examples/zenflow/try01/arts/tests/smoke.hurl
```

기대 결과: `Succeeded files: 1 (100.0%)`, 7/7 requests.

### 5. 정리

```bash
pkill -f zenflow-try01-server
docker stop zenflow-try01-pg
```

포트 / env 요약:

| 항목 | 값 |
|---|---|
| Postgres port | 55432 (host) → 5432 (container) |
| App port | 8080 |
| `DATABASE_URL` | `postgres://postgres:testpass@localhost:55432/zenflow?sslmode=disable` |
| `JWT_SECRET` | 32 자 이상 아무 문자열 |
| `OPA_POLICY_PATH` | `arts/backend/policy/authz.rego` 의 절대 경로 |
| `BACKEND_OBSERVABILITY_METRICS_ENABLED` | `false` (Prometheus 비활성; 선택) |

## 산출물 위치

- SSOT 19 파일: `examples/zenflow/try01/specs/`
- 생성물 202 파일: `examples/zenflow/try01/arts/` (backend + frontend + tests mirror + migrations)
- snapshot (yongol 관리): `examples/zenflow/try01/specs/db/.generated_schema.sql`
