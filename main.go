package main

import "github.com/polarysfoundation/polarysdb-cli/modules/cmd"

func main() {
	cli := cmd.NewCLI()
	defer cli.Shutdown()
	cli.Run()
}
