package agent

const validOpenAPIDoc = `openapi: "3.0.0"
info:
  title: T
  version: "1.0.0"
paths:
  /ping:
    get:
      operationId: Ping
      responses:
        '200':
          description: ok
`
