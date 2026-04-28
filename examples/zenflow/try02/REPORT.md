# zenflow/try02 완료 보고서

## 실행 정보
- 시작: 2026-04-25T12:58:00+09:00
- 종료: 2026-04-25T13:12:07+09:00
- 소요: 00:14:07 (약 14분 7초)
- Claude Code: 2.1.119
- Model: Opus 4.7 (1M context)
- yongol: v0.2.6 (@eval, S-64, smoke detector 포함 빌드)
- Go: go1.25.0 linux/amd64
- Hurl: 6.0.0
- Postgres: postgres:16-alpine (Docker, host port 15432)

## 범위
- 입력: examples/zenflow/zenflow.md 한 파일만 사용. add01~07.md / try01 어떤 것도 열지 않음.
- zenflow.md §2 DDL의 UUID PK는 매뉴얼(docs/ddl.md int64 표준)에 맞춰 BIGINT IDENTITY로 변환.
- §3 Mermaid의 Pause/Archive 전이는 본 try02 범위에서 SSaC/OpenAPI를 작성하지 않으므로 stateDiagram에서 제외 (operationId 매칭 규칙 준수). draft → active(Activate) + active 자기루프(Execute)만 유지.

## 작성한 SSOT 목록
모두 specs/ 기준.

- manifest.yaml — bearer JWT, claims ID/Email/Role/OrgID, roles [admin, member]
- api/openapi.yaml — Login, ListWorkflows, CreateWorkflow, ActivateWorkflow, ExecuteWorkflow + Workflow/Action/ExecutionLog/ErrorResponse 스키마
- db/organizations.sql — BIGINT IDENTITY + @sentinel id=0
- db/users.sql — org_id NOT NULL DEFAULT 0 FK + password_hash @sensitive + claims JSONB
- db/workflows.sql — status CHECK + @sentinel id=0 (actions/execution_logs FK 보호)
- db/actions.sql — workflow_id/sequence_order BIGINT
- db/execution_logs.sql — 실행 로그 + 인덱스 (org_id, executed_at DESC)
- db/refresh_tokens.sql — manifest auth backend canonical(XNA-90), revoked_at @nullable, @archived (XSD-55 면제)
- db/queries/users.sql — UserFindByEmail (@verify-password 코드젠), UserCreate
- db/queries/workflows.sql — Create / FindByID / ListByOrgIDPaged (BIGINT cast) / CountByOrgID / UpdateStatus / OwnerLookupWorkflow (XQP-30)
- db/queries/organizations.sql — FindByID, DeductCredits
- db/queries/actions.sql — ListByWorkflowID :many LIMIT 1000
- db/queries/execution_logs.sql — Create :one
- db/queries/auth.sql — LoginLookup + RefreshToken{Insert,FindByHash,Revoke,RevokeAll}
- db/sqlc.yaml — pgx/v5
- service/auth/login.ssac — @verify-password + auth.IssueToken
- service/workflow/list_workflows.ssac — @auth + offset 페이지네이션
- service/workflow/create_workflow.ssac — @auth (admin) + Workflow.Create
- service/workflow/activate_workflow.ssac — @auth + @empty wf + @call billing.CheckCredits 402 + @state Activate + @put + 재조회
- service/workflow/execute_workflow.ssac — @auth + @empty wf + @state Execute + CheckCredits 402 + @get []Action + worker.RunWorkflow + billing.DeductCredit + ExecutionLog.Create
- states/workflow.md — [*] → draft → active(Activate); active → active(Execute)
- policy/authz.rego — OPA v1, @ownership workflow: workflows.org_id, 5 allow rules (admin/member 명시)
- func/billing/check_credits.go — Balance ≤ 0 시 ErrInsufficientCredits 반환 (@error 402)
- func/billing/deduct_credit.go — 차감량 echo
- func/worker/run_workflow.go — status="success" 모킹
- frontend/pages/login.tsx — apiClient.Login + register('email','password')
- frontend/pages/workflows.tsx — List/Create/Activate/Execute 사용
- frontend/components/ui/index.tsx — @/components/ui import stub (T-1 우회)
- tests/smoke.hurl — Login → List → Create → Activate → Execute (5 requests)

