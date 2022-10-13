package log

import (
	"encoding/json"
	"io"
	"strings"
)

type Target interface {
	io.WriteCloser
	json.Unmarshaler
	Name() string
}

var registeredTargets = make(map[string]Target)

// RegisterTarget registers the provided Target for use with all Transport clients and servers.
func RegisterTarget(target Target) {
	if target == nil {
		panic("cannot register a nil Target")
	}
	if target.Name() == "" {
		panic("cannot register Target with empty string result for Name()")
	}
	contentSubtype := strings.ToLower(target.Name())
	registeredTargets[contentSubtype] = target
}

// GetTarget gets a registered Target by content-subtype, or nil if no Target is
// registered for the content-subtype.
//
// The content-subtype is expected to be lowercase.
func GetTarget(contentSubtype string) Target {
	contentSubtype = strings.ToLower(contentSubtype)
	return registeredTargets[contentSubtype]
}
