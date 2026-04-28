# pkg/generate/gogin/sqlcpost

## 변경이력

- 2026-04-28: 초기 작성

## 역할

sqlc 외부 실행 산출물 (row 구조체) 에 `LogValue()` 메소드를 후처리로 주입한다. DDL 의 `-- @sensitive` 마커 컬럼 1개 이상인 테이블에 대해 `<table>_log.go` 를 emit 하며, 민감 필드는 `[REDACTED]`, 그 외는 `slog.Any` 로 출력.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Generate` | `(fs *yongol.Fullstack, artifactsDir string) error` | DDL 테이블별 `<table>_log.go` (LogValue 메소드) emit |
| `StructNameForTable` | `(table string) string` | DDL 테이블명 → sqlc 생성 Go struct 이름 (단수화 + PascalCase) |

## 산출물

```
arts/backend/internal/db/
└── <table>_log.go    ← LogValue() 메소드 — slog.LogValuer 인터페이스 자동 매칭
```

sqlc 가 emit 한 `models.go` / `*.sql.go` 와 같은 패키지에 형제 파일로 들어가 메소드가 row 타입에 자동 부착된다.
