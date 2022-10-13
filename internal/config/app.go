package config

var (
	_ AppInterface = (*App)(nil)
)

type AppInterface interface {
	Reset()
}

func GetDebug(a AppInterface) bool {
	switch b := a.(type) {
	case *App:
		return b.GetDebug()
	default:
		if c, ok := b.(interface{ GetDebug() bool }); ok {
			return c.GetDebug()
		}
	}
	return false
}

func GetName(a AppInterface) string {
	switch b := a.(type) {
	case *App:
		return b.GetName()
	default:
		if c, ok := b.(interface{ GetName() string }); ok {
			return c.GetName()
		}
	}
	return ""
}

func GetEnv(a AppInterface) string {
	switch b := a.(type) {
	case *App:
		return b.GetEnv()
	default:
		if c, ok := b.(interface{ GetEnv() string }); ok {
			return c.GetEnv()
		}
	}
	return ""
}
