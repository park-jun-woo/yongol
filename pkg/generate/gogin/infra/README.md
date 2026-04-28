# pkg/generate/gogin/infra

## 변경이력

- 2026-04-28: 초기 작성

## 역할

ssac 백엔드 인프라 (auth.RefreshStore / cache.CacheModel / session.SessionModel / queue.Backend) 의 어댑터 (`postgres.go`) 를 사용자 sqlc Queries 로 구현해 emit 한다. manifest 의 `when:` 평가로 활성 port 만 산출.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `EmitAll` | `(fs *yongol.Fullstack, artifactsDir, modulePath string) error` | `fs.SsacInterfaces` 순회하여 활성 패키지의 `postgres.go` 작성 |

## 산출물

```
arts/backend/internal/infra/
├── auth/postgres.go      ← ssac/pkg/auth.RefreshStore 어댑터 (활성 시)
├── cache/postgres.go     ← ssac/pkg/cache.CacheModel
├── session/postgres.go   ← ssac/pkg/session.SessionModel
└── queue/postgres.go     ← ssac/pkg/queue.Backend
```

`manifest.<feature>.backend` 가 `postgres` 가 아니거나 모든 port `when:` 이 false 면 해당 패키지는 emit 되지 않는다.
