# pkg/rule/catalog

## 변경이력

- 2026-04-28: 초기 작성

## 역할

저장소 루트의 `rulebook.md` 를 `//go:embed` 로 내장하고 H2 섹션 + `| Rule ID | Level | Description | Source |` 테이블을 파싱해 Rule ID 기준 lookup table 로 노출한다. SARIF emitter 등이 카탈로그 전체 규칙 메타에 접근할 때 사용. `## Deprecated` 섹션은 자동으로 제외.

> 상위 문서: [`pkg/rule/README.md`](../README.md)

## 공개 함수 / 구조체

| 식별자 | 종류 | 설명 |
|---|---|---|
| `Load() (*Catalog, error)` | func | 내장된 `rulebook.md` 를 파싱해 `*Catalog` 반환 (`sync.Once` 캐시) |
| `MustLoad() *Catalog` | func | `Load` 래퍼 (실패 시 `log.Fatal`) |
| `Source() []byte` | func | 내장 rulebook 원본 바이트 노출 (테스트 재파싱용) |
| `Parse(io.Reader) ([]RuleMeta, error)` | func | 임의 reader 에서 rulebook 파싱 |
| `NewCatalog([]RuleMeta) *Catalog` | func | RuleMeta 슬라이스 → byID 인덱스 포함 Catalog |
| `Catalog.Lookup(id) (RuleMeta, bool)` | method | Rule ID 로 메타 조회 |
| `Catalog.Index(id) int` | method | Rule ID 의 0-based index (없으면 -1) |
| `Catalog.Len() int` | method | 카탈로그 내 규칙 수 |
| `Catalog` | type | rulebook 파싱 결과 lookup table |
| `RuleMeta` | type | 한 행 메타 (Rule ID / Level / Description / Source / Section) |

## 보조 헬퍼

`parse_rulebook` 의 상태 머신 (`rulebookParseState` + `feedLine` / `feedTableRow` / `handleSectionHeading` / `appendDataRow` / `resetTable`), 테이블 인식 (`isRuleTableHeader`, `isTableSeparator`, `splitRow`), GFM slug 변환 (`sectionAnchor`, `writeSectionAnchorRune`).
