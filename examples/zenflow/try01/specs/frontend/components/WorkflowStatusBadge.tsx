import React from 'react'
export type WorkflowStatusBadgeProps = { summary: { active_workflows: number; paused_workflows: number } }
export function WorkflowStatusBadge({ summary }: WorkflowStatusBadgeProps) {
  return <div>Active: {summary.active_workflows} | Paused: {summary.paused_workflows}</div>
}
