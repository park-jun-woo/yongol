import { useMutation } from '@tanstack/react-query'
import { useForm } from 'react-hook-form'
import { apiClient } from '@/lib/api'
import { Button, Card, Input, Form } from '@/components/ui'

export default function LoginPage() {
  const login = useMutation({ mutationFn: apiClient.Login })
  const { register, handleSubmit } = useForm()

  return (
    <Card>
      <Form onSubmit={handleSubmit((v: any) => login.mutate(v))}>
        <Input {...register('email', { required: true })} />
        <Input {...register('password', { required: true })} type="password" />
        <Button type="submit" variant="primary">Login</Button>
      </Form>
    </Card>
  )
}
