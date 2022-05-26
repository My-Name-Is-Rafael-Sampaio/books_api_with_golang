package migrations

import (
	"github.com/My-Name-Is-Rafael-Sampaio/webapi-with-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
