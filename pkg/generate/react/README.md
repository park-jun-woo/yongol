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
