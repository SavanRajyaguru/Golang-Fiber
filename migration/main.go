package migration

import (
	"log"

	"github.com/savan/database"
	associationtest "github.com/savan/services/association-test"
)

func Migration(models ...interface{}) {
	for _, model := range models {
		err := database.DB.AutoMigrate(&model)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func LoadAllSchema() {
	Migration(
		// &auth.UserModel{},
		// &auth.Company{},
		&associationtest.Emp{},
		&associationtest.Card{},
	)
	// func() {
	// 	database.DB.Migrator().DropColumn(&auth.UserModel{}, "company_id")
	// }()
}
