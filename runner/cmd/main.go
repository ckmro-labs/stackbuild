package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/drone/signal"
	"github.com/laidingqing/stackbuild/runner/executor"
	"github.com/laidingqing/stackbuild/runner/executor/docker"
	"github.com/laidingqing/stackbuild/runner/runtime"
	isatty "github.com/mattn/go-isatty"
)

var tty = isatty.IsTerminal(os.Stdout.Fd())

// This a stanlone main.
func main() {
	c := flag.String("config", "", "")
	t := flag.Duration("timeout", time.Hour, "")
	flag.Parse()

	var source string
	if flag.NArg() > 0 {
		source = flag.Args()[0]
	}

	config, err := executor.ParseFile(source)
	if err != nil {
		log.Fatalln(err)
	}
	if *c != "" {
		auths, err := docker.ParseFile(*c)
		if err != nil {
			log.Fatalln(err)
		}
		config.Docker.Auths = append(config.Docker.Auths, auths...)
	}

	executor, err := docker.NewEnv()
	if err != nil {
		log.Fatalln(err)
	}
	hooks := &runtime.Hook{}
	hooks.GotLine = runtime.WriteLine(os.Stdout)
	if tty {
		hooks.GotLine = runtime.WriteLinePretty(os.Stdout)
	}

	r := runtime.New(
		runtime.WithExecutor(executor),
		runtime.WithConfig(config),
		runtime.WithHooks(hooks),
	)

	ctx, cancel := context.WithTimeout(context.Background(), *t)
	ctx = signal.WithContext(ctx)
	defer cancel()
	err = r.Run(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
