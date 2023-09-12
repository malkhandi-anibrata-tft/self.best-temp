package seed

import (
	"log"

	"github.com/malkhandi-anibrata-tft/self.best-temp/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	err := db.AutoMigrate(
		&models.UserDetails{},
	)

	if err != nil {
		log.Fatalln(err)
		return
	}
}
