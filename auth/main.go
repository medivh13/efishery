package main

import (
	"context"

	usecases "efishery/auth/src/app/use_cases"

	"efishery/auth/src/infra/config"
	postgres "efishery/auth/src/infra/persistence/postgress"

	"efishery/auth/src/interface/rest"

	ms_log "efishery/auth/src/infra/log"

	userUc "efishery/auth/src/app/use_cases/user"

	_ "github.com/joho/godotenv/autoload"
)

//reupdate by Jody 24 Jan 2022
func main() {
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	postgresdb := postgres.New(conf.SqlDb, logger)

	userRepository := postgres.NewUsersRepository(postgresdb.DB)
	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		usecases.AllUseCases{
			UserUsecase: userUc.NewUserUseCase(userRepository),
		},
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
