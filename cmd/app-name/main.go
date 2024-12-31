package main

import (
	"fmt"
	"package-name/cmd/app-name/command"
	"package-name/pkg/global"
	"package-name/version"

	"github.com/pefish/go-commander"
)

func main() {
	global.Commander = commander.NewCommander(
		version.AppName,
		version.Version,
		fmt.Sprintf("%s is a template. Author: pefish", version.AppName),
	)
	global.Commander.RegisterDefaultSubcommand(&commander.SubcommandInfo{
		Desc:       "Use this command by default if you don't set subcommand.",
		Args:       nil,
		Subcommand: command.NewDefaultCommand(),
	})
	err := global.Commander.Run()
	if err != nil {
		global.Commander.Logger.Error(err)
	}
}
