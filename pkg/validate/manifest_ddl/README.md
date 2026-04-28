# pkg/validate/manifest_ddl

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

`manifest.yaml` `backend.auth` (user_table + claim mappings) ↔ DDL-parsed user table 교차 검증. JWT claims 가 가리키는 컬럼이 실제 user_table 에 존재하고 Go 타입까지 일치하는지 확인.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XDN-01 | `xdn01UserTableRequired` | `auth.type != "none"` 일 때 `backend.auth.user_table` 필수 (ERROR) | IF-ELSE | ✓ |
| XDN-02 | `xdn02UserTableExists` | `backend.auth.user_table` 가 DDL 파싱된 테이블에 실재 (ERROR) | IF-ELSE | ✓ |
| XDN-03 | `xdn03ClaimColumnExists` | 각 `backend.auth.claims.<Field>: <col>` 매핑의 컬럼이 user_table 에 존재 (ERROR) | IF-ELSE | ✓ |
| XDN-04 | `xdn04ClaimColumnType` | claim Go 타입 (`int64`/`string`/`bool`, 기본 `string`) ↔ user_table 컬럼 DDL Go 타입 일치 (ERROR) | IF-ELSE | ✓ |

## Defeater

없음.

## 단계적 falsification

- XDN-02 는 XDN-01 통과 (user_table 비어있지 않음) 후에만 점화.
- XDN-03 / XDN-04 는 XDN-02 통과 (실재 DDL 테이블) 후에만 점화.
- XDN-04 는 컬럼 부재 시 XDN-03 에 양보 — 작성자 의도당 진단 1건.
