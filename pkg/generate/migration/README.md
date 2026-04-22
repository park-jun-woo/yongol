# pkg/generate/migration

`yongol generate` 의 DDL 마이그레이션 자동 생성 파이프라인.

## 역할

- 직전 generate 시점 스냅샷 (`specs/db/.generated_schema.sql`) 과 현재 DDL(`specs/db/*.sql`) 을 비교
- 차이를 `Operation` 리스트로 변환 (CREATE/ALTER/DROP)
- `artifacts/db/migrations/NNNN_<desc>.sql` 로 방출
- 스냅샷 파일 갱신 (다음 generate 의 baseline)

## 파이프라인 위치

```
yongol generate
  ├─ DetectSSOTs
  ├─ ParseAll
  ├─ Validate (crosscheck + MIG-001~006)
  ├─ Migration emit          ← 본 패키지
  ├─ Backend codegen
  └─ Frontend codegen
```

## 파일 구성 (Phase 별 추가)

| Phase | 주요 파일 |
|-------|----------|
| 001 | `constants.go` — 경로/파일명 상수 |
| 002 | `ast.go`, `normalize_*.go`, `build_ast.go` |
| 003 | `operation.go`, `diff*.go`, `sort_ops.go`, `emit_sql.go`, `infer_description.go` |
| 004 | `hints.go`, `apply_rename_hints.go`, `data_migration_loader.go`, `check_safety.go` + `pkg/validate/migration/` |
| 005 | `generate.go`, `load_snapshot.go`, `write_snapshot.go`, `canonical_sql.go`, `next_sequence.go` |

## 설계 합의

- SSOT 는 DDL 그대로 (CREATE TABLE 상태 전체) — 사용자/AI 편집 대상
- 스냅샷은 직전 generate 의 정규화된 전체 스키마
- 최초 generate (스냅샷 부재) → `0001_initial.sql`
- 이후 generate → `NNNN_<desc>.sql` (incremental ALTER)
- 변경 없음 → migration 파일 생성 생략
- Git 무지 — 스냅샷은 일반 파일. 사용자가 커밋 여부 결정

## 범위 밖

- DOWN / rollback 스크립트 (v2 예정)
- PostgreSQL 외 DB 방언
- View / Function / Trigger diff
- 실제 DB 적용 (golang-migrate / flyway 등 표준 도구로 사용자가 수행)

## 참고

- 설계서: `plans/gen/migration/` (Phase001~005)
- 기존 DDL 파서: `pkg/parser/ddl/`
- 검증 규칙: `pkg/validate/migration/` (Phase004~)
- 규칙 카탈로그: `rulebook.md` → `MIG-*`
