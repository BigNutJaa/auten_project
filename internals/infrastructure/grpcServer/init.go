package grpcserver

import (
	"github.com/BigNutJaa/users/internals/config"
	"github.com/BigNutJaa/users/internals/controller"
	controllerProducts "github.com/BigNutJaa/users/internals/controller/products"
	controllerToken "github.com/BigNutJaa/users/internals/controller/token"
	controllerUsers "github.com/BigNutJaa/users/internals/controller/users"
	apiV1 "github.com/BigNutJaa/users/pkg/api/v1"
	grpcErrors "github.com/robowealth-mutual-fund/shared-utility/grpc_errors"
	validatorUtils "github.com/robowealth-mutual-fund/shared-utility/validator"
	"google.golang.org/grpc"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	Config       config.Configuration
	Server       *grpc.Server
	HealthCtrl   *controller.HealthZController
	PingPongCtrl *controller.PingPongController
	UsersCtrl    *controllerUsers.Controller
	TokenCtrl    *controllerToken.Controller
	ProductsCtrl *controllerProducts.Controller
}

// Configure ...
func (s *Server) Configure() {
	grpcHealthV1.RegisterHealthServer(s.Server, s.HealthCtrl)
	apiV1.RegisterPingPongServiceServer(s.Server, s.PingPongCtrl)
	apiV1.RegisterRegisterServiceServer(s.Server, s.UsersCtrl)
	apiV1.RegisterLoginServiceServer(s.Server, s.TokenCtrl)
	apiV1.RegisterProductsServiceServer(s.Server, s.ProductsCtrl)
}

func NewServer(
	config config.Configuration,
	healthCtrl *controller.HealthZController,
	pingPongCtrl *controller.PingPongController,
	usersCtrl *controllerUsers.Controller,
	tokenCtrl *controllerToken.Controller,
	productsCtrl *controllerProducts.Controller,
	validator *validatorUtils.CustomValidator,
) *Server {
	option := grpc.ChainUnaryInterceptor(
		grpcErrors.UnaryServerInterceptor(),
		validatorUtils.UnaryServerInterceptor(validator),
	)

	s := &Server{
		Server:       grpc.NewServer(option, grpc.MaxRecvMsgSize(10*10e6), grpc.MaxSendMsgSize(10*10e6)),
		Config:       config,
		HealthCtrl:   healthCtrl,
		PingPongCtrl: pingPongCtrl,
		UsersCtrl:    usersCtrl,
		TokenCtrl:    tokenCtrl,
		ProductsCtrl: productsCtrl,
	}

	s.Configure()

	return s
}
