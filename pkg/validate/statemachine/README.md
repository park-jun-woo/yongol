# pkg/validate/statemachine

Mermaid stateDiagram 자체 정합성 검증.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = pkg/rule 공통 함수 + defeater 그래프 / `IF-ELSE` = 단일 구조·흐름·휴리스틱 검사

## 고유 함수

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---------|--------|------|----------|---------|
| ST-1 | `StateDiagramParse` | 파싱 검증 (구조 오류) | IF-ELSE | ✓ |

## pkg/rule 사용

없음.
