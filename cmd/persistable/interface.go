package persistable

type Persistable interface {
	Save([]byte, ...any) error
	Load(...any) ([]byte, error)
}
