package config

var Config EnvConfig

type EnvConfig struct {
	AppName string
	Port    string
	JDoodle JDoodle
}

type JDoodle struct {
	ClientId     string
	ClientSecret string
	Host         string
}
