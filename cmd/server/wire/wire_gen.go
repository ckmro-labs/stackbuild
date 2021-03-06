// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire

import (
	"github.com/laidingqing/stackbuild/cmd/server/application"
	"github.com/laidingqing/stackbuild/cmd/server/config"
	"github.com/laidingqing/stackbuild/handler/api"
	"github.com/laidingqing/stackbuild/handler/web"
	"github.com/laidingqing/stackbuild/pubsub"
	"github.com/laidingqing/stackbuild/service/user"
)

// Injectors from wire.go:

func InitializeApplication(config2 config.Config) (application.Application, error) {
	sessionStore := provideDatabase(config2)
	repositoryStore := provideRepositoryStore(sessionStore)
	repositoryService := providerRepositoryService(config2)
	userStore := provideUserStore(sessionStore)
	syncer := provideSyncer(repositoryService, repositoryStore, userStore, config2)
	corePubsub := pubsub.New()
	stageStore := provideStageStore(sessionStore)
	session := provideSession(userStore, config2)
	server := api.New(repositoryStore, repositoryService, syncer, corePubsub, userStore, stageStore, session)
	userService := user.New()
	sourceAuthStore := provideSourceStore(sessionStore)
	webServer := web.New(repositoryStore, session, userStore, userService, syncer, sourceAuthStore)
	mux := provideRouter(server, webServer)
	serverServer := provideServer(mux, config2)
	applicationApplication := application.NewApplication(serverServer, userStore)
	return applicationApplication, nil
}
