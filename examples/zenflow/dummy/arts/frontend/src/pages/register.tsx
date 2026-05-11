'use client'

import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useForm } from 'react-hook-form'
import { api } from '@/lib/api'

export default function Register() {
  const queryClient = useQueryClient()

  const registerForm = useForm()
  const registerMutation = useMutation({
    mutationFn: (data: any) => api.Register(data),
    onSuccess: () => {
      queryClient.invalidateQueries()
    },
  })

  return (
    <main>
      <form onSubmit={registerForm.handleSubmit((data) => registerMutation.mutate(data))}>
        <h2>Register</h2>
        <input type="email" placeholder="Email" {...registerForm.register('email')} />
        <input type="password" placeholder="Password" {...registerForm.register('password')} />
        <input type="number" placeholder="Organization ID" {...registerForm.register('org_id', { valueAsNumber: true })} />
        <input type="text" placeholder="Role" {...registerForm.register('role')} />
        <button type="submit">Register</button>
      </form>
    </main>
  )
}
