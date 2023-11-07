package command

import (
	"flag"

	"github.com/pefish/go-commander"
	go_config "github.com/pefish/go-config"
	task_driver "github.com/pefish/go-task-driver"
	"github.com/pefish/template/pkg/global"
	"github.com/pefish/template/pkg/task"
)

type DefaultCommand struct {
}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{}
}

func (dc *DefaultCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	return nil
}

func (dc *DefaultCommand) Init(data *commander.StartData) error {
	err := go_config.ConfigManagerInstance.Unmarshal(&global.GlobalConfig)
	if err != nil {
		return err
	}

	//go_mysql.MysqlInstance.SetLogger(go_logger.Logger)
	//go_mysql.MysqlInstance.MustConnectWithConfiguration(go_mysql.Configuration{
	//	Host:     global.GlobalConfig.Db.Host,
	//	Username: global.GlobalConfig.Db.User,
	//	Password: global.GlobalConfig.Db.Pass,
	//	Database: global.GlobalConfig.Db.Db,
	//})

	return nil
}

func (dc *DefaultCommand) OnExited(data *commander.StartData) error {
	//go_mysql.MysqlInstance.Close()
	return nil
}

func (dc *DefaultCommand) Start(data *commander.StartData) error {
	taskDriver := task_driver.NewTaskDriver()
	taskDriver.Register(task.NewTest())
	taskDriver.RunWait(data.ExitCancelCtx)
	return nil
}
