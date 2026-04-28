import { useQuery, useMutation } from '@tanstack/react-query'
import { useForm } from 'react-hook-form'
import { apiClient } from '@/lib/api'
import { Button, Card, Input, Form, Table, Badge } from '@/components/ui'

export default function WorkflowsPage() {
  const { data } = useQuery({ queryKey: ['ListWorkflows'], queryFn: apiClient.ListWorkflows })
  const create = useMutation({ mutationFn: apiClient.CreateWorkflow })
  const activate = useMutation({ mutationFn: apiClient.ActivateWorkflow })
  const execute = useMutation({ mutationFn: apiClient.ExecuteWorkflow })
  const { register, handleSubmit } = useForm()

  return (
    <Card>
      <Form onSubmit={handleSubmit(v => create.mutate(v))}>
        <Input {...register('title', { required: true })} />
        <Input {...register('trigger_event', { required: true })} />
        <Button type="submit" variant="primary">Create Workflow</Button>
      </Form>
      <Table>
        {(data?.items ?? []).map((w: any) => (
          <tr key={w.id}>
            <td>{w.title}</td>
            <td><Badge>{w.status}</Badge></td>
            <td>
              <Button onClick={() => activate.mutate({ id: w.id })}>Activate</Button>
              <Button onClick={() => execute.mutate({ id: w.id })}>Execute</Button>
            </td>
          </tr>
        ))}
      </Table>
    </Card>
  )
}
