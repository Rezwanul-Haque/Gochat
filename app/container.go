package container

import (
	"gochat/app/http/controllers"
	repoImpl "gochat/app/repository/impl"
	svcImpl "gochat/app/svc/impl"
)

func Init(g interface{}) {
	// register all repos impl, services impl, controllers
	sysRepo := repoImpl.NewSystemRepository()
	userRepo := repoImpl.NewFirebaseUsersRepository()
	authRepo := repoImpl.NewFirebaseAuthRepository()

	sysSvc := svcImpl.NewSystemService(sysRepo)
	userSvc := svcImpl.NewUsersService(userRepo)
	authSvc := svcImpl.NewAuthService(authRepo)

	controllers.NewSystemController(g, sysSvc)
	controllers.NewAuthController(g, authSvc, userSvc)
	controllers.NewUsersController(g, userSvc)
	controllers.NewRoomsController(g)
}
