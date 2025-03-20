package main

import (
	"golang-ecommerce/internal/handler/rest"
	"golang-ecommerce/internal/repository"
	"golang-ecommerce/internal/service"
	"golang-ecommerce/pkg/bcrypt"
	"golang-ecommerce/pkg/config"
	"golang-ecommerce/pkg/database/mariadb"
	"golang-ecommerce/pkg/jwt"
	"golang-ecommerce/pkg/middleware"
	"log"
)

func main() {
	config.LoadEnvironment()

	db, err := mariadb.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	err = mariadb.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	bcrypt := bcrypt.Init()
	jwt := jwt.Init()
	svc := service.NewService(repo, bcrypt, jwt)
	middleware := middleware.Init(svc, jwt)

	r := rest.NewRest(svc, middleware)
	r.MountEndpoint()
	r.Run()

}
