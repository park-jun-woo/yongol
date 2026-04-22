# pkg/validate/ssac_ddl

SSaC @result/@input 이 DDL 테이블/컬럼과 일치하는지, 역으로 DDL 정의가 SSaC 에서 사용되는지 확인.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule 공통 함수 + defeater 그래프 / `IF-ELSE` = 단일 구조·흐름·휴리스틱 검사 — 본 폴더는 전부 TOULMIN (defeater 체인 밀도가 높음)

## CoverageCheck

| 규칙 ID | LookupKey | 설명 | 구현 방식 | 예외 |
|---------|-----------|------|----------|------|
| XSD-55 | `SSaC.modelRef` | DDL table → SSaC 참조 여부 | TOULMIN | `IsArchived` defeater 적용 |

## 고유 함수 (TOULMIN 권장 — defeater 다수)

| 규칙 ID | 함수명 | 설명 | 구현 방식 | 조건 / 예외 |
|---------|--------|------|----------|------|
| XDS-12 | `ResultNoDDLTable` | @result 타입 DDL 테이블 없음 (WARNING) | TOULMIN | primitive Go 타입 스킵, `seq.Type == "call"` 스킵, `seq.Package != ""` 스킵, sqlc 합성 row type 스킵 |

## Defeater

- `IsPkgModel` — `pkg/<pkg>/` 내장 모델 테이블 (session/cache/file 등)
- `IsArchived` — DDL `@archived` 테이블/컬럼 (XSD-55 면제)

## internal 필수 예외

- XDS-12: `seq.Type == "call"` 스킵, `seq.Package != ""` 스킵
- XDS-12: primitive 타입 + sqlc row type 스킵 — `check_seq_result_type.go`
