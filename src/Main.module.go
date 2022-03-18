package mainModule

import (
	"github.com/gofiber/fiber/v2"
	coreError "github.com/shinYeongHyeon/settlement-supporter/src/core/error"
	core "github.com/shinYeongHyeon/settlement-supporter/src/core/postgres"
	groupEntity "github.com/shinYeongHyeon/settlement-supporter/src/group/entity"
	userModule "github.com/shinYeongHyeon/settlement-supporter/src/user"
	userEntity "github.com/shinYeongHyeon/settlement-supporter/src/user/entity"
	"log"
)

// CreateModule is returned fiber.App for mounting
func CreateModule() *fiber.App {
	migratePostgres()

	mainModule := fiber.New()

	mainModule.Mount("/users", userModule.CreateModule())

	mainModule.Get("/*", coreError.NotFoundError)

	return mainModule
}

func migratePostgres() {
	manager := core.GetManager()

	// TODO: 여러 에러 동시처리
	errUsers := manager.Db.Table("users").AutoMigrate(&userEntity.User{})
	errGroups := manager.Db.Table("groups").AutoMigrate(&groupEntity.Group{})
	errGroupUsers := manager.Db.Table("group_users").AutoMigrate(&groupEntity.GroupUsers{})

	if errUsers != nil || errGroups != nil || errGroupUsers != nil {
		log.Fatal(errUsers, errGroups, errGroupUsers)
	}
}
