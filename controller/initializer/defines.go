package initializer

import (
	"CeobeBot-Backend/middleware/database"
	"github.com/gin-gonic/gin"
)

var (
	DatabaseService database.MysqlConnection
	ApiEngine *gin.Engine
	BasePath string
)
