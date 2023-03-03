package main

import (
	"gin-framework-gs/routers"
)

func main() {

	// routers.TestServer().Run(":8080")
	routers.NewServer().Run(":8080")
}
