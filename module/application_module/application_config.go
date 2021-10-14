package application_module

import (
	"CeobeBot-Backend/middleware/config"
	"os"
	"runtime"
	"time"
)

type ApplicationConfig struct {
	BotName        string    // 机器人名称
	BotStartAt     time.Time // 机器人启动时间
	BotVersion     string    // 机器人版本
	SysOS          string    // 操作系统
	SysVersion     string    // 操作系统版本
	SysArch        string    // 操作系统类型
	SysName        string    // 操作系统主机名
	BindUrl        string    // 绑定地址
}

func (c *ApplicationConfig) Init(config config.Config) error {
	c.BotName = config.String("application.bot_name")
	c.BotVersion = config.String("application.bot_version")
	c.BindUrl = config.String("application.bind_url")
	c.BotStartAt = time.Now()
	c.SysArch = runtime.GOARCH
	c.SysOS = runtime.GOOS
	c.SysVersion = runtime.Version()
	var err error
	if c.SysName, err = os.Hostname(); err != nil {
		return err
	}

	return nil
}