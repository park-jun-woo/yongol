# zenflow/try03 완료 보고서

## 실행 정보
- 시작: 2026-04-25T13:47:04+09:00
- 종료: 2026-04-25T14:03:37+09:00
- 소요: 16분 33초
- Claude Code 2.1.119 / **Sonnet 4.6** (Opus 4.7 가 아님)
- yongol v0.2.7 (@eval 동작, S-64, smoke detector, BIGINT 표준)
- Go: go1.25.0 linux/amd64
- Hurl: 6.0.0

## 범위
- 입력: examples/zenflow/zenflow.md (add*.md 제외)

## 작성한 SSOT 목록

### manifest
- `specs/specs/manifest.yaml` — 프로젝트 설정 (bearer 인증, claims: ID/Email/Role/OrgID int64, roles: admin/member)

### DDL
- `specs/specs/db/organizations.sql` — 조직 테이블 (BIGINT PK, credits_balance, sentinel row)
- `specs/specs/db/users.sql` — 사용자 테이블 (BIGINT PK, org_id FK, password_hash @sensitive, claims @sensitive)
- `specs/specs/db/workflows.sql` — 워크플로우 테이블 (BIGINT PK, status DEFAULT 'draft', sentinel row)
- `specs/specs/db/actions.sql` — 액션 테이블 (BIGINT PK, workflow_id FK, sequence_order BIGINT)
- `specs/specs/db/execution_logs.sql` — 실행 로그 테이블 (BIGINT PK, credits_spent BIGINT)
- `specs/specs/db/auth.sql` — refresh_tokens 테이블 (@archived, XNA-90 요구사항)
- `specs/specs/db/sqlc.yaml` — sqlc 설정 (pgx/v5, out: arts/backend/internal/db)

### sqlc queries
- `specs/specs/db/queries/auth.sql` — LoginLookup, RefreshToken* 쿼리
- `specs/specs/db/queries/organizations.sql` — OrganizationFindByID, Create, DeductCredits
- `specs/specs/db/queries/users.sql` — UserCreate, UserFindByEmail (+allow-sensitive), UserFindByID
- `specs/specs/db/queries/workflows.sql` — WorkflowCreate, FindByID, ListByOrgID, UpdateStatus, OwnerLookupWorkflow
- `specs/specs/db/queries/actions.sql` — ActionCreate, ListByWorkflowID (+no-pagination)
- `specs/specs/db/queries/execution_logs.sql` — ExecutionLogCreate, FindByID, ListByWorkflowID (+no-pagination)

### OpenAPI
- `specs/specs/api/openapi.yaml` — 11개 operation (Register, Login, ListWorkflows, CreateWorkflow, GetWorkflow, AddAction, ActivateWorkflow, PauseWorkflow, ArchiveWorkflow, ExecuteWorkflow, ListExecutionLogs), BIGINT IDs

### 상태 기계
- `specs/specs/states/workflow.md` — draft→active, active→paused, paused→active, active→archived, active→active (ExecuteWorkflow self-loop)

### OPA Rego
- `specs/specs/policy/authz.rego` — OPA v1 문법, @ownership workflow: workflows.org_id, 9개 allow rule

### SSaC
- `specs/specs/service/auth/register.ssac` — 사용자 등록 (HashPassword + User.Create)
- `specs/specs/service/auth/login.ssac` — 로그인 (@verify-password + IssueToken)
- `specs/specs/service/workflow/list_workflows.ssac` — 워크플로우 목록 (@no-pagination, @auth)
- `specs/specs/service/workflow/create_workflow.ssac` — 워크플로우 생성 (@auth admin)
- `specs/specs/service/workflow/get_workflow.ssac` — 워크플로우 조회 (@state-neutral, @auth, @empty)
- `specs/specs/service/workflow/add_action.ssac` — 액션 추가 (@state-neutral, @auth, @empty)
- `specs/specs/service/workflow/activate_workflow.ssac` — 활성화 (@auth, @call billing.CheckCredits 402, @state, @empty)
- `specs/specs/service/workflow/pause_workflow.ssac` — 일시정지 (@auth, @state, @empty)
- `specs/specs/service/workflow/archive_workflow.ssac` — 보관 (@auth, @state, @empty)
- `specs/specs/service/workflow/execute_workflow.ssac` — 실행 (@auth, @state, @call worker+billing)
- `specs/specs/service/workflow/list_execution_logs.ssac` — 실행 로그 목록 (@no-pagination, @auth)

### Func Spec
- `specs/specs/func/billing/is_zero_balance.go` — @func checkCredits (크레딧 잔액 검사, @error 402)
- `specs/specs/func/billing/deduct_credit.go` — @func deductCredit (크레딧 차감)
- `specs/specs/func/worker/process_action.go` — @func processAction (워크플로우 액션 처리 시뮬레이션)

