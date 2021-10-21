package statistic_controller

import "CeobeBot-Backend/middleware/database"

type StatisticController struct {
	db database.MysqlConnection
}

func (c *StatisticController) Init(db database.MysqlConnection) error {
	c.db = db
	return c.db.Ping()
}
