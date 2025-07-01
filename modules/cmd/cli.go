package cmd

import (
	"os"
	"os/signal"
	"syscall"

	polarysdb "github.com/polarysfoundation/polarys_db"
	"github.com/polarysfoundation/polarysdb-cli/modules/logger"
)

type CLI struct {
	db       *polarysdb.Database
	commands []Command
	logger   *logger.Logger
	version  string
}

func NewCLI() *CLI {
	return &CLI{
		db:       nil,
		commands: Commands,
		logger:   logger.NewLogger(),
		version:  "v1.0.0",
	}
}

func (c *CLI) Run() {
	c.logger.Init() // Initialize the logger

	c.logger.Info("PolarysDB CLI Version:", c.version)
	c.logger.Info("Available commands:")
	for _, cmd := range c.commands {
		c.logger.Info(cmd.Name+":", cmd.Description)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

}
