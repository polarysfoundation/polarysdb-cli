package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/chzyer/readline"
	polarysdb "github.com/polarysfoundation/polarysdb"
	"github.com/polarysfoundation/polarysdb-cli/modules/logger"
	"github.com/polarysfoundation/polarysdb/modules/common"
)

type CLI struct {
	db       *polarysdb.Database
	commands []Command
	logger   *logger.Logger
	version  string
}

func NewCLI() *CLI {

	log := logger.NewLogger(logger.Config{
		LogFilePath: "app.log",
		MinLevel:    logger.LevelInfo,
		ToConsole:   true,
		ToFile:      false,
	})
	return &CLI{
		db:       nil,
		commands: Commands,
		logger:   log,
		version:  "v1.0.0",
	}
}

func (c *CLI) Run() {
	c.logger.Info("PolarysDB CLI Version:", c.version)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		rl, err := readline.New("> ") // Create a CLI prompt
		if err != nil {
			c.logger.Fatal(err)
		}
		defer rl.Close()

		for {
			line, err := rl.Readline()
			if err != nil { // Handle CTRL+D (EOF)
				break
			}

			if line == "" {
				continue
			}

			args := strings.Fields(line)
			// Process command
			if err := c.handleCommands(args); err != nil {
				c.logger.Warn("Error:", err)
				continue
			}
		}
	}()

	<-stop
	c.logger.Info("Shutting down bot gracefully...")
}

func (c *CLI) handleCommands(args []string) error {
	if len(args) == 0 {
		return nil
	}

	commandName := args[0]
	switch commandName {
	case "init":
		if len(args) != 3 {
			return fmt.Errorf("usage: init <key> <path>")
		}
		keyStr := args[1]
		path := args[2]

		key := common.StringToKey(keyStr)
		if keyStr == `""` {
			c.logger.Warn("No key provided. Initializing with an empty key.")
			key = common.BytesToKey([]byte(""))
		}

		c.logger.Info("Initializing database with key:", key, "at path:", path)
		db, err := polarysdb.Init(key, path)
		if err != nil {
			return err
		}
		c.db = db
		c.logger.Info("Database initialized successfully.")
	case "export":
		if len(args) != 3 {
			return fmt.Errorf("usage: export <key> <path>")
		}
		keyStr := args[1]
		path := args[2]

		key := common.StringToKey(keyStr)
		if keyStr == `""` {
			c.logger.Warn("No key provided. Parsing an empty key.")
			key = common.BytesToKey([]byte(""))
		}

		c.logger.Info("Exporting database with key:", key, "to path:", path)
		if c.db == nil {
			return fmt.Errorf("database not initialized. Please run 'init' first")
		}
		err := c.db.Export(key, path)
		if err != nil {
			return err
		}
		c.logger.Info("Database exported successfully.")
	case "import":
		if len(args) != 3 {
			return fmt.Errorf("usage: import <key> <path>")
		}
		keyStr := args[1]
		path := args[2]

		key := common.StringToKey(keyStr)
		if keyStr == `""` {
			c.logger.Warn("No key provided. Parsing an empty key.")
			key = common.BytesToKey([]byte(""))
		}

		c.logger.Info("Importing database with key:", key, "from path:", path)
		if c.db == nil {
			return fmt.Errorf("database not initialized. Please run 'init' first")
		}
		err := c.db.Import(key, path)
		if err != nil {
			return err
		}
		c.logger.Info("Database imported successfully.")
	case "export-encrypted":
		if len(args) != 3 {
			return fmt.Errorf("usage: export-encrypted <key> <path>")
		}
		keyStr := args[1]
		path := args[2]

		key := common.StringToKey(keyStr)
		if keyStr == `""` {
			c.logger.Warn("No key provided. Parsing an empty key.")
			key = common.BytesToKey([]byte(""))
		}

		c.logger.Info("Exporting encrypted database with key:", key, "to path:", path)
		if c.db == nil {
			return fmt.Errorf("database not initialized. Please run 'init' first")
		}
		err := c.db.ExportEncrypted(key, path)
		if err != nil {
			return err
		}
		c.logger.Info("Encrypted database exported successfully.")
	case "import-encrypted":
		if len(args) != 3 {
			return fmt.Errorf("usage: import-encrypted <key> <path>")
		}
		keyStr := args[1]
		path := args[2]

		key := common.StringToKey(keyStr)
		if keyStr == `""` {
			c.logger.Warn("No key provided. Parsing an empty key.")
			key = common.BytesToKey([]byte(""))
		}

		c.logger.Info("Importing encrypted database with key:", key, "from path:", path)
		if c.db == nil {
			return fmt.Errorf("database not initialized. Please run 'init' first")
		}
		err := c.db.ImportEncrypted(key, path)
		if err != nil {
			return err
		}
		c.logger.Info("Encrypted database imported successfully.")
	case "new-key":
		c.logger.Info("Generating a new key...")
		key := polarysdb.GenerateKey()
		c.logger.Info("New key generated:", key.String())
	case "key-from":
		if len(args) != 2 {
			return fmt.Errorf("usage: key-from <string>")
		}
		keyStr := args[1]
		keyTemp := common.StringToKey(keyStr)
		keyByte := keyTemp.Bytes()
		key := polarysdb.GenerateKeyFromBytes(keyByte)
		c.logger.Info("Key generated from string:", key.String())
	case "change-key":
		if len(args) != 3 {
			return fmt.Errorf("usage: change-key <old-key> <new-key>")
		}
		oldKeyStr := args[1]
		oldKey := common.StringToKey(oldKeyStr)
		if oldKeyStr == `""` {
			c.logger.Warn("No key provided. Parsing an empty key.")
			oldKey = common.BytesToKey([]byte(""))
		}
		newKeyStr := args[2]
		newKey := common.StringToKey(newKeyStr)
		if newKeyStr == `""` {
			c.logger.Warn("No key provided. Parsing an empty key.")
			newKey = common.BytesToKey([]byte(""))
		}

		c.logger.Info("Changing database key to:", newKey)
		if c.db == nil {
			return fmt.Errorf("database not initialized. Please run 'init' first")
		}
		err := c.db.ChangeKey(oldKey, newKey)
		if err != nil {
			return err
		}
		c.logger.Info("Database key changed successfully.")
	case "exit":
		c.logger.Info("Exiting CLI.")
		os.Exit(0)
	case "version":
		c.logger.Info("PolarysDB CLI Version:", c.version)

	case "help":
		c.logger.Info("Available commands:")
		for _, cmd := range c.commands {
			c.logger.Info(cmd.Name+":", cmd.Description)
		}
	default:
		return fmt.Errorf("unknown command: %s", commandName)
	}
	return nil

}

func (c *CLI) Shutdown() {
	if c.db != nil {
		c.db.Close()
	}

	if err := c.logger.Close(); err != nil {
		fmt.Println("error closing logger:", err)
	}
}
