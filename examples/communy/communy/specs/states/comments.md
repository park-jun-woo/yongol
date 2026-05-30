```mermaid
stateDiagram-v2
    [*] --> published
    published --> hidden: HideComment
    published --> deleted: DeleteComment
```
