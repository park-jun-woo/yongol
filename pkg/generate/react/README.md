# pkg/generate/react

## 변경이력

- 2026-04-28: 초기 작성

## 역할

React 19 + Vite + TanStack Query + shadcn-like 프론트엔드 스캐폴드를 emit 한다. OpenAPI → TypeScript 타입 (openapi-typescript) + operationId 기반 `apiClient` + manifest.theme 반영 Tailwind 설정 일체를 한 sequence 에 산출.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Generate` | `(fs *yongol.Fullstack, artifactsDir string) error` | React + Vite + TanStack Query + shadcn-like 스캐폴드 전체 방출 |

## 산출물

```
arts/frontend/
├── package.json / vite.config.ts / tsconfig.json
├── tailwind.config.js / postcss.config.js / src/index.css   ← manifest.theme 반영
├── index.html / src/main.tsx / src/App.tsx
├── src/api.ts                                                ← operationId apiClient
├── src/types/api.d.ts                                        ← openapi-typescript 출력 (or stub)
├── src/lib/utils.ts                                          ← cn() 헬퍼
└── src/components/ui/*.tsx                                   ← shadcn-like 프리미티브 10종
```

`openapi-typescript` 바이너리는 env → PATH → npx 순으로 해결, 부재 시 fallback stub `api.d.ts` 기록.

## frontend tsc 게이트 (Phase041)

`runFrontend` 의 마지막 단계로 `RunTscCheck` 가 `arts/frontend/` 에서 `tsc --noEmit` 스모크를 돌린다 — backend 의 `go build` 에 대응하는 frontend 컴파일 게이트. 타입 에러가 있으면 그 출력과 함께 generate 가 실패한다("generate 성공 = 빌드 가능" 불변식). `tsc` 해석은 프로젝트 `node_modules/.bin` → `resolveTscArgv`(env → `npx -p typescript tsc`) 순이며, 해석 불가(`node_modules`/tsc/npx 부재) 시 **경고 후 skip**(generate 실패 아님) — 강제하려면 frontend deps 를 설치한다(CI/dev). `run_tsc_check.go` / `resolve_tsc_argv.go`.
