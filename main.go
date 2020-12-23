package main

import (
	"github.com/4everlynn/journal/cmd"
	"github.com/4everlynn/journal/config"
)

func main() {
	config.InstallDefaultPlugins()
	cmd.Execute()
}
