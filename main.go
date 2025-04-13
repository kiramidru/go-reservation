package main

import (
	"carbon/go-reservation/bootstrap"
	"time"
)

func main() {
	app := bootstrap.App()

	env := app.Env
	db := app.Database.Database(env.AppEnv)
	
	timeout := time.Duration(env.ContextTimeout) * time.Second
	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}
