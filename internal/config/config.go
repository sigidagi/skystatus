package config

type Config struct {
	General struct {
		LogLevel    int  `mapstructure:"log_level"`
		LogToSyslog bool `mapstructure:"log_to_syslog"`
	} `mapstructure:"general"`

	Device struct {
		Tty string `mapstructure:"tty"`
	} `mapstructure:"device"`
}

var C Config
