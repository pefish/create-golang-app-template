package main

import (
	"create_golang_app_template/cmd/main/command"
	"create_golang_app_template/version"
	"github.com/pefish/go-commander"
	go_logger "github.com/pefish/go-logger"
)

func main() {
	commanderInstance := commander.NewCommander(version.AppName, version.Version, version.AppName + " 是一个模板，祝你玩得开心。作者：pefish")
	commanderInstance.RegisterDefaultSubcommand("默认子命令", command.NewDefaultCommand())
	err := commanderInstance.Run()
	if err != nil {
		go_logger.Logger.Error(err)
	}
}
