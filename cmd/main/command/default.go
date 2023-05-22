package command

import (
	"create_golang_app_template/pkg/global"
	"flag"
	"github.com/pefish/go-commander"
	go_config "github.com/pefish/go-config"
	go_logger "github.com/pefish/go-logger"
	"net"
)

type DefaultCommand struct {
}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{}
}

func (dc *DefaultCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	flagSet.String("tcp-address", "0.0.0.0:8000", "<addr>:<port> to listen on for TCP clients")
	return nil
}

func (dc *DefaultCommand) OnExited(data *commander.StartData) error {
	return nil
}

func (dc *DefaultCommand) Start(data *commander.StartData) error {
	err := go_config.ConfigManagerInstance.Unmarshal(&global.MyConfig)
	if err != nil {
		return err
	}

	tcpListener, err := net.Listen("tcp", global.MyConfig.TcpAddress)
	if err != nil {
		return err
	}
	defer func() {
		tcpListener.Close()
		go_logger.Logger.InfoF("HTTP endpoint closed. url: %s", tcpListener.Addr())
	}()
	go_logger.Logger.InfoF("HTTP endpoint opened. url: %s", tcpListener.Addr())

	<-data.ExitCancelCtx.Done()
	return nil
}
