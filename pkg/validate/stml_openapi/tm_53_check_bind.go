//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what tm53CheckBind — 단일 data-bind에 TM-53 세 케이스(img 불일치·비스칼라·미지원 태그)를 적용한다

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm53UnsupportedBindTags are void/media tags codegen cannot bind as text and
// that are not the supported media tag (<img>). A data-bind on one of these
// renders nothing readable.
var tm53UnsupportedBindTags = map[string]bool{
	"input": true, "br": true, "hr": true, "video": true, "audio": true,
	"iframe": true, "embed": true, "source": true, "track": true,
	"area": true, "col": true, "wbr": true, "canvas": true,
}

// tm53CheckBind applies the three TM-53 cases to a single bind of known type:
// (c) <img> bound to a non-string field, (a) object/array bound as text, and
// (b) data-bind on an unsupported void/media tag.
func tm53CheckBind(b stml.FieldBind, typ, opID, file string) []diagnostic.Diagnostic {
	if b.Tag == "img" {
		if typ == "string" {
			return nil
		}
		return tm53Diag(file, fmt.Sprintf("[TM-53] <img data-bind=%q> in operationId %q binds a %q field, but an image src must be a string URL", b.Name, opID, typ),
			"Bind the <img> to a string URL field, or use a text tag (e.g. <span>) for a non-URL value")
	}
	if typ == "object" || typ == "array" {
		return tm53Diag(file, fmt.Sprintf("[TM-53] data-bind %q in operationId %q is a %q value rendered as text — React would show [object Object] or a comma-joined string", b.Name, opID, typ),
			"For an object bind a dotted path (e.g. data-bind=\"User.Name\"); for an array use data-each")
	}
	if tm53UnsupportedBindTags[b.Tag] {
		return tm53Diag(file, fmt.Sprintf("[TM-53] data-bind %q in operationId %q is on <%s>, which codegen cannot render as bound content (only <img> src and text tags are supported)", b.Name, opID, b.Tag),
			"Move the data-bind to a text tag (e.g. <span>) or, for an image URL, to <img>")
	}
	return nil
}
