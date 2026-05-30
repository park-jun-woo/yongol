//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what walkPaths/walkComponents/walkSchemas/indexPathItem/indexOperation/indexRequestBody/indexSchemaEntry/collectPropertyLines 직접 단위 검증

package openapi

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func newIdx() *LineIndex {
	return &LineIndex{
		Operations:       map[string]int{},
		RequestFields:    map[string]map[string]int{},
		ResponseFields:   map[string]map[string]int{},
		Schemas:          map[string]int{},
		SchemaProperties: map[string]map[string]int{},
		Paths:            map[string]int{},
	}
}

func docNode(t *testing.T, src string) *yaml.Node {
	t.Helper()
	var root yaml.Node
	if err := yaml.Unmarshal([]byte(src), &root); err != nil {
		t.Fatalf("yaml: %v", err)
	}
	return root.Content[0] // mapping node of document
}

func TestWalkPathsAndOperations(t *testing.T) {
	doc := docNode(t, `paths:
  /login:
    post:
      operationId: login
      requestBody:
        content:
          application/json:
            schema:
              properties:
                email:
                  type: string
      responses:
        "200":
          content:
            application/json:
              schema:
                properties:
                  token:
                    type: string
`)
	idx := newIdx()
	paths := mapValue(doc, "paths")
	walkPaths(paths, idx)

	if idx.Paths["/login"] == 0 {
		t.Errorf("path /login not indexed: %v", idx.Paths)
	}
	if idx.Operations["login"] == 0 {
		t.Errorf("operation login not indexed: %v", idx.Operations)
	}
	if idx.RequestFields["login"]["email"] == 0 {
		t.Errorf("request field email not indexed: %v", idx.RequestFields)
	}
	if idx.ResponseFields["login"]["token"] == 0 {
		t.Errorf("response field token not indexed: %v", idx.ResponseFields)
	}
}

func TestWalkPathsNonMapping(t *testing.T) {
	scalar := &yaml.Node{Kind: yaml.ScalarNode, Value: "x"}
	idx := newIdx()
	walkPaths(scalar, idx) // must not panic, no-op
	if len(idx.Paths) != 0 {
		t.Errorf("expected empty paths")
	}
}

func TestWalkComponentsAndSchemas(t *testing.T) {
	doc := docNode(t, `components:
  schemas:
    User:
      properties:
        id:
          type: integer
        name:
          type: string
`)
	idx := newIdx()
	comps := mapValue(doc, "components")
	walkComponents(comps, idx)

	if idx.Schemas["User"] == 0 {
		t.Errorf("schema User not indexed: %v", idx.Schemas)
	}
	props := idx.SchemaProperties["User"]
	if props["id"] == 0 || props["name"] == 0 {
		t.Errorf("schema properties not indexed: %v", props)
	}
}

func TestWalkComponentsNoSchemas(t *testing.T) {
	doc := docNode(t, `components:
  responses: {}
`)
	idx := newIdx()
	walkComponents(mapValue(doc, "components"), idx)
	if len(idx.Schemas) != 0 {
		t.Errorf("expected no schemas")
	}
}

func TestIndexPathItemDirect(t *testing.T) {
	doc := docNode(t, `/users/{id}:
  get:
    operationId: getUser
    responses:
      "200":
        content:
          application/json:
            schema:
              properties:
                id:
                  type: integer
`)
	// doc is the mapping; its first key/value pair is the path entry.
	pathKey := doc.Content[0]
	pathItem := doc.Content[1]
	idx := newIdx()
	indexPathItem(pathKey, pathItem, idx)
	if idx.Paths["/users/{id}"] == 0 || idx.Operations["getUser"] == 0 {
		t.Errorf("indexPathItem failed: paths=%v ops=%v", idx.Paths, idx.Operations)
	}
	if idx.ResponseFields["getUser"]["id"] == 0 {
		t.Errorf("response field not indexed: %v", idx.ResponseFields)
	}
}

func TestIndexOperationDirect(t *testing.T) {
	doc := docNode(t, `operationId: createUser
requestBody:
  content:
    application/json:
      schema:
        properties:
          email:
            type: string
responses:
  "201":
    content:
      application/json:
        schema:
          properties:
            id:
              type: integer
`)
	idx := newIdx()
	indexOperation(doc, idx)
	if idx.Operations["createUser"] == 0 {
		t.Errorf("operation not indexed: %v", idx.Operations)
	}
	if idx.RequestFields["createUser"]["email"] == 0 {
		t.Errorf("request field not indexed: %v", idx.RequestFields)
	}
	if idx.ResponseFields["createUser"]["id"] == 0 {
		t.Errorf("response field not indexed: %v", idx.ResponseFields)
	}
}

func TestIndexOperationNoOperationID(t *testing.T) {
	doc := docNode(t, `summary: no id here
`)
	idx := newIdx()
	indexOperation(doc, idx)
	if len(idx.Operations) != 0 {
		t.Errorf("expected no operations, got %v", idx.Operations)
	}
}

func TestIndexRequestBodyDirect(t *testing.T) {
	doc := docNode(t, `content:
  application/json:
    schema:
      properties:
        name:
          type: string
`)
	idx := newIdx()
	indexRequestBody(doc, "op1", idx)
	if idx.RequestFields["op1"]["name"] == 0 {
		t.Errorf("request body field not indexed: %v", idx.RequestFields)
	}
	// body without schema props -> no-op
	empty := docNode(t, `content: {}
`)
	idx2 := newIdx()
	indexRequestBody(empty, "op2", idx2)
	if len(idx2.RequestFields) != 0 {
		t.Errorf("expected no request fields")
	}
}

func TestIndexSchemaEntryDirect(t *testing.T) {
	doc := docNode(t, `User:
  properties:
    id:
      type: integer
`)
	nameKey := doc.Content[0]
	schemaNode := doc.Content[1]
	idx := newIdx()
	indexSchemaEntry(nameKey, schemaNode, idx)
	if idx.Schemas["User"] == 0 {
		t.Errorf("schema not indexed: %v", idx.Schemas)
	}
	if idx.SchemaProperties["User"]["id"] == 0 {
		t.Errorf("schema property not indexed: %v", idx.SchemaProperties)
	}
	// schema with no properties -> only schema line recorded
	doc2 := docNode(t, `Empty:
  type: object
`)
	idx2 := newIdx()
	indexSchemaEntry(doc2.Content[0], doc2.Content[1], idx2)
	if idx2.Schemas["Empty"] == 0 {
		t.Errorf("schema line should still be recorded")
	}
	if _, ok := idx2.SchemaProperties["Empty"]; ok {
		t.Errorf("no properties should mean no SchemaProperties entry")
	}
}

func TestIndexFirst2xxResponseDirect(t *testing.T) {
	doc := docNode(t, `"200":
  content:
    application/json:
      schema:
        properties:
          token:
            type: string
`)
	idx := newIdx()
	indexFirst2xxResponse(doc, "login", idx)
	if idx.ResponseFields["login"]["token"] == 0 {
		t.Errorf("first 2xx response field not indexed: %v", idx.ResponseFields)
	}
}

func TestCollectPropertyLines(t *testing.T) {
	doc := docNode(t, `properties:
  a:
    type: string
  b:
    type: integer
`)
	props := mapValue(doc, "properties")
	got := collectPropertyLines(props)
	if got["a"] == 0 || got["b"] == 0 {
		t.Errorf("collectPropertyLines = %v", got)
	}
	// nil / non-mapping returns empty map
	if g := collectPropertyLines(nil); len(g) != 0 {
		t.Errorf("nil props = %v", g)
	}
}
