# pkg/validate/funcspec

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

Func 스펙 (Go AST 기반 `@func` 선언) 자체 정합성 검증. built-in 패키지명 충돌, 본체 미구현, 금지 import 검사.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| F-1 | `f01BuiltinOverride` | built-in 패키지명(auth/session/cache/file 등) 충돌 (ERROR) | IF-ELSE | ✓ |
| XFF-40 | `xff40FuncBodyTodo` | func 본체 미구현 (HasBody=false / TODO / 빈 본체) (ERROR) | IF-ELSE | ✓ |
| XFF-41 | `xff41FuncForbiddenImport` | I/O 경계 패키지 import 금지 (`database/sql`, `net/http`, `net/rpc`, `grpc` 등) (ERROR) | IF-ELSE | ✓ |

## Defeater

없음.
