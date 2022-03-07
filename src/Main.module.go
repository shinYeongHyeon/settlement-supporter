package mainModule

import (
	"github.com/gofiber/fiber/v2"
	coreError "github.com/shinYeongHyeon/settlement-supporter/src/core/error"
	core "github.com/shinYeongHyeon/settlement-supporter/src/core/postgres"
	userModule "github.com/shinYeongHyeon/settlement-supporter/src/user"
	userEntity "github.com/shinYeongHyeon/settlement-supporter/src/user/entity"
	"log"
)

// CreateModule is returned fiber.App for mounting
func CreateModule() *fiber.App {
	migratePostgres()

	mainModule := fiber.New()

	mainModule.Mount("/user", userModule.CreateModule())

	mainModule.Get("/*", coreError.NotFoundError)

	return mainModule
}

func migratePostgres() {
	manager := core.GetManager()

	err := manager.Db.Table("users").AutoMigrate(&userEntity.User{})

	if err != nil {
		log.Fatal(err)
	}
}
