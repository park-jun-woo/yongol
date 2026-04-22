# pkg/validate/ssac_func

SSaC @call 과 Func 스펙(Request/Response) 간의 존재/시그니처/사용 일관성 검증.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 실 작동 또는 반례 확장 가능 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## RefExists (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XFS-39 | `Func.spec` | @call → func 구현 존재 | IF-ELSE |
| XFS-43 | `Func.request.<funcName>` | @call Input field → FuncRequest | IF-ELSE |

## CoverageCheck (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XSF-62 | `SSaC.callRef` | func spec → @call 사용 여부 | IF-ELSE |

## TypeMatch (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XFS-44 | `Func.request.<funcName>` | @call Input type ↔ Request field type | IF-ELSE |

## SchemaMatch (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XFS-42 | `Func.request.<funcName>` | @call Inputs count → FuncRequest fields | IF-ELSE |

## 고유 함수 (IF-ELSE)

| 규칙 ID | 함수명 | 설명 | 구현 방식 |
|---------|--------|------|----------|
| XFS-45 | `CallResultMissing` | @result 있지만 func Response 없음 (ERROR) | IF-ELSE |
| XSF-46 | `CallResultIgnored` | @result 없지만 func Response 있음 (WARNING) | IF-ELSE |

## Defeater

없음.
