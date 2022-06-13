package config

type Config struct {
	user     string
	password string
}

func NewConfig(user string, password string) *Config {
	return &Config{user: user, password: password}
}

func (c *Config) User() string {
	return c.user
}

func (c *Config) Password() string {
	return c.password
}
