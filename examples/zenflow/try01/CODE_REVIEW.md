# zenflow/try01 arts/ 코드 품질 리뷰

## 실행 정보
- 시작 시각: 2026-04-25T02:07:07+09:00
- 종료 시각: 2026-04-25T02:11:29+09:00
- 소요 시간: 00:04:22 (약 262초)
- Claude Code: 2.1.119
- Model: claude-opus-4-7 (Opus 4.7, 1M context)
- 검토 대상: `examples/zenflow/try01/arts/` (257 파일)
- 산출물 생성 시각: 2026-04-24T16:58:34Z (yongol v0.2.5)
- 대상 커밋: master @ 37d1d60 기준 (리뷰 중 파일 수정 없음)

## Executive Summary

생성물은 빌드/정적 분석 레벨에서는 깔끔하게 통과한다 (`go build ./...` / `go vet ./...` 무오류). 백엔드는 미들웨어 스택(Request-ID, 보안 헤더, rate limit, body limit, OpenAPI 런타임 검증), 트랜잭션 경계, 타이밍 공격 완화 로그인, OPA 기반 인가, slog 민감 필드 자동 마스킹 등 운영 수준의 기본기가 매우 견고하게 깔려 있다. 반면 **프론트엔드는 현재 상태로 쓸 수 없는 수준**이다 — 페이지가 router 에 연결되지 않았고, TanStack Query v5 와 호환되지 않는 v4 문법을 쓰고 있으며, UI 프리미티브 4종(`Button/Card/Input/Form`)이 `any` 타입 스텁으로 남아 strict TypeScript 의미가 사라졌다. Hurl 시나리오는 happy path 6개만 있어 zenflow.md §7 의 `@invariant` 2 건(403 Tenant Breach / 402 Insufficient Credits)을 전혀 검증하지 않는다. 비즈니스 로직 한 군데는 SSaC 명세에서부터 누락돼 있다 — `ExecuteWorkflow` 가 actions 를 조회만 하고 `processAction` 루프를 돌지 않는다.

**점수 (10점 만점)**
- Correctness: **7** (빌드/정적 분석 통과, 다만 ExecuteWorkflow 가 action 루프 미구현, Register 409 미처리, DeductCredit rows-affected 미확인)
- Security: **8** (미들웨어 스택 견고, 민감 필드 redact 자동, JWT 시크릿 길이 검증 / 다만 oapi-codegen 기본 ErrorHandler 가 내부 에러 문자열 노출)
- Go idiom: **8** (context 전파 일관, early return, errors.Is, gofmt 준수 / context key 가 string 이라 최소한의 흠)
- Frontend: **3** (페이지 미연결, v4/v5 혼용, UI 컴포넌트 any, 상태 처리 부재)
- Migration / DDL: **6** (BIGINT identity + sentinel 일관, 인덱스 적절 / down 은 의도적으로 빈 스텁, zenflow.md 요구 UUID 와 괴리)
- Tests (Hurl): **4** (happy path 만, 네거티브 0건)

**Top 3 Critical**
1. Frontend `App.tsx` 가 `LoginPage` / `WorkflowsPage` 를 라우트에 전혀 마운트하지 않음 — 생성된 페이지가 런타임에 절대 렌더링되지 않는다.
2. `ExecuteWorkflow` 서비스가 `actions` 만 조회하고 루프를 돌지 않음 — zenflow.md §5 "Loop: @call worker.processAction" 요구 미충족.
3. Hurl 시나리오에 네거티브 케이스가 0건 — zenflow.md §7 의 invariant 2 건(403/402) 미검증.

**Top 3 Quick Wins**
1. `App.tsx` 에 `<Route path="/login">` / `<Route path="/workflows">` 만 붙여도 라우팅이 살아난다.
2. `components/ui/{Button,Card,Input,Form}.tsx` 를 `Table.tsx` / `Badge.tsx` / `Modal.tsx` 수준의 타입드 구현으로 교체.
3. Hurl 에 `@invariant` 2건을 추가 (다른 org 토큰으로 403, credits=0 org 로 402).

