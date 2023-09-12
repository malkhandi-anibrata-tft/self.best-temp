package main

import (
	"github.com/malkhandi-anibrata-tft/self.best-temp/config"
	"github.com/malkhandi-anibrata-tft/self.best-temp/db"
	"github.com/malkhandi-anibrata-tft/self.best-temp/pkg/api/seed"
	"github.com/malkhandi-anibrata-tft/self.best-temp/pkg/server"
	"github.com/spf13/viper"
)

func main() {
	config.Init()
	dbHandler := db.Init()

	if viper.GetBool("database.migrate") {
		seed.Migrate(dbHandler)
	}

	server.Setup(dbHandler)
}
