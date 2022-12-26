package grpcserver

import (
	"github.com/BigNutJaa/users/internals/config"
	"github.com/BigNutJaa/users/internals/controller"
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
}

// Configure ...
func (s *Server) Configure() {
	grpcHealthV1.RegisterHealthServer(s.Server, s.HealthCtrl)
	apiV1.RegisterPingPongServiceServer(s.Server, s.PingPongCtrl)
	apiV1.RegisterAutenServiceServer(s.Server, s.UsersCtrl)
}

func NewServer(
	config config.Configuration,
	healthCtrl *controller.HealthZController,
	pingPongCtrl *controller.PingPongController,
	usersCtrl *controllerUsers.Controller,
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
	}

	s.Configure()

	return s
}
