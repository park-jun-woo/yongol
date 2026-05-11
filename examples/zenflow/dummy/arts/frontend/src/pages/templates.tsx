import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useParams } from 'react-router-dom'
import { useForm } from 'react-hook-form'
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

  const publishTemplateForm = useForm()
  const publishTemplateMutation = useMutation({
    mutationFn: (data: any) => api.PublishTemplate(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['ListTemplates'] })
      queryClient.invalidateQueries({ queryKey: ['GetTemplate'] })
    },
  })

  const cloneTemplateMutation = useMutation({
    mutationFn: (data: any) => api.CloneTemplate({ ...data, id: id }),
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
            {listTemplatesData.items?.map((item: any, index: number) => (
              <li key={index}>
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
        <input type="number" placeholder="Source Workflow ID" {...publishTemplateForm.register('source_workflow_id', { valueAsNumber: true })} />
        <input type="text" placeholder="Title" {...publishTemplateForm.register('title')} />
        <input type="text" placeholder="Description" {...publishTemplateForm.register('description')} />
        <input type="text" placeholder="Category" {...publishTemplateForm.register('category')} />
        <button type="submit">Publish</button>
      </form>
      <div><button onClick={() => cloneTemplateMutation.mutate({})}>Clone Template</button></div>
    </main>
  )
}
