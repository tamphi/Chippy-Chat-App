package main

import (
	"log"
	"server/db"
	"server/internal/user"
	"server/internal/ws"
	"server/router"
)

func main() {
	dbInit, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to establish database connection: %s", err)
	}
	dbConn := dbInit.GetDB()
	//repository
	var userRepo user.Repository = user.NewRepository(dbConn)
	//repository service
	var userService user.Service = user.NewService(userRepo)
	//service controller ptr
	var userController *user.Controller = user.NewController(userService)

	var board *ws.Board = ws.NewBoard()
	var wsController *ws.Controller = ws.NewController(board, dbConn)

	go board.Run()
	router.InitRouter(userController, wsController)
	router.RunRouter("0.0.0.0:8080")
}
