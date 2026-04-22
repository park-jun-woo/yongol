//ff:func feature=gen-react type=generator control=sequence
//ff:what main.tsx — BrowserRouter + QueryClientProvider + index.css import

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

const queryClient = new QueryClient()

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </QueryClientProvider>
  </React.StrictMode>,
)
`
	return os.WriteFile(filepath.Join(srcDir, "main.tsx"), []byte(src), 0644)
}
