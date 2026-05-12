//ff:func feature=gen-react type=generator control=sequence
//ff:what STML 페이지 미존재 시 placeholder App.tsx를 생성한다

package react

import (
	"os"
	"path/filepath"
)

// writeAppTSXPlaceholder emits a minimal placeholder App.tsx when no STML
// pages are available.
func writeAppTSXPlaceholder(srcDir string) error {
	const src = `import { Routes, Route } from 'react-router-dom'

// Add your pages under src/pages/*.tsx and wire them below.
//   e.g. <Route path="/workflows" element={<WorkflowsPage />} />
export default function App() {
  return (
    <Routes>
      <Route path="/" element={<div className="p-6">yongol scaffolded frontend — add pages under src/pages/</div>} />
    </Routes>
  )
}
`
	return os.WriteFile(filepath.Join(srcDir, "App.tsx"), []byte(src), 0644)
}
