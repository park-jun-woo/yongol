import { Routes, Route, Link } from 'react-router-dom'
import Login from './pages/login'
import Register from './pages/register'
import Workflows from './pages/workflows'
import WorkflowDetail from './pages/workflow-detail'
import Dashboard from './pages/dashboard'

export default function App() {
  return (
    <div className="min-h-screen">
      <nav className="border-b bg-white px-6 py-3 flex gap-4 items-center">
        <Link to="/" className="font-bold text-lg">ZenFlow</Link>
        <Link to="/workflows" className="text-sm hover:underline">Workflows</Link>
        <Link to="/dashboard" className="text-sm hover:underline">Dashboard</Link>
        <div className="ml-auto flex gap-2">
          <Link to="/login" className="text-sm hover:underline">Login</Link>
          <Link to="/register" className="text-sm hover:underline">Register</Link>
        </div>
      </nav>
      <div className="p-6">
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/workflows" element={<Workflows />} />
          <Route path="/workflows/:id" element={<WorkflowDetail />} />
          <Route path="/dashboard" element={<Dashboard />} />
        </Routes>
      </div>
    </div>
  )
}
