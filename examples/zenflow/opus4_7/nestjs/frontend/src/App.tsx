import { Routes, Route } from 'react-router-dom'

// Add your pages under src/pages/*.tsx and wire them below.
//   e.g. <Route path="/workflows" element={<WorkflowsPage />} />
export default function App() {
  return (
    <Routes>
      <Route path="/" element={<div className="p-6">yongol scaffolded frontend — add pages under src/pages/</div>} />
    </Routes>
  )
}
