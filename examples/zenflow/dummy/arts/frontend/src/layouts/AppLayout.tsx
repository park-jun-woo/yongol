import { Link, Outlet } from 'react-router-dom'

export default function AppLayout() {
  return (
    <div>
      <nav>
        <Link to="/workflows">Workflows</Link>
        <Link to="/templates">Templates</Link>
        <Link to="/dashboard">Dashboard</Link>
      </nav>
      <Outlet />
    </div>
  )
}
