```mermaid
stateDiagram-v2
    [*] --> draft
    draft --> active: ActivateWorkflow
    active --> paused: PauseWorkflow
    paused --> active: ActivateWorkflow
    active --> active: ExecuteWorkflow
    active --> active: ExecuteWithReport
    active --> archived: ArchiveWorkflow
```
