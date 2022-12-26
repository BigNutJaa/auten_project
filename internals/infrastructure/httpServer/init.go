package httpServer

import (
	"context"
	"net/http"
	"strconv"

	"github.com/BigNutJaa/users/internals/config"
	controllerUsers "github.com/BigNutJaa/users/internals/controller/users"
	apiV1 "github.com/BigNutJaa/users/pkg/api/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	Config    config.Configuration
	Server    *runtime.ServeMux
	HttpMux   *http.ServeMux
	UsersCtrl *controllerUsers.Controller
}

func (s *Server) Configure(ctx context.Context, opts []grpc.DialOption) {

	apiV1.RegisterAutenServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)

}

func NewServer(config config.Configuration, rmux *runtime.ServeMux, httpMux *http.ServeMux,
	usersCtrl *controllerUsers.Controller,

) *Server {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	s := &Server{
		Config:    config,
		Server:    rmux,
		HttpMux:   httpMux,
		UsersCtrl: usersCtrl,
	}
	s.Configure(context.Background(), opts)
	return s
}
