package global

type Config struct {
	TestTask struct {
		Interval uint64 `json:"interval"`
	} `json:"testTask"`
}

var GlobalConfig Config
