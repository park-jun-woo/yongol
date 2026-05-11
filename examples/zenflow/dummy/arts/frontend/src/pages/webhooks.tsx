import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useParams } from 'react-router-dom'
import { useForm } from 'react-hook-form'
import { api } from '@/lib/api'

export default function Webhooks() {
  const { id } = useParams()
  const queryClient = useQueryClient()

  const { data: listWebhooksData, isLoading: listWebhooksDataLoading, error: listWebhooksDataError } = useQuery({
    queryKey: ['ListWebhooks'],
    queryFn: () => api.ListWebhooks(),
  })

  const createWebhookForm = useForm()
  const createWebhookMutation = useMutation({
    mutationFn: (data: any) => api.CreateWebhook(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['ListWebhooks'] })
    },
  })

  const deleteWebhookMutation = useMutation({
    mutationFn: (data: any) => api.DeleteWebhook({ ...data, id: id }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['ListWebhooks'] })
    },
  })

  return (
    <main>
      {listWebhooksDataLoading && <div>로딩 중...</div>}
      {listWebhooksDataError && <div>오류가 발생했습니다</div>}
      {listWebhooksData && (
        <section>
          <h2>Webhooks</h2>
          <ul>
            {listWebhooksData.webhooks?.map((item: any, index: number) => (
              <li key={index}>
                <span>{item.url}</span>
                <span>{item.event_type}</span>
              </li>
            ))}
          </ul>
        </section>
      )}
      <form onSubmit={createWebhookForm.handleSubmit((data) => createWebhookMutation.mutate(data))}>
        <h3>Add Webhook</h3>
        <input type="text" placeholder="Webhook URL" {...createWebhookForm.register('url')} />
        <input type="text" placeholder="Event Type" {...createWebhookForm.register('event_type')} />
        <button type="submit">Create</button>
      </form>
      <div><button onClick={() => deleteWebhookMutation.mutate({})}>Delete Webhook</button></div>
    </main>
  )
}
