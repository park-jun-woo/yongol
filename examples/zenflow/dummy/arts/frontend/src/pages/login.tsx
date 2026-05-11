'use client'

import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useForm } from 'react-hook-form'
import { api } from '@/lib/api'

export default function Login() {
  const queryClient = useQueryClient()

  const loginForm = useForm()
  const loginMutation = useMutation({
    mutationFn: (data: any) => api.Login(data),
    onSuccess: () => {
      queryClient.invalidateQueries()
    },
  })

  return (
    <main>
      <form onSubmit={loginForm.handleSubmit((data) => loginMutation.mutate(data))}>
        <h2>Login</h2>
        <input type="email" placeholder="Email" {...loginForm.register('email')} />
        <input type="password" placeholder="Password" {...loginForm.register('password')} />
        <button type="submit">Login</button>
      </form>
    </main>
  )
}
