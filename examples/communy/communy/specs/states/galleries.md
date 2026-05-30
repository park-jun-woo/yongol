```mermaid
stateDiagram-v2
    [*] --> active
    active --> suspended: SuspendGallery
    suspended --> active: UnsuspendGallery
```
