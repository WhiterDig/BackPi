package main

import (
	"BackPi/global"
	"BackPi/repository"
	"BackPi/router"
	"BackPi/service"
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/juju/loggo"
	"net/http"
)

var (
	l   loggo.Logger
	ctx = context.Background()
)

func init() {
	l.SetLogLevel(loggo.DEBUG)
	//Create Repository
	cred := repository.NewCredentialRepo(ctx, l)
	//Create Service
	serv := service.NewCredential(ctx, cred, l)
	ctx = context.WithValue(ctx, global.CredentialService, serv)
}

func main() {
	fmt.Println("listening on port 8080")
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"POST"})
	router := router.NewRouter(ctx)
	l.Criticalf(http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)).Error())
}
