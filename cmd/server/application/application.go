package application

import "github.com/laidingqing/stackbuild/server"

// Application is the main struct for the server.
type Application struct {
	Server *server.Server
}

// NewApplication creates a new application struct.
func NewApplication(
	server *server.Server) Application {
	return Application{
		Server: server,
	}
}
