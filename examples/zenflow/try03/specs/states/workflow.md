# WorkflowState

```mermaid
stateDiagram-v2
    [*] --> draft
    draft --> active: ActivateWorkflow
    active --> paused: PauseWorkflow
    paused --> active: ActivateWorkflow
    active --> archived: ArchiveWorkflow
    active --> active: ExecuteWorkflow
    active --> active: ExecuteWithReport
```
