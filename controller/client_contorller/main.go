package client_contorller

import "CeobeBot-Backend/middleware/database"

type ClientStatus uint

const (
	normal ClientStatus = 1
	banned ClientStatus = 2
)

type ClientController struct {
	db database.MysqlConnection
}

func (c *ClientController) Init(db database.MysqlConnection) error {
	c.db = db
	return c.db.Ping()
}