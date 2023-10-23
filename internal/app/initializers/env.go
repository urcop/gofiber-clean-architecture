package initializers

import (
	"github.com/gobuffalo/envy"
	"github.com/urcop/go-fiber-template/pkg/logger"
)

func InitEnv() {
	if err := envy.Load(); err != nil {
		logger.Info("Can not load environment: ", err)
		envy.Reload()
	}
}
