import { useQuery, useMutation } from '@tanstack/react-query'
import { useForm } from 'react-hook-form'
import { apiClient } from '@/lib/api'
import { Button, Card, Input, Form } from '@/components/ui'

export default function WorkflowsPage() {
  const { data } = useQuery({ queryKey: ['listWorkflows'], queryFn: apiClient.ListWorkflows })
  const create = useMutation({ mutationFn: apiClient.CreateWorkflow })
  const { register, handleSubmit } = useForm()

  return (
    <Card>
      <Form onSubmit={handleSubmit(v => create.mutate(v))}>
        <Input {...register('title', { required: true })} />
        <Input {...register('trigger_event', { required: true })} />
        <Button type="submit" variant="primary">Create</Button>
      </Form>
    </Card>
  )
}
