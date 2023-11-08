package main

import (
	"github.com/pefish/go-commander"
	go_logger "github.com/pefish/go-logger"
	"github.com/pefish/template/cmd/template/command"
	"github.com/pefish/template/version"
)

func main() {
	commanderInstance := commander.NewCommander(version.AppName, version.Version, version.AppName+" is a template.")
	commanderInstance.RegisterDefaultSubcommand("Use this command by default if you don't set subcommand.", command.NewDefaultCommand())
	err := commanderInstance.Run()
	if err != nil {
		go_logger.Logger.Error(err)
	}
}
