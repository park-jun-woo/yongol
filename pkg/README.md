# yongol/pkg

## 변경이력

- 2026-04-28: 79 개 패키지 README 4 원칙 정비 + pkg/ 인덱스 신설

## 역할

yongol Go 모듈의 핵심 코드. 9 SSOT 파싱 → 자체·교차 검증 → Go+Gin 백엔드 / React 프론트엔드 / DDL 마이그레이션 코드젠 + 공용 인프라 (진단 / 규칙 엔진 / preserve 해시 / 리포트).

filefunc F1 (1 파일 1 함수) 규약 준수. 79 개 패키지 / 3,145 .go 파일 (코드 2,103 + 테스트 1,042).

## 도메인 진입점

| 도메인 | 패키지 수 | 역할 | 진입 README |
|---|---:|---|---|
| **parser** | 11 | 9 SSOT 파서. 공개 시그니처 `(result, []diagnostic.Diagnostic)` 통일 | [`parser/`](parser/README.md) |
| **validate** | 32 | 자체 검증 + 교차 검증 ~150 규칙 (TOULMIN / IF-ELSE) | [`validate/`](validate/README.md) |
| **generate** | 20 | Go+Gin / React / Migration / Hurl mirror / NestJS 코드젠 | [`generate/`](generate/README.md) |
| **report** | 3 | 콘솔 / JSON / SARIF 리포트 | [`report/`](report/README.md) |
| **rule** | 2 | Toulmin 규칙 인프라 + catalog | [`rule/`](rule/README.md) |

## 인프라 패키지

| 패키지 | 역할 |
|---|---|
| [`yongol/`](yongol/README.md) | 오케스트레이션 — `Fullstack` 컨테이너 + `DetectSSOTs` / `ParseAll` |
| [`chain/`](chain/README.md) | `yongol chain <operationId>` — SSOT 노드 체인 추적 |
| [`external/`](external/README.md) | `yongol import` — 외부 OpenAPI → Go 클라이언트 코드젠 |
| [`contract/`](contract/README.md) | `//yg:checked` preserve 해시 + `//yg:preserve` + PRV-* 가드 |
| [`diagnostic/`](diagnostic/README.md) | 공용 진단 (`Diagnostic`, `Phase`, `Level`, `Loc`) |
| [`ground/`](ground/README.md) | 파싱 결과 Lookup 테이블 (검증 공용) |
| [`ssacmeta/`](ssacmeta/README.md) | ssac repo `interface.yaml` 메타 로드 |
| [`cmd/`](cmd/README.md) · [`cmd/init/`](cmd/init/README.md) | `yongol init` 스캐폴드 보조 |
| [`util/`](util/README.md) · [`util/caseconv/`](util/caseconv/README.md) | 공용 유틸 (case 변환 등) |

## 코드 탐색 우선순위

1. **codebook** — `yongol/codebook.yaml` 의 feature / type / topic 키워드 확인
2. **상위 README** — 본 표에서 도메인 진입점 → 자식 README 의 한 줄 요약
3. **소스 헤더** — `//ff:func` / `//ff:what` / `//ff:type` 어노테이션 grep
4. **filefunc chain** — `filefunc chain func <Name> --chon 1 --meta what` 로 호출 관계 추적

## 외부 매뉴얼

- 문법: [`yongol/manual-for-ai.md`](../manual-for-ai.md)
- 검증 규칙 카탈로그: [`yongol/rulebook.md`](../rulebook.md)
- preserve 컨트랙트: [`yongol/docs/PRESERVE.md`](../docs/PRESERVE.md)
- 마이그레이션: [`yongol/docs/MIGRATION.md`](../docs/MIGRATION.md)

## README 컨벤션 (2026-04-28 도입)

각 패키지 README 는 다음 4 원칙을 따른다:

1. **변경이력** 섹션을 맨 처음에
2. **역할 요약 3 줄 이하** (변경이력 다음)
3. **공개 func / struct + 규칙 ID** 한 줄 요약 표
4. **총 길이 50 줄 이하** 권장 (규칙 ID 다수 / 표 자연 확장 시 예외)
