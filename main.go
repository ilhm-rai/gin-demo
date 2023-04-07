package main

import (
	"github.com/ilhm-rai/go-middleware/database"
	"github.com/ilhm-rai/go-middleware/router"
)

func main() {
	database.ConnectDB()
	r := router.Start()
	r.Run(":8080")
}