---

## 영역별 상세

### Backend (Go+Gin)

#### 빌드 결과
- `cd arts/backend && go build ./...` → **무오류**
- `go vet ./...` → **무오류**
- `go test ./...` → `no test files` (테스트 부재)

#### 잘 된 점
1. **보안 미들웨어 스택이 충실**: Request-ID, ErrorEnvelope, Prometheus, SecurityHeaders(HSTS/CSP/XFO/Referrer/Permissions), BodyLimit/MultipartLimit(`http.MaxBytesReader`), kin-openapi 런타임 요청 검증. (`arts/backend/internal/middleware/security_headers.go`, `arts/backend/internal/middleware/request_validator.go`, `arts/backend/internal/middleware/body_limit.go`)
2. **민감 필드 자동 마스킹**: sqlc row 에 `LogValue() slog.Value` 자동 생성(`arts/backend/internal/db/users_log.go`, `arts/backend/internal/db/refresh_tokens_log.go`) 로 slog 가 `password_hash` / `token_hash` 를 `[REDACTED]` 로 자동 치환.
3. **타이밍 공격 완화 로그인**: 사용자 미존재 시에도 `auth.VerifyPassword(..., DummyHash)` 호출(`arts/backend/internal/service/login.go:22`).
4. **트랜잭션 경계 & 롤백 규약**: 모든 쓰기 핸들러가 `tx.Begin → defer Rollback(IsNot ErrTxClosed) → tx.Commit` 패턴(`arts/backend/internal/service/create_workflow.go:26-44`, `activate_workflow.go:28-81`).
5. **JWT 시크릿 길이 검증**: main.go 에서 32자 미만이면 `os.Exit(1)`(`arts/backend/cmd/main.go:49-55`).
6. **Refresh Token 회전 / 재사용 탐지**: `Consume()` 가 revoked 행을 다시 만나면 `auth.ErrRefreshTokenReused` 를 반환하여 패밀리 전체 revoke 트리거 가능(`arts/backend/internal/infra/auth/postgres.go:61-79`).
7. **Graceful shutdown**: SIGINT/SIGTERM 에서 10s 타임아웃 shutdown(`arts/backend/cmd/run_server_with_graceful_shutdown.go`).

#### 이슈

##### Critical
- **C-B1 `ExecuteWorkflow` 가 action 루프를 돌지 않음**
  - 위치: `arts/backend/internal/service/execute_workflow.go:60-67`
  - 관찰: `actions, err := qtx.ActionListByWorkflowID(...)` 로 조회한 뒤 그대로 `OrganizationDeductCredit` 과 `ExecutionLogCreate` 호출. zenflow.md §5 의 "Loop: @call worker.processAction" 미구현.
  - 원인: SSaC (`specs/specs/service/workflow/execute_workflow.ssac`) 자체에 `@foreach` 구문 없음. processAction 관련 `@call` 도 없음.
  - 조치: SSaC 에 `@foreach actions as a { @call worker.processAction({...}) }` 추가 필요. codegen 측은 @foreach 지원이 이미 있다면 SSOT 보강만으로 해결.

- **C-B2 Register 409(Email 충돌) 미처리 → 500 leak**
  - 위치: `arts/backend/internal/service/register.go:37-38`
  - 관찰: `users_email_key` UNIQUE 제약이 있고 `api.Register409JSONResponse` 타입도 생성돼 있지만, 핸들러는 `UserCreate` 의 에러를 바로 `return nil, err` 로 상위에 던져 oapi-codegen 기본 ErrorHandler 가 `{"msg": err.Error()}` + 500 을 반환한다 (`arts/backend/internal/api/register_handlers_with_options.gen.go:18-21`).
  - 원인: sqlc unique-violation → SQLSTATE 23505 체크 패턴이 codegen 에 없음. SSaC 에 `@catch unique_violation 409` 같은 구문 필요.
  - 조치: SSaC 에서 duplicate 처리 선언 → codegen 에서 `pgconn.PgError` 인스펙션 → 409 envelope 로 변환.

