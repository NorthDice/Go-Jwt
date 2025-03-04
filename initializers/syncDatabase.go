package initializers

import "go-jwt/models"

func SyncDatabase() {

	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Error while migrating database")
	}

}
