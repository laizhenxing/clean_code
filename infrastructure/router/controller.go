package router

import (
	"cc/application"
	"cc/infrastructure/persistence"
	"cc/infrastructure/persistence/mongodb"
	"cc/infrastructure/persistence/mysqldb"
	"cc/infrastructure/persistence/orm"

	"cc/interfaces/handlers"
	"cc/interfaces/repository"
)

func getAppGinHandler() *handlers.AppGinHandler {
	// infrastructure
	//appHandler := mysqldb.NewAppHandler(persistence.Db.GetConn())
	appHandler := orm.NewAppOrmHandler(orm.ORM.GetConn())
	// interfaces
	appRepo := repository.NewAppRepo(appHandler)
	// application
	appInteractor := application.NewAppInteractor(appRepo)
	// interfaces
	controller := handlers.NewAppGinHandler(appInteractor)

	return controller
}

func getAppMongoHandler() *handlers.AppGinHandler {
	mongoHandler := mongodb.NewAppMongoHandler(persistence.MongoClient, "cc", "apps")
	appRepo := repository.NewAppRepo(mongoHandler)
	appInteractor := application.NewAppInteractor(appRepo)
	ctrl := handlers.NewAppGinHandler(appInteractor)

	return ctrl
}

func getAppHandler() *handlers.AppHandler {
	// infrastructure
	appHandler := mysqldb.NewAppHandler(persistence.Db.GetConn())
	//appHandler := orm.NewAppOrmHandler(orm.ORM.GetConn())
	// interfaces
	appRepo := repository.NewAppRepo(appHandler)
	// application
	appInteractor := application.NewAppInteractor(appRepo)
	// interfaces
	controller := handlers.NewAppHandler(appInteractor)

	return controller
}

func getUserHandler() *handlers.UserHandler {
	userHandler := orm.NewUserHandler(orm.ORM.GetConn())
	userRepo := repository.NewUserRepo(userHandler, nil)
	userInte := application.NewUserInteractor(userRepo)
	ctrl := handlers.NewUserHandler(userInte)

	return ctrl
}
