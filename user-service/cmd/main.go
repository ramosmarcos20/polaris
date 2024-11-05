package main

import (
	"fmt"
	"log"
	"polaris/user-service/config"
)

func main() {
	env := config.LoadEnv()
	DB, err := config.Connection(env)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	sqlDB, err := DB.DB()
	defer sqlDB.Close()

	fmt.Println("Successfully connected to DB")
}
