package main

import (
	migrate "Backend-Server/library/database"
	"Backend-Server/user_service/config"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"log"
	"os"
)

var cfg *config.Config
var err error
var logger logr.Logger

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	cfg, err = config.Load()
	if err != nil {
		return err
	}
	logger = newLog()
	app := cli.NewApp()
	app.Name = "Gateway service"

	app.Commands = []*cli.Command{
		{
			Name:   "server",
			Usage:  "Start service server in local",
			Action: serverAction,
		},
		{
			Name:        "migrate",
			Usage:       "Migrate SQL folder",
			Subcommands: migrate.CliCommand(cfg.MigrationFolder, cfg.MySQL.String()),
		},
		{
			Name:   "config",
			Usage:  "Show service config",
			Action: configAction,
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
	return nil
}

func newLog() logr.Logger {
	zapConfig := zap.NewProductionConfig()
	zapLog, err := zapConfig.Build()
	if err != nil {
		panic(fmt.Sprintf("who watches the watchmen (%v)?", err))
	}
	return zapr.NewLogger(zapLog)
}
