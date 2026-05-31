package features

const twoFeats = `features:
  - op: CreateTask
    path: POST /tasks
    desc: Create a task
  - op: DeleteTask
    path: DELETE /tasks/{id}
    desc: Delete a task
`
