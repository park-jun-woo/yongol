import { useQuery } from '@tanstack/react-query'
import { api } from '@/lib/api'

export default function Dashboard() {

  const { data: getDashboardData, isLoading: getDashboardDataLoading, error: getDashboardDataError } = useQuery({
    queryKey: ['GetDashboard'],
    queryFn: () => api.GetDashboard(),
  })

  return (
    <main>
      {getDashboardDataLoading && <div>로딩 중...</div>}
      {getDashboardDataError && <div>오류가 발생했습니다</div>}
      {getDashboardData && (
        <section>
          <h2>Dashboard</h2>
          <div>
            <span>{getDashboardData.summary.name}</span>
            <span>{getDashboardData.summary.plan_type}</span>
            <span>{getDashboardData.summary.credits_balance}</span>
          </div>
        </section>
      )}
    </main>
  )
}
