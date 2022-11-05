package config

import (
	kconfig "github.com/go-kratos/kratos/v2/config"
)

func New(c kconfig.Config) (*Config, error) {
	var config = Config{}
	if c != nil {
		if err := c.Scan(&config); err != nil {
			return nil, err
		}
	}
	return &config, nil
}

func GetApp(c *Config) {

}
