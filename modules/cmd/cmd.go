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
		Name:        "help",
		Description: "Show help",
		Args:        nil,
	},
}