##### Major
- **M-B1 `OrganizationDeductCredit` 의 RowsAffected 미확인**
  - 위치: `arts/backend/internal/db/organization_deduct_credit.sql.go:11-14` + `arts/backend/internal/db/queries_organization_deduct_credit.sql.go:15-18`
  - 관찰: SQL 은 `WHERE id=$2 AND credits_balance >= $1` 로 atomic guard 걸어놨으나, sqlc 생성 함수 시그니처가 `error` 만 반환 → 0 rows affected 에서도 nil error. 호출자(`execute_workflow.go:63`)가 감지 불가 → credits < 1 상태에서도 ExecutionLog 가 success 로 남는다.
  - 원인: sqlc `:exec` 타입이 pgconn.CommandTag 를 버림. `:execrows` 로 선언하면 int64 를 받을 수 있음.
  - 조치: sqlc 쿼리 `:exec` → `:execrows` 로 변경 + 핸들러에서 `if rowsAffected == 0 { return 402 }`.

- **M-B2 oapi-codegen 기본 ErrorHandler 가 내부 에러 문자열 노출**
  - 위치: `arts/backend/internal/api/register_handlers_with_options.gen.go:17-21`
  - 관찰: `c.JSON(statusCode, gin.H{"msg": err.Error()})` — ErrorEnvelope 스키마와 불일치하고, DB 에러 원문(테이블명 / SQLSTATE / 파라미터 등)이 500 응답으로 그대로 새어 나간다.
  - 원인: `RegisterHandlersWithOptions` 호출 시 `GinServerOptions.ErrorHandler` 를 주입하지 않음. codegen 템플릿이 기본 handler 를 fallback 으로 씀.
  - 조치: main.go 에서 `api.RegisterHandlersWithOptions(r, strict, api.GinServerOptions{ErrorHandler: middleware.StrictErrorHandler})` 형태로 ErrorEnvelope 래퍼 주입. 혹은 codegen 이 envelope-aware fallback 을 emit 하도록 수정.

- **M-B3 Context key 가 string (collision risk)**
  - 위치: `arts/backend/internal/middleware/bearerauth.go:96` + 모든 service 핸들러의 `ctx.Value("currentUser")`
  - 관찰: `context.Value` 에 string key "currentUser" 사용. go vet 은 `-composites` 기본 leniency 로 통과하지만 lint 도구(staticcheck SA1029)가 잡는 패턴. 서드파티 미들웨어와 key 충돌 가능.
  - 조치: `type contextKey string; const currentUserKey contextKey = "currentUser"` 로 언래핑.

##### Minor
- **m-B1 Pause / Archive 전이 부재**: `arts/backend/internal/statemachine/workflow.go:7-10` 의 전이 맵에 `paused`, `archived` 가 빠져 있음. REPORT.md 에서 XSM-23 회피 목적으로 의도적으로 제거한 것이라 명시됐으므로 add*.md 단계에서 복구 예정.
- **m-B2 CSP `default-src 'self'` 만 — 이미지·폰트·API 호출 소스 분리 없음**. 단순 SSR 이면 OK, SPA 에서는 `connect-src`, `img-src` 추가 필요.
- **m-B3 Rate limit 키축이 "ip" 하드코딩**: `arts/backend/internal/middleware/rate_limit.go:25` 에서 email / user_id 축이 별도 구현이 없다 — `/auth/login` 에 email-bucket 을 걸고 싶다면 별개 헬퍼 필요.
- **m-B4 `authMode()` 가 매 요청마다 `os.Getenv` 호출** (`bearerauth.go:23-30`) — 고빈도 라우트에서 비용은 작지만, startup 에 한 번 읽어서 closure 변수로 고정하면 cleaner.

