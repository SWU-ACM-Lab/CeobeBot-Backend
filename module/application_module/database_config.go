package application_module

import (
	"CeobeBot-Backend/middleware/config"
	"errors"
)

type DatabaseConfig struct {
	Database           string `default:"mysql"`     // DatabaseType
	Address            string `default:"127.0.0.1"` // DatabaseAddress
	Port               string `default:"3306"`      // DatabasePort
	Username           string `default:"root"`      // DatabaseUsername
	Password           string `default:""`          // DatabasePassword
	SshKey             string `default:""`          // DatabaseSshKey
	DatabaseName       string `default:""`          // DatabaseDBName
	AuthMethod         string `default:"password"`  // DatabaseAuthMethod
	MaxOpenConnections int    `default:"20"`        // DatabaseMaxOpenConnections
	MaxIdleConnections int    `default:"10"`        // DatabaseMaxIdleConnections
}

func (c *DatabaseConfig) Init(config config.Config) error {
	c.Database = config.String("database.database")
	c.Address = config.String("database.address")
	c.Port = config.String("database.port")
	c.Username = config.String("database.username")
	c.AuthMethod = config.String("database.auth_method")
	c.DatabaseName = config.String("database.database_name")

	// load password or ssh key
	if c.AuthMethod == "password" {
		c.Password = config.String("database.password")
	} else if c.AuthMethod == "ssh" {
		c.SshKey = config.String("database.ssh_key")
	} else {
		return errors.New("unknown database auth method")
	}

	var err error
	if c.MaxOpenConnections, err = config.Int("database.max_open_connections"); err != nil {
		return err
	}
	if c.MaxIdleConnections, err = config.Int("database.max_idle_connections"); err != nil {
		return err
	}

	return nil
}
