import { useQuery } from '@tanstack/react-query'
import { api } from '@/lib/api'

export default function AuditLogs() {

  const { data: listAuditLogsData, isLoading: listAuditLogsDataLoading, error: listAuditLogsDataError } = useQuery({
    queryKey: ['ListAuditLogs'],
    queryFn: () => api.ListAuditLogs(),
  })

  return (
    <main>
      {listAuditLogsDataLoading && <div>로딩 중...</div>}
      {listAuditLogsDataError && <div>오류가 발생했습니다</div>}
      {listAuditLogsData && (
        <section>
          <h2>Audit Logs</h2>
          <ul>
            {listAuditLogsData.items?.map((item: any, index: number) => (
              <li key={index}>
                <span>{item.action}</span>
                <span>{item.resource_type}</span>
                <span>{item.detail}</span>
                <span>{item.created_at}</span>
              </li>
            ))}
          </ul>
        </section>
      )}
    </main>
  )
}
