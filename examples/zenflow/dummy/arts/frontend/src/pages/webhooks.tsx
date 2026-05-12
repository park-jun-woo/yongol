import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useParams } from 'react-router-dom'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { api } from '@/lib/api'

export default function Webhooks() {
  const { id } = useParams()
  const queryClient = useQueryClient()

  const { data: listWebhooksData, isLoading: listWebhooksDataLoading, error: listWebhooksDataError } = useQuery({
    queryKey: ['ListWebhooks'],
    queryFn: () => api.ListWebhooks(),
  })

  const createWebhookSchema = z.object({
  event_type: z.string().min(1),
  url: z.string().min(1),
})
  const createWebhookForm = useForm({
    resolver: zodResolver(createWebhookSchema),
  })
  const createWebhookMutation = useMutation({
    mutationFn: (data) => api.CreateWebhook(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['ListWebhooks'] })
    },
  })

  const deleteWebhookMutation = useMutation({
    mutationFn: (data) => api.DeleteWebhook({ ...data, id: id }),
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
            {listWebhooksData.webhooks?.map((item) => (
              <li key={item.id}>
                <span>{item.url}</span>
                <span>{item.event_type}</span>
              </li>
            ))}
          </ul>
        </section>
      )}
      <form onSubmit={createWebhookForm.handleSubmit((data) => createWebhookMutation.mutate(data))}>
        <h3>Add Webhook</h3>
        <div>
          <label htmlFor="url">Url</label>
          <input id="url" type="text" placeholder="Webhook URL" {...createWebhookForm.register('url')} />
        </div>
        <div>
          <label htmlFor="event_type">Event Type</label>
          <input id="event_type" type="text" placeholder="Event Type" {...createWebhookForm.register('event_type')} />
        </div>
        <button type="submit" disabled={createWebhookMutation.isPending}>{createWebhookMutation.isPending ? '처리 중...' : 'Create'}</button>
      </form>
      <div><button onClick={() => deleteWebhookMutation.mutate({})} disabled={deleteWebhookMutation.isPending}>{deleteWebhookMutation.isPending ? '처리 중...' : 'Delete Webhook'}</button></div>
    </main>
  )
}
