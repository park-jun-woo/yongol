//ff:func feature=gen-react type=test control=sequence
//ff:what writeMainTSX main.tsx 생성 내용·에러경로 검증
package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteMainTSX(t *testing.T) {
	dir := t.TempDir()
	if err := writeMainTSX(dir); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "main.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import { BrowserRouter } from 'react-router-dom'")
	assertContains(t, content, "import { QueryClient, QueryClientProvider } from '@tanstack/react-query'")
	assertContains(t, content, "import './index.css'")
	assertContains(t, content, "const queryClient = new QueryClient()")
	assertContains(t, content, "document.getElementById('root')")

	// Phase045: ErrorBoundary wrapping
	assertContains(t, content, "class ErrorBoundary")
	assertContains(t, content, "<ErrorBoundary>")
	assertContains(t, content, "getDerivedStateFromError")
}
