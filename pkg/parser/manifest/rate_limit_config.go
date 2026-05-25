//ff:type feature=projectconfig type=model topic=rate-limit
//ff:what RateLimitConfig — operationId → RateLimitEntry 매핑 타입

package manifest

// RateLimitConfig is a map of operationId → RateLimitEntry.
// The map key is PascalCase operationId matching the OpenAPI operationId.
type RateLimitConfig map[string]RateLimitEntry
