package config

import (
	"fmt"
	"github.com/astaxie/beego/config"
)


type Config struct {
	IniConfig config.Configer
}

func NewConfig() *Config {
	this := new(Config)
	//iniConfig, err := config.NewConfig("ini", "conf/app.conf")
	iniConfig, err := config.NewConfig("ini", "../conf/app.conf")
	if err != nil {
		panic(fmt.Sprintf("fail to load conf file : %s", err))
	}
	this.IniConfig = iniConfig
	return this
}

func (c *Config) Get(key string) interface{} {
	return c.IniConfig.String(key)
}

func (c *Config) GetString(key string) string {
	return c.IniConfig.String(key)
}

func (c *Config) GetBool(key string) bool {
	b, _ := c.IniConfig.Bool(key)
	return b
}

func (c *Config) GetList(key string) []string {
	return c.IniConfig.Strings(key)
}


func (c *Config) Has(key string) bool {
	return true
}

