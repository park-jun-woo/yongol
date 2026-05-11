'use client'

import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useParams } from 'react-router-dom'
import { useForm } from 'react-hook-form'
import { api } from '@/lib/api'

export default function WorkflowDetail() {
  const { id } = useParams()
  const queryClient = useQueryClient()

  const { data: getWorkflowData, isLoading: getWorkflowDataLoading, error: getWorkflowDataError } = useQuery({
    queryKey: ['GetWorkflow', id],
    queryFn: () => api.GetWorkflow({ id: id }),
  })

  const { data: listActionsData, isLoading: listActionsDataLoading, error: listActionsDataError } = useQuery({
    queryKey: ['ListActions', id],
    queryFn: () => api.ListActions({ id: id }),
  })

  const { data: listExecutionLogsData, isLoading: listExecutionLogsDataLoading, error: listExecutionLogsDataError } = useQuery({
    queryKey: ['ListExecutionLogs', id],
    queryFn: () => api.ListExecutionLogs({ id: id }),
  })

  const createActionForm = useForm()
  const createActionMutation = useMutation({
    mutationFn: (data: any) => api.CreateAction({ ...data, id: id }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['GetWorkflow'] })
      queryClient.invalidateQueries({ queryKey: ['ListActions'] })
      queryClient.invalidateQueries({ queryKey: ['ListExecutionLogs'] })
    },
  })

  const activateWorkflowMutation = useMutation({
    mutationFn: (data: any) => api.ActivateWorkflow({ ...data, id: id }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['GetWorkflow'] })
      queryClient.invalidateQueries({ queryKey: ['ListActions'] })
      queryClient.invalidateQueries({ queryKey: ['ListExecutionLogs'] })
    },
  })

  const pauseWorkflowMutation = useMutation({
    mutationFn: (data: any) => api.PauseWorkflow({ ...data, id: id }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['GetWorkflow'] })
      queryClient.invalidateQueries({ queryKey: ['ListActions'] })
      queryClient.invalidateQueries({ queryKey: ['ListExecutionLogs'] })
    },
  })

  const archiveWorkflowMutation = useMutation({
    mutationFn: (data: any) => api.ArchiveWorkflow({ ...data, id: id }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['GetWorkflow'] })
      queryClient.invalidateQueries({ queryKey: ['ListActions'] })
      queryClient.invalidateQueries({ queryKey: ['ListExecutionLogs'] })
    },
  })

  const executeWorkflowMutation = useMutation({
    mutationFn: (data: any) => api.ExecuteWorkflow({ ...data, id: id }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['GetWorkflow'] })
      queryClient.invalidateQueries({ queryKey: ['ListActions'] })
      queryClient.invalidateQueries({ queryKey: ['ListExecutionLogs'] })
    },
  })

  return (
    <main>
      {getWorkflowDataLoading && <div>로딩 중...</div>}
      {getWorkflowDataError && <div>오류가 발생했습니다</div>}
      {getWorkflowData && (
        <article>
          <h2>{getWorkflowData.workflow.title}</h2>
          <p>{getWorkflowData.workflow.status}</p>
          <p>{getWorkflowData.workflow.trigger_event}</p>
          {getWorkflowData.workflow.status === 'draft' && (
            <footer>
              <button onClick={() => activateWorkflowMutation.mutate({})}>Activate</button>
            </footer>
          )}
          {getWorkflowData.workflow.status === 'active' && (
            <footer>
              <button onClick={() => pauseWorkflowMutation.mutate({})}>Pause</button>
              <button onClick={() => archiveWorkflowMutation.mutate({})}>Archive</button>
              <button onClick={() => executeWorkflowMutation.mutate({})}>Execute</button>
            </footer>
          )}
          {getWorkflowData.workflow.status === 'paused' && (
            <footer>
              <button onClick={() => activateWorkflowMutation.mutate({})}>Resume</button>
            </footer>
          )}
        </article>
      )}
      {listActionsDataLoading && <div>로딩 중...</div>}
      {listActionsDataError && <div>오류가 발생했습니다</div>}
      {listActionsData && (
        <section>
          <h3>Actions</h3>
          <ul>
            {listActionsData.actions?.map((item: any, index: number) => (
              <li key={index}>
                <span>{item.action_type}</span>
                <span>{item.sequence_order}</span>
              </li>
            ))}
          </ul>
        </section>
      )}
      <form onSubmit={createActionForm.handleSubmit((data) => createActionMutation.mutate(data))}>
        <h3>Add Action</h3>
        <input type="text" placeholder="Action Type" {...createActionForm.register('action_type')} />
        <input type="text" placeholder="Config" {...createActionForm.register('config')} />
        <input type="number" placeholder="Order" {...createActionForm.register('sequence_order', { valueAsNumber: true })} />
        <button type="submit">Add</button>
      </form>
      {listExecutionLogsDataLoading && <div>로딩 중...</div>}
      {listExecutionLogsDataError && <div>오류가 발생했습니다</div>}
      {listExecutionLogsData && (
        <section>
          <h3>Execution Logs</h3>
          <ul>
            {listExecutionLogsData.execution_logs?.map((item: any, index: number) => (
              <li key={index}>
                <span>{item.status}</span>
                <span>{item.credits_spent}</span>
                <span>{item.executed_at}</span>
              </li>
            ))}
          </ul>
        </section>
      )}
    </main>
  )
}
