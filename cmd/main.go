package main

import (
	"log"
	"os"

	"github.com/VladislavEF/todo-app"
	"github.com/VladislavEF/todo-app/pkg/handlers"
	"github.com/VladislavEF/todo-app/pkg/repository"
	"github.com/VladislavEF/todo-app/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatal("error initialazing configs:", err.Error())
	}

	if err := gotenv.Load(); err != nil {
		log.Fatal("error loading env variables:", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatal("failed to initialize database:", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handlers.NewHandler(service)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running server:", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
