package config


var Configs Config

func init() {
	Configs = *(NewConfig())
}
