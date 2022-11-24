package store

type Store struct {
	config *Config
}

func New(c *Config) *Store {
	return &Store{
		config: c,
	}
}

func (s *Store) Close() error {
	return nil
}
func (s *Store) Open() {

}
