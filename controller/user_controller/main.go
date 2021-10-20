package user_controller

import "CeobeBot-Backend/middleware/database"

type UserController struct {
	db database.MysqlConnection
}

func (c *UserController) Init(db database.MysqlConnection) error {
	c.db = db
	return c.db.Ping()
}
