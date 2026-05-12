import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { api } from '@/lib/api'

export default function Login() {
  const queryClient = useQueryClient()

  const loginSchema = z.object({
  email: z.string().email().min(1),
  password: z.string().min(1).min(8),
})
  const loginForm = useForm({
    resolver: zodResolver(loginSchema),
  })
  const loginMutation = useMutation({
    mutationFn: (data) => api.Login(data),
    onSuccess: () => {
      queryClient.invalidateQueries()
    },
  })

  return (
    <main>
      <form onSubmit={loginForm.handleSubmit((data) => loginMutation.mutate(data))}>
        <h2>Login</h2>
        <div>
          <label htmlFor="email">Email</label>
          <input id="email" type="email" placeholder="Email" {...loginForm.register('email')} />
        </div>
        <div>
          <label htmlFor="password">Password</label>
          <input id="password" type="password" placeholder="Password" {...loginForm.register('password')} />
        </div>
        <button type="submit" disabled={loginMutation.isPending}>{loginMutation.isPending ? '처리 중...' : 'Login'}</button>
      </form>
    </main>
  )
}
