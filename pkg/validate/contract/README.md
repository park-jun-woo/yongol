# pkg/validate/contract

## 변경이력

- 2026-04-28: 초기 작성

## 역할

`//yg:preserve` 로 보존된 사용자 편집 파일이 SSOT 와 표류(drift) 했는지, 또 런타임 안전 가드(panic / nil deref / 누락된 close 등)를 위반했는지 검증한다. PRV-* 규칙군은 generate 후 산출물 디렉토리(`artsDir`)의 Go 파일을 AST 파싱해 점검한다.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `IF-ELSE` = AST 패턴 매칭 / 단일 흐름 휴리스틱

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| PRV-01 | `prv_01_signature_drift` | preserved 함수 시그니처가 SSOT 기대치와 다름 | IF-ELSE | ✓ |
| PRV-02 | `prv_02_external_symbol_drift` | preserved 파일이 존재하지 않는 sqlc query / @call / DDL field 참조 | IF-ELSE | ✓ |
| PRV-10 | `prv_10_preserved_panic` | preserved 파일에 disallowed `panic(` (init / `// nolint:panic` 제외) | IF-ELSE | ✓ |
| PRV-11 | `prv_11_preserved_currentuser_assertion` | `ctx.Value("currentUser").(T)` 가 comma-ok 형태가 아님 | IF-ELSE | ✓ |
| PRV-12 | `prv_12_preserved_unmarshal_err` | `json.Unmarshal` / `yaml.Unmarshal` 의 err 무시 | IF-ELSE | ✓ |
| PRV-13 | `prv_13_preserved_scan_err` | `sql.Row(s).Scan` 의 err 무시 | IF-ELSE | ✓ |
| PRV-14 | `prv_14_preserved_slice_bounds` | `len` 가드 없이 `x[0]` 접근 | IF-ELSE | ✓ |
| PRV-15 | `prv_15_preserved_map_access` | comma-ok 가드 없이 `m[k].Field` 인라인 셀렉터 | IF-ELSE | ✓ |
| PRV-16 | `prv_16_preserved_nil_deref` | `Get*()`/`Find*()` 반환값 nil 가드 없이 필드 접근 | IF-ELSE | ✓ |
| PRV-17 | `prv_17_preserved_missing_close` | resource acquire 후 `defer Close` 누락 | IF-ELSE | ✓ |

## 주요 함수

| 함수 | 설명 |
|---|---|
| `Run(fs, artsDir)` | PRV-01 ~ PRV-17 전체 실행. 산출물 디렉토리의 preserved Go 파일을 AST 파싱 |

## pkg/rule 사용

없음. AST 패턴 매칭 기반 plain helper 로 구현.
