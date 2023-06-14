package config

type App struct {
	Env        string `mapstructure:"env" json:"env" yaml:"env"`
	Port       int    `mapstructure:"port" json:"port" yaml:"port"`
	AppName    string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl     string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
	DbType     string `mapstructure:"db_type" json:"db_type" yaml:"db_type"`
	DbUserName string `mapstructure:"db_user_name" json:"db_user_name" yaml:"db_user_name"`
	DbPassWord string `mapstructure:"db_password" json:"db_password" yaml:"db_password"`
	DbName     string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`
	DbUrl      string `mapstructure:"db_url" json:"db_url" yaml:"db_url"`
	DbPort     string `mapstructure:"db_port" json:"db_port" yaml:"db_port"`
}
