import { useMutation, useQuery } from '@tanstack/react-query'
import { useForm } from 'react-hook-form'
import { apiClient } from '@/api'

export default function WorkflowsPage() {
  const { data } = useQuery(['ListWorkflows'], () => apiClient.ListWorkflows({ page: 1, per_page: 20, sort_by: 'created_at', sort_dir: 'desc' }))
  const create = useMutation({ mutationFn: apiClient.CreateWorkflow })
  const { register, handleSubmit } = useForm()

  return (
    <div>
      <h1>Workflows</h1>
      <ul>
        {data?.items?.map((w: { id: number; title: string; status: string }) => (
          <li key={w.id}>{w.title} — {w.status}</li>
        ))}
      </ul>
      <form onSubmit={handleSubmit(v => create.mutate(v as any))}>
        <input {...register('title', { required: true })} />
        <input {...register('trigger_event', { required: true })} />
        <button type="submit">Create</button>
      </form>
    </div>
  )
}
