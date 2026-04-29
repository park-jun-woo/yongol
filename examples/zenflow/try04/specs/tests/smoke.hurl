POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "admin@zenflow.test",
  "password": "p@ssw0rd1"
}
HTTP 200
[Captures]
token: jsonpath "$.access_token"
[Asserts]
jsonpath "$.access_token" isString

GET {{host}}/workflows
Authorization: Bearer {{token}}
HTTP 200
[Asserts]
jsonpath "$.items" isCollection
jsonpath "$.total" isInteger