##### Nit
- **n-B1** `deref_*` / `ptr_of.go` / `str_ptr.go` 가 각각 6-line 1-func 파일로 쪼개짐 — 파일 수가 service/ 에서 23개로 부푼다. `conversions.go` 같은 한 파일로 묶으면 가독성 상승.
- **n-B2** `init_authz.go:22-24` 의 `OwnershipMapping{Resource:"workflow", Table:"workflows", Column:"org_id"}` 가 단일 항목. 범용성은 있지만 zenflow 범위에서 과해 보임.

---

### DB / Migration

#### 잘 된 점
1. **Sentinel 행 + `BIGINT IDENTITY`**: `id=0` 행을 미리 꽂아놓고 FK `DEFAULT 0` 로 연결해 orphan 방지 + 쿼리에서 `wf.ID == 0` 이 곧 "not found" 감지 — 일관된 패턴.
2. **인덱스 적절**: 모든 FK 컬럼(`workflows.org_id`, `actions.workflow_id`, `execution_logs.{workflow_id,org_id}`)에 단독 인덱스. `users.email` UNIQUE, `refresh_tokens.claims` GIN.
3. **Timestamptz 사용**: `TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP` — timezone-aware (naive `TIMESTAMP` 대비 낫다).
4. **CHECK 제약 분리 `ALTER TABLE ADD CONSTRAINT`**: 마이그레이션 나중에 CHECK 를 완화 / 추가하기 편한 구조.
5. **`.latest_schema.sql` 에 해시 스탬프**: `YONGOL_SCHEMA_HASH` 로 다음 마이그 diff 기준점 명시.

#### 이슈

##### Major
- **M-D1 Down 마이그레이션이 빈 스텁**
  - 위치: `arts/db/migrations/0001_initial.down.sql`
  - 관찰: "Down stub — intentionally empty" 주석만. golang-migrate 가 `migrate down` 을 실행해도 아무 일도 일어나지 않음.
  - 원인: 설계 의도(docs/MIGRATION.md 참조) — 이전 specs/ revision 을 체크아웃 후 forward migration 을 재생성하라는 전략.
  - 조치: 이 전략이 맞다고 판단되면 .down.sql 파일을 아예 emit 하지 않거나, 파일명을 `0001_initial.up.sql` 만 두고 migrate 의 `-path` 를 single-direction 모드로 쓰는 편이 혼동 없음. 지금처럼 파일은 있는데 비어 있으면 "rollback 해도 DB 가 그대로인데 migrate 는 성공했다고 응답" → 운영자 혼동 유발.

##### Minor
- **m-D1 zenflow.md 와 타입 괴리 (UUID ↔ BIGINT)**: zenflow.md §2 는 `UUID PRIMARY KEY DEFAULT gen_random_uuid()` 를 요구하지만 실제 DDL 은 `BIGINT IDENTITY`. REPORT.md Bug #2 에서 XDO-77 의 int64 강제 제약 때문에 우회했다고 명시. 이는 **yongol validator ↔ codegen 간 규칙 충돌이 사용자 에게 스키마 변경을 강요**한 케이스. 생성기 차원의 해결 필요.
- **m-D2 `refresh_tokens.revoked_at` 에 인덱스 없음**: `RefreshTokenRevoke` 가 `WHERE token_hash=$1 AND revoked_at IS NULL` 를 도는데 PK(`token_hash`) 만으로도 OK 지만, `WHERE revoked_at IS NULL` partial index 는 청소 작업용으로 쓸모 있음.
- **m-D3 `actions.sequence_order BIGINT`**: 워크플로당 수억 개 액션은 현실적으로 없음 → INT 로 충분. 하지만 XDO-77 충돌 회피 차원에서 BIGINT 로 통일된 상태.

---

### Frontend (React+TSX)

#### 빌드 결과
- `node_modules` 없음 → `npm run build` / `tsc --noEmit` 미실행 (지시에 따라 정적 검토만).

