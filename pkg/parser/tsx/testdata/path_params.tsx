import { useQuery } from '@tanstack/react-query'
import { useParams } from 'react-router-dom'
import { apiClient } from '@/api'

export default function WorkflowDetailPage() {
  const params = useParams()
  const { data } = useQuery(
    ['getWorkflow', params.id],
    () => apiClient.getWorkflow({ id: params.id!, version: 'v2' }),
  )
  return <div>{data?.workflow?.title}</div>
}
