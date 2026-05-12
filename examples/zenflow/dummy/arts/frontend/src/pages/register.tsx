import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useForm } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { api } from '@/lib/api'

export default function Register() {
  const queryClient = useQueryClient()

  const registerSchema = z.object({
  email: z.string().email().min(1),
  org_id: z.number().int(),
  password: z.string().min(1).min(8),
  role: z.enum(["admin", "member"]),
})
  const registerForm = useForm({
    resolver: zodResolver(registerSchema),
  })
  const registerMutation = useMutation({
    mutationFn: (data) => api.Register(data),
    onSuccess: () => {
      queryClient.invalidateQueries()
    },
  })

  return (
    <main>
      <form onSubmit={registerForm.handleSubmit((data) => registerMutation.mutate(data))}>
        <h2>Register</h2>
        <div>
          <label htmlFor="email">Email</label>
          <input id="email" type="email" placeholder="Email" {...registerForm.register('email')} />
        </div>
        <div>
          <label htmlFor="password">Password</label>
          <input id="password" type="password" placeholder="Password" {...registerForm.register('password')} />
        </div>
        <div>
          <label htmlFor="org_id">Org Id</label>
          <input id="org_id" type="number" placeholder="Organization ID" {...registerForm.register('org_id', { valueAsNumber: true })} />
        </div>
        <div>
          <label htmlFor="role">Role</label>
          <input id="role" type="text" placeholder="Role" {...registerForm.register('role')} />
        </div>
        <button type="submit" disabled={registerMutation.isPending}>{registerMutation.isPending ? '처리 중...' : 'Register'}</button>
      </form>
    </main>
  )
}
