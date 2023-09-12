package models

import (
	"log"
	"github.com/malkhandi-anibrata-tft/self.best-temp/db"

	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	return db.GetDB(), nil
}

func GetTx() (*gorm.DB, error) {
	dbHandler, err := GetDB()
	if err != nil {
		log.Println("Transaction Begin Error : ", err)
		return nil, err
	}
	tx := dbHandler.Begin()
	err = tx.Error
	if err != nil {
		log.Println("Transaction Begin Error : ", tx.Error)
	}
	return tx, err
}

func CompleteTx(tx *gorm.DB, err error) {
	if tx != nil {
		if err == nil {
			if commitErr := tx.Commit().Error; commitErr != nil {
				log.Println("Commit Error : ", commitErr)
			}
		} else {
			tx.Rollback()
		}
	}
}
