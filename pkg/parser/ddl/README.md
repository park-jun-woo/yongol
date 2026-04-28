# pkg/parser/ddl

## 변경이력

- 2026-04-28: 초기 작성

## 역할

PostgreSQL DDL (`*.sql`) 디렉토리를 파싱해 CREATE TABLE 메타 (컬럼 / FK / 인덱스 / PK / VARCHAR 길이 / CHECK enum / sentinel INSERT / `-- @<tag>` 힌트 주석) 를 추출한다. 구조화된 `Table` 슬라이스와 외부 `pg_query_go` AST 결과 두 가지를 함께 노출한다.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `ParseTables` | `(dir string) ([]Table, []diagnostic.Diagnostic)` | 디렉토리 내 `*.sql` 을 라인 단위로 파싱해 Table 메타 추출 |
| `ParseDir` | `(dir string) ([]*pg_query.ParseResult, []diagnostic.Diagnostic)` | 디렉토리 내 `*.sql` 을 `pg_query_go` 로 파싱해 SQL 문법 검증 |
| `ScanSentinelInserts` | `(content string) []SentinelScanResult` | 파일 본문에서 top-level INSERT 블록을 찾아 `@sentinel` 어노테이션과 함께 수집 |
| `SentinelHasOnConflictDoNothing` | `(body string) bool` | INSERT 본문에 `ON CONFLICT DO NOTHING` 포함 여부 판정 |
| `ExtractHintCommentsFromDir` | `(dir string) ([]HintComment, []diagnostic.Diagnostic)` | 디렉토리 내 `*.sql` 의 yongol 힌트 주석 (`-- @rename`, `-- @cast` 등) 수집 |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `Table` | DDL CREATE TABLE 에서 추출한 테이블 메타 (Name / Columns / ColumnOrder / ForeignKeys / Indexes / PrimaryKey / Sentinels / Archived / File / Line) |
| `Column` | 한 컬럼의 통합 메타 (Name / RawType / NotNull / NullableAnnot / HasDefault / DefaultLiteral / VarcharLen / CheckEnum / Archived / Sensitive) — Phase002 평행 맵 통합 |
| `ForeignKey` | 외래키 관계 (`Column / RefTable / RefColumn`) |
| `Index` | DDL 인덱스 (`Name / Columns / IsUnique`, USING method 보존) |
| `SentinelInsert` | `-- @sentinel` 어노테이션 INSERT 블록 verbatim 보존 |
| `SentinelScanResult` | `ScanSentinelInserts` 가 반환하는 외부 공개용 결과 |
| `HintComment` | DDL 주석 내 `-- @<tag> key=val ...` 힌트 레코드 |

## 비고

- PostgreSQL 타입 → Go 타입 매핑은 `pgTypeToGo` (private). 외부에서는 `Table.Columns` 만 사용.
- ParseTables 는 자체 라인 스캐너, ParseDir 은 `pg_query_go/v5` 외부 라이브러리에 위임.
