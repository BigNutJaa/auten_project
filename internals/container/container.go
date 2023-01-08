package container

import (
	controllerToken "github.com/BigNutJaa/users/internals/controller/token"
	controllerUsers "github.com/BigNutJaa/users/internals/controller/users"
	"net/http"

	"github.com/BigNutJaa/users/internals/service/token"
	"github.com/BigNutJaa/users/internals/service/users"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/robowealth-mutual-fund/shared-utility/validator"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"

	"github.com/BigNutJaa/users/internals/config"
	"github.com/BigNutJaa/users/internals/controller"

	"github.com/BigNutJaa/users/internals/infrastructure/database"
	grpcServer "github.com/BigNutJaa/users/internals/infrastructure/grpcServer"
	httpServer "github.com/BigNutJaa/users/internals/infrastructure/httpServer"
	"github.com/BigNutJaa/users/internals/infrastructure/jaeger"
	"github.com/BigNutJaa/users/internals/repository/postgres"

	"github.com/BigNutJaa/users/internals/utils"
	"github.com/BigNutJaa/users/internals/utils/logrus"
)

type Container struct {
	container *dig.Container
}

func (c *Container) Configure() error {

	servicesConstructors := []interface{}{
		config.NewConfiguration,
		grpcServer.NewServer,
		database.NewServerBase,
		http.NewServeMux,
		httpServer.NewServer,
		runtime.NewServeMux,
		jaeger.NewJaeger,
		logrus.NewLog,
		controller.NewHealthZController,
		controller.NewPingPongController,
		validator.NewCustomValidator,
		postgres.NewRepository,
		utils.NewUtils,
		utils.NewCustomValidator,
		users.NewService,
		token.NewService,
		controllerUsers.NewController,
		controllerToken.NewController,
	}
	for _, service := range servicesConstructors {
		if err := c.container.Provide(service); err != nil {
			return err
		}
	}
	appConfig := config.NewConfiguration()
	jaeger.NewJaeger(appConfig)
	logrus.NewLog()
	return nil
}

func (c *Container) Start() error {
	log.Info("Start Container")
	if err := c.container.Invoke(func(s *grpcServer.Server, h *httpServer.Server, v *validator.CustomValidator) {
		go func() {
			_ = h.Start()
		}()
		s.Start()

	}); err != nil {
		log.Errorf("%s", err)

		return err
	}

	return nil
}

// MigrateDB ...
func (c *Container) MigrateDB() error {
	log.Info("Start Container DB")

	if err := c.container.Invoke(func(d *database.DB) {
		d.MigrateDB()
	}); err != nil {
		return err
	}

	return nil
}

func NewContainer() (*Container, error) {
	d := dig.New()

	container := &Container{
		container: d,
	}

	if err := container.Configure(); err != nil {
		return nil, err
	}

	return container, nil
}
