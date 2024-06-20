package global

import "github.com/pefish/go-commander"

type Config struct {
	commander.BasicConfig
	//Db struct {
	//	Db   string `json:"db"`
	//	Host string `json:"host"`
	//	User string `json:"user"`
	//	Pass string `json:"pass"`
	//} `json:"db"`
	Test string `json:"test" default:"t" usage:"This is test param."`
}

var GlobalConfig Config
