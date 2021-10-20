package word_controller

import "CeobeBot-Backend/middleware/database"

//type ClientStatus uint
//
//const (
//	normal ClientStatus = 1
//	banned ClientStatus = 2
//)

type WordController struct {
	db database.MysqlConnection
}

func (c *WordController) Init(db database.MysqlConnection) error {
	c.db = db
	return c.db.Ping()
}
