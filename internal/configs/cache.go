package configs

type CacheType string

type Cache struct {
	Type     CacheType `yaml:"type"`
	Address  string    `yaml:"address"`
	Username string    `yaml:"username"`
	Password string    `yaml:"password"`
}
