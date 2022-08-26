package target

type DB struct {
	DB    interface{ Write(p []byte) }
	Table string
}
