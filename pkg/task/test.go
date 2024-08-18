package task

import (
	"context"
	"package-name/pkg/global"
	"time"

	i_logger "github.com/pefish/go-interface/i-logger"
	go_logger "github.com/pefish/go-logger"
)

type Test struct {
	logger i_logger.ILogger
}

func NewTest() *Test {
	w := &Test{}
	w.logger = go_logger.Logger.CloneWithPrefix(w.Name())
	return w
}

func (t *Test) Init(ctx context.Context) error {
	return nil
}

func (t *Test) Run(ctx context.Context) error {
	t.Logger().Info(global.GlobalConfig)
	return nil
}

func (t *Test) Stop() error {
	return nil
}

func (t *Test) Name() string {
	return "Test"
}

func (t *Test) Interval() time.Duration {
	return 3 * time.Second
}

func (t *Test) Logger() i_logger.ILogger {
	return t.logger
}
