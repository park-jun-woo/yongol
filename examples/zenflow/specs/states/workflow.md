# Workflow state machine

```mermaid
stateDiagram-v2
    [*] --> draft
    draft --> active: ActivateWorkflow
    paused --> active: ActivateWorkflow
    active --> paused: PauseWorkflow
    active --> archived: ArchiveWorkflow
    active --> active: ExecuteWorkflow
```
