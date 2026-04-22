# ZenFlow Add-on #07 — Dashboard + Relation Enrichment

## Overview
An operator dashboard summarizing org state on one screen. Exercises two yongol features together: **Func Response types** (aggregate response structs declared in `func/<pkg>/*.go` and consumed via `@call`) and **TSX custom components** (widgets too complex for plain data binding).

## Verification Points
- **Func Response types** — `SummarizeResponse`, `BuildExecutionDetailResponse` declared in `func/dashboard/*.go`; assembled via `@call dashboard.Summarize` / `@call dashboard.BuildExecutionDetail` and referenced by the OpenAPI response schema per the `@response` ↔ Func Response rule.
- **TSX custom components** — widgets (credits gauge, execution chart, status badge) live under `frontend/components/*.tsx` and are invoked by the page component.

## New Endpoints
- **GET /dashboard** (`GetDashboard`)
  - Response: `{ summary: SummarizeResponse }`
  - Flow: `@get Organization org` + `@call dashboard.Summarize({OrgID, ...})` → aggregated object.
- **GET /audit-logs/{id}** (`GetAuditLog`) — detail view; caller can enrich client-side by calling `GetUser` and `GetOrganization` separately.
- **GET /execution-logs/{id}/detail** (`GetExecutionDetail`)
  - Response type: `BuildExecutionDetailResponse` (Func Response, combines log + workflow + org).

## DDL
No changes — reuse existing `workflows`, `execution_logs`, `audit_logs`, `organizations`, `users`.

## Custom Functions (`func/dashboard/`)

Each Func declares its `Response` struct in the same file; SSaC pulls the type via `@call`.

```go
// func/dashboard/summarize.go
package dashboard

// @func summarize
// @description Assemble org dashboard summary in-memory.

type SummarizeRequest struct {
    OrgID          int64
    OrgName        string
    PlanType       string
    CreditsBalance int64
}

type SummarizeResponse struct {
    OrgName           string
    PlanType          string
    CreditsBalance    int64
    ActiveWorkflows   int64
    PausedWorkflows   int64
    TotalExecutions   int64
    TotalCreditsSpent int64
}

func Summarize(req SummarizeRequest) (SummarizeResponse, error) {
    return SummarizeResponse{
        OrgName:        req.OrgName,
        PlanType:       req.PlanType,
        CreditsBalance: req.CreditsBalance,
        // remaining fields simulated (constant or random) to stay purity-safe.
    }, nil
}
```

```go
// func/dashboard/build_execution_detail.go
package dashboard

// @func buildExecutionDetail
// @description Compose ExecutionDetail from log + workflow + org.

type BuildExecutionDetailRequest struct {
    LogID         int64
    WorkflowID    int64
    WorkflowTitle string
    OrgID         int64
    OrgName       string
    Status        string
    CreditsSpent  int64
    ExecutedAt    string
}

type BuildExecutionDetailResponse struct {
    ID            int64
    WorkflowID    int64
    WorkflowTitle string
    OrgID         int64
    OrgName       string
    Status        string
    CreditsSpent  int64
    ExecutedAt    string
}

func BuildExecutionDetail(req BuildExecutionDetailRequest) (BuildExecutionDetailResponse, error) {
    return BuildExecutionDetailResponse{
        ID:            req.LogID,
        WorkflowID:    req.WorkflowID,
        WorkflowTitle: req.WorkflowTitle,
        OrgID:         req.OrgID,
        OrgName:       req.OrgName,
        Status:        req.Status,
        CreditsSpent:  req.CreditsSpent,
        ExecutedAt:    req.ExecutedAt,
    }, nil
}
```

If you prefer a single sqlc JOIN query over an in-memory composer for `GetExecutionDetail`, that is also valid — declare `-- name: ExecutionLogGetDetail :one` with `LEFT JOIN workflows / organizations` and use the synthesized `ExecutionLogGetDetailRow` directly in `@get`. The `@call` path shown above is the one the SSaC below follows.

## SSaC

`service/dashboard/get_dashboard.ssac`:
```go
package service

import "github.com/park-jun-woo/zenflow/internal/dashboard"

// @get Organization org = Organization.FindByID({ID: currentUser.OrgID})
// @empty org "Organization not found" 404
// @call dashboard.SummarizeResponse summary = dashboard.Summarize({OrgID: currentUser.OrgID, OrgName: org.Name, PlanType: org.PlanType, CreditsBalance: org.CreditsBalance})
// @response { summary: summary }
func GetDashboard() {}
```

`service/audit/get_audit_log.ssac`:
```go
package service

// @get AuditLog audit_log = AuditLog.FindByID({ID: request.id})
// @empty audit_log "Audit log not found" 404
// @response { audit_log: audit_log }
func GetAuditLog() {}
```

`service/execution/get_execution_detail.ssac`:
```go
package service

import "github.com/park-jun-woo/zenflow/internal/dashboard"

// @get ExecutionLog log = ExecutionLog.FindByID({ID: request.id})
// @empty log "Execution log not found" 404
// @get Workflow wf = Workflow.FindByID({ID: log.WorkflowID})
// @empty wf "Workflow not found" 404
// @get Organization org = Organization.FindByID({ID: log.OrgID})
// @empty org "Organization not found" 404
// @call dashboard.BuildExecutionDetailResponse detail = dashboard.BuildExecutionDetail({LogID: log.ID, WorkflowID: wf.ID, WorkflowTitle: wf.Title, OrgID: org.ID, OrgName: org.Name, Status: log.Status, CreditsSpent: log.CreditsSpent, ExecutedAt: "now"})
// @response { detail: detail }
func GetExecutionDetail() {}
```

## TSX Dashboard Page

`frontend/pages/DashboardPage.tsx`:
```tsx
import { useQuery } from '@tanstack/react-query'
import { apiClient } from '@/lib/api'
import { Card } from '@/components/ui'
import { CreditsGauge } from '@/components/CreditsGauge'
import { ExecutionChart } from '@/components/ExecutionChart'
import { WorkflowStatusBadge } from '@/components/WorkflowStatusBadge'

export default function DashboardPage() {
  const { data } = useQuery(['GetDashboard'], apiClient.getDashboard)
  const summary = data?.summary
  if (!summary) return null

  return (
    <>
      <h1>Dashboard — {summary.org_name}</h1>
      <Card>
        <CreditsGauge summary={summary} />
        <ExecutionChart summary={summary} />
        <WorkflowStatusBadge summary={summary} />
      </Card>
    </>
  )
}
```

`frontend/components/CreditsGauge.tsx`:
```tsx
import React from 'react'
export type CreditsGaugeProps = { summary: { credits_balance: number; plan_type: string } }
export function CreditsGauge({ summary }: CreditsGaugeProps) {
  return <div>Credits: {summary.credits_balance} ({summary.plan_type})</div>
}
```

(Follow the same pattern for `ExecutionChart.tsx` and `WorkflowStatusBadge.tsx`.)

## AuditLog sqlc Query
```sql
-- name: AuditLogFindByID :one
SELECT * FROM audit_logs WHERE id = $1;
```

## E2E Scenario
- Create org + a few workflows + executions → `GET /dashboard` returns summary.
- `GET /execution-logs/{id}/detail` → `BuildExecutionDetailResponse` composite returned.
- TSX: `XOT-1` / `XOT-2` pass (apiClient call + args align with OpenAPI). Custom component imports resolve against `frontend/components/`.
