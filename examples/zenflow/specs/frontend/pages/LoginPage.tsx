import { useMutation } from '@tanstack/react-query'
import { useForm } from 'react-hook-form'
import { apiClient } from '@/api'

export default function LoginPage() {
  const { register, handleSubmit } = useForm()
  const login = useMutation({ mutationFn: apiClient.Login })

  return (
    <form onSubmit={handleSubmit(v => login.mutate(v as any))}>
      <input {...register('email', { required: true })} />
      <input {...register('password', { required: true })} type="password" />
      <button type="submit">Sign in</button>
    </form>
  )
}
