package mainModule

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	core "github.com/shinYeongHyeon/settlement-supporter/src/core/postgres"
	user3 "github.com/shinYeongHyeon/settlement-supporter/src/user/domain"
	user "github.com/shinYeongHyeon/settlement-supporter/src/user/entity"
	user2 "github.com/shinYeongHyeon/settlement-supporter/src/user/repository"
	"log"
)

// CreateModule is returned fiber.App for mounting
func CreateModule() *fiber.App {
	migratePostgres()

	mainModule := fiber.New()

	// TODO: Error to 404
	mainModule.Get("/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello ✋\nWe are Settlement Supporter API.\nRequest URL %s is not found.", c.Params("*"))
		return c.SendString(msg)
	})

	return mainModule
}

func migratePostgres() {
	gorm := core.PostgresConnect()

	err := gorm.Table("users").AutoMigrate(&user.Entity{})

	userId, _ := user3.IdCreate("a")
	userPassword, _ := user3.PasswordCreate("a")
	userNickName, _ := user3.NickNameCreate("a")
	userProps := user3.NewUserProps{
		Id:       userId,
		Password: userPassword,
		NickName: userNickName,
	}
	user, _ := user3.NewUserCreate(userProps)

	user2.Init(gorm.Table("users"))
	user2.Create(user)

	if err != nil {
		log.Fatal(err)
	}
}
