import { useMutation } from '@tanstack/react-query'
import { useForm } from 'react-hook-form'
import { apiClient } from '@/api'
import { Button } from '@/components/ui/Button'
import { Input } from '@/components/ui/Input'

export default function CreateWorkflowPage() {
  const create = useMutation({ mutationFn: apiClient.createWorkflow })
  const { register, handleSubmit } = useForm()
  return (
    <form onSubmit={handleSubmit(v => create.mutate(v))}>
      <Input {...register('title', { required: true })} />
      <Input {...register('trigger_event', { required: true })} />
      <Input {...register('description')} />
      <Button type="submit">Create</Button>
    </form>
  )
}
