package global

type Config struct {
	TestTask struct {
		Interval uint64 `json:"interval"`
	} `json:"testTask"`
	//Db struct {
	//	Db   string `json:"db"`
	//	Host string `json:"host"`
	//	User string `json:"user"`
	//	Pass string `json:"pass"`
	//} `json:"db"`
}

var GlobalConfig Config