#### 잘 된 점
1. **툴체인 현대적**: Vite 5 + React 19 + TanStack Query 5 + react-hook-form 7 + openapi-fetch + openapi-typescript → OpenAPI 를 type 의 SSOT 로 삼는 모범.
2. **Tailwind + shadcn/ui 스타일 cn() 헬퍼**: `src/lib/utils.ts` 에 `twMerge(clsx(...))` 컴포지션 — 업계 표준.
3. **타입드 UI 컴포넌트 일부**: `Badge.tsx` / `Modal.tsx` / `Checkbox.tsx` / `Select.tsx` / `Table.tsx` 는 `React.HTMLAttributes<...>` / `forwardRef` 등 제대로 된 타입.
4. **`path alias @/*`**: `tsconfig.json` + `vite.config.ts` 양쪽 일치.
5. **strict TypeScript ON**: `"strict": true`.

#### 이슈

##### Critical
- **C-F1 `App.tsx` 가 페이지를 마운트하지 않음**
  - 위치: `arts/frontend/src/App.tsx:7-10`
  - 관찰: 유일한 라우트가 `"/"` 에 placeholder div. `LoginPage` / `WorkflowsPage` 가 생성돼 있지만 import / Route 등록 없음.
  - 원인: codegen 이 TSX 페이지는 emit 하지만 `App.tsx` 는 초기 scaffold 그대로 둠. routes 자동 등록 템플릿 부재.
  - 조치: codegen 이 SSOT 의 `frontend/pages/*.tsx` 를 스캔해 `<Route>` 를 자동 주입하거나, 최소한 주석 placeholder 를 실제 `<Route>` 로 치환.

- **C-F2 TanStack Query v5 위반**
  - 위치: `arts/frontend/src/pages/WorkflowsPage.tsx:7`
  - 관찰: `useQuery(['listWorkflows'], apiClient.ListWorkflows)` — v4 문법. v5 는 `useQuery({ queryKey: ['listWorkflows'], queryFn: apiClient.ListWorkflows })` 필수. 런타임에 throw.
  - 원인: SSOT TSX 템플릿이 v4 예시로 쓰여 있음.
  - 조치: SSOT 페이지 샘플을 v5 object-form 으로 수정 + validate 단계 `tsx_openapi` 에서 패턴 매칭 규칙 추가 (`useQuery(Array, Fn)` 호출을 ERROR 로).

- **C-F3 UI 프리미티브 4종 `any` 타입 스텁**
  - 위치: `arts/frontend/src/components/ui/{Button,Card,Input,Form}.tsx`
  - 관찰: 각 파일이 `function Button(props: any) { return <button {...props}/> }` 수준. strict TypeScript 무력화.
  - 원인: SSOT (`specs/specs/frontend/components/ui/*.tsx`) 단계에서 placeholder 로 작성된 것이 그대로 arts 로 복사됨. codegen 이 이 파일을 "authored SSOT" 로 간주해 손대지 않음.
  - 조치: codegen 이 `Table.tsx` / `Badge.tsx` / `Modal.tsx` 수준의 타입드 구현을 기본 emit 하도록 전환. SSOT 는 "이름만" 선언하고 구현은 codegen 이 채워주는 모델이 낫다.

##### Major
- **M-F1 `src/lib/api.ts` placeholder 와 `src/api.ts` 실제 구현 이중 존재**
  - 위치: `arts/frontend/src/lib/api.ts:2` (placeholder) vs `arts/frontend/src/api.ts` (실제)
  - 관찰: 페이지가 `@/lib/api` 를 import 하지만 그 파일은 `export const apiClient: any = {}` 스텁. 실제 타입드 client 는 `src/api.ts` 에 있음 → 런타임에 `apiClient.Login is undefined`.
  - 원인: codegen 이 `src/api.ts` 에 emit 하는데 SSOT 페이지는 `@/lib/api` 를 기대.
  - 조치: codegen 출력 경로를 `src/lib/api.ts` 로 맞추거나 placeholder 를 `src/api.ts` 를 re-export 하도록 대체.

