package application

import "github.com/laidingqing/stackbuild/server"

// Application is the main struct for the server.
type Application struct {
	Server *server.GrpcServer
}

// NewApplication creates a new application struct.
func NewApplication(
	server *server.GrpcServer) Application {
	return Application{
		Server: server,
	}
}
