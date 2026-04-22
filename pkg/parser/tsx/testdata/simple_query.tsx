import { useQuery } from '@tanstack/react-query'
import { apiClient } from '@/api'

export default function WorkflowsPage() {
  const { data } = useQuery(['listWorkflows'], () => apiClient.listWorkflows())
  return <ul>{data?.workflows?.map(w => <li key={w.id}>{w.title}</li>)}</ul>
}
