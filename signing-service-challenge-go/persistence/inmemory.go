package persistence

// TODO: in-memory persistence ...
type InMemoryPersistence map[string]interface{}

func NewInMemoryPersistence() *InMemoryPersistence {
	return &InMemoryPersistence{}
}

func (p *InMemoryPersistence) Save(key string, value interface{}) {
	(*p)[key] = value
}

func (p *InMemoryPersistence) GetAll() []interface{} {
	values := make([]interface{}, 0, len(*p))
	for _, value := range *p {
		values = append(values, value)
	}
	return values
}