- **M-F2 폼에 로딩 / 에러 / 성공 피드백 전무**
  - 위치: `arts/frontend/src/pages/LoginPage.tsx:11-18`, `WorkflowsPage.tsx:11-19`
  - 관찰: `login.mutate(v)` 호출 후 `login.isPending` / `login.error` / `login.data` 렌더링 없음. `useQuery` 결과(`data`) 는 선언만 되고 JSX 에 꺾인 괄호 없이 아예 사용 안 함.
  - 조치: codegen 템플릿에 empty-state / loading / error 3-way 분기 주입. 또는 SSOT TSX 샘플을 완전한 예시로 재작성.

- **M-F3 접근성 (a11y) 기본 누락**
  - 관찰: `<Input {...register('email')} />` 에 `<label>` 연결 없음. type="email" 같은 시맨틱 힌트도 없음. 스크린리더 비친화.
  - 조치: UI 컴포넌트에 `label` prop 지원 추가, register 시 `type` 전달 템플릿화.

##### Minor
- **m-F1 `ui/index.tsx` 가 Button/Card/Input/Form 만 re-export**: `Badge`, `Modal`, `Checkbox`, `Select`, `Table` 이 emit 됐지만 barrel 누락. 페이지에서 쓸 수 없다.
- **m-F2 `(v: any)`**: `handleSubmit((v: any) => ...)` — strict 하에서 `any` 단언은 회피책. `handleSubmit<LoginForm>(...)` 으로 타입 바운드 가능.
- **m-F3 `any.d.ts` 생성 스크립트만 있고 실행 안 됨**: `gen:api` npm script 는 있으나 codegen 이 자동 실행 안 함 — `api.d.ts` 는 이미 커밋돼 있어서 무관하지만 수동 단계가 한 번 있다.

##### Nit
- **n-F1** `index.html` 에서 preconnect / meta theme-color 등 표준 PWA 힌트 없음 (SaaS 페이지라면 production 에서는 필요).

---

### Tests (Hurl)

#### 잘 된 점
1. **e2e 플로우는 올바른 순서**: register → login → createWorkflow → activate → execute → list, 각 단계 captures 체인이 끊기지 않음.
2. **assertion 일부는 필드 레벨**: `jsonpath "$.workflow.status" == "active"`, `"$.log.credits_spent" == 1`.

#### 이슈

##### Critical
- **C-T1 네거티브 케이스 0건**
  - 위치: `arts/tests/scenario-smoke.hurl`
  - 관찰: 401 (토큰 없음 / 만료), 403 (다른 org 리소스), 402 (credits=0), 404 (존재하지 않는 wf id), 409 (state conflict — 이미 active 인데 activate), 422 (invalid body) 등이 전부 빠짐.
  - 원인: SSOT `tests/scenario-smoke.hurl` 가 happy path 로만 작성. `@invariant` 파일이 별도로 존재해야 하는데 생성되지 않음 — zenflow.md §7 에 정의된 "Tenant Breach" / "Insufficient Credits" 가 SSOT 에 미반영.
  - 조치: yongol 이 OpenAPI 의 각 non-2xx 응답에 대해 최소 1개 `invariant-*.hurl` stub 을 자동 emit 하도록 확장. 지금은 `scenario-*.hurl` 만 1st-party.

##### Major
- **M-T1 assertion 세밀도 부족**
  - 관찰: `POST /auth/register` 의 `HTTP 201` 뒤에 `jsonpath "$.user.id" isInteger` 만. email / role 값 검증 없음. Login 도 `HTTP 200` + access_token 존재만.
  - 조치: codegen 이 OpenAPI response schema 의 required 필드 목록을 읽어 `jsonpath "$.user.email" == "{{email}}"` 등을 자동 생성.

