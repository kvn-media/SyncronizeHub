// main.go

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kvn-media/SyncronizeHub/configs"
	"github.com/kvn-media/SyncronizeHub/delivery"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = ""
	DB_NAME     = ""
	DB_DRIVER   = ""

	API_HOST = "localhost"
	API_PORT = ""

	APP_NAME = ""
)

func main() {
	setEnv()
	viewConfigs()

	appServer := delivery.NewAppServer()
	appServer.Run()
}

func setEnv() {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Auto setting environment variables")
	fmt.Println("You can disable this feature on main.go")

	os.Setenv("DB_HOST", DB_HOST)
	os.Setenv("DB_PORT", DB_PORT)
	os.Setenv("DB_USER", DB_USER)
	os.Setenv("DB_PASSWORD", DB_PASSWORD)
	os.Setenv("DB_NAME", DB_NAME)
	os.Setenv("DB_DRIVER", DB_DRIVER)

	os.Setenv("API_HOST", API_HOST)
	os.Setenv("API_PORT", API_PORT)

	os.Setenv("APP_NAME", APP_NAME)

	fmt.Println("Setting finished")
	fmt.Println(strings.Repeat("=", 50))
}

func viewConfigs() {
	config := configs.NewConfig()

	fmt.Println("configs: ")
	fmt.Println("db config:", config.DbConfig)
	fmt.Println("api config (port maybe auto set):", config.ApiConfig)

	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()
}
