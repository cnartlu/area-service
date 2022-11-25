package log

import (
	"io"
	"strings"
)

type Target interface {
	Name() string
	Register(data []byte) (io.Writer, error)
}

var registeredTargets = make(map[string]Target)

func RegisterTarget(target Target) {
	if target == nil {
		panic("cannot register a nil Target")
	}
	contentSubtype := strings.ToLower(target.Name())
	if contentSubtype == "" {
		panic("cannot register Target with empty string result for Name()")
	}
	registeredTargets[contentSubtype] = target
}

func GetTarget(name string) Target {
	name = strings.ToLower(name)
	return registeredTargets[name]
}
