```mermaid
stateDiagram-v2
    [*] --> published
    published --> hidden: HidePost
    published --> deleted: DeletePost
    hidden --> published: UnhidePost
```
