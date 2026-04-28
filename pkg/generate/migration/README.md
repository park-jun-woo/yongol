# pkg/generate/migration

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

이전 generate 스냅샷(`arts/db/.latest_schema.sql`) 과 현재 DDL(`specs/db/*.sql`) 을 비교해 `Operation` 리스트로 변환, `arts/db/migrations/NNNN_<desc>.up.sql` (+ `.down.sql` no-op stub, golang-migrate 호환) 을 emit 하고 스냅샷을 갱신한다.

> 상위: [`pkg/generate/README.md`](../README.md). validate 도메인: [`pkg/validate/migration/`](../../validate/migration). 규칙 카탈로그: `rulebook.md` → `MIG-*`.

## 공개 함수

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Generate` | `(fs, specsDir, artifactsDir) (*Result, error)` | 진입점. baseline 로드 → diff → safety check → emit → snapshot 갱신 |
| `BuildASTFromDir` / `BuildASTFromSQL` | `(...) (*Schema, []Diag)` | DDL 디렉토리 / 단일 SQL → AST |
| `Diff` | `(prev, cur *Schema, hints *Hints) []Operation` | 두 schema 비교 → ALTER/CREATE/DROP 정렬된 op 리스트 |
| `EmitSQL` / `EmitMigration` / `EmitDownStub` | `(ops, opts) (string, error)` | up/down 스크립트 렌더링 |
| `LoadSnapshot` / `LoadPrevSnapshot` / `WriteSnapshot` | `(...)` | baseline 파일 입출력 |
| `CanonicalSQL` / `CanonicalType` / `NormalizeType` / `NormalizeDefault` | `(...)` | 결정론적 정규화 (diff 안정화) |
| `ApplyHints` / `ApplyRenameHints` / `ParseHints` | `(...)` | rename/cast/backfill/data-migration hint 적용 |
| `CheckSafety` / `CheckSafetyForOp` | `(ops) []SafetyIssue` | 위험 op 감지 (drop, not-null without backfill, risky cast) |
| `InferDescription` | `(ops) string` | 파일명 mnemonic |
| `NextSequence` | `(dir) int` | 다음 NNNN 일련번호 |

## 공개 타입

| 타입 | 설명 |
|---|---|
| `Schema` / `Table` / `Column` / `Index` / `ForeignKey` / `CheckConstraint` | DDL AST |
| `Operation` (`OpKind`) | CREATE_TABLE / ADD_COLUMN / ALTER_COLUMN_* / DROP_* / RENAME_* / ADD_FK / DROP_FK / ADD_CHECK / DROP_CHECK / CREATE_INDEX / DROP_INDEX / INSERT_SENTINEL |
| `Hints` | rename / cast / backfill / data-migration / allow-destructive |
| `SafetyIssue` / `SafetyLevel` | `Info`/`Warning`/`Error` |
| `Result` | emit 된 파일 + diag |
| `Mode` | `Initial` / `Incremental` / `NoOp` |

## 파이프라인 위치

```
yongol generate
  ├─ DetectSSOTs / ParseAll / Validate (MIG-001~006 포함)
  ├─ migration.Generate            ← 본 패키지
  └─ Backend / Frontend codegen
```

## 산출물

| 경우 | emit |
|---|---|
| 스냅샷 부재 (최초) | `0001_initial.up.sql` + `0001_initial.down.sql` (stub) |
| 변경 있음 | `NNNN_<desc>.up.sql` (incremental ALTER/...) + `NNNN_<desc>.down.sql` (stub) |
| 변경 없음 | 파일 생성 생략 (NoOp) |

`.down.sql` 은 no-op stub (Phase007). 롤백은 이전 `specs/` 커밋을 체크아웃 후 `yongol generate` 재실행 → 새 forward migration 으로 수행.

## 설계 합의

- SSOT = DDL 그대로 (CREATE TABLE 전체) — 사용자/AI 편집 대상.
- 스냅샷 = 직전 generate 의 정규화된 전체 스키마 (일반 파일, Git 커밋 여부는 사용자 결정).
- Diff 는 결정론적 (CanonicalSQL/CanonicalType/Sort).

## 범위 밖

PostgreSQL 외 방언 / View, Function, Trigger diff / 의미 있는 DOWN 스크립트 (v2) / 실제 DB 적용 (golang-migrate, flyway 등 표준 도구가 사용자 책임).

## 참고

설계서 `plans/gen/migration/Phase001~007`. DDL 파서 `pkg/parser/ddl/`. 검증 `pkg/validate/migration/`.
