package persistence

type Store interface {
	GetMany(keys []string) (map[string]string, error)
	Create(key string, value string) error
}