### Hurl
- `specs/specs/tests/smoke.hurl` — 6개 요청 스모크 테스트 (Register, Login, CreateWorkflow, AddAction, GetWorkflow, ListWorkflows)

## 산출물 요약
- `arts/backend/` — Go+Gin 백엔드 (github.com/park-jun-woo/zenflow)
- `arts/db/migrations/0001_initial.up.sql` — 초기 스키마 마이그레이션 (17 ops)
- `arts/tests/smoke.hurl` — 스모크 테스트 (specs/tests/ 미러)
- `arts/frontend/` — React+Vite 프론트엔드 스캐폴드

## 검증 결과
- yongol validate: **0 errors / 0 warnings** (최종)
- go build: **성공** (3개 codegen 버그 수동 패치 후)
- hurl --test: **6/6 통과** (1 파일, 6 요청, 실패 0)

## 신규 기능 사용 사례

### `@eval` 사용
- 당초 `activate_workflow.ssac`에서 `@eval billing.IsZeroBalance(...)` 사용 계획
- XSF-62 버그로 `@eval`이 참조하는 bool-반환 func을 "@call 미참조" WARNING으로 오탐
- `@call billing.CheckCredits({...}) 402` 패턴(error-guard @call)으로 우회
- `@eval` 최종 사용 횟수: **0회** (우회 처리, 버그 리포트 참조)

### `smoke.hurl`
- `specs/specs/tests/smoke.hurl` — 직접 작성, `arts/tests/smoke.hurl`로 자동 미러됨. **사용함**

### BIGINT 적용 컬럼
- 모든 PK: BIGINT GENERATED ALWAYS AS IDENTITY (5개 테이블)
- FK 컬럼: org_id, workflow_id 등 6개
- 비즈니스 컬럼: credits_balance, credits_spent, sequence_order
- **총 BIGINT 컬럼 수**: 약 12개

### `@empty` 가드 사용 횟수: **11회**
- get_workflow(1), add_action(1), activate_workflow(3), pause_workflow(2), archive_workflow(2), execute_workflow(1), list_execution_logs(1)

## 버그 리포트

### BUG-1: UUID 기반 DDL에서 codegen 타입 불일치 (심각)
**재현**: UUID PRIMARY KEY 사용 시 `yongol generate` 후 `go build`
**증상**: `types.UUID`([16]byte) vs `pgtype.UUID`(struct) 불일치, `workflow.ID == 0` 비교 불가, UUID 직접 캐스트 실패
**우회**: UUID → BIGINT GENERATED ALWAYS AS IDENTITY 전환

### BUG-2: XSF-62 — `@eval` 사용 func spec을 "@call 미참조"로 오탐 (경고)
**재현**: bool-반환 func을 `@eval`로만 참조
**증상**: XSF-62 WARNING — `@eval`은 bool 반환 func 전용이지만 validator는 `@call` 참조만 확인
**우회**: error-반환 func으로 변경 후 `@call pkg.Func({...}) STATUS` 패턴 사용

### BUG-3: register.go codegen 타입 오류 (심각)
**증상**: JSONB claims 컬럼에 string 리터럴 대입 + UserCreateRow/User 타입 혼용
**우회**: `"{}"` → `[]byte("{}")`, `convertUserCreateRow` 헬퍼 수동 추가

### BUG-4: convert_action.go PayloadTemplate 포인터 오류 (심각)
**증상**: `map[string]interface{}` → `*map[string]interface{}` 직접 대입 시도
**우회**: `&payloadTemplateMap` 수동 패치

## Sonnet으로 진행한 체감

Opus 대비 매뉴얼 의존도가 높았으나 validate 에러 메시지를 따르면 회복이 빠른 편이었다. 가장 크게 막힌 부분은 UUID PK로 인한 codegen 타입 불일치로, SSOT 레벨에서 BIGINT로 전환하는 결정이 필요했다. `@eval` 기능은 XSF-62 버그로 인해 활용하지 못했다 (validate가 통과해도 실제 의도한 boolean-predicate guard를 사용할 수 없었음). validate → fix → validate 사이클은 5회 반복 후 0/0 달성. 전반적으로 매뉴얼을 정독한 뒤 SSaC 작성을 시작하면 Sonnet도 충분히 완주 가능한 수준.

## 메모
- Organization 생성 API 없음 → 스모크 테스트 전 psql INSERT 필요
- 스모크 테스트는 기본 CRUD만 커버 (Activate/Execute는 credits 설정 후 별도 시나리오 필요)
- BUG-1(UUID codegen) 해결 전까지 BIGINT PK 패턴 권장
