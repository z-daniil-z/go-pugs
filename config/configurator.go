package config

type Config struct {
	IP   string
	Port string
}

func GetConfig() (*Config, error) {
	ret := &Config{}
}

func (conf *Config) DebugConfig() error {

}

func (conf *Config)
