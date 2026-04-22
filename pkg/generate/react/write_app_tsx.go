//ff:func feature=gen-react type=generator control=sequence
//ff:what App.tsx — 빈 Routes 스캐폴드 (AI 가 pages/*.tsx 생성하여 여기에 추가)

package react

import (
	"os"
	"path/filepath"
)

// writeAppTSX emits a minimal App.tsx. Pages under src/pages/*.tsx are the
// user's (or AI's) responsibility per the TSX SSOT design — yongol does
// not scan or auto-wire them. This file is a placeholder; manual edits are
// preserved on subsequent `yongol generate` runs once preserve policy
// covers it (future Phase).
func writeAppTSX(srcDir string) error {
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
