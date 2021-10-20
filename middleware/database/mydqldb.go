package database

import (
	"CeobeBot-Backend/module/application_module"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"strings"
)

type MysqlConnection struct {
	Connection  *xorm.Engine
	isConnected bool
}

func (c *MysqlConnection) Connect(config application_module.DatabaseConfig) error {
	if c.isConnected {
		if err := c.Close(); err != nil {
			return err
		}
	}

	if config.AuthMethod == "password" {
		var err error
		connectionPath := strings.Join([]string{config.Username, ":", config.Password, "@tcp(", config.Address, ":", config.Port, ")/", config.DatabaseName}, "")
		if c.Connection, err = xorm.NewEngine("mysql", connectionPath); err != nil {
			return err
		} else {
			c.isConnected = true
			return nil
		}
	} else if config.AuthMethod == "ssh" {
		return errors.New("mysql ssh auth is not supported yet")
	} else {
		return errors.New("unknown mysql auth method")
	}
}

func (c MysqlConnection) Ping() error {
	if c.isConnected {
		return c.Connection.Ping()
	} else {
		return errors.New("database is not connected")
	}
}

func (c *MysqlConnection) Close() error {
	if c.isConnected {
		return c.Connection.Close()
	} else {
		c.isConnected = false
		return nil
	}
}
