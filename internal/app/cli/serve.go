package cli

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	application "github.com/urcop/go-fiber-template/internal/app"
	"github.com/urcop/go-fiber-template/internal/config"
	"log"
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
			log.Printf("Server started")

			ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
			defer stop()

			app := fiber.New()
			cfg := config.GetConfig()

			application.InitApplication(app)

			go func() {
				if err := app.Listen(":" + cfg.HttpServer.Port); err != nil {
					fmt.Printf("Failed to start server %v\n", err)
				}
			}()

			select {
			case <-ctx.Done():
				if err := app.Shutdown(); err != nil {
					fmt.Printf("Failed to stop server %v\n", err)
				}
			}

			log.Println("Stopping server")
			time.Sleep(time.Second * cliCmdExecFinishDelaySeconds)

			log.Println("Server stopped")
		},
	}
}
