import { useMutation } from '@tanstack/react-query'
import { useForm } from 'react-hook-form'
import { apiClient } from '@/lib/api'
import { Button, Card, Input, Form } from '@/components/ui'

export default function LoginPage() {
  const login = useMutation({ mutationFn: apiClient.Login })
  const { register, handleSubmit } = useForm()

  return (
    <Card>
      <Form onSubmit={handleSubmit(v => login.mutate(v))}>
        <Input {...register('email', { required: true })} />
        <Input type="password" {...register('password', { required: true })} />
        <Button type="submit" variant="primary">Login</Button>
      </Form>
    </Card>
  )
}
