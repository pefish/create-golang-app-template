package main

import (
	"create_golang_app_template/pkg/command"
	"create_golang_app_template/version"
	"github.com/pefish/go-commander"
	"log"
)

func main() {
	commanderInstance := commander.NewCommander(version.AppName, version.Version, version.AppName + " 是一个模板，祝你玩得开心。作者：pefish")
	//commanderInstance.RegisterSubcommand("client", client.NewClient())
	//commanderInstance.RegisterSubcommand("server", server.NewServer())
	commanderInstance.RegisterDefaultSubcommand(command.NewDefaultCommand())
	err := commanderInstance.Run()
	if err != nil {
		log.Fatal(err)
	}
}
