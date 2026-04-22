# pkg/validate/funcspec

Func 스펙 (Go AST 기반 @func 선언) 자체 정합성 검증.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule 공통 함수 + defeater 그래프 / `IF-ELSE` = 단일 구조·흐름·휴리스틱 검사

## 고유 함수 (구조 검증)

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---------|--------|------|----------|---------|
| F-1 | `BuiltinOverride` | built-in 패키지명(auth/session/cache/file) 충돌 | IF-ELSE | ✓ |

## 고유 함수 (본체/import 검증 — 구 pkg/crosscheck/func)

| 규칙 ID | 함수명 | 설명 | 구현 방식 | 조건 |
|---------|--------|------|----------|------|
| XFF-40 | `FuncBodyTodo` | func 본체 미구현 (ERROR) | IF-ELSE | `panic("TODO")` / `// TODO` / 빈 본체 / 단순 zero-return 휴리스틱 |
| XFF-41 | `FuncForbiddenImport` | func I/O 패키지 import 금지 (ERROR) | IF-ELSE | `database/sql`, `net/http`, `net/rpc`, `grpc` 계열 금지. `io`, `bufio`, `os` 허용 |

## pkg/rule 사용

없음.

## Defeater

없음.
