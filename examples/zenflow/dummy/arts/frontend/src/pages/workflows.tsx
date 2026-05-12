import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { api } from '@/lib/api'

export default function Workflows() {
  const queryClient = useQueryClient()

  const { data: listWorkflowsData, isLoading: listWorkflowsDataLoading, error: listWorkflowsDataError } = useQuery({
    queryKey: ['ListWorkflows'],
    queryFn: () => api.ListWorkflows(),
  })

  const createWorkflowSchema = z.object({
  title: z.string().min(1),
  trigger_event: z.string().min(1),
})
  const createWorkflowForm = useForm({
    resolver: zodResolver(createWorkflowSchema),
  })
  const createWorkflowMutation = useMutation({
    mutationFn: (data) => api.CreateWorkflow(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['ListWorkflows'] })
    },
  })

  return (
    <main>
      {listWorkflowsDataLoading && <div>로딩 중...</div>}
      {listWorkflowsDataError && <div>오류가 발생했습니다</div>}
      {listWorkflowsData && (
        <section>
          <h2>Workflows</h2>
          <ul>
            {listWorkflowsData.workflows?.map((item) => (
              <li key={item.id}>
                <span>{item.title}</span>
                <span>{item.status}</span>
              </li>
            ))}
          </ul>
        </section>
      )}
      <form onSubmit={createWorkflowForm.handleSubmit((data) => createWorkflowMutation.mutate(data))}>
        <h3>New Workflow</h3>
        <div>
          <label htmlFor="title">Title</label>
          <input id="title" type="text" placeholder="Title" {...createWorkflowForm.register('title')} />
        </div>
        <div>
          <label htmlFor="trigger_event">Trigger Event</label>
          <input id="trigger_event" type="text" placeholder="Trigger Event" {...createWorkflowForm.register('trigger_event')} />
        </div>
        <button type="submit" disabled={createWorkflowMutation.isPending}>{createWorkflowMutation.isPending ? '처리 중...' : 'Create'}</button>
      </form>
    </main>
  )
}
