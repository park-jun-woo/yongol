# pkg/validate/ssac_ddl

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

SSaC `@result`/`@input` 이 DDL 테이블/컬럼과 일치하는지, 역으로 DDL 정의가 SSaC `@model` 에서 참조되는지 확인 (XDS-*, XSD-*).

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 그래프 / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XDS-12 | `ResultNoDDLTable` | `@result` 타입에 대응하는 DDL 테이블 없음 (WARNING) | TOULMIN | ✓ |
| XSD-55 | `DDLToModelRef` | DDL table → SSaC `@model` 참조 (ERROR, coverage) | TOULMIN | ✓ |

## Defeater

| defeater | 면제 warrant | 조건 |
|---|---|---|
| `IsPkgModel` | XSD-55 | `pkg/<pkg>/` 빌트인 모델 테이블 (session/cache/file) |
| `IsArchived` | XSD-55 | DDL `@archived` 테이블/컬럼 |
| `isSqlcRowType` | XDS-12 | `fs.SQLcQueries[*].RowType` 합성 row struct |
| primitive 타입 | XDS-12 | Go 기본 타입은 DDL 매칭 제외 |
| `seq.Type=="call"` / `seq.Package!=""` | XDS-12 | 외부 패키지 @call 결과는 DDL 비매핑 |

## internal 일치성 메모

- XDS-12: `check_seq_result_type.go` 가 primitive + sqlc row type + @call 스킵을 모두 묶어 처리.
