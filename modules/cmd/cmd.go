package cmd

type Command struct {
	Name        string
	Description string
	Args        []string
}

var Commands = []Command{
	{
		Name:        "init",
		Description: "Initialize database",
		Args:        []string{"key", "path"},
	},
	{
		Name:        "export",
		Description: "Export data to a file .json",
		Args:        []string{"key", "path"},
	},
	{
		Name:        "import",
		Description: "Import data from a file .json",
		Args:        []string{"key", "path"},
	},
	{
		Name:        "export-encrypted",
		Description: "Export encrypted data to a file .json",
		Args:        []string{"key", "path"},
	},
	{
		Name:        "import-encrypted",
		Description: "Import encrypted data from a file .json",
		Args:        []string{"key", "path"},
	},
	{
		Name:        "new-key",
		Description: "Generate a new key",
		Args:        nil,
	},
	{
		Name:        "key-from",
		Description: "Generate a key from a string",
		Args:        []string{"string"},
	},
	{
		Name:        "exit",
		Description: "Exit CLI",
		Args:        nil,
	},
	{
		Name:        "version",
		Description: "Show version",
		Args:        nil,
	},
	{
		Name:        "help",
		Description: "Show help",
		Args:        nil,
	},
}
