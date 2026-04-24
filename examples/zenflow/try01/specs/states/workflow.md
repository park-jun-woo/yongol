# Workflow

```mermaid
stateDiagram-v2
    [*] --> draft
    draft --> active: ActivateWorkflow
    paused --> active: ActivateWorkflow
    active --> active: ExecuteWorkflow
```
