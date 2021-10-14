package application_module

import (
	"CeobeBot-Backend/middleware/config"
)

type RedisConfig struct {
	Address            string `default:"127.0.0.1"` // RedisAddress
	Password           string `default:""`          // RedisPassword
	Port               string `default:"6379"`      // RedisPort
	Database           int    `default:"0"`         // RedisDatabaseIndex
	MaxOpenConnections int    `default:"20"`        // RedisMaxOpenConnections
	MaxIdleConnections int    `default:"10"`        // RedisMaxOpenConnections
}

func (c *RedisConfig) Init(config config.Config) error {
	c.Address = config.String("redis.address")
	c.Password = config.String("redis.password")
	c.Port = config.String("redis.port")

	var err error
	if c.Database, err = config.Int("redis.database"); err != nil {
		return err
	}

	if c.MaxOpenConnections, err = config.Int("database.max_open_connections"); err != nil {
		return err
	}
	if c.MaxIdleConnections, err = config.Int("database.max_idle_connections"); err != nil {
		return err
	}

	return nil
}
