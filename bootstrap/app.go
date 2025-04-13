package bootstrap

import "go.mongodb.org/mongo-driver/v2/mongo"

type Application struct {
	Env *Env
	Database mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Database = NewDatabase(app.Env)
	return *app
}
