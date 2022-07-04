package cache

type RedisConf struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Timeout  int    `yaml:"timeout"`
}
