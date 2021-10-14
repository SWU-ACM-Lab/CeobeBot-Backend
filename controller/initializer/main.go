package initializer

import (
	"CeobeBot-Backend/api"
	"CeobeBot-Backend/api/v1/clients"
	utils "CeobeBot-Backend/middleware/config"
	"CeobeBot-Backend/module"
	"CeobeBot-Backend/module/orm_module/orm_client"
	"CeobeBot-Backend/module/orm_module/orm_fortune"
	"CeobeBot-Backend/module/orm_module/orm_problem"
	"CeobeBot-Backend/module/orm_module/orm_quote"
	"CeobeBot-Backend/module/orm_module/orm_user"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
	"path/filepath"
	"xorm.io/core"
)

func Init() error {
	// 获取基础路径
	if path, err := exec.LookPath(os.Args[0]); err != nil {
		return err
	} else {
		path = filepath.Dir(path)
		module.BasePath = path
	}

	// 初始化配置文件
	if config, err := utils.NewFileConf(module.BasePath + "/config.ini"); err != nil {
		return err
	} else {
		if err := module.ApplicationConfig.Init(*config); err != nil {
			return err
		}
		if err := module.DatabaseConfig.Init(*config); err != nil {
			return err
		}
		if err := module.RedisConfig.Init(*config); err != nil {
			return err
		}
		//if err := module.RobotConfig.Init(*config); err != nil {
		//	return err
		//}
	}

	// 建立数据库链接
	if err := module.DatabaseService.Connect(module.DatabaseConfig); err != nil {
		return err
	}

	// 初始化Api引擎
	module.ApiEngine = gin.Default()

	return nil
}

// SyncDatabase 同步ORM映射
func SyncDatabase() error {
	if err := module.DatabaseService.Ping(); err != nil {
		return err
	} else {
		module.DatabaseService.Connection.SetMapper(core.SameMapper{})
		if err := module.DatabaseService.Connection.Sync2(
			new(orm_client.Clients),
			new(orm_module.Fortunes),
			new(orm_problem.Problems),
			new(orm_problem.ProblemSolvedRecords),
			new(orm_problem.Subjects),
			new(orm_quote.Quotes),
			new(orm_user.Users),
			new(orm_user.UserStatisticData),
			); err != nil {
			return err
		}
	}
	return nil
}

func BindApiEngine() error {
	// 声明API接口
	var ClientApi api.ServiceInterface = clients.ClientInterface{}

	// 绑定API路由
	if err := ClientApi.BindApi(module.ApiEngine, module.DatabaseService); err != nil {
		return err
	}

	return nil
}

func StartApiEngine() error {
	return module.ApiEngine.Run(module.ApplicationConfig.BindUrl)
}