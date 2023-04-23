package config

type Server struct {
	App    App    `mapstructure:"app" json:"app" yaml:"app"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
}
