package config

import "github.com/kelseyhightower/envconfig"

type (
	// Config provides the system configuration.
	Config struct {
		Logging Logging
		Server  Server
	}

	// Server provides the server configuration.
	Server struct {
		Port string `envconfig:"STACK_BUILD_SERVER_PORT" default:":7788"`
		TLS  bool   `envconfig:"STACK_BUILD_TLS_AUTOCERT"`
		Cert string `envconfig:"STACK_BUILD_TLS_CERT"`
		Key  string `envconfig:"STACK_BUILD_TLS_KEY"`
	}

	// Logging provides the logging configuration.
	Logging struct {
		Debug  bool `envconfig:"STACK_BUILD_LOGS_DEBUG"`
		Trace  bool `envconfig:"STACK_BUILD_LOGS_TRACE"`
		Color  bool `envconfig:"STACK_BUILD_LOGS_COLOR"`
		Pretty bool `envconfig:"STACK_BUILD_LOGS_PRETTY"`
		Text   bool `envconfig:"STACK_BUILD_LOGS_TEXT"`
	}
)

// Environ returns the settings from the environment.
func Environ() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	// defaultAddress(&cfg)
	return cfg, err
}

// func defaultAddress(c *Config) {
// 	if c.Server.Key != "" || c.Server.Cert != "" || c.Server.TLS {
// 		c.Server.Port = ":443"
// 		c.Server.Proto = "https"
// 	}
// 	c.Server.Addr = c.Server.Proto + "://" + c.Server.Host
// }
