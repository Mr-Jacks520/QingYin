//总配置结构体
package config

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// gorm
	Mysql MySql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// oss
	Local     Local     `mapstructure:"local" json:"local" yaml:"local"`
	AliyunOSS AliyunOSS `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`
}
