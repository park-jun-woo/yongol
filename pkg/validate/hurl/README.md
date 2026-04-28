# pkg/validate/hurl

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

Hurl 시나리오 파일(`tests/*.hurl`) 자체 정합성 검증. deprecated `.feature` 확장자 감지, 빈 `tests/` 디렉토리 경고.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| H-1 | `h01DeprecatedFeature` | `tests/*.feature` deprecated 확장자 (ERROR) | IF-ELSE | ✓ |
| H-2 | `h02EmptyTestsDir` | `tests/` 존재하나 `.hurl` 없음 (WARNING) | IF-ELSE | ✓ |

## Defeater

없음.
