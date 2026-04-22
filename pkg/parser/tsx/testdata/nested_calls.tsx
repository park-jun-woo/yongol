import { useQuery } from '@tanstack/react-query'
import { apiClient } from '@/api'

export default function DashboardPage() {
  const workflows = useQuery(['listWorkflows'], () => apiClient.listWorkflows())
  const executions = useQuery(['listExecutions'], () => apiClient.listExecutions({ status: 'running' }))
  return (
    <div>
      <p>workflows: {workflows.data?.total}</p>
      <p>executions: {executions.data?.total}</p>
    </div>
  )
}
