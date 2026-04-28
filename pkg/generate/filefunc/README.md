# pkg/generate/filefunc

## 변경이력

- 2026-04-28: 초기 작성

## 역할

생성된 백엔드 아티팩트 디렉토리에 `codebook.yaml` (filefunc feature/type/topic/ssot/pattern 색인) 을 자동 생성한다. SSOT 메타 + `arts/backend/internal/` 디렉토리 + 고정 baseline 을 병합.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Generate` | `(fs *yongol.Fullstack, artifactsDir string) error` | `arts/backend/codebook.yaml` 작성. 모든 백엔드 산출물 emit 이후 호출 |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `Codebook` | codebook.yaml 최상위 구조 (Required + Optional) |
| `Required` | required 섹션 (feature + type 필수 키) |
| `Optional` | optional 섹션 (topic / ssot / pattern 선택 키) |

## 산출물

```
arts/backend/codebook.yaml
  ├─ required.feature   ← SSOT 후보 + internal/ 디렉토리 + 고정 baseline 병합
  ├─ required.type      ← 고정 (generator/util/test/model/...)
  └─ optional.{topic,ssot,pattern}  ← 고정 baseline 맵
```
