package initializers

import (
	"github.com/gobuffalo/envy"
	"log"
)

func InitEnv() {
	if err := envy.Load(); err != nil {
		log.Printf("can not load environment")
		envy.Reload()
	}
}
