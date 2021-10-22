package fortune_controller

import "CeobeBot-Backend/middleware/database"



type FortuneController struct {
	db database.MysqlConnection
}

func (c *FortuneController) Init(db database.MysqlConnection) error {
	c.db = db
	return c.db.Ping()
}