## 산출물 요약 (arts/)
- arts/backend/cmd/ : main.go (pgx pool, JWT, OPA, gin)
- arts/backend/internal/api/ : oapi-codegen strict server interface
- arts/backend/internal/service/ : SSaC → Go handler 5개
- arts/backend/internal/db/ : sqlc 생성물
- arts/backend/internal/middleware/ : BearerAuth, RequestID, ErrorEnvelope, Prometheus, SecurityHeaders, BodyLimit, RequestValidator
- arts/backend/internal/billing|worker/ : func 미러
- arts/backend/internal/model/ : UserClaim
- arts/backend/internal/statemachine/ : workflow 상태 전이표
- arts/backend/internal/infra/auth/ : postgres RefreshStore (auth.Init 주입)
- arts/backend/policy/authz.rego : Rego 정책 미러
- arts/db/migrations/0001_initial.up.sql : 22 ops 초기 마이그레이션
- arts/db/.latest_schema.sql : 정규화 baseline 스냅샷
- arts/frontend/ : Vite + React + apiClient + 타입 + UI primitives
- arts/tests/smoke.hurl : specs/tests/ 미러

## 검증 결과
- yongol validate specs → 0 errors, 0 warnings (28개 phase 모두 ✓)
- yongol generate specs arts → "0 errors, 0 warnings", "[migration] mode=initial file=0001_initial.up.sql ops=22", "artifacts written to arts (backend=go-gin, frontend=react)"
- cd arts/backend && go build ./... → 정상 종료(0)
- 마이그레이션: psql … < 0001_initial.up.sql → COMMIT
- 시드: AcmeCo(credits 100) 조직 + admin@zenflow.test/p@ssw0rd
- hurl --test smoke.hurl → Success (5 requests / 66 ms), 100% 통과, 실패 0
  - Login 200, GET /workflows 200, POST /workflows 200, /activate 200(status="active"), /execute 200(credits_spent==1)

## 신규 기능 사용 사례
- @eval 사용: 0 회 — 사용 시 후술 BUG-001로 인해 S-25 ERROR 발생, 같은 의도의 매뉴얼 권장 우회(@call pkg.CheckX(...) 402, error 반환 Func)로 두 곳 대체.
- smoke.hurl 그대로 사용: rename 없이 specs/tests/smoke.hurl 작성 → arts/tests/smoke.hurl 정상 미러 (이전 강제 rename 회귀 없음).
- BIGINT/int64 표준 적용: 원안 INTEGER 컬럼 3개(credits_balance, sequence_order, credits_spent) + UUID PK 5개를 BIGINT IDENTITY로 변환. 추가 BIGINT FK 컬럼(org_id, workflow_id)을 포함해 정수 컬럼 약 12개. XDO-77 ERROR 없이 통과.
- S-64 (@empty/@exists 모델 한정) 4곳: @empty wf 2회, @empty org 2회 — 모두 Model 변수 가드. 스칼라 가드는 @call로 분리해 매뉴얼 권장 패턴 그대로 따름.

## 버그 리포트

### BUG-001 — @eval 사용 시 S-25 "unknown sequence type" ERROR
- 재현
  - SSaC 파일에 `// @eval billing.IsZeroBalance({Balance: org.CreditsBalance}) "Insufficient credits" 402` 추가 후 yongol validate specs.
- 실제 출력
  - [ssac] activate_workflow.ssac:10: [S-25] unknown sequence type: @eval
    Advice: Use one of: @get/@post/@put/@delete/@call/@empty/@exists/@state/@auth/@publish/@verify-password
