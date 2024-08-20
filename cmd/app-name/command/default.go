package command

import (
	"package-name/pkg/global"
	"package-name/pkg/task"

	"github.com/pefish/go-commander"
	task_driver "github.com/pefish/go-task-driver"
)

type PersistenceDataType struct {
}

type DefaultCommand struct {
	persistenceData *PersistenceDataType
}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{
		persistenceData: &PersistenceDataType{},
	}
}

func (dc *DefaultCommand) Config() interface{} {
	return &global.GlobalConfig
}

func (dc *DefaultCommand) Data() interface{} {
	return dc.persistenceData
}

func (dc *DefaultCommand) Init(command *commander.Commander) error {
	// global.MysqlInstance = go_mysql.NewMysqlInstance(command.Logger)
	// err := global.MysqlInstance.ConnectWithConfiguration(t_mysql.Configuration{
	// 	Host:     global.GlobalConfig.DbHost,
	// 	Username: global.GlobalConfig.DbUser,
	// 	Password: global.GlobalConfig.DbPass,
	// 	Database: global.GlobalConfig.DbDb,
	// })
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (dc *DefaultCommand) OnExited(command *commander.Commander) error {
	//global.MysqlInstance.Close()
	return nil
}

func (dc *DefaultCommand) Start(command *commander.Commander) error {
	taskDriver := task_driver.NewTaskDriver()
	taskDriver.Register(task.NewTest(command.Logger))
	taskDriver.RunWait(command.Ctx)
	return nil
}
