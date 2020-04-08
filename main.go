package main

import (
	"hackathon/initDB"
	"hackathon/routers"
)

func main() {
	router := routers.SetupRouter()
	defer initDB.Db.Close()

	panic(router.Run())
}

