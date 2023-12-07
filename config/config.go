package config

import "fmt"

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Driver   string
}

type ApiConfig struct {
	ApiHost string
}

type Config struct {
	DbConfig
	ApiConfig
}

func (c *Config) ConfigConfiguration() error {
	c.DbConfig = DbConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "stanners2020",
		Name:     "task_management_db",
		Driver:   "postgres",
	}

	c.ApiConfig = ApiConfig{ApiHost: "8080"}

	if c.Host == "" || c.Port == "" || c.User == "" || c.Name == "" || c.Driver == "" || c.ApiHost == "" {
		return fmt.Errorf("missing required environment")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cfg.ConfigConfiguration(); err != nil {
		return nil, err
	}
	return cfg, nil
}
