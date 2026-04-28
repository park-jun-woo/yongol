# pkg/cmd

## 변경이력

- 2026-04-28: 초기 작성

## 역할

CLI 서브명령의 구현 패키지를 모은 디렉토리. `cmd/yongol/` 의 cobra 엔트리는 본 디렉토리 하위 패키지의 `Run(...)` 류 진입점만 호출한다.

## 서브패키지

| 경로 | 대응 명령 | 설명 |
|---|---|---|
| `init/` | `yongol init <ProjectID> "<description>"` | 최소 SSOT skeleton (manifest / 빈 OpenAPI / sqlc.yaml / authz stub) 을 새 디렉토리에 scaffold |

## 비고

현재는 `init` 한 패키지만 산다. 다른 명령 (`validate`, `generate`, `chain`, `import`, `version`) 은 `cmd/yongol/` 에서 직접 `pkg/yongol`, `pkg/chain`, `pkg/external` 등을 호출하므로 별도 서브패키지가 없다.
