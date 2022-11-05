package filesystem

import (
	"context"
	"io"
	"strings"
)

// Storage defines the interface filesystem.  Note
// that implementations of this interface must be thread safe; a Storage's
// methods can be called from concurrent goroutines.
type Storage interface {
	Name() string
	Upload(ctx context.Context, key string, data io.Reader, handlers ...HookFunc) error
	Url(key string, handlers ...HookFunc) string
}

var registeredStorages = make(map[string]Storage)

// RegisterStorage registers the provided Storage for use with all Transport clients and servers.
func RegisterStorage(storage Storage) {
	if storage == nil {
		panic("cannot register a nil Storage")
	}
	if storage.Name() == "" {
		panic("cannot register Storage with empty string result for Name()")
	}
	contentSubtype := strings.ToLower(storage.Name())
	registeredStorages[contentSubtype] = storage
}

// GetStorage gets a registered Storage by content-subtype, or nil if no Storage is
// registered for the content-subtype.
//
// The content-subtype is expected to be lowercase.
func GetStorage(contentSubtype string) Storage {
	return registeredStorages[contentSubtype]
}
