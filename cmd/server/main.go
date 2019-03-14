// Copyright 2019 CAIKONG LTD, Inc.

package main

import (
	"context"
	"flag"

	"github.com/drone/signal"
	"github.com/joho/godotenv"
	"github.com/laidingqing/stackbuild/cmd/server/config"
	"github.com/laidingqing/stackbuild/server"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func main() {
	var envfile string
	flag.StringVar(&envfile, "env-file", ".env", "Read in a file of environment variables")
	flag.Parse()

	godotenv.Load(envfile)
	config, err := config.Environ()
	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: invalid configuration")
	}
	initLogging(config)
	ctx := signal.WithContext(
		context.Background(),
	)

	app, err := InitializeApplication(config)
	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: cannot initialize server")
	}

	g := errgroup.Group{}
	g.Go(func() error {
		logrus.WithFields(
			logrus.Fields{
				"proto": config.Server.Proto,
				"host":  config.Server.Host,
				"port":  config.Server.Port,
				"url":   config.Server.Addr,
				"tls":   config.Server.TLS,
			},
		).Infoln("starting the http server")
		return app.server.ListenAndServe(ctx)
	})
}

func initLogging(c config.Config) {
	if c.Logging.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if c.Logging.Text {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   c.Logging.Color,
			DisableColors: !c.Logging.Color,
		})
	}
}

// application is the main struct for the server.
type application struct {
	server *server.Server
}

// newApplication creates a new application struct.
func newApplication(
	server *server.Server) application {
	return application{
		server: server,
	}
}
