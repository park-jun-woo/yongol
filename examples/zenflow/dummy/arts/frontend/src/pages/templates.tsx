import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useParams } from 'react-router-dom'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { api } from '@/lib/api'

export default function Templates() {
  const { id } = useParams()
  const queryClient = useQueryClient()

  const { data: listTemplatesData, isLoading: listTemplatesDataLoading, error: listTemplatesDataError } = useQuery({
    queryKey: ['ListTemplates'],
    queryFn: () => api.ListTemplates(),
  })

  const { data: getTemplateData, isLoading: getTemplateDataLoading, error: getTemplateDataError } = useQuery({
    queryKey: ['GetTemplate', id],
    queryFn: () => api.GetTemplate({ id: id }),
  })

  const publishTemplateSchema = z.object({
  category: z.string().min(1),
  description: z.string().min(1),
  source_workflow_id: z.number().int(),
  title: z.string().min(1),
})
  const publishTemplateForm = useForm({
    resolver: zodResolver(publishTemplateSchema),
  })
  const publishTemplateMutation = useMutation({
    mutationFn: (data) => api.PublishTemplate(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['ListTemplates'] })
      queryClient.invalidateQueries({ queryKey: ['GetTemplate'] })
    },
  })

  const cloneTemplateMutation = useMutation({
    mutationFn: (data) => api.CloneTemplate({ ...data, id: id }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['ListTemplates'] })
      queryClient.invalidateQueries({ queryKey: ['GetTemplate'] })
    },
  })

  return (
    <main>
      {listTemplatesDataLoading && <div>로딩 중...</div>}
      {listTemplatesDataError && <div>오류가 발생했습니다</div>}
      {listTemplatesData && (
        <section>
          <h2>Templates</h2>
          <ul>
            {listTemplatesData.items?.map((item) => (
              <li key={item.id}>
                <span>{item.title}</span>
                <span>{item.category}</span>
                <span>{item.clone_count}</span>
              </li>
            ))}
          </ul>
        </section>
      )}
      {getTemplateDataLoading && <div>로딩 중...</div>}
      {getTemplateDataError && <div>오류가 발생했습니다</div>}
      {getTemplateData && (
        <article>
          <h3>{getTemplateData.template.title}</h3>
          <p>{getTemplateData.template.description}</p>
          <p>{getTemplateData.template.category}</p>
        </article>
      )}
      <form onSubmit={publishTemplateForm.handleSubmit((data) => publishTemplateMutation.mutate(data))}>
        <h3>Publish Template</h3>
        <div>
          <label htmlFor="source_workflow_id">Source Workflow Id</label>
          <input id="source_workflow_id" type="number" placeholder="Source Workflow ID" {...publishTemplateForm.register('source_workflow_id', { valueAsNumber: true })} />
        </div>
        <div>
          <label htmlFor="title">Title</label>
          <input id="title" type="text" placeholder="Title" {...publishTemplateForm.register('title')} />
        </div>
        <div>
          <label htmlFor="description">Description</label>
          <input id="description" type="text" placeholder="Description" {...publishTemplateForm.register('description')} />
        </div>
        <div>
          <label htmlFor="category">Category</label>
          <input id="category" type="text" placeholder="Category" {...publishTemplateForm.register('category')} />
        </div>
        <button type="submit" disabled={publishTemplateMutation.isPending}>{publishTemplateMutation.isPending ? '처리 중...' : 'Publish'}</button>
      </form>
      <div><button onClick={() => cloneTemplateMutation.mutate({})} disabled={cloneTemplateMutation.isPending}>{cloneTemplateMutation.isPending ? '처리 중...' : 'Clone Template'}</button></div>
    </main>
  )
}
