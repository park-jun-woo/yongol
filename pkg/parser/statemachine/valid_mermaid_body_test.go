package statemachine

// validMermaidBody is the shared mermaid fixture for ParseDir tests. Extracted
// to a const-only file (filefunc exemption) so per-test files can import it
// without each annotating a synthetic helper func.
const validMermaidBody = "# Gig\n\n```mermaid\nstateDiagram-v2\n    [*] --> draft\n    draft --> open: PublishGig\n    open --> closed: CloseGig\n```\n"
