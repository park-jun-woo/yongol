# pkg/validate/ssac_func

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

SSaC `@call`/`@eval` 과 Func 스펙 (Request/Response) 간 존재·시그니처·사용 일관성 검증 (XFS-*, XSF-*).

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 작동 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XFS-39 | `CallToFuncSpec` | `@call` → Func 구현 존재 (ERROR) | IF-ELSE | ✓ |
| XFS-42 | `CallInputsCount` | `@call` Inputs count ↔ FuncRequest 필드 수 (ERROR) | IF-ELSE | ✓ |
| XFS-43 | `CallInputFields` | `@call` Input field → FuncRequest (ERROR) | IF-ELSE | ✓ |
| XFS-44 | `CallInputType` | `@call` Input type ↔ FuncRequest field type (ERROR) | IF-ELSE | ✓ |
| XFS-45 | `CallResultMissing` | `@result` 있으나 Func Response 없음 (ERROR) | IF-ELSE | ✓ |
| XSF-46 | `CallResultIgnored` | Func Response 있으나 `@result` 없음 (WARNING) | IF-ELSE | ✓ |
| XSF-62 | `FuncSpecUsed` | Func spec → `@call`/`@eval` 사용 (WARNING, coverage) | IF-ELSE | ✓ |

## Defeater

없음. XSF-62 는 `@eval` 만으로 참조되는 Func Spec 도 사용으로 인정 (BUG-002 회귀 방지 — `test_xsf_62_eval_only_ref_test.go`).
