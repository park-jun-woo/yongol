# pkg/validate/hurl

Hurl 시나리오 파일 자체 정합성 검증 (deprecated 확장자 감지 등).

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule 공통 함수 + defeater 그래프 / `IF-ELSE` = 단일 구조·흐름·휴리스틱 검사

## 고유 함수

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---------|--------|------|----------|---------|
| H-1 | `DeprecatedFeature` | `.feature` 파일 존재 (deprecated) | IF-ELSE | ✓ |

## pkg/rule 사용

없음.
