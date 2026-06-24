//ff:func feature=gen-react type=generator control=sequence
//ff:what main.tsx — ErrorBoundary + BrowserRouter + QueryClientProvider + index.css import

package react

import (
	"os"
	"path/filepath"
)

func writeMainTSX(srcDir string) error {
	const src = `import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import App from './App'
import './index.css'

class ErrorBoundary extends React.Component<
  { children: React.ReactNode },
  { error: Error | null }
> {
  state = { error: null as Error | null }
  static getDerivedStateFromError(error: Error) { return { error } }
  render() {
    if (this.state.error) {
      return (
        <div style={{ padding: '2rem', textAlign: 'center' }}>
          <h1>오류가 발생했습니다</h1>
          <pre style={{ color: 'red' }}>{this.state.error.message}</pre>
          <button onClick={() => window.location.reload()}>새로고침</button>
        </div>
      )
    }
    return this.props.children
  }
}

const queryClient = new QueryClient()

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ErrorBoundary>
      <QueryClientProvider client={queryClient}>
        <BrowserRouter>
          <App />
        </BrowserRouter>
      </QueryClientProvider>
    </ErrorBoundary>
  </React.StrictMode>,
)
`
	return os.WriteFile(filepath.Join(srcDir, "main.tsx"), []byte(src), 0644)
}
