# ZenFlow Add-on #07 — Dashboard + Relation Enrichment

## Overview
An operator dashboard summarizing org state on one screen. Exercises two yongol features together: **`@dto`** (aggregate models with no backing table) and **TSX custom components** (widgets too complex for plain data binding).

## Verification Points
- **`@dto` models** — `DashboardSummary`, `ExecutionDetail`; no DDL table; assembled via `@call dashboard.Summarize` / `BuildExecutionDetail`.
- **TSX custom components** — widgets (credits gauge, execution chart, status badge) live under `frontend/components/*.tsx` and are invoked by the page component.

## New Endpoints
- **GET /dashboard** (`GetDashboard`)
  - Response: `{ summary: DashboardSummary }`
  - Flow: `@get Organization org` + `@call dashboard.Summarize({OrgID, ...})` → aggregated object.
- **GET /audit-logs/{id}** (`GetAuditLog`) — detail view; caller can enrich client-side by calling `GetUser` and `GetOrganization` separately.
- **GET /execution-logs/{id}/detail** (`GetExecutionDetail`)
  - Response type: `ExecutionDetail` (`@dto`, combines log + workflow + org).

## `@dto` Models (`model/`)

```go
// @dto
type DashboardSummary struct {
    OrgName           string
    PlanType          string
    CreditsBalance    int64
    ActiveWorkflows   int64
    PausedWorkflows   int64
    TotalExecutions   int64
    TotalCreditsSpent int64
}

// @dto
type ExecutionDetail struct {
    ID            int64
    WorkflowID    int64
    WorkflowTitle string
    OrgID         int64
    OrgName       string
    Status        string
    CreditsSpent  int64
    ExecutedAt    string
}
```

## DDL
No changes — reuse existing `workflows`, `execution_logs`, `audit_logs`, `organizations`, `users`.

## Custom Functions
- `dashboard.Summarize(OrgID, OrgName, PlanType, CreditsBalance)` — compose the aggregate object in-memory. Stats are simulated (random or constant) to stay purity-safe.
- `dashboard.BuildExecutionDetail(LogID, WorkflowID, OrgID, Status, CreditsSpent, ExecutedAt, WorkflowTitle, OrgName)` — assemble the `ExecutionDetail` DTO.

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
- `GET /execution-logs/{id}/detail` → `ExecutionDetail` composite returned.
- TSX: `XOT-1` / `XOT-2` pass (apiClient call + args align with OpenAPI). Custom component imports resolve against `frontend/components/`.