- 기대 동작
  - @eval 은 매뉴얼·docs/ssac.md @eval — predicate guard 섹션과 rulebook S-67/68/69 정식 규칙으로 지원되므로 검증기도 인지해야 함.
- 원인 (소스 분석)
  - pkg/parser/ssac/parse_annotation.go + parse_eval.go 는 SeqEval = "eval" Sequence 를 정상 생성하지만, pkg/validate/ssac/helpers.go knownSeqTypes map 에 "eval": true 항목이 누락되어 s_25_unknown_seq_type.go 가 모든 @eval 라인에 ERROR.
  - 빌드된 v0.2.6 바이너리도 동일 누락.
- 우회
  - @eval pkg.Predicate({...}) "msg" STATUS → @call pkg.CheckX({...}) STATUS (Func 를 error 반환 형태로 작성). 본 try02 두 곳(activate_workflow.ssac, execute_workflow.ssac)에 적용. 의미와 응답 코드는 동일하게 402 Insufficient credits.

### BUG-002 (참고) — TSX @/components/ui import 가 항상 T-1 WARNING
- 재현
  - frontend/pages/*.tsx 에 `import { Button } from '@/components/ui'` 만 두고 yongol validate.
- 실제 출력
  - [T-1] imported component file not found: @/components/ui (페이지의 import 토큰 수만큼 반복). WARNING 이지만 generate 가 거절.
- 기대 동작
  - @/components/ui/* 는 yongol-owned. validate 단계에서 arts/frontend/src/components/ui/ 가 아직 없는 게 정상이므로 alias 라면 자동 면제되어야 자연스러움.
- 우회
  - specs/frontend/components/ui/index.tsx 에 stub export 를 두면 alias root 가 specs/frontend 로 잡혀 T-1 가 사라짐. components/ui/ 는 isYongolManaged 미복사이므로 산출물에 영향 없음.

## try01 대비 체감 차이
이전 시도와 비교하지 않으려고 try01/try-02/try-03 의 어떤 파일도 열지 않았음. 매뉴얼·docs·rulebook·codebook 만으로 작업한 인상:
- yongol validate 의 진단이 매번 정확한 룰 ID + 라인 + Advice + 캐노니컬 SQL 패치까지 동봉돼서 SSOT 9개를 다뤄도 피드백 루프가 짧음. XNA-90 의 캐노니컬 DDL/queries 추천(refresh_tokens + RefreshToken* 4종)은 그대로 붙여넣기로 해결.
- BIGINT/int64 강제(XDO-77) 덕분에 sqlc → Go → OpenAPI int 폭 캐스팅 이슈가 처음부터 차단됨. 단, sqlc.arg() 의 LIMIT/OFFSET 내부 캐스트는 직접 ::bigint 로 바꿔야 list 핸들러 빌드 통과 (sqlc.arg(page)::int → int32 vs OpenAPI int64).
- 막힘: @eval 신기능을 매뉴얼대로 시도하다 BUG-001 로 막힘. 룰북/docs 와 검증기 사이 한 줄 누락 형태라 우회는 즉시 가능했지만, 가장 깔끔한 표현을 그대로 쓰지 못한 게 아쉬움.
- @empty 가 모델 한정(S-64)이 되어 "스칼라 가드 = Func 분리" 관습이 시작부터 강제됨, 코드가 깔끔.

## 메모
- 백엔드 환경변수: DATABASE_URL, JWT_SECRET(32자+), OPA_POLICY_PATH, BACKEND_AUTH_MODE(bearer/cookie/hybrid). JWT_SECRET 짧으면 startup fail-fast.
- 시드: organizations(name='AcmeCo', credits_balance=100) + users(role='admin', email='admin@zenflow.test', bcrypt('p@ssw0rd')). smoke 는 admin 으로만.
- mail/queue/cache/session/file backend 는 manifest 미선언이라 외부 인프라(SMTP dummy 등) 불필요.
- yongol 소스는 단 한 줄도 수정하지 않음. 우회는 모두 SSOT 측에서 처리.
