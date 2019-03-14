package application

import (
	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/server"
)

// Application is the main struct for the server.
type Application struct {
	Server *server.Server
	users  core.UserStore
}

// NewApplication creates a new application struct.
func NewApplication(
	server *server.Server, users core.UserStore) Application {
	return Application{
		Server: server,
		users:  users,
	}
}
