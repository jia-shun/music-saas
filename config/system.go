package config

type System struct {
	Env      string `mapstructure:"env" json:"env" yaml:"env"`
	Addr     int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	UseCache bool   `mapstructure:"use-cache" json:"useCache" yaml:"use-cache"`
}
