package config

import "github.com/kelseyhightower/envconfig"

type (
	// Config provides the system configuration.
	Config struct {
		Logging  Logging
		Server   Server
		Database Database

		// Remote configurations
		Github Github
	}

	// Server provides the server configuration.
	Server struct {
		Addr  string `envconfig:"-"`
		Host  string `envconfig:"STACK_BUILD_SERVER_HOST" default:"localhost:7788"`
		Port  string `envconfig:"STACK_BUILD_SERVER_PORT" default:":7788"`
		Proto string `envconfig:"STACK_BUILD_SERVER_PROTO" default:"http"`
		TLS   bool   `envconfig:"STACK_BUILD_TLS_AUTOCERT"`
		Cert  string `envconfig:"STACK_BUILD_TLS_CERT"`
		Key   string `envconfig:"STACK_BUILD_TLS_KEY"`
	}

	// Logging provides the logging configuration.
	Logging struct {
		Debug  bool `envconfig:"STACK_BUILD_LOGS_DEBUG"`
		Trace  bool `envconfig:"STACK_BUILD_LOGS_TRACE"`
		Color  bool `envconfig:"STACK_BUILD_LOGS_COLOR"`
		Pretty bool `envconfig:"STACK_BUILD_LOGS_PRETTY"`
		Text   bool `envconfig:"STACK_BUILD_LOGS_TEXT"`
	}

	// Github client configuration.
	Github struct {
		ClientID     string `envconfig:"STACK_BUILD_GITHUB_CLIENT_ID" default:"7b8a75b008712044e7cf"`
		ClientSecret string `envconfig:"STACK_BUILD_GITHUB_CLIENT_SECRET" default:"908c916f3998603622e070e7d8b6bfc47f733d01"`
		CallbackURL  string `envconfig:"STACK_BUILD_GITHUB_CALLBACK_URL" default:"http://localhost:7788/login/"`
	}

	// Database provides the database configuration.
	Database struct {
		Database   string `envconfig:"STACK_BUILD_DATABASE_DRIVER"     default:"test"`
		Datasource string `envconfig:"STACK_BUILD_DATABASE_DATASOURCE" default:"mongodb://localhost:27017/stackbuild"`
	}
)

// Environ returns the settings from the environment.
func Environ() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	defaultAddress(&cfg)
	return cfg, err
}

func defaultAddress(c *Config) {
	if c.Server.Key != "" || c.Server.Cert != "" || c.Server.TLS {
		c.Server.Port = ":443"
		c.Server.Proto = "https"
	}
	c.Server.Addr = c.Server.Proto + "://" + c.Server.Host
}
