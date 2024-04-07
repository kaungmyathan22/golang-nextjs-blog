package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	logger "github.com/kaungmyathan22/golang-blog/cmd/common"
	"github.com/kaungmyathan22/golang-blog/cmd/database"
	"github.com/kaungmyathan22/golang-blog/cmd/routes"
)

func (app *Application) BootstrapApp() *pgx.Conn {
	logger.Init()
	app.LoadConfig()
	err, conn := app.SetupDatabase()
	if err != nil {
		panic(err)
		return nil
	}
	return conn
}

func (app *Application) SetupDatabase() (error, *pgx.Conn) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, app.config.DATABASE_URL)
	if err != nil {
		logger.Fatal("Error connecting to the database.. ")
		return err, nil
	}
	logger.Info("successfully connected to database.")
	app.db = database.New(conn)
	return nil, conn
}

func (app *Application) Serve() error {
	r := gin.Default()
	routes.InitRoutes(r, app.db)
	err := r.Run(app.config.PORT)
	if err != nil {
		logger.Error(err.Error())
		logger.Error("server quit")
	}
	return nil
}
