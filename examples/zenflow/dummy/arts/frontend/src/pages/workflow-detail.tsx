import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useParams } from 'react-router-dom'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
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

  const { data: listWorkflowVersionsData, isLoading: listWorkflowVersionsDataLoading, error: listWorkflowVersionsDataError } = useQuery({
    queryKey: ['ListWorkflowVersions', id],
    queryFn: () => api.ListWorkflowVersions({ id: id }),
  })

  const createActionSchema = z.object({
  action_type: z.string().min(1),
  config: z.string().min(1),
  sequence_order: z.number().int(),
})
  const createActionForm = useForm({
    resolver: zodResolver(createActionSchema),
  })
  const createActionMutation = useMutation({
    mutationFn: (data) => api.CreateAction({ ...data, id: id }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['GetWorkflow'] })
      queryClient.invalidateQueries({ queryKey: ['ListActions'] })
      queryClient.invalidateQueries({ queryKey: ['ListExecutionLogs'] })
      queryClient.invalidateQueries({ queryKey: ['ListWorkflowVersions'] })
    },
  })

  const createWorkflowVersionMutation = useMutation({
    mutationFn: (data) => api.CreateWorkflowVersion({ ...data, id: id }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['GetWorkflow'] })
      queryClient.invalidateQueries({ queryKey: ['ListActions'] })
      queryClient.invalidateQueries({ queryKey: ['ListExecutionLogs'] })
      queryClient.invalidateQueries({ queryKey: ['ListWorkflowVersions'] })
    },
  })

  const activateWorkflowMutation = useMutation({
    mutationFn: (data) => api.ActivateWorkflow({ ...data, id: id }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['GetWorkflow'] })
    },
  })

  const pauseWorkflowMutation = useMutation({
    mutationFn: (data) => api.PauseWorkflow({ ...data, id: id }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['GetWorkflow'] })
    },
  })

  const archiveWorkflowMutation = useMutation({
    mutationFn: (data) => api.ArchiveWorkflow({ ...data, id: id }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['GetWorkflow'] })
    },
  })

  const executeWorkflowMutation = useMutation({
    mutationFn: (data) => api.ExecuteWorkflow({ ...data, id: id }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['GetWorkflow'] })
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
              <button onClick={() => activateWorkflowMutation.mutate({})} disabled={activateWorkflowMutation.isPending}>{activateWorkflowMutation.isPending ? '처리 중...' : 'Activate'}</button>
            </footer>
          )}
          {getWorkflowData.workflow.status === 'active' && (
            <footer>
              <button onClick={() => pauseWorkflowMutation.mutate({})} disabled={pauseWorkflowMutation.isPending}>{pauseWorkflowMutation.isPending ? '처리 중...' : 'Pause'}</button>
              <button onClick={() => archiveWorkflowMutation.mutate({})} disabled={archiveWorkflowMutation.isPending}>{archiveWorkflowMutation.isPending ? '처리 중...' : 'Archive'}</button>
              <button onClick={() => executeWorkflowMutation.mutate({})} disabled={executeWorkflowMutation.isPending}>{executeWorkflowMutation.isPending ? '처리 중...' : 'Execute'}</button>
            </footer>
          )}
          {getWorkflowData.workflow.status === 'paused' && (
            <footer>
              <button onClick={() => activateWorkflowMutation.mutate({})} disabled={activateWorkflowMutation.isPending}>{activateWorkflowMutation.isPending ? '처리 중...' : 'Resume'}</button>
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
            {listActionsData.actions?.map((item) => (
              <li key={item.id}>
                <span>{item.action_type}</span>
                <span>{item.sequence_order}</span>
              </li>
            ))}
          </ul>
        </section>
      )}
      <form onSubmit={createActionForm.handleSubmit((data) => createActionMutation.mutate(data))}>
        <h3>Add Action</h3>
        <div>
          <label htmlFor="action_type">Action Type</label>
          <input id="action_type" type="text" placeholder="Action Type" {...createActionForm.register('action_type')} />
        </div>
        <div>
          <label htmlFor="config">Config</label>
          <input id="config" type="text" placeholder="Config" {...createActionForm.register('config')} />
        </div>
        <div>
          <label htmlFor="sequence_order">Sequence Order</label>
          <input id="sequence_order" type="number" placeholder="Order" {...createActionForm.register('sequence_order', { valueAsNumber: true })} />
        </div>
        <button type="submit" disabled={createActionMutation.isPending}>{createActionMutation.isPending ? '처리 중...' : 'Add'}</button>
      </form>
      {listExecutionLogsDataLoading && <div>로딩 중...</div>}
      {listExecutionLogsDataError && <div>오류가 발생했습니다</div>}
      {listExecutionLogsData && (
        <section>
          <h3>Execution Logs</h3>
          <ul>
            {listExecutionLogsData.execution_logs?.map((item) => (
              <li key={item.id}>
                <span>{item.status}</span>
                <span>{item.credits_spent}</span>
                <span>{item.executed_at}</span>
              </li>
            ))}
          </ul>
        </section>
      )}
      {listWorkflowVersionsDataLoading && <div>로딩 중...</div>}
      {listWorkflowVersionsDataError && <div>오류가 발생했습니다</div>}
      {listWorkflowVersionsData && (
        <section>
          <h3>Versions</h3>
          <ul>
            {listWorkflowVersionsData.workflows?.map((item) => (
              <li key={item.id}>
                <span>{item.version}</span>
                <span>{item.status}</span>
                <span>{item.created_at}</span>
              </li>
            ))}
          </ul>
        </section>
      )}
      <div><button onClick={() => createWorkflowVersionMutation.mutate({})} disabled={createWorkflowVersionMutation.isPending}>{createWorkflowVersionMutation.isPending ? '처리 중...' : 'New Version'}</button></div>
    </main>
  )
}
