package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"polaris/user-service/database/migrations" // Importa el paquete de migraciones
	"polaris/user-service/database/seeders"
)

var rootCmd = &cobra.Command{
	Use:   "console",
	Short: "CLI para ejecutar comandos en user-service",
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Ejecuta las migraciones de la base de datos",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ejecutando migración...")
		migrations.RunMigration() // Llama a la función de migración
		fmt.Println("Migración completada.")
	},
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Ejecuta el seeder para cargar datos iniciales",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ejecutando seeder...")
		seeders.RunSeed()                 // Llama a la función de seeder
		fmt.Println("Seeder completado.") // Mensaje de éxito
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd) // Agrega el comando de migración al rootCmd
	rootCmd.AddCommand(seedCmd)    // Agrega el comando de seeder al rootCmd
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
