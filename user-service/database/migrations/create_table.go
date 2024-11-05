package migrations

import (
	"log"
	"polaris/user-service/config"
	"polaris/user-service/internal/models/entities"
)

func RunMigration() {
	env := config.LoadEnv()
	db, err := config.Connection(env)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	err = db.AutoMigrate(&entities.User{})
	if err != nil {
		log.Fatalf("No se pudo apregar a de datos: %v", err)
	}

	log.Println("Migrate Complete")
}
