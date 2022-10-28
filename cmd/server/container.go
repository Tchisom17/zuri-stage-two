package server

import (
	"gorm.io/gorm"
	"zuri-stage-one/internal/adapters/api"
	"zuri-stage-one/internal/adapters/repository"
	"zuri-stage-one/internal/core/helpers"
	"zuri-stage-one/internal/core/service"
)

// Injection inject all dependencies
func Injection(db *gorm.DB) {
	userRepository := repository.NewUser(db)
	userService := service.NewUserService(userRepository)
	Handler := api.NewHTTPHandler(userService)
	router := SetupRouter(Handler)

	_ = router.Run(":" + helpers.Instance.Port)
}
