import { Routes, Route } from 'react-router-dom'
import AuditLogs from './pages/audit-logs'
import Dashboard from './pages/dashboard'
import Login from './pages/login'
import Register from './pages/register'
import Templates from './pages/templates'
import Webhooks from './pages/webhooks'
import Workflows from './pages/workflows'
import WorkflowDetail from './pages/workflow-detail'
import AppLayout from './layouts/AppLayout'
import AuthLayout from './layouts/AuthLayout'

export default function App() {
  return (
    <Routes>
      <Route element={<AppLayout />}>
        <Route path="/audit-logs" element={<AuditLogs />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/templates" element={<Templates />} />
        <Route path="/templates/:id" element={<Templates />} />
        <Route path="/webhooks" element={<Webhooks />} />
        <Route path="/webhooks/:id" element={<Webhooks />} />
        <Route path="/workflows" element={<Workflows />} />
        <Route path="/workflows/:id" element={<WorkflowDetail />} />
      </Route>
      <Route element={<AuthLayout />}>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
      </Route>
    </Routes>
  )
}
