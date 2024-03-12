package config

type Config struct {
	General struct {
		LogLevel    int  `mapstructure:"log_level"`
		LogToSyslog bool `mapstructure:"log_to_syslog"`
	} `mapstructure:"general"`

	Device struct {
		Name string `mapstructure:"name"`
		Baud int    `mapstructure:"baud"`
	} `mapstructure:"device"`

	Project struct {
		Directory string `mapstructure:"directory"`
		Preset    string `mapstructure:"preset"`
		Cores     int    `mapstructure:"cores"`
		Plugins   string `mapstructure:"plugins"`
	} `mapstructure:"project"`
}

var C Config
