# pkg/generate/gogin/types

## 변경이력

- 2026-04-29: 신설 (Phase001) — `MapPGType(ddl.Column) → GoTypeBinding` 단일 출처. 17 PG family 전수 커버

## 역할

DDL `Column` 의 raw PostgreSQL 타입 토큰을 Go 측 표현 (sqlc go_type / api struct field / convert / insert / response 템플릿) 로 매핑하는 단일 출처. `pkg/generate/gogin/ssac` 의 3 emit 사이트와 `pkg/validate/query` 의 Q-12 + per-type Q-NN 룰이 본 모듈을 공유한다.

> 상위: [`pkg/generate/gogin/README.md`](../README.md).

## 공개 API

| 식별자 | 시그니처 | 설명 |
|---|---|---|
| `MapPGType` | `(col ddl.Column) GoTypeBinding` | 디스패처 — RawType + nullability + CheckEnum 으로 8 Kind 중 하나 반환 |
| `Expand` | `(tmpl, row, field, varName string) string` | 템플릿 placeholder 치환 (`{row}` / `{field}` / `{var}`, `{{`/`}}` escape) |
| `GoTypeBinding` | struct | 매핑 결과 (sqlc / api / convert / insert / response / DefaultLiteral / Kind / Supported) |
| `BindingKind` | int enum | Native / Pointer / Pgtype / JSONB / Bytea / Array / Enum / Unsupported |

## 매트릭스 (17 family × 3 nullability)

| family | RawType 토큰 | NOT NULL Kind / Go | NULLABLE Kind / Go | NeedsOverride |
|---|---|---|---|---|
| Integer | BIGINT, BIGSERIAL, INTEGER, INT, INT2, INT4, INT8, SMALLINT, SERIAL, SMALLSERIAL | Native / int64 | Pointer / *int64 | — |
| Float | REAL, FLOAT, FLOAT4, FLOAT8, DOUBLE | Native / float64 | Pointer / *float64 | — |
| String | VARCHAR(N), TEXT, CHAR, BPCHAR | Native / string | Pointer / *string | — |
| Boolean | BOOLEAN, BOOL | Native / bool | Pointer / *bool | — |
| UUID | UUID | Pgtype / pgtype.UUID | Pgtype / pgtype.UUID | ✓ (Q-12) |
| Numeric | NUMERIC(p,s), DECIMAL(p,s) | Pgtype / pgtype.Numeric | Pgtype / pgtype.Numeric | ✓ |
| Timestamp | TIMESTAMPTZ, TIMESTAMP, DATE | Pgtype / pgtype.Timestamptz/Timestamp/Date | (동상) | ✓ |
| Inet | INET, CIDR | Pgtype / pgtype.Inet | (동상) | ✓ |
| Interval | INTERVAL | Pgtype / pgtype.Interval | (동상) | ✓ |
| JSONB | JSONB, JSON | JSONB / map[string]interface{} | JSONB / *map[string]interface{} | — |
| Bytea | BYTEA | Bytea / []byte | Bytea / []byte (slice 자체 nullable) | — |
| Array | TEXT[], BIGINT[], … | Array / []T | Array / []T (slice 자체 nullable) | — |
| Enum | VARCHAR(N) + CHECK IN (…) | Enum / string + apiCast | Enum / *string + apiCast | — |
| Unsupported | DOUBLE PRECISION, TIMESTAMP WITH TIME ZONE, CREATE TYPE 사용자 ENUM | Unsupported | Unsupported | — (D-11 거절) |

## 사용 가이드 (4 원칙)

1. **단일 출처** — Row → Model 변환 / INSERT params / 응답 emit / validate Q-NN 모두 본 모듈 호출. 이외에서 raw 타입 → Go 타입 분기 작성 금지.
2. **string template** — convert / insert / response 표현은 `Expand` 템플릿. `{row}` (row 변수), `{field}` (PascalCase 컬럼명), `{var}` (소스 변수). 누락 placeholder 는 빈 문자열.
3. **선언적 거절** — `Supported=false` 컬럼은 D-11 이 generate 전에 ERROR. silent fallback 0.
4. **Kind 분기 최소** — emit 사이트는 Kind 로 import 결정만. 표현 자체는 항상 템플릿 경유.

## 파일 분할 (filefunc F1)

- `types.go` — `MapPGType` 디스패처
- `binding_type.go` / `binding_kind.go` / `expand.go` / `parse_raw_type.go`
- `native_*.go` (4) + `pointer.go`
- `pgtype_*.go` (5: uuid / numeric / timestamp / inet / interval)
- `jsonb.go` / `bytea.go` / `array.go` / `enum.go` / `unsupported.go`
- 각 family `*_test.go` 또는 통합 `test_map_pg_type_test.go` (table-driven)

## 의존성

- `github.com/park-jun-woo/yongol/pkg/parser/ddl` (Column)
- 외부 추가 의존 없음

## 관련 검증 룰

- `validate/query/Q-12` — UUID 컬럼이 있으면 sqlc.yaml 에 pgtype.UUID override 강제
- `validate/query/Q-NN` (per-type, NeedsOverride=true 마다 1 룰) — 같은 정책을 다른 타입으로 일반화
- `validate/ddl/D-11` — Supported=false 컬럼 거절

## 미루기 (Out of Scope)

- 다중 토큰 PG 타입 실제 매핑 (`DOUBLE PRECISION`, `TIMESTAMP WITH TIME ZONE`) — parser 측 보강 별 Phase
- `CREATE TYPE` 사용자 정의 ENUM — 별 도메인
- `pkg/generate/react/`, `pkg/generate/nestjs/` 에서 본 모듈 소비 — 별 Phase
