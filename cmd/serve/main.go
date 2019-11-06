package main

import (
	log "github.com/blockchain-abstraction-middleware/project-template/pkg/logger"
	"github.com/blockchain-abstraction-middleware/project-template/pkg/routes"
	"github.com/blockchain-abstraction-middleware/project-template/pkg/server"

	config "github.com/blockchain-abstraction-middleware/game-jam-abstraction/pkg/config"
	gm "github.com/blockchain-abstraction-middleware/game-jam-abstraction/pkg/routes"
)

func main() {
	abstractionConfig := config.NewConfig()

	serverConfig := server.Config{
		BasePath:       "/api/v1",
		Name:           abstractionConfig.Name,
		Port:           abstractionConfig.Port,
		MetricsEnabled: abstractionConfig.MetricsEnabled,
	}

	srv := server.New(&serverConfig)

	healthResource := routes.HealthResource{}
	swaggerResource := routes.SwaggerResource{}
	gameManagerResource := gm.ManagerResource{}

	srv.RegisterResource(healthResource.NewResource("/health"))
	srv.RegisterResource(swaggerResource.NewResource("/swagger"))

	srv.RegisterResource(gameManagerResource.NewResource("/manager", abstractionConfig))

	if err := srv.Run(); err != nil {
		log.WithError(err).Fatal("Serving failed")
	}
}
