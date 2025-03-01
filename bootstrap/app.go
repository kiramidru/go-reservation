package bootstrap

import "go.mongodb.org/mongo-driver/mongo"

type Application struct {
    Env *Env
    Database mongo.Client
}

func App() Application {
    app := &Application{}
    app.Env = NewEnv()
    app.Database = NewMongoDatabase(app.Env)
    return *app
}

func (app *Application) CloseDBConnection() {
    CloseMongoDBConnection(app.Mongo)
}