- **M-T2 `newUuid` 가 Hurl 빌트인인가 의문**
  - 관찰: `email: "smoke+{{newUuid}}@example.com"` — Hurl 의 `{{newUuid}}` 는 존재하지만 `{{newUuid}}` 가 매 request 마다 새로 평가되는지 환경 의존적. 실제 REPORT 에서 6/6 통과했다니 OK.

##### Minor
- **m-T1 파일명 리네임 강제 (`smoke.hurl` → `scenario-smoke.hurl`)**: REPORT Bug #3 로 이미 알려진 detector 패턴 한계.

---

## 생성기 개선 제안 (yongol codegen 측)

> **Bug #N** 표기는 zenflow/try01/REPORT.md 의 기존 버그 번호와의 중복 여부.

1. **[SSaC 확장] `@foreach` 루프 구문 지원** — Critical C-B1 해결. SSaC 에서 `@foreach actions as a { @call worker.processAction({type: a.ActionType, payload: a.PayloadTemplate}) }` 를 허용하면 ExecuteWorkflow 의 action 처리가 복원된다.

2. **[SSaC 확장] `@catch` 에러 매핑** — Critical C-B2 해결. `@post User user = User.Create({...}) @catch unique_violation 409 "Email already registered"` 형태로 SQLSTATE 기반 409 분기 선언.

3. **[sqlc 템플릿] `:exec` → `:execrows`** — Major M-B1 해결. `OrganizationDeductCredit` 처럼 atomic guard 가 있는 쿼리는 rows affected 를 반환해야 호출자가 0 을 감지하고 402 를 낼 수 있다. 선언적으로 "@guard" 어노테이션이 있으면 codegen 이 자동 `:execrows` 로 승격.

4. **[oapi-codegen 래핑] ErrorHandler 주입** — Major M-B2 해결. main.go 에서 `api.GinServerOptions{ErrorHandler: middleware.EnvelopeErrorHandler}` 를 기본 주입. 현재는 `{"msg": err.Error()}` fallback 이 노출.

5. **[TSX 템플릿] App.tsx 라우트 자동 조립** — Critical C-F1 해결. pages/ 아래의 `export default` 를 스캔해 `<Route path="..." element={<...Page/>}/>` 를 자동 주입. URL path 는 pages 파일의 `// @route /login` 주석으로 선언.

