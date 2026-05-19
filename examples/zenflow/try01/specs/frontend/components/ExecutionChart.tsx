import React from 'react'
export type ExecutionChartProps = { summary: { total_executions: number } }
export function ExecutionChart({ summary }: ExecutionChartProps) {
  return <div>Total Executions: {summary.total_executions}</div>
}
