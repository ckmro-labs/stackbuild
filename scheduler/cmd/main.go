package main

import (
	"context"

	"github.com/laidingqing/stackbuild/cmd/server/config"
	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/scheduler/queue"
	"github.com/laidingqing/stackbuild/store/shared/db"
	"github.com/laidingqing/stackbuild/store/stage"
	"github.com/sirupsen/logrus"
)

// this is a queue test main

func main() {
	config, err := config.Environ()
	if err != nil {
		logger := logrus.WithError(err)
		logger.Fatalln("main: invalid configuration")
	}

	session := db.NewSessionStore(config.Database.Datasource, config.Database.Database)
	stageStore := stage.New(session)
	ctx := context.Background()

	item1 := &core.Stage{
		ID: "1",
	}
	item2 := &core.Stage{
		ID: "2",
	}
	q := queue.NewQueue(stageStore)
	q.Schedule(ctx, item1)
	q.Schedule(ctx, item2)
	stage, err := q.Request(ctx)
	if err != nil {
		logrus.Errorf("request next stage err : %v", err.Error())
	}
	logrus.Infof("next: %v", stage.ID)
}
