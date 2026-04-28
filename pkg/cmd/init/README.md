# pkg/cmd/init

## 변경이력

- 2026-04-28: 초기 작성

## 역할

`yongol init <ProjectID> "<description>"` 명령 구현. ProjectID 검증 → 모듈 경로 추론 → 디렉토리 비움 검사 → embed.FS 템플릿을 렌더링해 manifest / OpenAPI / sqlc.yaml / authz stub / README / .gitignore 를 새 디렉토리에 작성한다. 패키지명은 `cliinit`.

## 공개 함수

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Run` | `Run(out, errOut io.Writer, opts Options) error` | skeleton 전체 생성 진입점. 검증 → 디렉토리 준비 → 템플릿 렌더 → 안내 출력 순으로 실행 |
| `ValidateProjectID` | `ValidateProjectID(id string) error` | ProjectID 가 `^[A-Za-z][A-Za-z0-9_]*$` 를 만족하는지 검사 (PascalCase / snake_case 허용) |
| `NormalizeProjectID` | `NormalizeProjectID(id string) string` | ProjectID 를 manifest.metadata.name 용 lowercase snake_case 로 정규화 |
| `DetectModule` | `DetectModule(projectID string) (module, warning string)` | `--module` 미지정 시 `GITHUB_USER` → `GH_USER` → `git config user.name` → placeholder 순으로 Go 모듈 경로를 추론 |

## 공개 구조체 / 상수

| 식별자 | 설명 |
|---|---|
| `Options` | `Run` 입력. `ProjectID`, `Description`, `Dir`, `Module`, `Force` |
| `ModulePlaceholder` | 사용자 식별 실패 시 모듈 경로에 들어가는 `"github.com/REPLACE_ME"` |

## 비고

스캐폴드용 템플릿은 `templates/` (embed.FS) 에 있다 — `manifest.yaml.tmpl`, `openapi.yaml.tmpl`, `authz.rego.tmpl`, `sqlc.yaml`, `README.md.tmpl`, `gitignore`.
