//ff:type feature=orchestrator type=test-helper
//ff:what minimalManifest — 테스트 공용 최소 manifest.yaml 문자열 상수
package yongol

// minimalManifest is the smallest manifest.yaml accepted by the manifest
// loader. Used by SSOT detection tests that need KindConfig to be populated
// but do not exercise manifest semantics themselves.
const minimalManifest = `apiVersion: yongol/v1
kind: Project
metadata:
  name: unit-test
backend:
  lang: go
  framework: gin
  module: example.com/unit-test
frontend:
  lang: typescript
  framework: react
  bundler: vite
  name: unit-test-web
`
