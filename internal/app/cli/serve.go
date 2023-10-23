package cli

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	application "github.com/urcop/go-fiber-template/internal/app"
	"github.com/urcop/go-fiber-template/internal/config"
	"github.com/urcop/go-fiber-template/pkg/logger"
	"os/signal"
	"syscall"
	"time"
)

func NewServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "serve",
		Aliases: []string{"s"},
		Short:   "Start server",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("Server started")

			ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
			defer stop()

			app := fiber.New()
			cfg := config.GetConfig()

			application.InitApplication(app)

			go func() {
				if err := app.Listen(":" + cfg.HttpServer.Port); err != nil {
					logger.Fatal("Failed to start server: ", err)
				}
			}()

			select {
			case <-ctx.Done():
				if err := app.Shutdown(); err != nil {
					logger.Fatal("Failed to stop server: ", err)
				}
			}

			logger.Info("Stopping server")
			time.Sleep(time.Second * cliCmdExecFinishDelaySeconds)

			logger.Info("Server stopped")
		},
	}
}
