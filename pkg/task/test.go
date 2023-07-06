package task

import (
	"context"
	"create_golang_app_template/pkg/global"
	go_logger "github.com/pefish/go-logger"
	"time"
)

type Test struct {
	logger go_logger.InterfaceLogger
}

func NewTest() *Test {
	w := &Test{}
	w.logger = go_logger.Logger.CloneWithPrefix(w.GetName())
	return w
}

func (t *Test) Run(ctx context.Context) error {
	t.GetLogger().Info("test")
	return nil
}

func (t *Test) Stop() error {
	return nil
}

func (t *Test) GetName() string {
	return "Test"
}

func (t *Test) GetInterval() time.Duration {
	return time.Duration(global.GlobalConfig.TestTask.Interval) * time.Second
}

func (t *Test) GetLogger() go_logger.InterfaceLogger {
	return t.logger
}
