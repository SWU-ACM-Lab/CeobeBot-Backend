package api

import (
	"CeobeBot-Backend/middleware/database"
	"github.com/gin-gonic/gin"
)

type ServiceInterface interface {
	BindApi(engine *gin.Engine, db database.MysqlConnection) error
}
