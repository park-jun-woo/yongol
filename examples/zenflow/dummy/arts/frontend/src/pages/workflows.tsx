'use client'

import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useForm } from 'react-hook-form'
import { api } from '@/lib/api'

export default function Workflows() {
  const queryClient = useQueryClient()

  const { data: listWorkflowsData, isLoading: listWorkflowsDataLoading, error: listWorkflowsDataError } = useQuery({
    queryKey: ['ListWorkflows'],
    queryFn: () => api.ListWorkflows(),
  })

  const createWorkflowForm = useForm()
  const createWorkflowMutation = useMutation({
    mutationFn: (data: any) => api.CreateWorkflow(data),
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
            {listWorkflowsData.workflows?.map((item: any, index: number) => (
              <li key={index}>
                <span>{item.title}</span>
                <span>{item.status}</span>
              </li>
            ))}
          </ul>
        </section>
      )}
      <form onSubmit={createWorkflowForm.handleSubmit((data) => createWorkflowMutation.mutate(data))}>
        <h3>New Workflow</h3>
        <input type="text" placeholder="Title" {...createWorkflowForm.register('title')} />
        <input type="text" placeholder="Trigger Event" {...createWorkflowForm.register('trigger_event')} />
        <button type="submit">Create</button>
      </form>
    </main>
  )
}
