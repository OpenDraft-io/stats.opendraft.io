package initializers

import (
	"../../app/models"
	"../globals"
)

func InitDB(debug bool) {
	global.InitDB()
	if debug {
		global.DB.Debug()
	}
	migrateDB()
}

// Private Methods
func migrateDB() {
	global.DB.Connection.AutoMigrate(&models.Article{})
}
