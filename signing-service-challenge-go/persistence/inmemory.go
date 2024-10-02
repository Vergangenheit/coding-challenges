package persistence

// TODO: in-memory persistence ...
type InMemoryPersistence map[string]interface{}

func NewInMemoryPersistence() *InMemoryPersistence {
	return &InMemoryPersistence{}
}
