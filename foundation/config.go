package foundation

import (
	"flag"
)

// LoadConfig loads config from flags provided to the process.
func LoadConfig() (*Config, error) {
	config := &Config{}

	flag.StringVar(&config.Web.Host, "host", "localhost", "Server host.")
	flag.StringVar(&config.Web.Port, "port", "8080", "Server port.")

	flag.Parse()

	err := config.validate()
	if err != nil {
		return nil, err
	}

	return config, nil
}

type Config struct {
	Web WebConfig
}

func (c *Config) validate() error {
	if err := c.Web.validate(); err != nil {
		return err
	}
	return nil
}

type WebConfig struct {
	Host string
	Port string
}

func (c *WebConfig) validate() error {
	return nil
}
