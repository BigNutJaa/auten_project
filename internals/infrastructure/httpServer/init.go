package httpServer

import (
	"context"
	controllerProducts "github.com/BigNutJaa/users/internals/controller/products"
	controllerToken "github.com/BigNutJaa/users/internals/controller/token"
	"net/http"
	"strconv"

	"github.com/BigNutJaa/users/internals/config"
	controllerUsers "github.com/BigNutJaa/users/internals/controller/users"
	apiV1 "github.com/BigNutJaa/users/pkg/api/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	Config       config.Configuration
	Server       *runtime.ServeMux
	HttpMux      *http.ServeMux
	UsersCtrl    *controllerUsers.Controller
	TokenCtrl    *controllerToken.Controller
	ProductsCtrl *controllerProducts.Controller
}

func (s *Server) Configure(ctx context.Context, opts []grpc.DialOption) {

	apiV1.RegisterRegisterServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
	apiV1.RegisterLoginServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
	apiV1.RegisterProductsServiceHandlerFromEndpoint(ctx, s.Server, "0.0.0.0:"+strconv.Itoa(s.Config.Port), opts)
}

func NewServer(config config.Configuration, rmux *runtime.ServeMux, httpMux *http.ServeMux,
	usersCtrl *controllerUsers.Controller,
	tokenCtrl *controllerToken.Controller,
	productsCtrl *controllerProducts.Controller,

) *Server {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	s := &Server{
		Config:       config,
		Server:       rmux,
		HttpMux:      httpMux,
		UsersCtrl:    usersCtrl,
		TokenCtrl:    tokenCtrl,
		ProductsCtrl: productsCtrl,
	}
	s.Configure(context.Background(), opts)
	return s
}
