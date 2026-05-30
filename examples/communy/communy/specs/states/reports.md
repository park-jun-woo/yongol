```mermaid
stateDiagram-v2
    [*] --> pending
    pending --> resolved: ResolveReport
    pending --> dismissed: DismissReport
```
