# WorkflowState

```mermaid
stateDiagram-v2
    [*] --> draft
    draft --> active: ActivateWorkflow
    active --> active: ExecuteWorkflow
```
