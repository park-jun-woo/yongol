# zenflow/try01 완료 보고서

## 실행 정보
- 시작 시각: 2026-04-25T01:47:12+09:00
- 종료 시각: 2026-04-25T01:59:46+09:00
- 소요 시간: 00:12:34 (754s)
- Claude Code: 2.1.119
- Model: claude-opus-4-7 (Opus 4.7, 1M context)
- yongol: v0.2.5
- Go: go1.25.0 linux/amd64
- Hurl: 6.0.0 (x86_64-pc-linux-gnu)

## 범위
- 입력 문서: `examples/zenflow/zenflow.md` (add*.md 제외)
- 작업 디렉토리: `examples/zenflow/try01/`
- SSOT 루트: `try01/specs/specs/` (`yongol init` 의 기본 레이아웃을 그대로 유지)
- 산출물 루트: `try01/arts/`

## 작성한 SSOT 목록
- `specs/specs/manifest.yaml` — JWT bearer, claims(ID/Email/Role/OrgID), roles admin/member
- `specs/specs/api/openapi.yaml` — 6 ops: Register, Login, CreateWorkflow, ListWorkflows, ActivateWorkflow, ExecuteWorkflow
- `specs/specs/db/*.sql` — organizations, users, workflows, actions, execution_logs, refresh_tokens(@archived)
- `specs/specs/db/sqlc.yaml` + `specs/specs/db/queries/*.sql` — sqlc 설정/쿼리
- `specs/specs/states/workflow.md` — Mermaid stateDiagram (draft→active, active self-loop on ExecuteWorkflow)
- `specs/specs/policy/authz.rego` — @ownership workflow: workflows.org_id
- `specs/specs/service/auth/{register,login}.ssac` — 인증 서비스 시퀀스
- `specs/specs/service/workflow/{create,list,activate,execute}_workflow.ssac` — 워크플로 서비스 시퀀스
- `specs/specs/func/billing/check_credits.go` — Bug #1 우회용 순수 Func
- `specs/specs/frontend/pages/{LoginPage,WorkflowsPage}.tsx` — React 페이지
- `specs/specs/frontend/components/ui/{Button,Card,Input,Form,index}.tsx` — 공용 컴포넌트
- `specs/specs/frontend/lib/api.ts` — apiClient
- `specs/specs/tests/scenario-smoke.hurl` — 6 스텝 happy path

## 산출물 요약
- `arts/backend/` — Go+Gin 백엔드 (handlers, services, policy, db, main)
- `arts/db/migrations/` — golang-migrate up/down 페어 + `.latest_schema.sql`
- `arts/frontend/` — React+Vite 프론트엔드
- `arts/tests/scenario-smoke.hurl` — specs 의 Hurl 미러

## 검증 결과
- `yongol validate specs/specs` → **0 errors, 0 warnings** (27개 검증기 모두 통과)
- `yongol generate specs/specs arts` → `mode=initial ops=23 | artifacts written`
- `go build ./...` (arts/backend) → 클린 빌드
- `migrate up` (postgres:16-alpine @ :15432) → `1/u initial`
- `hurl --test arts/tests/scenario-smoke.hurl` → **6/6 requests succeeded, 0 failures in 110 ms**

## 버그 리포트

### Bug #1 — `@empty <scalar>` 가 모델 비교 코드로 생성됨 (codegen)
- 증상: SSaC 에서 `@empty org.CreditsBalance` 사용 시 생성 코드가 `org.CreditsBalance.ID == 0` 처럼 스칼라 필드에 `.ID` 를 붙여 `go build` 실패.
- 기대 동작: 스칼라 타입이면 해당 타입의 제로 값(`int64(0)`)과 비교하거나, validate 단계에서 스칼라 대상 `@empty` 를 거절.
- 재현: `@empty <스칼라 경로>` 포함한 SSaC 파일로 `yongol generate`.
- 우회: 순수 Func `billing.CheckCredits({Balance: int64}) -> err` 로 대체하고 SSaC 에서 `// @call billing.CheckCredits({...}) 402` 로 호출.

### Bug #2 — DDL INTEGER vs XDO-77 의 int64 강제가 상호 모순 (validator ↔ codegen)
- 증상: DDL `INTEGER` 는 sqlc 에서 `int32` 로 매핑되는데, 교차검증 규칙 XDO-77 (openapi_ddl) 은 OpenAPI `integer` 에 `format: int64` 를 강제. 두 제약이 충돌해 어느 한 쪽이 항상 실패.
- 기대 동작: XDO-77 이 DDL 타입에 따라 `int32`/`int64` 포맷을 분기하도록 완화하거나, sqlc 타입 매핑을 규칙에 맞춰 `int64` 로 통일.
- 우회: 비-ID 수치 컬럼(`credits_balance`, `sequence_order`, `credits_spent`)을 모두 `BIGINT` 로 선언.

### Bug #3 — `smoke.hurl` 이 detector 에 잡히지 않음 (parser)
- 증상: `tests/smoke.hurl` 단독 배치 시 KindScenario 가 `SSOTDeclared` 로 떠서 H-2 WARNING 이 발생. `directory_ssots.go` 가 `scenario-*.hurl` / `invariant-*.hurl` 글롭만 스캔.
- 기대 동작: `docs/scenario.md` 에 `smoke.hurl` 이 1st-party 명시돼 있으므로 detector 에도 포함해야 함.
- 우회: 파일명을 `scenario-smoke.hurl` 로 변경.

## 메모 (add*.md 진행 전 준비사항)
- `yongol init` 이 `try01/specs/specs/` 중첩 레이아웃을 만든다. `yongol validate` 와 `yongol generate` 모두 `try01/specs/specs` 를 인자로 준다.
- `db/sqlc.yaml` 의 `gen.go.out` 은 `../../../arts/backend/internal/db` (db/ 기준 3단계 상향) 로 맞춰야 `go build` 가 통과.
- `refresh_tokens` 테이블은 XNA-90 이 요구하지만 SSaC 에서 미사용 → `CREATE TABLE` 상단에 `-- @archived` 주석.
- 상태머신에 `PauseWorkflow` / `ArchiveWorkflow` 전이가 아직 없다 (XSM-23 회피 목적). 대응 SSaC 함수가 추가되는 add*.md 단계에서 상태 전이도 함께 복원 필요.
- 백엔드 런타임 환경변수: `DATABASE_URL`, `JWT_SECRET` (≥32자), `OPA_POLICY_PATH=arts/backend/policy/authz.rego`, `BACKEND_AUTH_MODE=bearer`.
- 본 보고서는 서브 에이전트가 산출한 요약을 검증(`yongol validate` 재실행 → 0/0 확인) 후 부모 세션에서 작성.