6. **[TSX 검증] TanStack Query v5 강제** — Critical C-F2 해결. `tsx_openapi` 또는 신규 `tsx_lib` 규칙에서 `useQuery\(\[` (positional key) 매칭 → ERROR. (Bug #4 후보)

7. **[TSX 템플릿] UI 프리미티브 기본 구현 제공** — Critical C-F3 해결. `specs/frontend/components/ui/*.tsx` 가 placeholder 이면 codegen 이 덮어쓰도록 "authored vs generated" 마커를 명확히.

8. **[TSX 경로 정합] `src/lib/api.ts` ↔ `src/api.ts` 통일** — Major M-F1 해결. codegen output path 를 `src/lib/api.ts` 로 이동하거나 `src/lib/api.ts` 를 re-export 로 자동 생성.

9. **[Migration 정책] down.sql 전략 재검토** — Major M-D1. 빈 down 파일을 emit 하면 운영자에게 오해를 준다. 파일 자체를 생성하지 않거나, "DROP TABLE IF EXISTS ..." 를 실제로 emit 하는 모드를 제공.

10. **[XDO-77 완화] DDL ↔ OpenAPI int 타입 분기 허용** — Minor m-D1 해결. Bug #2 와 중복. sqlc 의 int32 매핑 vs OpenAPI int64 강제가 상호 모순이라 사용자에게 BIGINT 로 강제 이전을 요구 → UUID 지정도 못 쓰게 만듦.

11. **[Hurl 템플릿] invariant-*.hurl 자동 emit** — Critical C-T1 해결. OpenAPI 의 각 non-2xx response 에 대해 stub 케이스를 생성하고, `@invariant` 주석이 붙은 zenflow.md 항목은 별도 파일 강제.

12. **[context key 규약] typed context key emit** — Minor M-B3. `type contextKey string` 를 service/ 패키지에 한 번 emit 하고 전 핸들러가 사용.

---

## zenflow.md 요구사항 커버리지

| 요구 | 상태 | 비고 |
|---|---|---|
| §1 Multi-tenant SaaS 도메인 모델 | 충족 | organizations/users/workflows/actions/execution_logs + org_id 분리 |
| §2 DDL 전체 (UUID PK) | **부분** | 테이블 6개 전부 있으나 PK 를 BIGINT IDENTITY 로 변경 (Bug #2 우회) |
| §3 State Machine (draft/active/paused/archived) | **부분** | draft→active / active→active(Execute) 만 존재. paused / archived 전이 누락 |
| §4 Authz (CreateWorkflow admin / ListWorkflows same_org / ActivateWorkflow admin+same_org) | 충족 | authz.rego 에 3 규칙 + ExecuteWorkflow 까지 확장 |
| §5 ActivateWorkflow: credits check → 402 | 충족 | billing.CheckCredits 분기 (Bug #1 우회 경로) |
| §5 ExecuteWorkflow: @auth, @state, actions load, loop processAction, deductCredit, log | **부분** | loop processAction 누락 (C-B1) |
| §6 Custom funcs (processAction/checkCredits/deductCredit) | **부분** | checkCredits = billing.CheckCredits. processAction 미구현 |
| §7 @scenario Happy Path | 충족 | scenario-smoke.hurl 6 스텝 |
| §7 @invariant Tenant Breach (403) | **미충족** | Hurl 테스트 없음 |
| §7 @invariant Insufficient Credits (402) | **미충족** | Hurl 테스트 없음 |
| §8.5 Postgres / migrate 적용 | 충족 | golang-migrate up 페어, REPORT 에서 적용 확인됨 |
| §8.5 Dummy SMTP | 해당 없음 | `@call mail.*` 미사용 |

**종합**: 12 항목 중 충족 6 / 부분 4 / 미충족 2. Critical C-B1 (processAction loop) + C-T1 (invariant 테스트) 이 채워지면 충족 8 / 부분 3 / 미충족 1 로 개선.

---

## 결론 및 다음 단계

백엔드는 **production-ready 한 기본기를 이미 갖춘 수준** (미들웨어 스택, 트랜잭션 규약, 민감 필드 마스킹, 타이밍 공격 완화, JWT 검증). 프론트엔드는 **generator 가 손을 놓은 영역** — 페이지/UI/api 경로의 연결이 끊겨 있다. Hurl 은 happy-path-only 로 zenflow.md §7 의 핵심 invariant 를 검증하지 못한다.

**다음 권장 액션 (우선순위 순)**
1. yongol codegen 개선 제안 #5, #6, #7, #8 을 반영해 프론트엔드를 "빌드는 되는 상태" 에서 "라우팅·타입·상태 처리가 살아 있는 상태" 로 끌어올림.
2. SSaC 에 `@foreach` / `@catch` 구문을 도입 (제안 #1, #2) — ExecuteWorkflow loop 복원 + Register 409 정합화.
3. Hurl codegen 을 OpenAPI-driven 으로 확장해 invariant 자동 emit (제안 #11).
4. Bug #2 의 XDO-77 ↔ sqlc int32 충돌을 해결 (제안 #10) — zenflow.md 의 UUID 요구를 그대로 수용할 수 있게 함.
5. Migration down 전략 확정 (제안 #9) — 빈 스텁을 emit 하지 말거나 실제 DROP 을 emit 하거나.

본 리뷰는 파일을 전혀 수정하지 않고 정적 관찰만으로 작성됐다. 구체적 patch 는 해당 제안들이 승인된 뒤 plans/gen/ 혹은 plans/validate/ 에 Phase 계획으로 분할해 진행 권장.
