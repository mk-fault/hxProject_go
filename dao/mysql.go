package dao

import (
	"fmt"
	"hx/asset"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Mysql MysqlConfig `yaml:"mysql"`
}

type MysqlConfig struct {
	Hx    HxMysqlConfig    `yaml:"hx"`
	HxGBD HxGBDMysqlConfig `yaml:"hx_gbd"`
}

type HxMysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type HxGBDMysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

var DBHX, DBHXGBD *gorm.DB

func InitDB() {
	data, err := asset.ConfigFile.ReadFile("database.yaml")
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic("数据库配置文件解析失败")
	}

	dsnHx := fmt.Sprintf("%s:%s@tcp(%s:%d)/hx?charset=utf8mb4&parseTime=True&loc=Local",
		config.Mysql.Hx.User,
		config.Mysql.Hx.Password,
		config.Mysql.Hx.Host,
		config.Mysql.Hx.Port,
	)

	dsnHxGBD := fmt.Sprintf("%s:%s@tcp(%s:%d)/hx_gbd?charset=utf8mb4&parseTime=True&loc=Local",
		config.Mysql.HxGBD.User,
		config.Mysql.HxGBD.Password,
		config.Mysql.HxGBD.Host,
		config.Mysql.HxGBD.Port,
	)

	DBHX, err = gorm.Open(mysql.Open(dsnHx), &gorm.Config{})
	if err != nil {
		panic("数据库`hx`连接失败:")
	}

	DBHXGBD, err = gorm.Open(mysql.Open(dsnHxGBD), &gorm.Config{})
	if err != nil {
		panic("数据库`hx_gbd`连接失败:")
	}
}
