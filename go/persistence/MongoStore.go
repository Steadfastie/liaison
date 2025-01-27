package persistence

type MongoStore struct {
}

func NewMongoStore() *MongoStore {
	return &MongoStore{}
}

func (m *MongoStore) GetMany(keys []string) (map[string]string, error) {
	return nil, nil
}

func (m *MongoStore) Create(key string, value string) error {
	return nil
}
