package users

import (
	"CeobeBot-Backend/middleware/database"
	"github.com/gin-gonic/gin"
)

type UserInterface struct {
}

func (u UserInterface) BindApi(engine *gin.Engine, db database.MysqlConnection) error {
	// route := engine.Group("/v1/users")

	return nil
}
