//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what otelServiceName — service.name 값 결정 (tracing.service_name → metadata.name → "service")

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// otelServiceName resolves the `service.name` resource attribute: manifest
// tracing.service_name when set, else metadata.name (the project name), else
// "service" as a last-resort fallback so the exporter never receives an
// empty identifier.
func otelServiceName(fs *yongol.Fullstack) string {
	if fs == nil || fs.Manifest == nil {
		return "service"
	}
	obs := fs.Manifest.Backend.Observability
	if obs != nil && obs.Tracing != nil && obs.Tracing.ServiceName != "" {
		return obs.Tracing.ServiceName
	}
	if fs.Manifest.Metadata.Name != "" {
		return fs.Manifest.Metadata.Name
	}
	return "service"
}
