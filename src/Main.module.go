package mainModule

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
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

	// TODO: Error to 404
	mainModule.Get("/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello ✋\nWe are Settlement Supporter API.\nRequest URL %s is not found.", c.Params("*"))
		return c.SendString(msg)
	})

	return mainModule
}

func migratePostgres() {
	manager := core.GetManager()

	err := manager.Db.Table("users").AutoMigrate(&userEntity.User{})

	if err != nil {
		log.Fatal(err)
	}
}
