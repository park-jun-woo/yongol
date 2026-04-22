# pkg/validate/manifest

`manifest.yaml` 로드 및 최소 스키마 검증.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule 공통 함수 + defeater 그래프 / `IF-ELSE` = 단일 구조·흐름·휴리스틱 검사

## 고유 함수

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---------|--------|------|----------|---------|
| C-1 | `ManifestLoad` | `manifest.yaml` 로드 실패 | IF-ELSE | ✓ |

## pkg/rule 사용

없음. 파싱 단계 단일 검사.
